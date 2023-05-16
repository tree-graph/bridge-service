package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/tree-graph/bridge-service/infra/evm"
	"github.com/tree-graph/bridge-service/models"
	"net/http"
)

type BridgeController struct{}

func (ctrl *BridgeController) CrossRequest(ctx *gin.Context) {
	bean := new(models.CrossRequest)

	err := ctx.ShouldBindJSON(&bean)
	if err != nil {
		logrus.Errorf("bind json error: %v", err)
		SendError(ctx, 400, err.Error())
		return
	}
	_, err = evm.AddRequest(bean.ChainId, bean.Hash)
	if err != nil {
		logrus.Errorf("save cross request hash error: %v", err)
		SendError(ctx, 400, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, &bean)
}
