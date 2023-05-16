package routers

import (
	"github.com/Conflux-Chain/go-conflux-util/api"
	"github.com/gin-gonic/gin"
	"github.com/tree-graph/bridge-service/controllers"
)

func BridgeRoutes(route *gin.Engine) {
	v1 := route.Group("/bridge")
	v1.POST("crossRequest", api.Wrap(controllers.CrossRequest))
}
