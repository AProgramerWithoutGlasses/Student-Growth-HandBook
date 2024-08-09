package routes

import (
	"github.com/gin-gonic/gin"
	"studentGrow/controller/message"
	"studentGrow/utils/token"
)

func routesMsg(r *gin.Engine) {
	gp := r.Group("/report_box")
	// 查看举报信息
	gp.POST("/getlist", token.AuthMiddleware(), message.GetUnreadReportsController)

	// 确认举报信息
	gp.POST("/ack", token.AuthMiddleware(), message.AckUnreadReportsController)
}
