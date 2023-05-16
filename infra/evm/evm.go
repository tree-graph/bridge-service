package evm

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
)

type HandlerEvm struct {
	Client *ethclient.Client
}

var clients map[int64]HandlerEvm

func GetClient(chainId int64) HandlerEvm {
	return clients[chainId]
}
func AddChainClient(chain models.Chain) {
	if clients == nil {
		clients = make(map[int64]HandlerEvm)
	}
	client, err := ethclient.Dial(chain.Rpc)
	helpers.CheckError(
		fmt.Sprintf("evm client fail %v %v %v", chain.Id, chain.Name, chain.Rpc),
		err)
	clients[chain.Id] = HandlerEvm{
		Client: client,
	}
}
func AddRequest(chainId int64, txHash string) (*models.CrossRequest, error) {
	evm := clients[chainId]
	if evm.Client == nil {
		return nil, fmt.Errorf("chain not found: %v", chainId)
	}
	_, _, err := evm.Client.TransactionByHash(context.Background(), common.HexToHash(txHash))
	if err != nil {
		return nil, fmt.Errorf("TransactionByHash error: [%v], [%v]", txHash, err.Error())
	}
	var bean = models.CrossRequest{ChainId: chainId, Hash: txHash}
	err = database.DB.Create(&bean).Error
	if err != nil {
		return nil, err
	}
	logrus.Infof("crossing request added %v %v\n", chainId, txHash)
	return &bean, nil
}
func (evm *HandlerEvm) ParseCrossChainRequest(hash string) ([]*vault.VaultCrossRequest, error) {
	rcpt, err0 := evm.Client.TransactionReceipt(context.Background(), common.HexToHash(hash))
	if err0 != nil {
		return nil, err0
	}
	if rcpt.Status != 1 {
		return nil, nil
	}
	cf, err := vault.NewVaultFilterer(common.HexToAddress(""), evm.Client)
	if err != nil {
		logrus.Println("create contract fail", err.Error())
		return nil, err
	}
	var requests []*vault.VaultCrossRequest
	for _, v := range rcpt.Logs {
		parseLog, err := cf.ParseCrossRequest(*v)
		if err != nil {
			logrus.Println("parse log fail", err.Error())
			continue
		}
		logrus.Infoln("parsed log Asset", parseLog.Asset)
		logrus.Infoln("parsed log From", parseLog.From)
		logrus.Infoln("parsed log TargetChain", parseLog.ToChainId)
		logrus.Infoln("parsed log TokenIds", parseLog.TokenIds)
		logrus.Infoln("parsed log Amounts", parseLog.Amounts)
		logrus.Infoln("parsed log Uris", parseLog.Uris)
		logrus.Infoln("parsed log TargetContract", parseLog.TargetContract)
		logrus.Infoln("parsed log UserNonce", parseLog.UserNonce)
		requests = append(requests, parseLog)
	}
	return requests, nil
}
