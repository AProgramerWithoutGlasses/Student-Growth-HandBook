package routes

import (
	"github.com/gin-gonic/gin"
	"studentGrow/controller/user"
	"studentGrow/logger"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	//r.POST("/article/content", article.GetArticleId)
	r.POST("/user/getSelfCotnent", user.GetSelfContentContro)
	r.POST("/user/updateSelfContent", user.UpdateSelfContentContro)

	return r
}
