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
		logrus.Errorf("bind json error: %v", err)
		return nil, err
	}
	_, err = blockchain.AddCrossRequest(bean.ChainId, bean.Hash)
	if err != nil {
		logrus.Errorf("save cross request hash error: %v", err)
		return nil, err
	}
	return bean, nil
}
