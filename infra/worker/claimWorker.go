package worker

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	sdk "github.com/Conflux-Chain/go-conflux-sdk"
	"github.com/Conflux-Chain/go-conflux-sdk/types/cfxaddress"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/sirupsen/logrus"
	"github.com/tree-graph/bridge-service/infra/blockchain"
	"github.com/tree-graph/bridge-service/infra/contracts/vault"
	"github.com/tree-graph/bridge-service/infra/database"
	"github.com/tree-graph/bridge-service/models"
	"gorm.io/gorm"
	"math/big"
	"strings"
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
type IClaimWorker interface {
	Init()
	DoWork() (int, error)
	GetChainId() int64
}

type ClaimWorker struct {
	keyPair       *ecdsa.PrivateKey
	address       *common.Address
	evmHandler    blockchain.EvmHandler
	ChainType     string
	DelayForError int
	Chain         models.Chain
}

func RunAllClaimWorker() error {
	workers, err := setupClaimWorkers()
	if err != nil {
		logrus.WithError(err).Error("setupClaimWorkers fail")
		return err
	}
	logrus.Info("claim workers count ", len(workers))
	for _, worker := range workers {
		go Run(*worker)
	}
	return nil
}

func Run(worker IClaimWorker) {
	for {
		sleepT, err := worker.DoWork()
		if err != nil {
			logrus.WithError(err).
				WithField("chainId", worker.GetChainId()).
				Error("unhandled claim worker error")
			time.Sleep(time.Second)
		} else if sleepT > 0 {
			time.Sleep(time.Duration(sleepT) * time.Second)
		}
	}
}

func (worker ClaimWorker) Init() {

}
func (worker ClaimWorker) GetChainId() int64 {
	return worker.Chain.ChainId
}

func HasPooledClaim(chainDbId int64, name string) (bool, models.ClaimPool, error) {
	// check whether there is a undergoing task
	var pooledClaim models.ClaimPool
	hasPooledClaim := true
	if err := database.DB.Where("target_chain=?", chainDbId).
		Take(&pooledClaim).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			hasPooledClaim = false
		} else {
			logrus.WithError(err).
				WithField("chain", name).
				Error("check pooled claim error")
			return false, pooledClaim, err
		}
	}
	return hasPooledClaim, pooledClaim, nil
}

func (worker ClaimWorker) DoWork() (int, error) {
	hasPooledClaim, pooledClaim, err := HasPooledClaim(worker.Chain.Id, worker.Chain.Name)
	if err != nil {
		return 0, err
	}
	if hasPooledClaim {
		return worker.checkPooledClaim(pooledClaim)
	}
	return CheckClaimTask(worker.Chain.Id, worker.Chain.Name)
}

func CheckClaimTask(chainDbId int64, chainName string) (int, error) {
	// Fetch cursor each time, in case an operator changes it directly in the database.
	var claimCursor models.ClaimCursor
	if err := database.DB.Where("target_chain=?", chainDbId).
		Take(&claimCursor).Error; err != nil {
		return 5, err
	}
	var crossInfo models.CrossInfo
	if err := database.DB.Where("target_chain = ? and id > ?", chainDbId, claimCursor.CrossInfoId).
		Order("id asc").
		Take(&crossInfo).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logrus.WithFields(logrus.Fields{
				"chain.id": chainDbId, "name": chainName,
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
			From:           "",
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

func setupClaimWorkers() ([]*IClaimWorker, error) {
	var workers []*IClaimWorker
	var chains []models.Chain
	if err := database.DB.Find(&chains).Error; err != nil {
		return nil, err
	}
	logrus.Debug("chain count ", len(chains))
	for _, chain := range chains {

		if err := blockchain.AddChainClient(chain); err != nil {
			return nil, err
		}

		secret, err := models.GetSecret(chain.Id)
		if err != nil {
			logrus.WithFields(logrus.Fields{"chain": chain.Id, "name": chain.Name}).
				Error("secret not found")
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
		logrus.Debug("create claiming worker for chain ", chain.Name)
		var claimWorker *IClaimWorker
		var errTmp error
		if chain.ChainType == models.EvmChain {
			claimWorker, errTmp = createEvmWorker(chain, secret)
		} else if chain.ChainType == models.CfxChain {
			claimWorker, errTmp = createCfxWorker(chain, secret)
		} else {
			errTmp = fmt.Errorf("unsupported chain type [%v], id %v",
				chain.ChainType, chain.Id)
		}
		if errTmp != nil {
			logrus.Debug("create worker fail, chain ", chain.Name)
			return nil, errTmp
		}
		workers = append(workers, claimWorker)
	}
	return workers, nil
}

func createCfxWorker(chain models.Chain, secret models.Secret) (*IClaimWorker, error) {
	cfxClient, err := sdk.NewClient(chain.Rpc)
	if err != nil {
		logrus.Debug("create cfx client fail")
		return nil, err
	}
	cfxClient.SetNetworkId(uint32(chain.ChainId))
	am := sdk.NewAccountManager("./keystore/", uint32(chain.ChainId))
	cfxClient.SetAccountManager(am)

	address, err := am.ImportKey(secret.Private, "")
	if err == nil {
		// things go well
	} else if strings.Contains(err.Error(), keystore.ErrAccountAlreadyExists.Error()) {
		// nothing
		addressD := cfxaddress.MustNewFromBase32(secret.Address)
		address = addressD
	} else {
		logrus.Debug("import cfx key fail")
		return nil, err
	}
	logrus.WithFields(logrus.Fields{
		"chain.Id": chain.Id, "chain": chain.Name,
	}).Info("cfx address is ", address.String())

	balanceTmp, errGetBalance := cfxClient.GetBalance(address)
	var balance *big.Int
	if balanceTmp != nil {
		balance = balanceTmp.ToInt()
	}
	if err := afterCreateAccount(secret, chain, address.String(), balance, errGetBalance); err != nil {
		return nil, err
	}
	var worker IClaimWorker
	worker = CfxClaimWorker{
		Chain:         chain,
		ChainId:       chain.ChainId,
		address:       address,
		CfxClient:     cfxClient,
		DelayForError: 60 * 10,
	}
	worker.Init()
	return &worker, nil
}

func createEvmWorker(chain models.Chain, secret models.Secret) (*IClaimWorker, error) {
	// prepare private key
	privateKey, address, err := blockchain.CreateKeyPair(secret.Private[2:])
	if err != nil {
		logrus.WithError(err).WithFields(logrus.Fields{
			"chainId": chain.ChainId,
		}).Error("create key pair error")
		return nil, err
	}
	evmHandler := blockchain.GetEvmHandler(chain.Id)
	balance, errGetBalance := evmHandler.Client.BalanceAt(context.Background(), *address, nil)

	if err := afterCreateAccount(secret, chain, address.Hex(), balance, errGetBalance); err != nil {
		return nil, err
	}
	var claimWorker IClaimWorker
	claimWorker = ClaimWorker{
		Chain:         chain,
		keyPair:       privateKey,
		address:       address,
		evmHandler:    evmHandler,
		DelayForError: 60 * 10, // delay 10 minutes
	}
	return &claimWorker, nil
}

func afterCreateAccount(secret models.Secret, chain models.Chain,
	accountStr string,
	balance *big.Int, errGetBalance error) error {
	if secret.Address == "" {
		if err := database.DB.Model(&secret).Where("id=?", secret.Id).
			Update("address", accountStr).Error; err != nil {
			logrus.WithError(err).WithFields(logrus.Fields{
				"chainId": chain.Id, "secretId": secret.Id,
			}).Error("update address for secret fail")
			return err
		}
		if errGetBalance != nil {
			logrus.WithError(errGetBalance).WithFields(logrus.Fields{
				"chainId": chain.Id, "address": accountStr,
			}).Error("get balance fail")
			return errGetBalance
		}
		logrus.Info("balance ", balance)
		if (*balance).Cmp(big.NewInt(0)) == 0 {
			logrus.WithFields(logrus.Fields{
				"chainId": chain.Id, "address": accountStr,
			}).Fatal("balance is zero")
		}
	}
	return nil
}

func BuildCrossRequest(crossInfo models.CrossInfo) (*vault.VaultCrossRequest, error) {
	var items []models.CrossItem
	if err := database.DB.Where("cross_info_id=?", crossInfo.Id).Find(&items).Error; err != nil {
		logrus.WithError(err).WithFields(logrus.Fields{
			"crossId": crossInfo.Id, "chain": crossInfo.TargetChain,
		}).Error("querying CrossItems in DB fail")
		return nil, err
	}

	logrus.Debug("cross item count ", len(items))

	var tokenIds []*big.Int
	var amounts []*big.Int
	var uris []string
	for _, item := range items {
		tokenIds = append(tokenIds, common.HexToHash(item.TokenId).Big())
		amounts = append(amounts, common.HexToHash(item.Amount).Big())
		uris = append(uris, item.Uri)
	}
	logrus.Debug("build token ids ", tokenIds)
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
	return &request, nil
}

func (worker ClaimWorker) Claim(crossInfo models.CrossInfo) (string, *uint64, error) {
	request, err := BuildCrossRequest(crossInfo)
	if err != nil {
		return "", nil, err
	}

	targetChain, err := models.GetChain(crossInfo.TargetChain)
	if err != nil {
		return "", nil, err
	}

	evmHandler := blockchain.GetEvmHandler(crossInfo.TargetChain)
	claimTxHash, nonce, err := blockchain.SendClaimTx(worker.keyPair, worker.address, targetChain,
		big.NewInt(crossInfo.SourceChain), *request, *evmHandler.Client)
	if err != nil {
		return "", nil, err
	}

	return claimTxHash, nonce, nil
}

func (worker ClaimWorker) checkPooledClaim(claim models.ClaimPool) (int, error) {
	if claim.Step == models.ClaimStepSendingTx {
		return worker.sendClaimTx(claim)
	}
	if claim.Step == models.ClaimStepWaitingTx {
		receipt, err := worker.evmHandler.Client.TransactionReceipt(
			context.Background(),
			common.HexToHash(claim.TxnHash),
		)
		if errors.Is(err, ethereum.NotFound) {
			logrus.
				WithFields(logrus.Fields{"chainDbId": worker.Chain.Id, "crossInfoId": claim.CrossInfoId,
					"txnHash": claim.TxnHash}).
				Info("receipt not found")
			return 5, nil
		}
		if err != nil {
			return 0, err
		}
		return worker.checkReceipt(claim, receipt)
	}
	worker.notifyError("unknown claiming step "+claim.Step, "")
	return worker.DelayForError, nil
}

func (worker ClaimWorker) sendClaimTx(claim models.ClaimPool) (int, error) {
	var crossInfo models.CrossInfo
	if err := database.DB.Where("id=?", claim.CrossInfoId).
		Take(&crossInfo).Error; err != nil {
		logrus.WithError(err).
			WithFields(logrus.Fields{"chainDbId": worker.Chain.Id, "crossInfoId": claim.CrossInfoId}).
			Error("querying cross info failed when sending claim tx")
		return 0, err
	}
	claimTxHash, claimNonce, err := worker.Claim(crossInfo)
	if err != nil {
		errorType := blockchain.ParseRpcError(err.Error())
		resend, _ := errorType.CheckTxErrorStatus()
		if resend {
			_ = MoveClaimFromPoolToHistory(claim, uint64(errorType), "sending tx fail")
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
			WithFields(logrus.Fields{"chain.Id": worker.Chain.Id, "crossInfoId": claim.CrossInfoId}).
			Error("update claim info after sending tx fail")
		return 0, err
	}
	return 0, nil
}

func (worker ClaimWorker) checkReceipt(claim models.ClaimPool, receipt *types.Receipt) (int, error) {
	if receipt.Status == 1 {
		logrus.Debug("claiming succeeded ", claim.TxnHash)
		if err := MoveClaimFromPoolToHistory(claim, receipt.Status, "OK"); err != nil {
			return 0, err
		}
		return 0, nil
	}
	worker.notifyError(fmt.Sprintf("tx status fail: [%v]", receipt.Status), claim.TxnHash)
	// reach code here, we have problem
	return worker.DelayForError, nil
}

func (worker ClaimWorker) notifyError(errorInfo string, txnHash string) {
	logrus.WithFields(logrus.Fields{
		"chain.Id": worker.Chain.Id, "tx": txnHash, "errorInfo": errorInfo,
	}).Error("claim transaction fail")
}

func MoveClaimFromPoolToHistory(claim models.ClaimPool, status uint64, comment string) error {
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
