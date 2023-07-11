///*
//
// */
package cmd

import (
	"context"
	sdk "github.com/Conflux-Chain/go-conflux-sdk"
	"github.com/Conflux-Chain/go-conflux-sdk/bind"
	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/Conflux-Chain/go-conflux-sdk/types/cfxaddress"
	"github.com/Conflux-Chain/go-conflux-sdk/types/unit"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tree-graph/bridge-service/helpers"
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
		cfxCore(cmd)
	},
}

func dbTest() {
	database.Init()
	if err := worker.SetupChains(); err != nil {
		logrus.WithError(err).Error("")
		return
	}
}
func cfxCore(cmd *cobra.Command) {
	url := viper.GetString("RPC")
	logrus.Info("use rpc ", url)

	cfxOption := sdk.ClientOption{}
	if viper.GetBool("CFX_LOG") {
		cfxOption.Logger = os.Stdout
	}

	client, err := sdk.NewClient(url, cfxOption)
	if err != nil {
		logrus.Fatal("create cfx core client fail:", err.Error())
		return
	}
	forceNetworkId := blockchain.PatchChain(client)
	am := sdk.NewAccountManager("./keystore/", forceNetworkId)
	client.SetAccountManager(am)
	if epochNumber(err, client) {
		return
	}

	// tx
	status, err := client.GetStatus()
	if err != nil {
		logrus.WithError(err).Error("get chain id fail")
		return
	}
	networkID := status.NetworkID
	logrus.Info("chain id ", status.ChainID.String(), " network id ", networkID.String())

	if checkAccounts(client) {
		return
	}

	from, err := client.GetAccountManager().GetDefault()
	if err != nil {
		logrus.WithError(err).Error("get account fail")
		return
	}
	logrus.Info("default account is ", *from, " hex ", from.GetHexAddress())

	if showBalance(err, client, from) {
		return
	}

	client.AccountManager.Unlock(*from, "")

	if dev, _ := cmd.Flags().GetBool("dev"); dev {
		logrus.Info("dev sdk test")
		testCfxSdk(networkID, client, from)
	} else if test, err := cmd.Flags().GetBool("test"); test {
		deploy, _ := cmd.Flags().GetBool("deploy")
		claim, _ := cmd.Flags().GetBool("claim")
		reverse, _ := cmd.Flags().GetBool("reverse")
		tokenId, _ := cmd.Flags().GetInt64("tid")
		remote, _ := cmd.Flags().GetString("remote")
		chainDbId, _ := cmd.Flags().GetInt64("cid")
		err = blockchain.Test721(*client, blockchain.InfoFilename, tokenId, deploy, claim, remote, reverse, chainDbId)
		helpers.CheckFatalError("test fail", err)
		return
	} else if register, _ := cmd.Flags().GetBool("register"); register {
		isPegged, _ := cmd.Flags().GetBool("pegged")
		chainDbId, _ := cmd.Flags().GetInt64("cid")
		local, _ := cmd.Flags().GetString("local")
		remote, _ := cmd.Flags().GetString("remote")
		blockchain.RegisterRouter(*client, blockchain.InfoFilename, local, remote, chainDbId, isPegged)
	} else if deploy, err := cmd.Flags().GetBool("deploy"); deploy {
		logrus.Info("test deployment")
		err = blockchain.DeployReal(*client)
		if err != nil {
			logrus.WithError(err).Error("deploy cfx fail")
		}
		return
	} else {
		logrus.Warn("use -dev -test or -deploy to run")
	}
}

func testCfxSdk(networkID hexutil.Uint64, client *sdk.Client, from *types.Address) {
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
		logrus.WithError(err).Error("create tx fail")
		return
	}
	logrus.WithFields(logrus.Fields{
		"ChainID": utx.ChainID, "Nonce": utx.Nonce, "epoch": utx.EpochHeight,
	}).Info("tx info")

	txHash, err := client.SendTransaction(utx)
	if err != nil {
		logrus.WithError(err).Error("send tx fail")
		return
	}
	logrus.Info("tx hash is ", txHash)
	client.WaitForTransationReceipt(txHash, time.Second)

	//
	//dUtx, hash, inst, err := tokens.DeployPeggedERC721(nil, client)
	dUtx, hash, _, err := tokens.DeployLock(&bind.TransactOpts{
		Gas: types.NewBigInt(0x19197),
	}, client)
	if err != nil {
		logrus.WithError(err).Error("deploy fail")
		return
	}
	deployer := dUtx.From.String()
	logrus.WithFields(logrus.Fields{
		"hash": *hash, "deployer": deployer,
	}).Info("deployment tx sent Lock")

	receipt, err := client.WaitForTransationReceipt(*hash, time.Second)
	if err != nil {
		logrus.WithError(err).Error("deploy receipt fail")
		return
	}
	logrus.WithFields(logrus.Fields{
		"hash":             hash,
		"contract address": receipt.ContractCreated, "outcomeStatus": receipt.OutcomeStatus,
	}).Info("tokens.DeployLock done.")
}

func epochNumber(err error, client *sdk.Client) bool {
	epoch, err := client.GetEpochNumber()
	if err != nil {
		logrus.Fatal("GetEpochNumber fail:", err.Error())
		return true
	}
	logrus.Info("epoch is ", epoch.ToInt())
	return false
}

func checkAccounts(client *sdk.Client) bool {
	accounts := client.AccountManager.List()
	if len(accounts) == 0 {
		passphrase := viper.GetString("PK")
		newAddr, err := client.AccountManager.ImportKey(passphrase, "")
		if err != nil {
			logrus.WithError(err).Error("create account fail")
			return true
		}
		logrus.Info("created account is ", newAddr, " hex ", newAddr.GetHexAddress())
	}
	return false
}

func showBalance(err error, client *sdk.Client, from *types.Address) bool {
	balance, err := client.GetBalance(*from)
	if err != nil {
		logrus.WithError(err).Error("GetBalance error")
		return true
	}
	logrus.Info("balance is ", unit.NewDrip(balance.ToInt()).FormatCFX(), " CFX")
	return false
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
	devCmd.PersistentFlags().String("bar", "", "A help for bar")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	//devCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	devCmd.Flags().BoolP("dev", "d", false, "dev code test")
	devCmd.Flags().BoolP("test", "t", false, "test")
	devCmd.Flags().BoolP("deploy", "D", false, "deploy")
	devCmd.Flags().BoolP("claim", "c", false, "claim")
	devCmd.Flags().BoolP("reverse", "r", false, "reverse, cross back")
	devCmd.Flags().Int64("tid", 0, "token id")
	devCmd.Flags().Int64("cid", 0, "logic chain id")

	devCmd.Flags().Bool("register", false, "register router")
	devCmd.Flags().Bool("pegged", false, "register on pegged")
	devCmd.Flags().String("local", "", "local contract")
	devCmd.Flags().String("remote", "", "remote contract")
}
