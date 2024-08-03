package routes

import (
	"github.com/gin-gonic/gin"
	"studentGrow/logger"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	routesArticle(r)
	routesTopic(r)
	return r
}
