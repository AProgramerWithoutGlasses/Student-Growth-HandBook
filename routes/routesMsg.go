package routes

import (
	"github.com/gin-gonic/gin"
	"studentGrow/controller/message"
	"studentGrow/utils/middleWare"
	"studentGrow/utils/token"
)

func routesMsg(r *gin.Engine) {
	gp := r.Group("/report_box")
	// 查看举报信息
	gp.POST("/getlist", token.AuthMiddleware(), message.GetUnreadReportsController)

	// 确认举报信息
	gp.POST("/ack", token.AuthMiddleware(), message.AckUnreadReportsController)

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
	msg.POST("/ack_interactMsg", middleWare.CORSMiddleware(), message.AckInterMsgController)

	// 确认系统消息
	msg.POST("/ack_systemMsg", middleWare.CORSMiddleware(), message.AckSystemMsgController)

	// 确认管理员消息
	msg.POST("/ack_managerMsg", middleWare.CORSMiddleware(), message.AckManagerMsgController)

	// 发布管理员通知
	msg.POST("/publish_managerMsg", middleWare.CORSMiddleware(), message.PublishManagerMsgController)

	// 发布系统通知
	msg.POST("/publish_systemMsg", middleWare.CORSMiddleware(), message.PublishSystemMsgController)

}
