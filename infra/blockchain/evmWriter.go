package blockchain

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
	"github.com/tree-graph/bridge-service/infra/contracts/vault"
	"github.com/tree-graph/bridge-service/models"
	"math/big"
	"os"
	"strings"
)

func SendClaimTx(privateKey *ecdsa.PrivateKey, address *common.Address, targetChain models.Chain,
	srcChainId *big.Int, request vault.VaultCrossRequest, client ethclient.Client) (string, *uint64, error) {
	vaultAddr := targetChain.VaultAddr
	nonce, gasPrice, err := GetNonceAndGas(client, address)
	if err != nil {
		return "", nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(targetChain.ChainId))
	if err != nil {
		return "", nil, err
	}
	auth.Nonce = big.NewInt(int64(*nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	valutContract, err := vault.NewVaultTransactor(common.HexToAddress(vaultAddr), &client)
	if err != nil {
		return "", nil, err
	}
	srcContract := request.Asset
	localContract := request.TargetContract
	tokenIds := request.TokenIds
	amounts := request.Amounts
	issuer := request.From
	userNonce := request.UserNonce
	uris := request.Uris

	logrus.WithFields(logrus.Fields{
		"userNonce": userNonce,
		"tokenIds":  tokenIds,
		"amounts":   amounts,
		"uris":      uris,
	}).Debug("sending claiming tx")
	transaction, err := valutContract.ClaimByAdmin(auth, srcChainId, srcContract, localContract, tokenIds, amounts, uris, issuer, userNonce)
	if err != nil {
		if strings.Index(err.Error(), "'bad user claim nonce") >= 0 {
			checkUserNonceInContract(targetChain, client, issuer, srcChainId)
		}
		return "", nil, err
	}

	return transaction.Hash().Hex(), nonce, nil
}

func checkUserNonceInContract(chain models.Chain, client ethclient.Client,
	issuer common.Address, srcChainId *big.Int) {
	caller, err := vault.NewVaultCaller(common.HexToAddress(chain.VaultAddr), &client)
	if err != nil {
		logrus.WithError(err).Error("vault.NewVaultCaller fail")
		return
	}
	n, err := caller.GetUserNextClaimNonce(&bind.CallOpts{}, issuer, srcChainId)
	if err != nil {
		logrus.WithError(err).Error("GetUserNextClaimNonce fail")
		return
	}
	logrus.WithFields(logrus.Fields{
		"fromChain.id": srcChainId, "toChain.id": chain.Id, "toChain.name": chain.Name,
	}).Info("user next nonce ", n)
	os.Exit(1)
}

func GetNonceAndGas(client ethclient.Client, fromAddress *common.Address) (*uint64, *big.Int, error) {
	nonce, err := client.PendingNonceAt(context.Background(), *fromAddress)
	if err != nil {
		return nil, nil, err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, nil, err
	}
	return &nonce, gasPrice, nil
}

func CreateKeyPair(hexKey string) (*ecdsa.PrivateKey, *common.Address, error) {
	privateKey, err := crypto.HexToECDSA(hexKey)
	if err != nil {
		return nil, nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, nil, fmt.Errorf("generate public key fail")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	return privateKey, &fromAddress, nil
}
