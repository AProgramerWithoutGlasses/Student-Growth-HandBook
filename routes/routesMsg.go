package routes

import (
	"github.com/gin-gonic/gin"
	"studentGrow/controller/message"
)

func routesMsg(r *gin.Engine) {
	// 获取未读举报信息列表
	r.POST("report_box/get_list", message.GetUnreadReportsController)
}
