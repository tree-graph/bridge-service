package worker

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/sirupsen/logrus"
	"github.com/tree-graph/bridge-service/infra/blockchain"
	"github.com/tree-graph/bridge-service/infra/contracts/vault"
	"github.com/tree-graph/bridge-service/infra/database"
	"github.com/tree-graph/bridge-service/models"
	"gorm.io/gorm"
	"math/big"
	"time"
)

/*
 * Workflow:
 *	0 check whether there is a pooled(undergoing) claiming record.
 *	1 check pooled record: send tx or check receipt, then finish it, say, delete it from pool and save it in history.
 *	2 find one crossInfo in DB with id > previous claiming cursor, order by id asc
 *	3 put a record in pooled claiming table and update cursor.
 *	4 back to step 0.
 */

type ClaimWorker struct {
	ChainId       int64
	keyPair       *ecdsa.PrivateKey
	address       *common.Address
	evmHandler    blockchain.EvmHandler
	DelayForError int
}

func RunAllClaimWorker() error {
	workers, err := setupClaimWorkers()
	if err != nil {
		logrus.WithError(err).Error("setupClaimWorkers fail")
		return err
	}
	for _, worker := range workers {
		go worker.Run()
	}
	return nil
}

func (worker ClaimWorker) Run() {
	for {
		sleepT, err := worker.doWork()
		if err != nil {
			logrus.WithError(err).
				WithField("chainId", worker.ChainId).
				Error("unhandled claim worker error")
			time.Sleep(time.Second)
		} else if sleepT > 0 {
			time.Sleep(time.Duration(sleepT) * time.Second)
		}
	}
}

func (worker ClaimWorker) doWork() (int, error) {
	// check whether there is a undergoing task
	var pooledClaim models.ClaimPool
	hasPooledClaim := true
	if err := database.DB.Where("target_chain=?", worker.ChainId).
		Take(&pooledClaim).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			hasPooledClaim = false
		} else {
			logrus.WithError(err).
				WithField("chainId", worker.ChainId).
				Error("check pooled claim error")
			return 0, err
		}
	}

	if hasPooledClaim {
		return worker.checkPooledClaim(pooledClaim)
	}
	// Fetch cursor each time, in case an operator changes it directly in the database.
	var claimCursor models.ClaimCursor
	if err := database.DB.Where("target_chain=?", worker.ChainId).
		Take(&claimCursor).Error; err != nil {
		return 5, err
	}

	var crossInfo models.CrossInfo
	if err := database.DB.Where("target_chain = ? and id > ?", worker.ChainId, claimCursor.CrossInfoId).
		Order("id asc").
		Take(&crossInfo).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logrus.WithFields(logrus.Fields{
				"chain": worker.ChainId,
			}).Debug("no claim task")
			return 5, nil
		}
		return 5, err
	}
	// put a pooled claim into DB and update cursor.
	// next step, sending tx
	if allTxErr := database.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(claimCursor).
			Where("target_chain=?", claimCursor.TargetChain).
			Update("cross_info_id", crossInfo.Id).Error; err != nil {
			return err
		}
		if err := tx.Create(&models.ClaimPool{
			CrossInfoId:    crossInfo.Id,
			TargetChain:    crossInfo.TargetChain,
			TargetContract: crossInfo.TargetContract,
			TxnHash:        fmt.Sprintf("placeholder %v", crossInfo.Id),
			From:           worker.address.Hex(),
			Nonce:          0,
			Step:           models.ClaimStepSendingTx,
			Status:         nil,
		}).Error; err != nil {
			return err
		}
		return nil
	}); allTxErr != nil {
		return 5, allTxErr
	}
	return 0, nil
}

func setupClaimWorkers() ([]*ClaimWorker, error) {
	var workers []*ClaimWorker
	var chains []models.Chain
	if err := database.DB.Find(&chains).Error; err != nil {
		return nil, err
	}
	for _, chain := range chains {

		if err := blockchain.AddChainClient(chain); err != nil {
			return nil, err
		}

		secret, err := models.GetSecret(chain.Id)
		if err != nil {
			logrus.WithFields(logrus.Fields{"chain": chain.Id, "name": chain.Name}).Error("secret not found")
			return nil, err
		}
		// prepare claiming cursor
		var claimCursor models.ClaimCursor
		err = database.DB.Where("target_chain=?", chain.Id).Take(&claimCursor).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if err = database.DB.Create(&models.ClaimCursor{TargetChain: chain.Id, CrossInfoId: 0}).Error; err != nil {
				logrus.WithFields(logrus.Fields{"chainId": chain.Id}).
					Error("create claim cursor error")
				return nil, err
			}
		} else if err != nil {
			logrus.WithFields(logrus.Fields{"chainId": chain.Id}).
				Error("take claim cursor error")
			return nil, err
		}

		// prepare private key
		privateKey, address, err := blockchain.CreateKeyPair(secret.Private[2:])
		if err != nil {
			logrus.WithError(err).WithFields(logrus.Fields{
				"chainId": chain.Id,
			}).Error("create key pair error")
			return nil, err
		}
		evmHandler := blockchain.GetEvmHandler(chain.Id)

		if secret.Address == "" {
			if err := database.DB.Model(&secret).Where("id=?", secret.Id).
				Update("address", address.Hex()).Error; err != nil {
				logrus.WithError(err).WithFields(logrus.Fields{
					"chainId": chain.Id, "secretId": secret.Id,
				}).Error("update address for secret fail")
				return nil, err
			}
			balance, err := evmHandler.Client.BalanceAt(context.Background(), *address, nil)
			if err != nil {
				logrus.WithError(err).WithFields(logrus.Fields{
					"chainId": chain.Id, "address": *address,
				}).Error("get balance fail")
				return nil, err
			}
			if (*balance).Cmp(big.NewInt(0)) == 0 {
				logrus.WithError(err).WithFields(logrus.Fields{
					"chainId": chain.Id, "address": *address,
				}).Fatal("balance is zero")
			}
		}

		workers = append(workers, &ClaimWorker{
			ChainId:       chain.Id,
			keyPair:       privateKey,
			address:       address,
			evmHandler:    evmHandler,
			DelayForError: 60 * 10, // delay 10 minutes
		})
	}
	return workers, nil
}

func (worker ClaimWorker) Claim(crossInfo models.CrossInfo) (string, *uint64, error) {
	var items []models.CrossItem
	if err := database.DB.Where("cross_info_id=?", crossInfo.Id).Find(&items).Error; err != nil {
		logrus.WithError(err).WithFields(logrus.Fields{
			"crossId": crossInfo.Id, "chain": crossInfo.TargetChain,
		}).Error("cross items not found")
		return "", nil, err
	}

	targetChain, err := models.GetChain(crossInfo.TargetChain)
	if err != nil {
		return "", nil, err
	}

	var tokenIds []*big.Int
	var amounts []*big.Int
	var uris []string
	for _, item := range items {
		tokenIds = append(tokenIds, common.HexToHash(item.TokenId).Big())
		amounts = append(tokenIds, common.HexToHash(item.Amount).Big())
		uris = append(uris, item.Uri)
	}
	request := vault.VaultCrossRequest{
		Asset:          common.HexToAddress(crossInfo.Asset),
		From:           common.HexToAddress(crossInfo.From),
		TokenIds:       tokenIds,
		Amounts:        amounts,
		Uris:           uris,
		ToChainId:      big.NewInt(crossInfo.TargetChain),
		TargetContract: common.HexToAddress(crossInfo.TargetContract),
		UserNonce:      big.NewInt(crossInfo.UserNonce),
		Raw:            types.Log{},
	}

	evmHandler := blockchain.GetEvmHandler(crossInfo.TargetChain)

	claimTxHash, nonce, err := blockchain.SendClaimTx(worker.keyPair, worker.address, targetChain.VaultAddr,
		big.NewInt(crossInfo.SourceChain), request, *evmHandler.Client)
	if err != nil {
		return "", nil, err
	}

	return claimTxHash, nonce, nil
}

func (worker ClaimWorker) checkPooledClaim(claim models.ClaimPool) (int, error) {
	if claim.Step == models.ClaimStepSendingTx {
		return worker.sendClaimTx(claim)
	} else if claim.Step == models.ClaimStepWaitingTx {
		receipt, err := worker.evmHandler.Client.TransactionReceipt(
			context.Background(),
			common.HexToHash(claim.TxnHash),
		)
		if errors.Is(err, ethereum.NotFound) {
			logrus.
				WithFields(logrus.Fields{"chainId": worker.ChainId, "crossInfoId": claim.CrossInfoId,
					"txnHash": claim.TxnHash}).
				Info("receipt not found")
			return 5, nil
		}
		if err != nil {
			return 0, err
		}
		return worker.checkReceipt(claim, receipt)
	} else {
		worker.notifyError("unknown claiming step "+claim.Step, "")
		return worker.DelayForError, nil
	}
}

func (worker ClaimWorker) sendClaimTx(claim models.ClaimPool) (int, error) {
	var crossInfo models.CrossInfo
	if err := database.DB.Where("id=?", claim.CrossInfoId).
		Take(&crossInfo).Error; err != nil {
		logrus.WithError(err).
			WithFields(logrus.Fields{"chainId": worker.ChainId, "crossInfoId": claim.CrossInfoId}).
			Error("querying cross info failed when sending claim tx")
		return 0, err
	}
	claimTxHash, claimNonce, err := worker.Claim(crossInfo)
	if err != nil {
		errorType := blockchain.GetRpcErrorType(err.Error())
		resend, _ := errorType.NeedResendWhenSentErr()
		if resend {
			_ = moveClaimFromPoolToHistory(claim, uint64(errorType), "sending tx fail")
			return 0, err
		}
		worker.notifyError("sending tx fail:"+err.Error(), "")
		return worker.DelayForError, err
	}
	if err := database.DB.Model(claim).Where("id=?", claim.Id).
		Updates(&models.ClaimPool{
			TxnHash: claimTxHash,
			From:    worker.address.Hex(),
			Nonce:   *claimNonce,
			Step:    models.ClaimStepWaitingTx,
			Status:  nil,
		}).Error; err != nil {
		logrus.WithError(err).
			WithFields(logrus.Fields{"chainId": worker.ChainId, "crossInfoId": claim.CrossInfoId}).
			Error("update claim info after sending tx fail")
		return 0, err
	}
	return 0, nil
}

func (worker ClaimWorker) checkReceipt(claim models.ClaimPool, receipt *types.Receipt) (int, error) {
	if receipt.Status == 1 {
		if err := moveClaimFromPoolToHistory(claim, receipt.Status, "OK"); err != nil {
			return 0, err
		}
		return 0, nil
	}
	worker.notifyError(fmt.Sprintf("tx status fail: [%v]", receipt.Status), claim.TxnHash)
	// reach code here, we have problem
	return worker.DelayForError, nil
}

func (worker ClaimWorker) notifyError(info string, txnHash string) {
	logrus.WithFields(logrus.Fields{
		"chainId": worker.ChainId, "tx": txnHash, "info": info,
	}).Error("claim transaction fail")
}

func moveClaimFromPoolToHistory(claim models.ClaimPool, status uint64, comment string) error {
	return database.DB.Transaction(func(tx *gorm.DB) error {
		// delete pooled claim
		if err := tx.Delete(claim).Error; err != nil {
			return err
		}
		// add history
		return tx.Create(&models.ClaimHistory{
			CrossInfoId:    claim.CrossInfoId,
			TargetChain:    claim.TargetChain,
			TargetContract: claim.TargetContract,
			TxnHash:        claim.TxnHash,
			From:           claim.From,
			Nonce:          claim.Nonce,
			Comment:        comment,
			Status:         status,
		}).Error
	})
}
