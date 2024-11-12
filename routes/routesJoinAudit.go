package routes

import (
	"github.com/gin-gonic/gin"
	"studentGrow/controller/routesJoinAudit"
	"studentGrow/utils/token"
)

func RoutsJoinAudit(router *gin.Engine) {
	r := router.Group("/routesJoinAudit")
	r.Use(token.AuthMiddleware())
	r.POST("/isOpen", routesJoinAudit.OpenMsg)
	r.GET("/StuForm", routesJoinAudit.GetStudForm)
	r.POST("/StuForm", routesJoinAudit.SaveStudForm)
	r.GET("/StudFile", routesJoinAudit.GetStuFile)
	r.POST("/StudFile", routesJoinAudit.SaveStuFile)
	r.POST("/DelStudFile", routesJoinAudit.DelStuFile)
	r.GET("activityManage", routesJoinAudit.GetActivityMsg)
}
