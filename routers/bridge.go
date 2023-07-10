package routers

import (
	"github.com/Conflux-Chain/go-conflux-util/api"
	"github.com/gin-gonic/gin"
	"github.com/tree-graph/bridge-service/controllers"
)

func BridgeRoutes(route *gin.Engine) {
	bridgeRoute := route.Group("/bridge")
	bridgeRoute.POST("crossRequest", api.Wrap(controllers.CrossRequest))
	bridgeRoute.GET("pluginChainInfo", api.Wrap(controllers.PluginChainInfo))
	bridgeRoute.GET("eventCursor", api.Wrap(controllers.EventCursor))
	bridgeRoute.GET("claimCursor", api.Wrap(controllers.ClaimCursor))
	bridgeRoute.POST("reportCrossRequest", api.Wrap(controllers.ReportCrossRequest))
	bridgeRoute.GET("getPooledClaim", api.Wrap(controllers.GetPooledClaim))
	bridgeRoute.GET("moveToHistory", api.Wrap(controllers.MoveToHistory))
	bridgeRoute.GET("updateClaimPool", api.Wrap(controllers.UpdateClaimPool))
	bridgeRoute.GET("checkClaimTask", api.Wrap(controllers.CheckClaimTask))
	bridgeRoute.GET("buildCrossRequest", api.Wrap(controllers.BuildCrossRequest))
}
