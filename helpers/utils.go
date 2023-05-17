package helpers

import (
	"github.com/sirupsen/logrus"
)

func CheckFatalError(msg string, err error) {
	if err != nil {
		logrus.Fatalln(msg, err.Error())
	}
}
