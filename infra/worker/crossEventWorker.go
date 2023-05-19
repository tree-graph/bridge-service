package worker

import (
	"context"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/tree-graph/bridge-service/helpers"
	"github.com/tree-graph/bridge-service/infra/blockchain"
	"github.com/tree-graph/bridge-service/infra/database"
	"github.com/tree-graph/bridge-service/models"
	"gorm.io/gorm"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type CrossEventWorker struct {
	fetcher   *blockchain.EventFetcher
	vaultAddr string
}

// It will block current thread.
// Run a go routine for each chain.
func RunAllCrossEventWorker() error {
	workers, err := setup()
	if err != nil {
		logrus.Error("setup fail")
		return err
	}
	logrus.Info("worker count ", len(workers))

	for _, worker := range workers {
		logrus.WithField("chain", worker.ChainId()).
			WithField("vault addr", worker.vaultAddr).
			Info("start chain")
		go worker.Run()
	}

	// block main thread
	exitSignal := make(chan os.Signal)
	signal.Notify(exitSignal, syscall.SIGINT, syscall.SIGTERM)
	<-exitSignal
	return nil
}

// setup all chains, create a worker for each chain.
func (worker CrossEventWorker) ChainId() int64 {
	return worker.fetcher.Handler.ChainId
}
func setup() ([]*CrossEventWorker, error) {
	var workers []*CrossEventWorker
	var chains []models.Chain
	if err := database.DB.Find(&chains).Error; err != nil {
		return nil, err
	}
	for _, chain := range chains {
		if len(chain.VaultAddr) == 0 {
			return nil, fmt.Errorf("vault address is empty, chain %v", chain.Id)
		}
		var cursor models.ChainCursor
		if err := database.DB.Where("id = ?", chain.Id).Take(&cursor).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// create cursor
				if errCreate := database.DB.Create(&models.ChainCursor{
					Id:          chain.Id,
					Block:       0,
					LatestBlock: 0,
				}).Error; errCreate != nil {
					return nil, errCreate
				}
			}
			return nil, err
		}

		if err := blockchain.AddChainClient(chain); err != nil {
			return nil, err
		}

		fetcher, err := blockchain.NewEventFetcher(chain.Id, chain.VaultAddr)
		if err != nil {
			return nil, err
		}

		worker := &CrossEventWorker{
			fetcher:   fetcher,
			vaultAddr: chain.VaultAddr,
		}
		workers = append(workers, worker)
	}

	return workers, nil
}

func (worker CrossEventWorker) Run() {
	for {
		sleepT, err := worker.doWork()
		if err != nil {
			logrus.WithError(err).
				WithField("chainId", worker.fetcher.Handler.ChainId).
				Error("unhandled error")
			time.Sleep(time.Second)
		} else if sleepT > 0 {
			time.Sleep(time.Duration(sleepT) * time.Second)
		}
	}
}
func (worker CrossEventWorker) doWork() (int, error) {
	// Fetch cursor each time, in case an operator changes it directly in the database.
	var cursor models.ChainCursor
	var fetcher = worker.fetcher
	chainId := fetcher.Handler.ChainId
	if err := database.DB.Where("id = ?", chainId).Take(&cursor).Error; err != nil {
		return 0, err
	}
	chain, err := models.GetChain(chainId)
	if err != nil {
		return 0, err
	}

	currentBlock := cursor.Block + 1
	logEntry := logrus.WithFields(logrus.Fields{
		"ChainId": chainId, "currentBlock": currentBlock,
	})

	if cursor.LatestBlock < currentBlock+int64(chain.DelayBlock) {
		n, err := fetcher.Handler.Client.BlockNumber(context.Background())
		if err != nil {
			return 0, err
		}
		logEntry.Info("latest block ", n)
		if n != uint64(cursor.LatestBlock) {
			logrus.WithFields(logrus.Fields{
				"latest_block": n, "chain": chainId,
			}).Info("update latest_block")
			return 0, database.DB.Model(cursor).
				Where("id = ?", chainId).
				Update("latest_block", n).Error
		}
		return 1, nil
	}

	parsedLogs, err := fetcher.Fetch(uint64(currentBlock))
	if err != nil {
		return 0, err
	}
	if len(parsedLogs) == 0 {
		logEntry.Debug("no logs")
		return 0, database.DB.Model(&cursor).Where("id = ?", chainId).Update("block", currentBlock).Error
	}

	block, blockErr := fetcher.Handler.BlockByNumber(helpers.Uint64ToBigInt(parsedLogs[0].Raw.BlockNumber))
	if blockErr != nil {
		logEntry.WithError(blockErr).Error("fetch block fail")
		return 1, nil
	}

	crossInfoArr, crossItemsArr := BuildCrossInfo(block, parsedLogs, chainId)

	allItemCount := 0
	allTxErr := database.DB.Transaction(func(tx *gorm.DB) error {
		if infoE := tx.Create(crossInfoArr).Error; infoE != nil {
			return infoE
		}
		var allItems []models.CrossItem
		for i, info := range crossInfoArr {
			for _, item := range crossItemsArr[i] {
				item.CrossInfoId = info.Id
				allItems = append(allItems, item)
			}
			allItemCount += len(crossItemsArr[i])
		}
		if itemE := tx.Create(allItems).Error; itemE != nil {
			return itemE
		}
		return tx.Model(&cursor).Where("id = ?", chainId).Update("block", currentBlock).Error
	})

	if allTxErr != nil {
		logEntry.WithError(allTxErr).Error("save crossing info error")
		return 1, nil
	}

	logEntry.WithFields(logrus.Fields{
		"eventCount": len(crossInfoArr), "itemCount": allItemCount,
	}).Infof("crossing info saved successfully")

	return 0, nil
}
