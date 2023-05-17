package blockchain

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
	"github.com/tree-graph/bridge-service/helpers"
	"github.com/tree-graph/bridge-service/infra/contracts/vault"
	"github.com/tree-graph/bridge-service/infra/database"
	"github.com/tree-graph/bridge-service/models"
	"sync"
)

type EvmHandler struct {
	Client  *ethclient.Client
	ChainId int64
}

var clients map[int64]EvmHandler
var clientsMapLock sync.Mutex

// Call it at the application entry point
func SetupEvmEnv() {
	clients = make(map[int64]EvmHandler)
}
func GetEvmHandler(chainId int64) EvmHandler {
	return clients[chainId]
}
func AddChainClient(chain models.Chain) {
	clientsMapLock.Lock()
	defer clientsMapLock.Unlock()

	client, err := ethclient.Dial(chain.Rpc)
	helpers.CheckFatalError(
		fmt.Sprintf("evm client fail %v %v %v", chain.Id, chain.Name, chain.Rpc),
		err)
	clients[chain.Id] = EvmHandler{
		Client:  client,
		ChainId: chain.Id,
	}
}
func AddCrossRequest(chainId int64, txHash string) (*models.CrossRequest, error) {
	evm, ok := clients[chainId]
	if !ok {
		return nil, fmt.Errorf("chain not found: %v", chainId)
	}

	if _, _, err := evm.Client.TransactionByHash(context.Background(), common.HexToHash(txHash)); err != nil {
		return nil, fmt.Errorf("TransactionByHash error: [%v], [%v]", txHash, err.Error())
	}

	var bean = models.CrossRequest{ChainId: chainId, Hash: txHash}
	if err := database.DB.Create(&bean).Error; err != nil {
		return nil, err
	}
	logrus.WithFields(logrus.Fields{
		"chainId": chainId, "txHash": txHash,
	}).Infoln("crossing request added")
	return &bean, nil
}
func (evm *EvmHandler) GetCrossChainRequest(hash string) ([]*vault.VaultCrossRequest, error) {
	rcpt, err0 := evm.Client.TransactionReceipt(context.Background(), common.HexToHash(hash))
	if err0 != nil {
		return nil, err0
	}
	if rcpt.Status != 1 {
		return nil, nil
	}
	contract, err := vault.NewVaultFilterer(common.HexToAddress(""), evm.Client)
	if err != nil {
		logrus.Println("create contract fail", err.Error())
		return nil, err
	}
	var requests []*vault.VaultCrossRequest
	for _, eachLog := range rcpt.Logs {
		parsedLog, err := contract.ParseCrossRequest(*eachLog)
		if err != nil {
			logrus.WithError(err).
				WithField("txHash", hash).
				WithField("chainId", evm.ChainId).
				Error("parse log fail")
			continue
		}
		logrus.WithFields(logrus.Fields{
			"Asset":          parsedLog.Asset,
			"From":           parsedLog.From,
			"TargetChain":    parsedLog.ToChainId,
			"TokenIds":       parsedLog.TokenIds,
			"Amounts":        parsedLog.Amounts,
			"Uris":           parsedLog.Uris,
			"TargetContract": parsedLog.TargetContract,
			"UserNonce":      parsedLog.UserNonce,
		}).Debug("parsed log")
		requests = append(requests, parsedLog)
	}
	return requests, nil
}
