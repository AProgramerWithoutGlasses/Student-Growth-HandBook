package routes

import (
	"github.com/gin-gonic/gin"
	"studentGrow/controller/article"
	"studentGrow/logger"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.POST("/article/content", article.GetArticleId)

	return r
}
