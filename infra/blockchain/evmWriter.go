package blockchain

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/tree-graph/bridge-service/infra/contracts/vault"
	"math/big"
)

func SendClaimTx(pkStr string, vaultAddr string,
	srcChainId *big.Int,
	request vault.VaultCrossRequest, client ethclient.Client) (string, error) {
	privateKey, address, err := CreateKeyPair(pkStr)
	if err != nil {
		return "", err
	}
	nonce, gasPrice, err := GetNonceAndGas(client, address)
	if err != nil {
		return "", err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, request.ToChainId)
	if err != nil {
		return "", err
	}
	auth.Nonce = big.NewInt(int64(*nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	valutContract, err := vault.NewVaultTransactor(common.HexToAddress(vaultAddr), &client)
	if err != nil {
		return "", err
	}
	srcContract := request.Asset
	localContract := request.TargetContract
	tokenIds := request.TokenIds
	amounts := request.Amounts
	issuer := request.From
	userNonce := request.UserNonce
	uris := request.Uris

	transaction, err := valutContract.ClaimByAdmin(auth, srcChainId, srcContract, localContract, tokenIds, amounts, uris, issuer, userNonce)
	if err != nil {
		return "", err
	}

	return transaction.Hash().Hex(), nil
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
