package blockchain

import (
	"encoding/json"
	"fmt"
	sdk "github.com/Conflux-Chain/go-conflux-sdk"
	"github.com/Conflux-Chain/go-conflux-sdk/bind"
	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/Conflux-Chain/go-conflux-sdk/types/cfxaddress"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/sirupsen/logrus"
	"github.com/tree-graph/bridge-service/helpers"
	"github.com/tree-graph/bridge-service/infra/contracts/tokens"
	"io/ioutil"
	"math/big"
	"time"
)

func Test721(client sdk.Client, infoFile string, tokenId int64, deploy bool, claim bool, reverse bool) error {
	chainId := big.NewInt(1029)
	account, _ := client.AccountManager.GetDefault()
	if tokenId < 1 {
		tokenId = time.Now().Unix()
	}
	var erc721a *types.Address
	var erc721b *types.Address
	pegInfoFile := "./pegInfo.json"
	var vaultProxy, _ = cfxaddress.New(ReadInfo(infoFile, "proxyTokenVault"))
	DumpTokenVaultInfo(client, vaultProxy, *account)

	if deploy {
		tag := infoFile
		erc721aTmp, err := CreatePegged721(client, tag, "721-0", "p721-0", "https://baidu.com/")
		erc721a = erc721aTmp
		showOwner(client, erc721a)

		erc721bTmp, err := CreatePegged721(client, tag, "721-01", "p721-01", "")
		erc721b = erc721bTmp
		showOwner(client, erc721b)
		// 0->1
		_ = RegisterArrival(client, tag, *erc721a, chainId, *erc721b, tokens.MINT, tokens.STORAGE)
		_ = RegisterDeparture(client, tag, *erc721a, chainId, *erc721b, tokens.OpNotSet, tokens.STORAGE)
		// 1->0
		_ = RegisterArrival(client, tag, *erc721b, chainId, *erc721a, tokens.TRANSFER, tokens.UriModeNotSet)
		_ = RegisterDeparture(client, tag, *erc721b, chainId, *erc721a, tokens.BURN721, tokens.UriModeNotSet)
		// transfer ownership to token vault
		logrus.Info("transfer ownership of ", erc721b, " to ", vaultProxy)
		erc721aContract, _ := tokens.NewPeggedERC721Transactor(*erc721b, &client)
		_, ownerTxHash, err := erc721aContract.TransferOwnership(buildGas(), vaultProxy.MustGetCommonAddress())
		rcpt, err := client.WaitForTransationReceipt(*ownerTxHash, time.Second)
		CheckTxFail("TransferOwnership ", rcpt, err)
		//
		err = WriteJson(map[string]string{
			"erc721a": erc721a.String(),
			"erc721b": erc721b.String(),
		}, pegInfoFile)
		helpers.CheckFatalError("WriteJson", err)
	} else {
		erc721aRef := cfxaddress.MustNew(ReadInfo(pegInfoFile, "erc721a"))
		erc721a = &erc721aRef
		erc721bRef := cfxaddress.MustNew(ReadInfo(pegInfoFile, "erc721b"))
		erc721b = &erc721bRef
	}
	//
	erc721aContract, _ := tokens.NewPeggedERC721Transactor(*erc721a, &client)
	if reverse {
		erc721a, erc721b = erc721b, erc721a
	} else {
		_, mintTxHash, _ := erc721aContract.SafeMint(buildGas(), account.MustGetCommonAddress(), big.NewInt(tokenId), "storage-uri-test721.json")
		mintRcpt, _ := client.WaitForTransationReceipt(*mintTxHash, time.Second)
		if mintRcpt.OutcomeStatus != 0 {
			logrus.Error("mint fail, token id ", tokenId)
			return nil
		}
	}
	Cross721(client, &vaultProxy, erc721a, chainId, erc721b, tokenId, claim)
	if claim {
		Cross721(client, &vaultProxy, erc721b, chainId, erc721a, tokenId, claim)
	}
	return nil
}

func Cross721(client sdk.Client,
	vaultProxy, erc721a *types.Address,
	dstChain *big.Int, erc721b *types.Address, tokenId int64, claim bool) {
	account, _ := client.AccountManager.GetDefault()

	data, err := encode(dstChain, erc721b.MustGetCommonAddress())
	logrus.Info("abi encoded data ", hexutil.Encode(data))
	erc721, err := tokens.NewPeggedERC721Transactor(*erc721a, &client)
	logrus.Info("transfer from ", account.GetHexAddress(), " to ", vaultProxy.GetHexAddress())

	showOwner(client, erc721b)
	logrus.Info("token vault : ", vaultProxy, " ", vaultProxy.GetHexAddress())

	_, transferTxHash, err := erc721.SafeTransferFrom0(buildGas(), account.MustGetCommonAddress(), vaultProxy.MustGetCommonAddress(), big.NewInt(tokenId), data)
	helpers.CheckFatalError("erc721 SafeTransferFrom", err)
	rcpt, err := client.WaitForTransationReceipt(*transferTxHash, time.Second)
	helpers.CheckFatalError("WaitForTransationReceipt", err)
	if rcpt.OutcomeStatus != 0 {
		logrus.Error("SafeTransferFrom0 tx failed with error : ", rcpt.TxExecErrorMsg, " hash ", transferTxHash)
		return
	}

	req := GetCrossRequest(client, vaultProxy, rcpt, transferTxHash)
	if req == nil {
		return
	}

	if !claim {
		logrus.Info("do not claim, skip.")
		return
	}
	// claim
	vault, err := tokens.NewTokenVaultTransactor(*vaultProxy, &client)
	helpers.CheckFatalError("NewTokenVaultTransactor", err)

	_, claimTxHash, err := vault.ClaimByAdmin(buildGas(),
		req.ToChainId, req.Asset, req.TargetContract, req.TokenIds,
		req.Amounts, req.Uris, req.From, req.UserNonce)
	helpers.CheckFatalError("ClaimByAdmin", err)
	claimRcpt, err := client.WaitForTransationReceipt(*claimTxHash, time.Second)
	helpers.CheckFatalError("wait receipt for ClaimByAdmin", err)
	if claimRcpt.OutcomeStatus != 0 {
		logrus.WithFields(logrus.Fields{
			"message": claimRcpt.TxExecErrorMsg, "hash": claimTxHash,
		}).Error("claim by admin tx failed")
		DumpTokenVaultInfo(client, *vaultProxy, *account)
		return
	}
	logrus.Info("claim by tx succeeded, hash ", claimTxHash)
}

func CfxClaim() {

}

func DumpTokenVaultInfo(client sdk.Client, vaultAddr, account types.Address) {
	logrus.Info("DumpTokenVaultInfo,  account ", account)
	vault, _ := tokens.NewTokenVaultCaller(vaultAddr, &client)

	ClaimOnVault, _ := hexutil.Decode("0x41eff8af0daebb3a8ffbb22fbb01e7b90d33fb63b6b9d6e729d25a3cf8c903ea")
	var b32 [32]byte
	copy(b32[:], ClaimOnVault)
	has, err := vault.HasRole(buildCallOpt(client), b32, account.MustGetCommonAddress())
	logrus.Info("has role CLAIM_ON_VAULT ", has, " ", err)

	n, err := vault.GetUserNextClaimNonce(buildCallOpt(client), account.MustGetCommonAddress(), big.NewInt(1029))
	logrus.Info("GetUserNextClaimNonce ", n, err)
}

func GetCrossRequest(client sdk.Client, vaultProxy *types.Address, rcpt *types.TransactionReceipt, transferTxHash *types.Hash) *tokens.TokenVaultCrossRequest {
	var req *tokens.TokenVaultCrossRequest
	eventFilter, _ := tokens.NewTokenVaultFilterer(*vaultProxy, &client)
	wantEventId := tokens.TokenVaultInterface.Events["CrossRequest"].ID
	logrus.Info("want event id ", wantEventId)
	for _, log := range rcpt.Logs {
		if *log.Topics[0].ToCommonHash() == wantEventId {
			err := error(nil)
			req, err = eventFilter.ParseCrossRequest(log)
			if err != nil {
				continue
			}
			break
		} else {
			logrus.Debug("mismatch event ", log.Topics[0].String())
		}
	}
	if req == nil {
		logrus.Error("crossRequest event not found in tx ", transferTxHash)
		return nil
	}
	logrus.WithFields(logrus.Fields{"epoch": rcpt.EpochNumber}).Info("found crossRequest event , tx ", transferTxHash)
	return req
}

func showOwner(client sdk.Client, erc721 *types.Address) {
	erc721c1, err := tokens.NewPeggedERC721Caller(*erc721, &client)
	helpers.CheckFatalError("NewPeggedERC721Caller", err)
	opts := buildCallOpt(client)
	targetOwner, err := erc721c1.Owner(opts)
	helpers.CheckFatalError("call owner fail ", err)
	logrus.Info("target contract owner: ", targetOwner)
}

func buildCallOpt(client sdk.Client) *bind.CallOpts {
	epoch, err := client.GetEpochNumber()
	helpers.CheckFatalError("GetEpochNumber", err)
	logrus.Info("Use epoch ", epoch)
	opts := &bind.CallOpts{
		EpochNumber: types.NewEpochNumber(epoch),
	}
	return opts
}

func encode(bi *big.Int, addr common.Address) ([]byte, error) {
	var (
		structThing, _ = abi.NewType("tuple", "struct thing", []abi.ArgumentMarshaling{
			{Name: "field_one", Type: "uint256"},
			{Name: "field_two", Type: "address"},
		})

		args = abi.Arguments{
			{Type: structThing, Name: "param_one"},
		}
	)

	record := struct {
		FieldOne *big.Int
		FieldTwo common.Address
	}{
		bi, addr,
	}

	return args.Pack(&record)
}

func RegisterDeparture(client sdk.Client, infoFile string, localContract types.Address,
	dstChain *big.Int, dstContract types.Address, op, uriMode uint8) error {
	logrus.Info(`registerDeparture`)
	var addr, _ = cfxaddress.New(ReadInfo(infoFile, "proxyTokenVault"))
	vault, err := tokens.NewTokenVaultTransactor(addr, &client)
	helpers.CheckFatalError("NewTokenVaultTransactor", err)
	//await vault.registerDeparture(localContract, dstChain, op, uriMode, dstContract).then(waitTx)

	_, hash, _ := vault.RegisterDeparture(buildGas(),
		localContract.MustGetCommonAddress(),
		dstChain, op, uriMode, dstContract.MustGetCommonAddress())
	receipt, err := client.WaitForTransationReceipt(*hash, time.Second)
	helpers.CheckFatalError("registerDeparture tx fail.", err)

	if receipt.OutcomeStatus != 0 {
		retE := fmt.Errorf("RegisterDeparture tx fail with msg: %v", receipt.TxExecErrorMsg)
		logrus.Error(retE.Error())
		return retE
	}
	return nil
}

func RegisterArrival(client sdk.Client, infoFile string,
	srcContract types.Address, srcChain *big.Int, localContract types.Address, op uint8, uriMode uint8) error {
	logrus.Info(`registerArrival`)
	var addr, _ = cfxaddress.New(ReadInfo(infoFile, "proxyTokenVault"))
	vault, err := tokens.NewTokenVaultTransactor(addr, &client)
	helpers.CheckFatalError("NewTokenVaultTransactor", err)

	_, hash, _ := vault.RegisterArrival(buildGas(),
		srcContract.MustGetCommonAddress(),
		srcChain, op, uriMode, localContract.MustGetCommonAddress())
	receipt, err := client.WaitForTransationReceipt(*hash, time.Second)
	helpers.CheckFatalError("RegisterArrival tx fail.", err)

	if receipt.OutcomeStatus != 0 {
		retE := fmt.Errorf("RegisterArrival tx fail with msg: %v", receipt.TxExecErrorMsg)
		logrus.Error(retE.Error())
		return retE
	}
	return nil
}

func ReadInfo(infoFile, key string) string {
	bytes, err := ioutil.ReadFile(infoFile)
	helpers.CheckFatalError("read file", err)

	m := make(map[string]string)

	err = json.Unmarshal(bytes, &m)
	helpers.CheckFatalError("unmarshal info file", err)
	return m[key]
}

func CreatePegged721(client sdk.Client, infoFile string, name string, symbol string, baseUri string) (*types.Address, error) {
	var addr, _ = cfxaddress.New(ReadInfo(infoFile, "proxyTokenFactory"))
	logrus.Info("use proxyTokenFactory ", addr)
	tokenFactory, err := tokens.NewTokenFactoryTransactor(addr, &client)
	helpers.CheckFatalError("new token factory transactor", err)

	_, hash, err := tokenFactory.DeployERC721(buildGas(), name, symbol, baseUri)
	helpers.CheckFatalError("build raw tx", err)

	logrus.WithFields(logrus.Fields{
		"hash": hash, "name": name,
	}).Info("send tx to create pegged 721")

	return WaitDeployTx(client, err, hash, "create pegged 721: "+name)
}

func buildGas() *bind.TransactOpts {
	return &bind.TransactOpts{Gas: types.NewBigInt(GAS)}
}
