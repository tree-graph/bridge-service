package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/tree-graph/bridge-service/infra/blockchain"
	"github.com/tree-graph/bridge-service/models"
)

func CrossRequest(ctx *gin.Context) (interface{}, error) {
	bean := new(models.CrossRequest)

	err := ctx.ShouldBindJSON(&bean)
	if err != nil {
		logrus.WithError(err).Error("bind json error")
		return nil, err
	}
	_, err = blockchain.AddCrossRequest(bean.ChainId, bean.Hash)
	if err != nil {
		logrus.WithError(err).Error("save cross request hash error")
		return nil, err
	}
	return bean, nil
}
