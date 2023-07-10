package helpers

import (
	"crypto/ecdsa"
	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/Conflux-Chain/go-conflux-sdk/utils/addressutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/sirupsen/logrus"
	"math/big"
)

func CheckFatalError(msg string, err error) {
	if err != nil {
		logrus.WithError(err).Fatalln(msg)
	}
}
func Uint64ToBigInt(v uint64) *big.Int {
	return new(big.Int).SetUint64(v)
}

func CfxAddressOfPrivate(key string, networkID uint32) (*types.Address, error) {
	privateKey, err := crypto.HexToECDSA(key)
	if err != nil {
		return nil, err
	}
	return CfxAddressOfPrivateKey(privateKey, networkID), nil
}
func CfxAddressOfPrivateKey(key *ecdsa.PrivateKey, networkID uint32) *types.Address {
	ethAddr := crypto.PubkeyToAddress(key.PublicKey)
	cfxAddr := addressutil.EtherAddressToCfxAddress(ethAddr, false, networkID)
	return &cfxAddr
}
