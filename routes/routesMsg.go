package routes

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"studentGrow/controller/message"
	"studentGrow/dao/mysql"
	"studentGrow/models/casbinModels"
	"studentGrow/utils/middleWare"
	"studentGrow/utils/token"
)

func routesMsg(r *gin.Engine) {

	casbinService, err := casbinModels.NewCasbinService(mysql.DB)
	if err != nil {
		zap.L().Error("routesMsg() routes.routesArticle.NewCasbinService err=", zap.Error(err))
		return
	}
	gp := r.Group("/report_box")
	// 查看举报信息
	gp.POST("/getlist", middleWare.NewCasbinAuth(casbinService), token.AuthMiddleware(), message.GetUnreadReportsController)

	// 确认举报信息
	gp.POST("/ack", middleWare.NewCasbinAuth(casbinService), token.AuthMiddleware(), message.AckUnreadReportsController)

	msg := r.Group("/message")

	// 获取系统消息
	msg.POST("/get_system", message.GetSystemMsgController)

	// 获取管理员消息
	msg.POST("/get_manager", message.GetManagerMsgController)

	// 获取点赞消息
	msg.POST("/get_thumbList", message.GetLikeMsgController)

	// 获取收藏消息
	msg.POST("/get_starList", message.GetCollectMsgController)

	// 获取评论消息
	msg.POST("/get_comList", message.GetCommentMsgController)

	// 确认互动消息
	msg.POST("/ack_interactMsg", token.AuthMiddleware(), message.AckInterMsgController)

	// 确认系统消息
	msg.POST("/ack_systemMsg", token.AuthMiddleware(), message.AckSystemMsgController)

	// 确认管理员消息
	msg.POST("/ack_managerMsg", token.AuthMiddleware(), message.AckManagerMsgController)

	// 发布管理员通知
	msg.POST("/publish_managerMsg", token.AuthMiddleware(), message.PublishManagerMsgController)

	// 发布系统通知
	msg.POST("/publish_systemMsg", token.AuthMiddleware(), message.PublishSystemMsgController)

}
