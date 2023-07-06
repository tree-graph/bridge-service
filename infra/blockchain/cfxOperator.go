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

func RegisterRouter(client sdk.Client, tag, local string, remote string, remoteDbId int64, isPegged bool) {
	logrus.WithFields(logrus.Fields{
		"local": local, "remote": remote, "remoteDbId": remoteDbId, "isPegged": isPegged,
	}).Info("RegisterRouter")

	remoteAddr, err := cfxaddress.New(remote)
	helpers.CheckFatalError("invalid remote address "+remote, err)

	localAddr, err := cfxaddress.New(local)
	helpers.CheckFatalError("invalid local address "+local, err)

	remoteId := big.NewInt(remoteDbId)
	// native(not pegged as) default
	arrivalOp := tokens.TRANSFER           // unlock from vault
	arrivalUriMode := tokens.UriModeNotSet // native contract holds its uri
	departureOp := tokens.OpNotSet         // nothing
	departureUriMode := tokens.STORAGE     // track uri when leaving
	if isPegged {                          // register on pegged chain
		arrivalOp = tokens.MINT
		departureOp = tokens.BURN721
		arrivalUriMode = tokens.STORAGE         // set uri when reaching pegged chain
		departureUriMode = tokens.UriModeNotSet // nothing when leaving pegged chain
	}
	logrus.Info("show var ", arrivalOp, arrivalUriMode)
	err = RegisterArrival(client, tag, remoteAddr, remoteId, localAddr, arrivalOp, arrivalUriMode)
	helpers.CheckFatalError("RegisterArrival ", err)

	err = RegisterDeparture(client, tag, localAddr, remoteId, remoteAddr, departureOp, departureUriMode)
	helpers.CheckFatalError("RegisterDeparture ", err)
}

func Test721(client sdk.Client, infoFile string, tokenId int64, deploy bool, claim bool,
	remote string,
	reverse bool, chDbId int64) error {
	// consortium chain may have duplicate chain id. we use a logic db id instead.
	chainDbId := big.NewInt(chDbId)
	//netId, _ := client.GetNetworkID()
	//chainId := big.NewInt(int64(netId))
	account, _ := client.AccountManager.GetDefault()
	if tokenId < 1 {
		tokenId = time.Now().Unix()
	}
	var erc721a *types.Address
	var erc721b *types.Address
	pegInfoFile := "./pegInfo.json"
	var vaultProxy, _ = cfxaddress.New(ReadInfo(infoFile, "proxyTokenVault"))
	DumpTokenVaultInfo(client, vaultProxy, *account, chainDbId)

	if deploy {
		tag := infoFile
		erc721aTmp, err := CreatePegged721(client, tag, "721-0", "p721-0", "")
		erc721a = erc721aTmp
		showOwner(client, erc721a)

		erc721bTmp, err := CreatePegged721(client, tag, "721-01", "p721-01", "")
		erc721b = erc721bTmp
		showOwner(client, erc721b)
		// 0->1
		_ = RegisterArrival(client, tag, *erc721a, chainDbId, *erc721b, tokens.MINT, tokens.STORAGE)
		_ = RegisterDeparture(client, tag, *erc721a, chainDbId, *erc721b, tokens.OpNotSet, tokens.STORAGE)
		// 1->0
		_ = RegisterArrival(client, tag, *erc721b, chainDbId, *erc721a, tokens.TRANSFER, tokens.UriModeNotSet)
		_ = RegisterDeparture(client, tag, *erc721b, chainDbId, *erc721a, tokens.BURN721, tokens.UriModeNotSet)
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
	if remote != "" {
		remoteAddr := cfxaddress.MustNew(remote)
		if reverse {
			// b->a, a is remote
			erc721a = &remoteAddr
		} else {
			// a->b, b is remote
			erc721b = &remoteAddr
		}
	}
	//
	if reverse {
		erc721a, erc721b = erc721b, erc721a
	} else {
		erc721aContract, _ := tokens.NewPeggedERC721Transactor(*erc721a, &client)
		uri := "https://api.nftrainbow.cn/assets/metadata/102/nft/5eaee62bba91528fedb3acbc78e40ddb7fac8bf1639386c48698ceeff2ad792f.json"
		_, mintTxHash, err := erc721aContract.SafeMint(buildGas(), account.MustGetCommonAddress(), big.NewInt(tokenId), uri)
		if err != nil {
			logrus.WithError(err).Error("mint fail")
			return err
		}
		mintRcpt, _ := client.WaitForTransationReceipt(*mintTxHash, time.Second)
		if mintRcpt.OutcomeStatus != 0 {
			logrus.Error("mint fail, token id ", tokenId)
			return nil
		}
	}
	Cross721(client, &vaultProxy, erc721a, chainDbId, erc721b, tokenId, claim)
	if claim {
		Cross721(client, &vaultProxy, erc721b, chainDbId, erc721a, tokenId, claim)
	}
	return nil
}

func Cross721(client sdk.Client,
	vaultProxy, erc721a *types.Address,
	dstChain *big.Int, erc721b *types.Address, tokenId int64, claim bool) {
	account, _ := client.AccountManager.GetDefault()

	remoteHex := erc721b.MustGetCommonAddress()
	data, err := encode(dstChain, remoteHex)
	//logrus.Info("abi encoded data ", hexutil.Encode(data))
	erc721, _ := tokens.NewPeggedERC721Transactor(*erc721a, &client)
	erc721caller, _ := tokens.NewPeggedERC721Caller(*erc721a, &client)

	logrus.Info("token contract ", *erc721a, " hex ", erc721a.MustGetCommonAddress())
	logrus.Info("remote contract ", *erc721b, " hex ", erc721b.MustGetCommonAddress())
	logrus.Info("transfer from ", account.GetHexAddress(), " to ", vaultProxy.GetHexAddress())

	tokenOwner, err := erc721caller.OwnerOf(buildCallOpt(client), big.NewInt(tokenId))
	logrus.Infof("token %v owner is %v \n", tokenId, tokenOwner)
	//showOwner(client, erc721b)
	logrus.Info("token vault : ", vaultProxy, " ", vaultProxy.GetHexAddress())

	vaultReader, err := tokens.NewTokenVaultCaller(*vaultProxy, &client)
	helpers.CheckFatalError("NewTokenVaultCaller", err)

	//check departure
	localHex := erc721a.MustGetCommonAddress()
	info, err := vaultReader.GetDepartureInfo(buildCallOpt(client), localHex, dstChain, remoteHex)
	helpers.CheckFatalError("GetDepartureInfo", err)
	if info.Timestamp.Cmp(big.NewInt(0)) == 0 {
		logrus.Warn("departure info not exists , local ", localHex, " remote ", remoteHex)
		DumpDeparture(client, vaultReader)
	} else {
		logrus.Info("departure exists ")
	}

	_, transferTxHash, err := erc721.SafeTransferFrom0(buildGas(), account.MustGetCommonAddress(), vaultProxy.MustGetCommonAddress(), big.NewInt(tokenId), data)
	helpers.CheckFatalError("erc721 SafeTransferFrom", err)
	rcpt, err := client.WaitForTransationReceipt(*transferTxHash, time.Second)
	helpers.CheckFatalError("WaitForTransationReceipt", err)
	if rcpt.OutcomeStatus != 0 {
		logrus.Error("SafeTransferFrom0 tx failed with error : ", rcpt.TxExecErrorMsg, " hash ", transferTxHash)
		DumpDeparture(client, vaultReader)
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
		chainId, _ := client.GetNetworkID()
		DumpTokenVaultInfo(client, *vaultProxy, *account, big.NewInt(int64(chainId)))
		return
	}
	logrus.Info("claim by tx succeeded, hash ", claimTxHash)
}

func CfxClaim() {

}

func DumpTokenVaultInfo(client sdk.Client, vaultAddr, account types.Address, chainId *big.Int) {
	logrus.Info("DumpTokenVaultInfo,  account ", account)
	vault, _ := tokens.NewTokenVaultCaller(vaultAddr, &client)

	ClaimOnVault, _ := hexutil.Decode("0x41eff8af0daebb3a8ffbb22fbb01e7b90d33fb63b6b9d6e729d25a3cf8c903ea")
	var b32 [32]byte
	copy(b32[:], ClaimOnVault)
	has, err := vault.HasRole(buildCallOpt(client), b32, account.MustGetCommonAddress())
	logrus.Info("has role CLAIM_ON_VAULT ", has, " ", err)

	n, err := vault.GetUserNextClaimNonce(buildCallOpt(client), account.MustGetCommonAddress(), chainId)
	logrus.WithFields(logrus.Fields{"remote chain id": chainId}).Info("GetUserNextClaimNonce ", n, err)
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
	//epoch, err := client.GetEpochNumber()
	//helpers.CheckFatalError("GetEpochNumber", err)
	//logrus.Info("Use epoch ", epoch.ToInt())
	opts := &bind.CallOpts{
		EpochNumber: types.EpochLatestState,
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
	logrus.Info(`registerDeparture to chain `, *dstChain)
	var addr, _ = cfxaddress.New(ReadInfo(infoFile, "proxyTokenVault"))
	vault, err := tokens.NewTokenVaultTransactor(addr, &client)
	helpers.CheckFatalError("NewTokenVaultTransactor", err)
	vaultReader, err := tokens.NewTokenVaultCaller(addr, &client)
	helpers.CheckFatalError("NewTokenVaultCaller", err)

	localHex := localContract.MustGetCommonAddress()
	remoteHex := dstContract.MustGetCommonAddress()

	info, err := vaultReader.GetDepartureInfo(buildCallOpt(client), localHex, dstChain, remoteHex)
	marshal, _ := json.Marshal(info)
	logrus.Info("check peer info exists ", string(marshal), " error ", err)
	if info.Timestamp.Cmp(big.NewInt(0)) > 0 {
		logrus.Info("already registered")
		return nil
	}

	_, hash, _ := vault.RegisterDeparture(buildGas(),
		localHex,
		dstChain, op, uriMode, remoteHex)
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
	logrus.Info(`registerArrival from chain `, *srcChain)
	var addr, _ = cfxaddress.New(ReadInfo(infoFile, "proxyTokenVault"))
	vault, err := tokens.NewTokenVaultTransactor(addr, &client)
	helpers.CheckFatalError("NewTokenVaultTransactor", err)
	vaultReader, err := tokens.NewTokenVaultCaller(addr, &client)
	helpers.CheckFatalError("NewTokenVaultCaller", err)

	eip, err := vaultReader.DetectLocalEIP(buildCallOpt(client), localContract.MustGetCommonAddress())
	logrus.Debug("local eip ", eip, " error ", err)

	srcHex := srcContract.MustGetCommonAddress()
	localHex := localContract.MustGetCommonAddress()

	info, err := vaultReader.GetArrivalInfo(buildCallOpt(client), srcHex, srcChain, localHex)
	marshal, _ := json.Marshal(info)
	logrus.Info("check peer info exists ", string(marshal), " error ", err)
	if info.Timestamp.Cmp(big.NewInt(0)) > 0 {
		logrus.Info("already registered")
		return nil
	}

	logrus.WithFields(logrus.Fields{
		"src": srcHex, "local": localHex,
	}).Info("hex address")

	_, hash, err := vault.RegisterArrival(buildGas(),
		srcHex,
		srcChain, op, uriMode, localHex)
	if err != nil {
		DumpArrival(client, vaultReader)
	}
	helpers.CheckFatalError("RegisterArrival tx.", err)
	receipt, err := client.WaitForTransationReceipt(*hash, time.Second)
	helpers.CheckFatalError("RegisterArrival tx fail.", err)

	if receipt.OutcomeStatus != 0 {
		retE := fmt.Errorf("RegisterArrival tx fail with msg: %v", receipt.TxExecErrorMsg)
		logrus.Error(retE.Error())
		return retE
	}
	return nil
}

func DumpDeparture(client sdk.Client, reader *tokens.TokenVaultCaller) {
	opt := buildCallOpt(client)
	big0 := big.NewInt(0)
	big100 := big.NewInt(100)
	remoteArr, cnt, err := reader.ListDepartureIndex(opt, big0, big100)
	helpers.CheckFatalError("ListDepartureIndex", err)
	logrus.Info("dump Departure")
	for i := int64(0); i < cnt.Int64(); i++ {
		remoteAddr := remoteArr[i]
		logrus.Info("remoteAddr ", remoteAddr)
		chains, _, err := reader.ListDepartureChainIndex(opt, remoteAddr, big0, big100)
		helpers.CheckFatalError("ListDepartureChainIndex", err)
		for _, ch := range chains {
			logrus.Infof("\t chain %v \n", ch.Int64())
			peer, _, err := reader.ListDeparturePeerIndex(opt, remoteAddr, ch, big0, big100)
			helpers.CheckFatalError("ListDeparturePeerIndex", err)
			for _, p := range peer {
				logrus.Infof("\t\t peer %v \n", p)
				info, err := reader.GetDepartureInfo(opt, remoteAddr, ch, p)
				helpers.CheckFatalError("GetDepartureInfo", err)
				marshal, _ := json.Marshal(info)
				logrus.Info("\t\t\t peer ", string(marshal))
			}
		}
	}
}

func DumpArrival(client sdk.Client, reader *tokens.TokenVaultCaller) {
	opt := buildCallOpt(client)
	big0 := big.NewInt(0)
	big100 := big.NewInt(100)
	remoteArr, cnt, err := reader.ListArrivalIndex(opt, big0, big100)
	helpers.CheckFatalError("ListArrivalIndex", err)
	logrus.Info("dump arrival")
	for i := int64(0); i < cnt.Int64(); i++ {
		remoteAddr := remoteArr[i]
		logrus.Info("remoteAddr ", remoteAddr)
		chains, _, err := reader.ListArrivalChainIndex(opt, remoteAddr, big0, big100)
		helpers.CheckFatalError("ListArrivalChainIndex", err)
		for _, ch := range chains {
			logrus.Infof("\t chain %v \n", ch.Int64())
			peer, _, err := reader.ListArrivalPeerIndex(opt, remoteAddr, ch, big0, big100)
			helpers.CheckFatalError("ListArrivalPeerIndex", err)
			for _, p := range peer {
				logrus.Infof("\t\t peer %v \n", p)
				info, err := reader.GetArrivalInfo(opt, remoteAddr, ch, p)
				helpers.CheckFatalError("GetArrivalInfo", err)
				marshal, _ := json.Marshal(info)
				logrus.Info("\t\t\t peer ", string(marshal))
			}
		}
	}
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
	return &bind.TransactOpts{Gas: types.NewBigInt(GAS),
		GasPrice: types.NewBigInt(90_000_000_000),
		//Nonce: types.NewBigInt(66),
	}
}
