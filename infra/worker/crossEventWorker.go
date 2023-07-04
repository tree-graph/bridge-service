package worker

import (
	"fmt"
	sdk "github.com/Conflux-Chain/go-conflux-sdk"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/tree-graph/bridge-service/infra/blockchain"
	"github.com/tree-graph/bridge-service/infra/database"
	"github.com/tree-graph/bridge-service/models"
	"gorm.io/gorm"
	"time"
)

type CrossEventWorker struct {
	fetcher   *blockchain.IEventFetcher
	vaultAddr string
	Chain     models.Chain
}

// Run a go routine for each chain.
func RunAllCrossEventWorker() error {
	workers, err := setup()
	if err != nil {
		logrus.Error("setup fail")
		return err
	}
	logrus.Info("worker count ", len(workers))

	for _, worker := range workers {
		logrus.WithFields(logrus.Fields{
			"chain":      worker.Chain.Name,
			"vault addr": worker.vaultAddr,
		}).Info("start chain")
		go worker.Run()
	}

	return nil
}

// setup all chains, create a worker for each chain.
func setup() ([]*CrossEventWorker, error) {
	var workers []*CrossEventWorker
	var chains []models.Chain
	if err := database.DB.Find(&chains).Error; err != nil {
		return nil, err
	}
	for _, chain := range chains {
		if !chain.Enabled {
			logrus.WithFields(logrus.Fields{"name": chain.Name}).Info("chain is disabled")
			continue
		}
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
			return nil, errors.WithMessage(err, "check event cursor fail")
		}
		var worker *CrossEventWorker
		var err error
		if chain.ChainType == models.CfxChain {
			worker, err = buildCfxWorker(chain)
		} else if chain.ChainType == models.EvmChain {
			worker, err = buildEvmWorker(chain)
		} else {
			err = fmt.Errorf("unsportted chain type [%v]", chain.ChainType)
		}
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"name": chain.Name, "type": chain.ChainType,
			}).Debug("build chain fetcher fail")
			return nil, err
		}
		workers = append(workers, worker)
	}

	return workers, nil
}

func buildCfxWorker(chain models.Chain) (*CrossEventWorker, error) {
	cfxClient, err := sdk.NewClient(chain.Rpc)
	if err != nil {
		logrus.Debug("create cfx client fail")
		return nil, err
	}

	cfxFetcher, err := blockchain.NewCfxFetcher(chain.ChainId, chain.VaultAddr, cfxClient)
	if err != nil {
		return nil, err
	}

	var iFetcher blockchain.IEventFetcher
	iFetcher = cfxFetcher

	worker := &CrossEventWorker{
		fetcher:   &iFetcher,
		vaultAddr: chain.VaultAddr,
		Chain:     chain,
	}
	return worker, nil
}

func buildEvmWorker(chain models.Chain) (*CrossEventWorker, error) {
	if err := blockchain.AddChainClient(chain); err != nil {
		return nil, err
	}

	fetcher, err := blockchain.NewEventFetcher(chain.Id, chain.VaultAddr)
	if err != nil {
		return nil, err
	}

	var iFetcher blockchain.IEventFetcher
	iFetcher = fetcher
	worker := &CrossEventWorker{
		fetcher:   &iFetcher,
		vaultAddr: chain.VaultAddr,
		Chain:     chain,
	}
	return worker, nil
}

func (worker CrossEventWorker) Run() {
	for {
		sleepT, err := worker.doWork()
		if err != nil {
			logrus.WithError(err).
				WithField("chain", worker.Chain.Name).
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
	fetcher := *worker.fetcher
	chainDbId := worker.Chain.Id
	if err := database.DB.Where("id = ?", chainDbId).Take(&cursor).Error; err != nil {
		return 0, err
	}
	chain, err := models.GetChain(chainDbId)
	if err != nil {
		return 0, err
	}

	currentBlock := cursor.Block + 1
	logEntry := logrus.WithFields(logrus.Fields{
		"ChainId": chainDbId, "currentBlock": currentBlock, "latest": cursor.LatestBlock,
	})

	if cursor.LatestBlock < currentBlock+int64(chain.DelayBlock) {
		n, err := fetcher.BlockNumber()
		if err != nil {
			return 0, err
		}
		logEntry.Info("latest block ", n)
		if n != uint64(cursor.LatestBlock) {
			logrus.WithFields(logrus.Fields{
				"latest_block": n, "chain": chainDbId,
			}).Info("update latest_block")
			return 0, database.DB.Model(cursor).
				Where("id = ?", chainDbId).
				Update("latest_block", n).Error
		}
		return 1, nil
	}

	blockTime, parsedLogs, err := fetcher.Fetch(uint64(currentBlock))
	if err != nil {
		return 0, err
	}
	if len(parsedLogs) == 0 {
		logEntry.Debug("no logs")
		return 0, database.DB.Model(&cursor).Where("id = ?", chainDbId).Update("block", currentBlock).Error
	}

	crossInfoArr, crossItemsArr := BuildCrossInfo(*blockTime, parsedLogs, chainDbId)

	err = SaveCrossRequest(crossInfoArr, crossItemsArr, chainDbId, currentBlock)
	if err != nil {
		return 1, err
	}

	return 0, nil
}

func SaveCrossRequest(crossInfoArr []models.CrossInfo, crossItemsArr [][]models.CrossItem,
	chainDbId int64, currentBlock int64) error {
	allItemCount := 0
	allTxErr := database.DB.Transaction(func(tx *gorm.DB) error {
		if len(crossItemsArr) > 0 {
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
		}
		return tx.Model(&models.ChainCursor{}).Where("id = ?", chainDbId).Update("block", currentBlock).Error
	})

	if allTxErr != nil {
		return errors.WithMessage(allTxErr, "save crossing info error")
	}

	logrus.WithFields(logrus.Fields{
		"eventCount": len(crossInfoArr), "itemCount": allItemCount, "cursor": currentBlock,
	}).Infof("crossing info saved successfully")

	return nil
}
