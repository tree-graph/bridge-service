/*

 */
package cmd

import (
	"context"
	sdk "github.com/Conflux-Chain/go-conflux-sdk"
	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/Conflux-Chain/go-conflux-sdk/types/cfxaddress"
	"github.com/Conflux-Chain/go-conflux-sdk/types/unit"
	"github.com/ethereum/go-ethereum"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tree-graph/bridge-service/infra/blockchain"
	"github.com/tree-graph/bridge-service/infra/contracts/tokens"
	"github.com/tree-graph/bridge-service/infra/database"
	"github.com/tree-graph/bridge-service/infra/worker"
	"math/big"
	"os"
	"time"
)

// devCmd represents the dev command
var devCmd = &cobra.Command{
	Use: "dev",
	Run: func(cmd *cobra.Command, args []string) {
		//dbTest()
		cfxCore()
		//chainId := 31337
		//h := blockchain.GetEvmHandler(chainId)
		//getLogs(h)
		//var key = "fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19"
		//blockchain.SendClaimTx(key, chainId, )
	},
}

func dbTest() {
	database.Init()
	if err := worker.SetupChains(); err != nil {
		logrus.WithError(err).Error("")
		return
	}
}
func cfxCore() {
	url := "https://test.confluxrpc.com"
	url = "http://47.92.143.10/12550"
	forceNetworkId := 1029
	client, err := sdk.NewClient(url, sdk.ClientOption{
		Logger: os.Stdout,
	})
	/*, sdk.ClientOption{
		KeystorePath: "./keystore/",
	}*/
	if err != nil {
		logrus.Fatal("create cfx core client fail:", err.Error())
		return
	}
	client.SetNetworkId(uint32(forceNetworkId))
	am := sdk.NewAccountManager("./keystore/", uint32(forceNetworkId))
	client.SetAccountManager(am)
	epoch, err := client.GetEpochNumber()
	if err != nil {
		logrus.Fatal("GetEpochNumber fail:", err.Error())
		return
	}
	logrus.Info("epoch is ", epoch.ToInt())

	// tx
	status, err := client.GetStatus()
	if err != nil {
		logrus.WithError(err).Error("get chain id fail")
		return
	}
	networkID := status.NetworkID
	logrus.Info("chain id ", status.ChainID, " network id ", networkID)

	accounts := client.AccountManager.List()
	if len(accounts) == 0 {
		//passphrase := "0x46b9e861b63d3509c88b7817275a30d22d62c8cd8fa6486ddee35ef0d8e0495f"
		passphrase := "9a6d3ba2b0c7514b16a006ee605055d71b9edfad183aeb2d9790e9d4ccced471"
		newAddr, err := client.AccountManager.ImportKey(passphrase, "")
		if err != nil {
			logrus.WithError(err).Error("create account fail")
			return
		}
		logrus.Info("created account is ", newAddr, " hex ", newAddr.GetHexAddress())
	}

	from, err := client.GetAccountManager().GetDefault()
	if err != nil {
		logrus.WithError(err).Error("get account fail")
		return
	}
	logrus.Info("default account is ", *from, " hex ", from.GetHexAddress())

	balance, err := client.GetBalance(*from)
	if err != nil {
		logrus.WithError(err).Error("GetBalance error")
		return
	}
	logrus.Info("balance is ", unit.NewDrip(balance.ToInt()).FormatCFX(), " CFX")

	toAddr := cfxaddress.MustNewFromHex("0x1cad0b19bb29d4674531d6f115237e16afce377d", uint32(networkID))

	//utx := new(types.UnsignedTransaction)
	//utx.From = from
	//utx.To = &toAddr
	//utx.Value = types.NewBigInt(1)
	//utx.Data = nil
	//utx.GasPrice = types.NewBigInt(2)
	//utx.Gas = types.NewBigInt(2100)
	//utx.StorageLimit = types.NewUint64(3)

	utx, err := client.CreateUnsignedTransaction(*from, toAddr, types.NewBigInt(1), nil)
	//err = client.ApplyUnsignedTransactionDefault(utx)
	if err != nil {
		logrus.WithError(err).Error("create tx fail.")
		return
	}
	logrus.WithFields(logrus.Fields{
		"ChainID": utx.ChainID, "Nonce": utx.Nonce,
	}).Info("tx info")

	client.AccountManager.Unlock(*from, "")
	txHash, err := client.SendTransaction(utx)
	if err != nil {
		logrus.WithError(err).Error("send tx fail")
		return
	}
	logrus.Info("tx hash is ", txHash)

	client.AccountManager.Unlock(*from, "")
	//
	//dUtx, hash, inst, err := tokens.DeployPeggedERC721(nil, client)
	dUtx, hash, inst, err := tokens.DeployLock(nil, client)
	if err != nil {
		logrus.WithError(err).Error("deploy fail")
		return
	}
	deployer := dUtx.From.String()
	logrus.WithFields(logrus.Fields{
		"hash": *hash, "deployer": deployer, "instance": inst,
	}).Info("deployed")

	receipt, err := client.WaitForTransationReceipt(*hash, time.Second)
	if err != nil {
		logrus.WithError(err).Error("deploy receipt fail.")
		return
	}
	logrus.WithFields(logrus.Fields{
		"tx":               dUtx,
		"hash":             hash,
		"contract address": receipt.ContractCreated,
	}).Info("deploy token done")
}

func getLogs(h blockchain.EvmHandler) {
	var blockNo *big.Int
	blockNo = big.NewInt(39)
	logs, err := h.Client.FilterLogs(context.Background(), ethereum.FilterQuery{
		FromBlock: blockNo,
		ToBlock:   blockNo,
		//Addresses: nil,
		//Topics:    nil,
	})
	if err != nil {
		logrus.WithError(err).Error("filter logs")
		return
	}
	logrus.Info("count ", len(logs))
	for _, log := range logs {
		logrus.WithFields(logrus.Fields{
			"Address":   log.Address,
			"Topics[0]": log.Topics[0],
		}).Info("log")
	}
}

func init() {
	rootCmd.AddCommand(devCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	devCmd.PersistentFlags().String("foo", "", "A help for foo")
	devCmd.PersistentFlags().String("bar", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	devCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
