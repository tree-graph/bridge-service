package blockchain

import (
	"encoding/json"
	"fmt"
	sdk "github.com/Conflux-Chain/go-conflux-sdk"
	"github.com/Conflux-Chain/go-conflux-sdk/bind"
	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/Conflux-Chain/go-conflux-sdk/types/cfxaddress"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/sirupsen/logrus"
	"github.com/tree-graph/bridge-service/helpers"
	"github.com/tree-graph/bridge-service/infra/contracts/openzeppelin"
	"github.com/tree-graph/bridge-service/infra/contracts/tokens"
	"io/ioutil"
	"math/big"
	"time"
)

type CfxClient sdk.Client

// consortium chain spec
const GAS = uint64(1400_0000)
const InfoFilename = "./cfx-deploy.json"

type deployImpl func(client sdk.Client) (*types.Hash, error)

func DeployReal(client sdk.Client) error {
	impl721, beacon721, err := DeployImplAndBeacon(Deploy721impl, "721", client)
	if err != nil {
		logrus.Error("deploy 721 impl and beacon fail")
		return err
	}
	implTokenFactory, beaconTokenFactory, err := DeployImplAndBeacon(DeployTokenFactoryImpl, "tokenFactory", client)
	if err != nil {
		logrus.Error("deploy token factory impl and beacon fail")
		return err
	}
	tokenFactoryProxy, _ := DeployBeaconProxy(client, beaconTokenFactory.MustGetCommonAddress(), buildTokenFactoryInitData(client, beacon721.MustGetCommonAddress()))
	implTokenVault, beaconTokenVault, err := DeployImplAndBeacon(DeployTokenVaultImpl, "tokenVault", client)
	if err != nil {
		logrus.Error("deploy token vault impl and beacon fail")
		return err
	}
	tokenVaultProxy, _ := DeployBeaconProxy(client, beaconTokenVault.MustGetCommonAddress(), buildTokenVaultInitData(client))

	m := make(map[string]string)
	m["impl721"] = impl721.String()
	m["beacon721"] = beacon721.String()
	m["implTokenFactory"] = implTokenFactory.String()
	m["beaconTokenFactory"] = beaconTokenFactory.String()
	m["proxyTokenFactory"] = tokenFactoryProxy.String()
	m["implTokenVault"] = implTokenVault.String()
	m["beaconTokenVault"] = beaconTokenVault.String()
	m["proxyTokenVault"] = tokenVaultProxy.String()

	logrus.WithFields(logrus.Fields{"InfoFilename": InfoFilename}).Info("all deployed")

	return WriteJson(m, InfoFilename)
}

func WriteJson(m interface{}, fileTo string) error {
	file, _ := json.MarshalIndent(m, "", " ")
	return ioutil.WriteFile(fileTo, file, 0644)
}

func DeployTokenFactoryImpl(client sdk.Client) (*types.Hash, error) {
	makeConfigGas := &bind.TransactOpts{Gas: types.NewBigInt(GAS)}
	_, hash, _, err := tokens.DeployTokenFactory(makeConfigGas, &client)
	return hash, err
}

func DeployTokenVaultImpl(client sdk.Client) (*types.Hash, error) {
	_, hash, _, err := tokens.DeployTokenVault(&bind.TransactOpts{Gas: types.NewBigInt(GAS)}, &client)
	return hash, err
}

func DeployImplAndBeacon(fn deployImpl, name string, client sdk.Client) (*types.Address, *types.Address, error) {
	hash, err := fn(client)
	if err != nil {
		logrus.WithFields(logrus.Fields{"name": name}).Error("deploy impl fail")
		return nil, nil, err
	}
	logrus.WithFields(logrus.Fields{"name": name, "hash": hash.String()}).Info("deployment tx sent : ", name)
	erc721, err := WaitDeployTx(client, err, hash, name)
	if err != nil {
		logrus.WithError(err).Error("wait impl receipt fail")
		return nil, nil, err
	}
	return DeployBeacon(client, hash, err, erc721)
}

func Deploy721impl(client sdk.Client) (*types.Hash, error) {
	_, hash, _, err := tokens.DeployPeggedERC721(&bind.TransactOpts{Gas: types.NewBigInt(GAS)}, &client)
	return hash, err
}

func DeployOZ721impl(client sdk.Client) (*types.Hash, error) {
	logrus.Info("openzeppelin.DeployERC721")
	_, hash, _, err := openzeppelin.DeployERC721(&bind.TransactOpts{Gas: types.NewBigInt(GAS)}, &client, "oz721", "oz721")
	return hash, err
}

func DeployOZ20impl(client sdk.Client) (*types.Hash, error) {
	logrus.Info("openzeppelin.DeployERC20")
	_, hash, _, err := openzeppelin.DeployERC20(&bind.TransactOpts{Gas: types.NewBigInt(GAS)}, &client, "oz20", "oz20")
	return hash, err
}

func buildTokenVaultInitData(client sdk.Client) hexutil.Bytes {
	_, address, _ := createZeroAddr()
	tokenVault, err := tokens.NewTokenVaultBulkTransactor(address, &client)
	helpers.CheckFatalError("NewTokenVaultBulkTransactor", err)
	rawTx := tokenVault.Initialize(&bind.TransactOpts{})
	return rawTx.Data
}

func buildTokenFactoryInitData(client sdk.Client, addr721 common.Address) hexutil.Bytes {
	zeroAddr, address, err := createZeroAddr()

	caller, err := tokens.NewTokenFactoryBulkTransactor(address, &client)
	helpers.CheckFatalError("NewTokenFactoryBulkTransactor", err)

	rawTx := caller.Initialize(&bind.TransactOpts{}, zeroAddr, addr721, zeroAddr)

	return rawTx.Data
}

func createZeroAddr() (common.Address, cfxaddress.Address, error) {
	zeroAddr := common.BigToAddress(big.NewInt(0))
	address, err := cfxaddress.New(zeroAddr.Hex(), 0)
	return zeroAddr, address, err
}

func DeployBeaconProxy(client sdk.Client, beacon common.Address, data []byte) (*types.Address, error) {
	logrus.Info("DeployBeaconProxy beacon ", beacon, " init data ", hexutil.Encode(data))
	_, hash, _, err := openzeppelin.DeployBeaconProxy(&bind.TransactOpts{Gas: types.NewBigInt(GAS)}, &client, beacon, data)
	proxy, err := WaitDeployTx(client, err, hash, "beacon proxy")
	if err != nil {
		logrus.WithError(err).Fatal("DeployBeaconProxy fail.")
	}
	return proxy, err
}

func DeployBeacon(client sdk.Client, hash *types.Hash, err error, impl *types.Address) (*types.Address, *types.Address, error) {
	_, hash, _, err = openzeppelin.DeployUpgradeableBeacon(&bind.TransactOpts{Gas: types.NewBigInt(GAS)}, &client, impl.MustGetCommonAddress())
	if err != nil {
		logrus.WithError(err).Error("deploy beacon fail")
		return nil, nil, err
	}
	logrus.Info("deploy beacon sent")
	upgradeableBeacon, err := WaitDeployTx(client, err, hash, "beacon")
	if err != nil {
		logrus.WithError(err).Error("wait beacon receipt fail")
		return nil, nil, err
	}
	return impl, upgradeableBeacon, nil
}

func WaitDeployTx(client sdk.Client, err error, hash *types.Hash, name string) (*types.Address, error) {
	receipt, err := client.WaitForTransationReceipt(*hash, time.Second)
	if err != nil {
		logrus.WithFields(logrus.Fields{"hash": hash.String()}).Error("receipt error")
		return nil, err
	}
	if receipt.OutcomeStatus != 0 {
		logrus.Error("tx failed, error ", receipt.TxExecErrorMsg)
		return nil, fmt.Errorf("tx failed with error: %v", receipt.TxExecErrorMsg)
	}
	if receipt.ContractCreated == nil {
		// check event log
		_, address, _ := createZeroAddr()
		filter, err := tokens.NewTokenFactoryFilterer(address, nil)
		helpers.CheckFatalError("NewTokenFactoryFilterer", err)
		eventID := tokens.GetEventID("ERC721_CREATED")
		for _, log := range receipt.Logs {
			log721, err := filter.ParseERC721CREATED(log)
			if err == nil && *log721.Raw.Topics[0].ToCommonHash() == eventID {
				contractCreated, err := cfxaddress.New(log721.Artifact.Hex(), 1029)
				helpers.CheckFatalError("ParseERC721CREATED ", err)
				receipt.ContractCreated = &contractCreated
			}
		}
	}
	logrus.WithFields(logrus.Fields{
		"name":            name,
		"hash":            hash,
		"ContractCreated": receipt.ContractCreated,
		"OutcomeStatus":   receipt.OutcomeStatus,
		"TxExecErrorMsg":  receipt.TxExecErrorMsg,
	}).Info("deployment info: ", name)

	if receipt.ContractCreated == nil {
		return nil, fmt.Errorf("contract not created. ")
	}

	return receipt.ContractCreated, nil
}

func CheckTxFail(info string, rcpt *types.TransactionReceipt, err error) {
	if err != nil {
		logrus.WithError(err).Error("transaction receipt error")
		return
	}
	if rcpt.OutcomeStatus != 0 {
		logrus.Error(info, " transaction failed , hash ", rcpt.TransactionHash)
	}
}
