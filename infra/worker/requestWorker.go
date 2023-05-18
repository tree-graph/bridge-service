package worker

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/sirupsen/logrus"
	"github.com/tree-graph/bridge-service/helpers"
	"github.com/tree-graph/bridge-service/infra/blockchain"
	"github.com/tree-graph/bridge-service/infra/contracts/vault"
	"github.com/tree-graph/bridge-service/infra/database"
	"github.com/tree-graph/bridge-service/models"
	"gorm.io/gorm"
	"strconv"
	"time"
)

func InitCrossReqWorker() {
	var config models.Config
	err := database.DB.Where("name = ? ", models.CrossReqId).Take(&config).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		config.Name = models.CrossReqId
		config.Content = "0"
		if errCreate := database.DB.Create(&config); errCreate != nil {
			logrus.WithError(err).Fatal("InitCrossReqWorker, create config fail")
		}
	} else if err != nil {
		logrus.WithError(err).Fatal("InitCrossReqWorker fail")
	}
	SetupChains()
}
func SetupChains() {
	var chains []models.Chain
	err := database.DB.Find(&chains).Error
	helpers.CheckFatalError("load all chain", err)
	for _, c := range chains {
		blockchain.AddChainClient(c)
	}
}

// This worker fetches transaction receipt, parses logs, and saves crossing information to DB.
func CrossRequestWorker() {
	for {
		sleepT, err := runRequestWorker()
		if err != nil {
			logrus.WithError(err).Error("unhandled error")
			time.Sleep(time.Second)
		} else if sleepT > 0 {
			time.Sleep(time.Duration(sleepT) * time.Second)
		}
	}
}

var notFoundTimes uint64 = 0

// Processes crossing requests one by one, according to the cursor in the database.
// returns n as the desired number of seconds to sleep.
func runRequestWorker() (int, error) {
	// Fetch cursor each time, in case an operator changes it directly in the database.
	var config models.Config
	if err := database.DB.Take(&config, "name", models.CrossReqId).Error; err != nil {
		logrus.WithError(err).Error("fetch crossing record cursor fail")
		// wait a moment and allow automatic recovery
		return 1, nil
	}

	preId, err := strconv.ParseUint(config.Content, 10, 64)
	if err != nil {
		logrus.WithError(err).WithField("value", config.Content).
			Error("invalid cursor in DB")
		return 1, nil
	}

	var req models.CrossRequest
	err = database.DB.Where("id > ?", preId).Order("id asc").First(&req).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		if notFoundTimes%10 == 0 {
			logrus.Debug("no task")
		}
		notFoundTimes += 1
		return 1, nil
	}
	logEntry := logrus.WithFields(logrus.Fields{
		"reqId": req.Id, "txHash": req.Hash, "chainId": req.ChainId,
	})
	logEntry.Info("handle task")
	if len(req.Hash) != 66 {
		logEntry.Warn("invalid hash")
		err := saveAsInvalidHash(req, config, "invalid hash", models.InvalidHash)
		return 0, err
	}
	evmClient := blockchain.GetEvmHandler(req.ChainId)
	parsedLogs, e := evmClient.GetCrossChainRequest(req.Hash)
	if errors.Is(e, ethereum.NotFound) {
		_, _, errTx := evmClient.TransactionByHash(req.Hash)
		if errors.Is(errTx, ethereum.NotFound) {
			logEntry.Warn("both transaction and it's receipt are missing")
			err := saveAsInvalidHash(req, config, "tx not found", models.TxNotFound)
			return 0, err
		}
		if errTx != nil {
			logEntry.WithError(e).Error("check tx existence error")
			return 1, nil
		}
	}
	if e != nil {
		logEntry.WithError(e).Error("parse request fail")
		return 1, nil
	}
	if len(parsedLogs) == 0 {
		logEntry.Warn("event is empty")
		err := saveAsInvalidHash(req, config, "empty event", models.EmptyEvent)
		return 0, err
	}
	block, blockErr := evmClient.BlockByNumber(helpers.Uint64ToBigInt(parsedLogs[0].Raw.BlockNumber))
	if blockErr != nil {
		logEntry.WithError(blockErr).Error("fetch block fail")
		return 1, nil
	}
	crossInfoArr, crossItemsArr := buildCrossInfo(block, parsedLogs, req)
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
		configE := tx.Model(&config).Updates(&models.Config{
			Name:    models.CrossReqId,
			Content: fmt.Sprint(req.Id),
		}).Error
		return configE
	})
	if allTxErr != nil {
		logEntry.WithError(allTxErr).Error("save crossing info error")
		return 1, nil
	}
	logEntry.WithFields(logrus.Fields{
		"count": len(crossInfoArr), "itemCount": allItemCount,
	}).Infof("crossing info saved successfully")
	return 0, nil
}

func saveAsInvalidHash(req models.CrossRequest, config models.Config, result string, status int) error {
	return database.DB.Transaction(func(tx *gorm.DB) error {
		err1 := tx.Model(&req).Updates(&models.CrossRequest{
			Id:     req.Id,
			Status: status,
			Result: result,
		}).Error
		if err1 != nil {
			return err1
		}
		err2 := tx.Model(&config).Updates(&models.Config{
			Name:    models.CrossReqId,
			Content: fmt.Sprint(req.Id),
		}).Error
		return err2
	})
}

func buildCrossInfo(block *types.Block, parseLogs []*vault.VaultCrossRequest, req models.CrossRequest) ([]models.CrossInfo, [][]models.CrossItem) {
	blockTime := time.Unix(int64(block.Time()), 0)
	var crossInfoArr []models.CrossInfo
	var crossItemsArr = make([][]models.CrossItem, len(parseLogs))

	for infoIndex, crossLog := range parseLogs {
		var crossInfo = models.CrossInfo{
			SourceChain:    req.ChainId,
			TxnHash:        req.Hash,
			Asset:          crossLog.Asset.Hex(),
			From:           crossLog.From.Hex(),
			TargetChain:    crossLog.ToChainId.Int64(),
			TargetContract: crossLog.TargetContract.Hex(),
			UserNonce:      crossLog.UserNonce.Int64(),
			BlockNumber:    crossLog.Raw.BlockNumber,
			BlockTime:      &blockTime,
		}
		crossInfoArr = append(crossInfoArr, crossInfo)
		var itemArr = make([]models.CrossItem, len(crossLog.TokenIds))
		for i, tokenId := range crossLog.TokenIds {
			var item = models.CrossItem{
				TokenId: common.BigToHash(tokenId).Hex(),
				Amount:  common.BigToHash(crossLog.Amounts[i]).Hex(),
				Uri:     crossLog.Uris[i],
			}
			itemArr[i] = item
		}
		crossItemsArr[infoIndex] = itemArr
	}
	return crossInfoArr, crossItemsArr
}
