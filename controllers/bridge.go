package controllers

import (
	"fmt"
	"github.com/Conflux-Chain/go-conflux-util/api"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/tree-graph/bridge-service/infra/blockchain"
	"github.com/tree-graph/bridge-service/infra/database"
	"github.com/tree-graph/bridge-service/infra/worker"
	"github.com/tree-graph/bridge-service/models"
	"strconv"
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

func UpdateClaimPool(ctx *gin.Context) (interface{}, error) {
	id := ctx.Query("id")
	claimTxHash := ctx.Query("claimTxHash")
	txFrom := ctx.Query("txFrom")
	nonce, err := strconv.ParseUint(ctx.Query("nonce"), 10, 64)
	if err != nil {
		return nil, errors.WithMessage(err, "invalid nonce")
	}
	var pooledClaim models.ClaimPool
	if err := database.DB.Where("id=?", id).
		Take(&pooledClaim).Error; err != nil {
		return nil, errors.WithMessage(err, "claimPool query failed")
	}
	err = worker.UpdateClaimPool(pooledClaim, claimTxHash, txFrom, nonce)
	if err != nil {
		return nil, errors.WithMessage(err, "claimPoll update failed")
	}
	return "OK", nil
}
func MoveToHistory(ctx *gin.Context) (interface{}, error) {
	id := ctx.Query("id")
	statusStr := ctx.Query("status")
	status, err := strconv.ParseUint(statusStr, 10, 64)
	if err != nil {
		return nil, err
	}
	comment := ctx.Query("comment")
	var pooledClaim models.ClaimPool
	if err := database.DB.Where("id=?", id).
		Take(&pooledClaim).Error; err != nil {
		return nil, err
	}
	if err := worker.MoveClaimFromPoolToHistory(pooledClaim, status, comment); err != nil {
		return nil, err
	}
	return "OK", nil
}

func BuildCrossRequest(ctx *gin.Context) (interface{}, error) {
	id := ctx.Query("id")
	var bean models.CrossInfo
	if err := database.GetDB().Where("id=?", id).Take(&bean).Error; err != nil {
		return nil, err
	}
	request, err := worker.BuildCrossRequest(bean)
	if err != nil {
		return nil, err
	}
	var items []models.CrossItem
	for idx, id := range request.TokenIds {
		items = append(items, models.CrossItem{
			TokenId: hexutil.EncodeBig(id),
			Amount:  hexutil.EncodeBig(request.Amounts[idx]),
			Uri:     request.Uris[idx],
		})
	}
	report := models.ReportCrossRequest{
		Infos: []models.CrossInfo{
			{
				Id:             0,
				SourceChain:    bean.SourceChain,
				TxnHash:        "",
				Asset:          request.Asset.Hex(),
				From:           request.From.Hex(),
				TargetChain:    request.ToChainId.Int64(),
				TargetContract: request.TargetContract.Hex(),
				UserNonce:      request.UserNonce.Int64(),
				BlockNumber:    0,
				BlockTime:      nil,
			},
		},
		Items: [][]models.CrossItem{
			items,
		},
	}
	return report, nil
}
func CheckClaimTask(ctx *gin.Context) (interface{}, error) {
	id := ctx.Query("id")
	clientName := ctx.Query("clientName")
	uid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	delaySeconds, err := worker.CheckClaimTask(uid, clientName)
	if err != nil {
		return nil, err
	}
	return delaySeconds, nil
}
func GetPooledClaim(ctx *gin.Context) (interface{}, error) {
	id := ctx.Query("id")
	uid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	hasPooledClaim, pooledClaim, err := worker.HasPooledClaim(uid, "")
	if err != nil {
		return 0, err
	}
	ret := make(map[string]interface{})
	ret["hasPooledClaim"] = hasPooledClaim
	if hasPooledClaim {
		ret["claim"] = pooledClaim
	}
	return ret, nil
}

func ReportCrossRequest(ctx *gin.Context) (interface{}, error) {
	bean := new(models.ReportCrossRequest)
	err := ctx.ShouldBindJSON(&bean)
	if err != nil {
		return nil, errors.WithMessage(err, "invalid parameter")
	}
	err = worker.SaveCrossRequest(bean.Infos, bean.Items, bean.ChainId, bean.Cursor)
	if err != nil {
		return "fail", err
	}
	return "OK", err
}

func ClaimCursor(ctx *gin.Context) (interface{}, error) {
	id := ctx.Param("id")
	uid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}

	var bean models.ClaimCursor
	err = database.DB.Where("id=?", uid).Take(&bean).Error
	return bean, err
}

func EventCursor(ctx *gin.Context) (interface{}, error) {
	id := ctx.Query("id")
	uid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}

	var bean models.ChainCursor
	err = database.DB.Where("id=?", uid).Take(&bean).Error
	return bean, err
}

func PluginChainInfo(ctx *gin.Context) (interface{}, error) {
	id := ctx.Query("id")
	uid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	chain, err := models.GetChain(uid)
	if err != nil {
		return nil, err
	}
	obj := make(map[string]interface{})
	obj["name"] = chain.Name
	obj["id"] = chain.Id
	obj["vaultAddr"] = chain.VaultAddr
	obj["rpc"] = chain.Rpc
	obj["chainId"] = chain.ChainId
	obj["chainType"] = chain.ChainType
	obj["delayBlock"] = chain.DelayBlock

	return obj, nil
}
