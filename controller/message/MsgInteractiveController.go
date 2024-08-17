package message

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	myErr "studentGrow/pkg/error"
	res "studentGrow/pkg/response"
	"studentGrow/service/message"
	"studentGrow/utils/token"
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

	// 获取用户身份
	role, err := token.GetRole(c.GetHeader("token"))
	if err != nil {
		zap.L().Error("GetManagerMsgController() controller.message.GetRole err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}

	msg, UnreadCount, err := message.GetManagerMsgService(in.Limit, in.Page, in.Username)
	if err != nil {
		zap.L().Error("GetManagerMsgController() controller.message.GetManagerMsgService err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}

	res.ResponseSuccess(c, map[string]any{
		"manager_info": msg,
		"role":         role,
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

	list, num, err := message.GetArticleAndCommentLikedMsgService(in.Username, in.Page, in.Limit)
	if err != nil {
		zap.L().Error("GetLikeMsgController() controller.message.GetArticleAndCommentLikedMsgService err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}

	res.ResponseSuccess(c, map[string]any{
		"thumbsUp":     list,
		"unread_count": num,
	})
}

// GetCollectMsgController 获取收藏消息
func GetCollectMsgController(c *gin.Context) {
	in := struct {
		Username string `json:"username"`
		Page     int    `json:"page"`
		Limit    int    `json:"limit"`
	}{}

	err := c.ShouldBindJSON(&in)
	if err != nil {
		zap.L().Error("GetCollectMsgController() controller.message.ShouldBindJSON err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}

	list, num, err := message.GetCollectMsgService(in.Username, in.Page, in.Limit)
	if err != nil {
		zap.L().Error("GetCollectMsgController() controller.message.GetCollectMsgService err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}

	res.ResponseSuccess(c, map[string]any{
		"star":         list,
		"unread_count": num,
	})

}

// GetCommentMsgController 获取评论消息
func GetCommentMsgController(c *gin.Context) {
	in := struct {
		Username string `json:"username"`
		Page     int    `json:"page"`
		Limit    int    `json:"limit"`
	}{}

	err := c.ShouldBindJSON(&in)
	if err != nil {
		zap.L().Error("GetCollectMsgController() controller.message.ShouldBindJSON err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}

	comments, err := message.GetCommentMsgService(in.Username, in.Page, in.Limit)
	if err != nil {
		zap.L().Error("GetCollectMsgController() controller.message.GetCommentMsgService err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}

	res.ResponseSuccess(c, comments)
}

// AckInterMsgController 确认互动消息通知
func AckInterMsgController(c *gin.Context) {
	in := struct {
		MsgId int `json:"msg_id"`
		Type  int `json:"type"`
	}{}

	err := c.ShouldBindJSON(&in)
	if err != nil {
		zap.L().Error("AckInterMsgController() controller.message.ShouldBindJSON err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}

	// 确认消息
	err = message.AckInterMsgService(in.MsgId, in.Type)
	if err != nil {
		zap.L().Error("AckInterMsgController() controller.message.AckInterMsgService err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}

	res.ResponseSuccess(c, nil)
}

// AckManagerMsgController 确认管理员消息
func AckManagerMsgController(c *gin.Context) {
	// 获取username
	username, err := token.GetUsername(c.GetHeader("token"))
	if err != nil {
		zap.L().Error("AckManagerMsgController() controller.message.GetUsername err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}

	// 确认消息
	err = message.AckManagerMsgService(username)
	if err != nil {
		zap.L().Error("AckManagerMsgController() controller.message.AckManagerMsgService err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}
	res.ResponseSuccess(c, nil)
}

// AckSystemMsgController 确认系统消息
func AckSystemMsgController(c *gin.Context) {
	// 获取username
	username, err := token.GetUsername(c.GetHeader("token"))
	if err != nil {
		zap.L().Error("AckSystemMsgController() controller.message.GetUsername err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}

	err = message.AckSystemMsgService(username)
	if err != nil {
		zap.L().Error("AckSystemMsgController() controller.message.AckSystemMsgService err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}
	res.ResponseSuccess(c, nil)
}
