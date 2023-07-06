package worker

import (
	"fmt"
	sdk "github.com/Conflux-Chain/go-conflux-sdk"
	"github.com/Conflux-Chain/go-conflux-sdk/bind"
	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/Conflux-Chain/go-conflux-sdk/types/cfxaddress"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/tree-graph/bridge-service/infra/blockchain"
	"github.com/tree-graph/bridge-service/infra/contracts/tokens"
	"github.com/tree-graph/bridge-service/infra/database"
	"github.com/tree-graph/bridge-service/models"
	"math/big"
	"strings"
	"time"
)

type CfxClaimWorker struct {
	Chain         models.Chain
	ChainId       int64
	CfxClient     *sdk.Client
	address       types.Address
	DelayForError int
}

func (worker CfxClaimWorker) GetChainId() int64 {
	return worker.ChainId
}

func (worker CfxClaimWorker) Init() {
	_ = worker.CfxClient.AccountManager.UnlockDefault("")
}

func (worker CfxClaimWorker) DoWork() (int, error) {
	hasPooledClaim, pooledClaim, err := HasPooledClaim(worker.Chain.Id, worker.Chain.Name)
	if err != nil {
		return 0, err
	}
	if hasPooledClaim {
		return worker.checkPooledClaim(pooledClaim)
	}
	return CheckClaimTask(worker.Chain.Id, worker.Chain.Name)
}

func (worker CfxClaimWorker) checkPooledClaim(claim models.ClaimPool) (int, error) {
	if claim.Step == models.ClaimStepSendingTx {
		return worker.sendClaimTx(claim)
	}
	if claim.Step == models.ClaimStepWaitingTx {
		receipt, err := worker.CfxClient.WaitForTransationReceipt(types.Hash(claim.TxnHash), time.Second)
		if err != nil {
			return 0, err
		}
		return worker.checkReceipt(claim, receipt)
	}
	worker.notifyError(worker.Chain.Name+" : unknown claiming step "+claim.Step, "")
	return worker.DelayForError, nil
}

func (worker CfxClaimWorker) Claim(crossInfo models.CrossInfo) (string, *uint64, error) {
	var items []models.CrossItem
	if err := database.DB.Where("cross_info_id=?", crossInfo.Id).Find(&items).Error; err != nil {
		logrus.Debug("cfx claim query cross info fail")
		return "", nil, err
	}
	vaultAddr, err := cfxaddress.New(worker.Chain.VaultAddr)
	if err != nil {
		return "", nil, errors.WithMessage(err, "parse vault addr fail")
	}
	vault, err := tokens.NewTokenVaultTransactor(vaultAddr, worker.CfxClient)
	if err != nil {
		return "", nil, errors.WithMessage(err, "tokens.NewTokenVaultTransactor fail")
	}

	var tokenIds []*big.Int
	var amounts []*big.Int
	var uris []string
	for _, item := range items {
		tokenIds = append(tokenIds, common.HexToHash(item.TokenId).Big())
		amounts = append(amounts, common.HexToHash(item.Amount).Big())
		uris = append(uris, item.Uri)
	}
	req := crossInfo
	hexIssuer := common.HexToAddress(req.From)
	srcChainId := big.NewInt(req.SourceChain)
	utx, claimTxHash, err := vault.ClaimByAdmin(buildGas(),
		srcChainId, common.HexToAddress(req.Asset), common.HexToAddress(req.TargetContract),
		tokenIds, amounts, uris,
		hexIssuer, big.NewInt(req.UserNonce))
	if err != nil {
		if strings.Contains(err.Error(), "bad user claim nonce") {
			vaultCaller, _ := tokens.NewTokenVaultCaller(vaultAddr, worker.CfxClient)
			n, err := vaultCaller.GetUserNextClaimNonce(&bind.CallOpts{
				EpochNumber: types.EpochLatestState,
			}, hexIssuer, srcChainId)
			logrus.Info("GetUserNextClaimNonce ", n, " error ", err)
			logrus.Info("off-chain claim nonce ", req.UserNonce, " src chain ", req.SourceChain)
		}
		return "", nil, errors.WithMessage(err, "cfx claim by admin fail")
	}

	nonce := utx.Nonce.ToInt().Uint64()
	logrus.WithFields(logrus.Fields{
		"nonce": nonce, "chain": worker.Chain.Name, "tokenIds": tokenIds,
	}).Info("cfx claim sent")
	return claimTxHash.String(), &nonce, nil
}

func buildGas() *bind.TransactOpts {
	return &bind.TransactOpts{Gas: types.NewBigInt(blockchain.GAS)}
}

func (worker CfxClaimWorker) sendClaimTx(claim models.ClaimPool) (int, error) {
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
	txFrom := worker.address.String()
	if err := UpdateClaimPool(claim, claimTxHash, txFrom, *claimNonce); err != nil {
		logrus.WithError(err).
			WithFields(logrus.Fields{"chain.Id": worker.Chain.Id, "crossInfoId": claim.CrossInfoId}).
			Error("update claim info after sending tx fail")
		return 0, err
	}
	return 0, nil
}

func UpdateClaimPool(claim models.ClaimPool, claimTxHash string, txFrom string, claimNonce uint64) error {
	return database.DB.Model(claim).Where("id=?", claim.Id).
		Updates(&models.ClaimPool{
			TxnHash: claimTxHash,
			From:    txFrom,
			Nonce:   claimNonce,
			Step:    models.ClaimStepWaitingTx,
			Status:  nil,
		}).Error
}

func (worker CfxClaimWorker) checkReceipt(claim models.ClaimPool, receipt *types.TransactionReceipt) (int, error) {
	if receipt.OutcomeStatus == 0 {
		logrus.WithFields(logrus.Fields{
			"chain": worker.Chain.Name, "CrossInfoId": claim.CrossInfoId,
		}).Debug("claiming succeeded ", claim.TxnHash)
		if err := MoveClaimFromPoolToHistory(claim, uint64(receipt.OutcomeStatus), "OK"); err != nil {
			return 0, err
		}
		return 0, nil
	}
	worker.debugFailedTx(claim, receipt)
	worker.notifyError(fmt.Sprintf("tx status fail: [%v]", receipt.OutcomeStatus), claim.TxnHash)
	// reach code here, we have problem
	return worker.DelayForError, nil
}

func (worker CfxClaimWorker) debugFailedTx(claim models.ClaimPool, receipt *types.TransactionReceipt) {
	logrus.WithFields(logrus.Fields{
		"TxExecErrorMsg": receipt.TxExecErrorMsg,
	}).Info("debug failed cfx tx")
	var crossInfo models.CrossInfo
	if err := database.DB.Where("id = ?", claim.CrossInfoId).Take(&crossInfo).Error; err != nil {
		logrus.WithError(err).Error("query cross info fail")
		return
	}
	vaultAddr, _ := cfxaddress.New(worker.Chain.VaultAddr)
	vault, _ := tokens.NewTokenVaultCaller(vaultAddr, worker.CfxClient)

	epoch, _ := worker.CfxClient.GetEpochNumber()
	n, err := vault.GetUserNextClaimNonce(&bind.CallOpts{
		EpochNumber: types.NewEpochNumber(epoch),
	}, common.HexToAddress(claim.From), big.NewInt(crossInfo.SourceChain))
	logrus.WithFields(logrus.Fields{
		"useUserNonce": crossInfo.UserNonce, "onChain.nonce": n, "call.error": err,
	}).Info("debug failed claim tx on cfx")
}

func (worker CfxClaimWorker) notifyError(errorInfo string, txnHash string) {
	logrus.WithFields(logrus.Fields{
		"chain.Id": worker.Chain.Id, "tx": txnHash, "errorInfo": errorInfo,
	}).Error("claim transaction fail")
}
