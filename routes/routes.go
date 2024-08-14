package routes

import (
	"github.com/gin-gonic/gin"
	"studentGrow/logger"
	"studentGrow/utils/middleWare"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.Use(middleWare.CORSMiddleware())

	// 勋
	routesHomepage(r)
	routesStudentManage(r)

	return r
}
