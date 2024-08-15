package message

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	myErr "studentGrow/pkg/error"
	res "studentGrow/pkg/response"
	"studentGrow/service/article"
	"studentGrow/service/message"
)

// GetSystemMsgController 获取系统消息
func GetSystemMsgController(c *gin.Context) {
	in := struct {
		Limit    int    `json:"limit"`
		Page     int    `json:"page"`
		Username string `json:"username"`
	}{}

	err := c.ShouldBindJSON(&in)
	if err != nil {
		zap.L().Error("GetSystemMsgController() controller.message.ShouldBindJSON err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}

	msg, UnreadCount, err := message.GetSystemMsgService(in.Limit, in.Page, in.Username)
	if err != nil {
		zap.L().Error("GetSystemMsgController() controller.message.GetSystemMsgService err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}

	res.ResponseSuccess(c, map[string]any{
		"admin_info":   msg,
		"unread_count": UnreadCount,
	})
}

// GetManagerMsgController 获取管理员消息
func GetManagerMsgController(c *gin.Context) {
	in := struct {
		Limit    int    `json:"limit"`
		Page     int    `json:"page"`
		Username string `json:"username"`
	}{}

	err := c.ShouldBindJSON(&in)
	if err != nil {
		zap.L().Error("GetManagerMsgController() controller.message.ShouldBindJSON err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}

	msg, UnreadCount, err := message.GetSystemMsgService(in.Limit, in.Page, in.Username)
	if err != nil {
		zap.L().Error("GetManagerMsgController() controller.message.GetSystemMsgService err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}

	res.ResponseSuccess(c, map[string]any{
		"manager_info": msg,
		"unread_count": UnreadCount,
	})
}

// GetLikeMsgController 获取点赞消息
func GetLikeMsgController(c *gin.Context) {
	in := struct {
		Username string `json:"username"`
		Page     int    `json:"page"`
		Limit    int    `json:"limit"`
	}{}

	err := c.ShouldBindJSON(&in)
	if err != nil {
		zap.L().Error("GetLikeMsgController() controller.message.ShouldBindJSON err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}

	list, num, err := article.GetArticleAndCommentLikedList(in.Username, in.Page, in.Limit)
	if err != nil {
		zap.L().Error("GetLikeMsgController() controller.message.GetArticleAndCommentLikedList err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}

	res.ResponseSuccess(c, map[string]any{
		"thumbsUp":     list,
		"unread_count": num,
	})
}
