package controllers

import (
	"fmt"
	"github.com/Conflux-Chain/go-conflux-util/api"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/tree-graph/bridge-service/infra/blockchain"
	"github.com/tree-graph/bridge-service/models"
	"strings"
)

func CrossRequest(ctx *gin.Context) (interface{}, error) {
	bean := new(models.CrossRequest)

	err := ctx.ShouldBindJSON(&bean)
	if err != nil {
		logrus.WithError(err).Debug("bind json error")
		return nil, err
	}
	_, err = blockchain.AddCrossRequest(bean.ChainId, bean.Hash)
	if err == nil {
		return bean, nil
	}
	if strings.HasPrefix(err.Error(), "Error 1062: Duplicate entry") {
		return nil, api.NewBusinessError(DuplicateEntryError, err.Error(), "")
	}
	logrus.WithFields(logrus.Fields{
		"bean": fmt.Sprintf("%+v", bean),
	}).WithError(err).Error("save cross request hash error")
	return nil, err
}
