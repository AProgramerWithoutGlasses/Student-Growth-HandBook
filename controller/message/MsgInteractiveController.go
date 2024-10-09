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
		Limit int `json:"limit"`
		Page  int `json:"page"`
	}{}
	username, err := token.GetUsername(c.GetHeader("token"))
	if err != nil {
		zap.L().Error("GetSystemMsgController() controller.message.GetUsername err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}

	err = c.ShouldBindJSON(&in)
	if err != nil {
		zap.L().Error("GetSystemMsgController() controller.message.ShouldBindJSON err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}

	msgs, UnreadCount, err := message.GetSystemMsgService(in.Limit, in.Page, username)
	if err != nil {
		zap.L().Error("GetSystemMsgController() controller.message.GetSystemMsgService err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}

	list := make([]map[string]any, 0)
	for _, msg := range msgs {
		list = append(list, map[string]any{
			"ID":            msg.ID,
			"msg_content":   msg.Content,
			"msg_time":      msg.Time,
			"username":      username,
			"user_headshot": msg.OwnUser.HeadShot,
			"is_read":       msg.IsRead,
		})
	}

	res.ResponseSuccess(c, map[string]any{
		"admin_info":   list,
		"unread_count": UnreadCount,
	})
}

// GetManagerMsgController 获取管理员消息
func GetManagerMsgController(c *gin.Context) {
	in := struct {
		Limit int `json:"limit"`
		Page  int `json:"page"`
	}{}

	username, err := token.GetUsername(c.GetHeader("token"))
	if err != nil {
		zap.L().Error("GetSystemMsgController() controller.message.GetUsername err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}

	err = c.ShouldBindJSON(&in)
	if err != nil {
		zap.L().Error("GetManagerMsgController() controller.message.ShouldBindJSON err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}

	msgs, UnreadCount, err := message.GetManagerMsgService(in.Limit, in.Page, username)
	if err != nil {
		zap.L().Error("GetManagerMsgController() controller.message.GetManagerMsgService err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}

	list := make([]map[string]any, 0)
	for _, msg := range msgs {
		list = append(list, map[string]any{
			"ID":            msg.ID,
			"msg_content":   msg.Content,
			"msg_time":      msg.Time,
			"username":      username,
			"user_headshot": msg.OwnUser.HeadShot,
			"is_read":       msg.IsRead,
		})
	}

	res.ResponseSuccess(c, map[string]any{
		"manager_info": list,
		"unread_count": UnreadCount,
	})
}

// GetLikeMsgController 获取点赞消息
func GetLikeMsgController(c *gin.Context) {
	in := struct {
		Page  int `json:"page"`
		Limit int `json:"limit"`
	}{}

	username, err := token.GetUsername(c.GetHeader("token"))
	if err != nil {
		zap.L().Error("GetSystemMsgController() controller.message.GetUsername err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}

	err = c.ShouldBindJSON(&in)
	if err != nil {
		zap.L().Error("GetLikeMsgController() controller.message.ShouldBindJSON err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}

	list, num, err := message.GetArticleAndCommentLikedMsgService(username, in.Page, in.Limit)
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
		Page  int `json:"page"`
		Limit int `json:"limit"`
	}{}

	username, err := token.GetUsername(c.GetHeader("token"))
	if err != nil {
		zap.L().Error("GetSystemMsgController() controller.message.GetUsername err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}

	err = c.ShouldBindJSON(&in)
	if err != nil {
		zap.L().Error("GetCollectMsgController() controller.message.ShouldBindJSON err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}

	list, num, err := message.GetCollectMsgService(username, in.Page, in.Limit)
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
		Page  int `json:"page"`
		Limit int `json:"limit"`
	}{}

	username, err := token.GetUsername(c.GetHeader("token"))
	if err != nil {
		zap.L().Error("GetSystemMsgController() controller.message.GetUsername err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}

	err = c.ShouldBindJSON(&in)
	if err != nil {
		zap.L().Error("GetCollectMsgController() controller.message.ShouldBindJSON err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}

	comments, num, err := message.GetCommentMsgService(username, in.Page, in.Limit)
	if err != nil {
		zap.L().Error("GetCollectMsgController() controller.message.GetCommentMsgService err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}

	res.ResponseSuccess(c, map[string]any{
		"comments":     comments,
		"unread_count": num,
	})
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

	res.ResponseSuccess(c, struct{}{})
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

	in := struct {
		MsgId int `json:"msg_id"`
	}{}
	err = c.ShouldBindJSON(&in)
	if err != nil {
		zap.L().Error("AckManagerMsgController() controller.message.ShouldBindJSON err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}
	// 确认消息
	err = message.AckManagerMsgService(username, in.MsgId)
	if err != nil {
		zap.L().Error("AckManagerMsgController() controller.message.AckManagerMsgService err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}
	res.ResponseSuccess(c, struct{}{})
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

	in := struct {
		MsgId int `json:"msg_id"`
	}{}

	err = c.ShouldBindJSON(&in)
	if err != nil {
		zap.L().Error("AckSystemMsgController() controller.message.GetUsername err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}

	err = message.AckSystemMsgService(username, in.MsgId)
	if err != nil {
		zap.L().Error("AckSystemMsgController() controller.message.AckSystemMsgService err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}
	res.ResponseSuccess(c, struct{}{})
}

// PublishManagerMsgController 发布管理员通知
func PublishManagerMsgController(c *gin.Context) {
	in := struct {
		Content string `json:"msg_content"`
	}{}
	err := c.ShouldBindJSON(&in)
	if err != nil {
		zap.L().Error("PublishManagerMsgController() controller.message.ShouldBindJSON err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}

	username, err := token.GetUsername(c.GetHeader("token"))
	if err != nil {
		zap.L().Error("PublishManagerMsgController() controller.message.GetUsername err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}

	// 获取角色
	role, err := token.GetRole(c.GetHeader("token"))
	if err != nil {
		zap.L().Error("PublishManagerMsgController() controller.message.GetRole err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}

	err = message.PublishManagerMsgService(username, in.Content, role)
	if err != nil {
		zap.L().Error("PublishManagerMsgController() controller.message.PublishManagerMsgService err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}

	err = message.PublishSystemMsgService(in.Content, role, username)
	if err != nil {
		zap.L().Error("PublishSystemMsgController() controller.message.PublishSystemMsgService err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}

	res.ResponseSuccess(c, struct{}{})
}

// PublishSystemMsgController 发布系统通知
func PublishSystemMsgController(c *gin.Context) {
	in := struct {
		Content string `json:"msg_content"`
	}{}
	err := c.ShouldBindJSON(&in)
	if err != nil {
		zap.L().Error("PublishSystemMsgController() controller.message.ShouldBindJSON err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}

	role, err := token.GetRole(c.GetHeader("token"))
	if err != nil {
		zap.L().Error("PublishSystemMsgController() controller.message.GetRole err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}

	username, err := token.GetUsername(c.GetHeader("token"))
	if err != nil {
		zap.L().Error("PublishSystemMsgController() controller.message.GetUsername err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}

	err = message.PublishSystemMsgService(in.Content, role, username)
	if err != nil {
		zap.L().Error("PublishSystemMsgController() controller.message.PublishSystemMsgService err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}

	res.ResponseSuccess(c, struct{}{})
}
