package helpers

import (
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
