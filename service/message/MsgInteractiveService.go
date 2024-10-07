package message

import (
	"go.uber.org/zap"
	"studentGrow/dao/mysql"
	"studentGrow/models/constant"
	"studentGrow/models/gorm_model"
	"studentGrow/models/nzx_model"
	myErr "studentGrow/pkg/error"
	"studentGrow/pkg/sse"
	NotificationPush "studentGrow/service/notificationPush"
	"studentGrow/utils/timeConverter"
)

// GetSystemMsgService 获取系统消息通知
func GetSystemMsgService(limit, page int, username string) ([]gorm_model.MsgRecord, int, error) {
	// 获取uid
	user, err := mysql.GetUserByUsername(username)
	if err != nil {
		return nil, 0, err
	}

	msgs, err := mysql.QuerySystemMsg(page, limit, int(user.ID))
	if err != nil {
		zap.L().Error("GetSystemMsgService() service.message.QuerySystemMsg", zap.Error(err))
		return nil, -1, err
	}

	// 查询未读消息条数
	count, err := mysql.QueryUnreadSystemMsg(int(user.ID))
	if err != nil {
		zap.L().Error("GetSystemMsgService() service.message.QueryUnreadSystemMsg", zap.Error(err))
		return nil, -1, err
	}

	for i := 0; i < len(msgs); i++ {
		msgs[i].Time = timeConverter.IntervalConversion(msgs[i].CreatedAt)
	}

	return msgs, count, nil
}

// GetManagerMsgService 获取管理员消息
func GetManagerMsgService(limit, page int, username string) ([]gorm_model.MsgRecord, int, error) {
	// 获取uid
	uid, err := mysql.GetIdByUsername(username)
	if err != nil {
		zap.L().Error("GetManagerMsgService() service.message.GetIdByUsername", zap.Error(err))
		return nil, 0, err
	}

	msgs, err := mysql.QueryManagerMsg(page, limit, uid)
	if err != nil {
		zap.L().Error("GetManagerMsgService() service.message.QueryManagerMsg", zap.Error(err))
		return nil, -1, err
	}

	// 查询未读消息条数
	count, err := mysql.QueryUnreadManagerMsg(uid)
	if err != nil {
		zap.L().Error("GetManagerMsgService() service.message.QueryUnreadManagerMsg", zap.Error(err))
		return nil, -1, err
	}

	for i := 0; i < len(msgs); i++ {
		msgs[i].Time = timeConverter.IntervalConversion(msgs[i].CreatedAt)
	}

	return msgs, count, nil
}

// GetArticleAndCommentLikedMsgService  获取点赞消息
func GetArticleAndCommentLikedMsgService(username string, page, limit int) ([]nzx_model.Out, int, error) {
	// 获取uid
	uid, err := mysql.GetIdByUsername(username)
	if err != nil {
		zap.L().Error("GetArticleAndCommentLikedMsgService() service.article.likeService.GetIdByUsername err=", zap.Error(err))
		return nil, -1, err
	}

	// 获取点赞列表
	likes, err := mysql.QueryLikeRecordByUser(uid, page, limit)
	if err != nil {
		zap.L().Error("GetArticleAndCommentLikedMsgService() service.article.likeService.QueryLikeRecordByUserArticle err=", zap.Error(err))
		return nil, -1, err
	}

	// 获取文章点赞未读消息总数
	sum, err := mysql.QueryLikeRecordNumByUser(uid)
	if err != nil {
		zap.L().Error("GetArticleAndCommentLikedMsgService() service.article.likeService.QueryLikeRecordNumByUser err=", zap.Error(err))
		return nil, -1, err
	}

	list := make([]nzx_model.Out, 0)

	for _, like := range likes {
		// 判断文章点赞还是评论点赞
		usernameL := like.User.Username
		name := like.User.Name
		content := like.Article.Content
		userHeadshot := like.User.HeadShot
		likeType := 0
		articleId := like.ArticleID
		if like.Type == 1 {
			content = like.Comment.Content
			likeType = 1
			articleId = like.Comment.ArticleID
		}

		list = append(list, nzx_model.Out{
			Username:     usernameL,
			Name:         name,
			Content:      content,
			UserHeadshot: userHeadshot,
			PostTime:     timeConverter.IntervalConversion(like.CreatedAt),
			IsRead:       like.IsRead,
			Type:         likeType,
			ArticleId:    articleId,
			MsgId:        like.ID,
		})
	}

	return list, sum, nil
}

// GetCollectMsgService 获取收藏消息
func GetCollectMsgService(username string, page, limit int) ([]map[string]any, int, error) {
	// 获取uid
	uid, err := mysql.GetIdByUsername(username)
	if err != nil {
		zap.L().Error("GetCollectMsgService() service.article.likeService.GetIdByUsername err=", zap.Error(err))
		return nil, -1, err
	}

	// 获取收藏消息列表
	articleCollects, err := mysql.QueryCollectRecordByUserArticles(uid, page, limit)
	if err != nil {
		zap.L().Error("GetCollectMsgService() service.article.likeService.QueryCollectRecordByUserArticles err=", zap.Error(err))
		return nil, -1, err
	}

	// 获取未读收藏消息数量
	collectNum, err := mysql.QueryCollectRecordNumByUserArticle(uid)
	if err != nil {
		zap.L().Error("GetCollectMsgService() service.article.likeService.QueryCollectRecordNumByUserArticle err=", zap.Error(err))
		return nil, -1, err
	}

	list := make([]map[string]any, 0)

	for _, collect := range articleCollects {
		list = append(list, map[string]any{
			"username":        collect.User.Username,
			"name":            collect.User.Name,
			"article_content": collect.Article.Content,
			"user_headshot":   collect.User.HeadShot,
			"post_time":       timeConverter.IntervalConversion(collect.CreatedAt),
			"is_read":         collect.IsRead,
			"article_id":      collect.ArticleID,
			"msg_id":          collect.ID,
		})
	}

	return list, collectNum, nil

}

// GetCommentMsgService 获取评论消息
func GetCommentMsgService(username string, page, limit int) (nzx_model.CommentMsgs, int, error) {
	// 获取uid
	uid, err := mysql.GetIdByUsername(username)
	if err != nil {
		zap.L().Error("GetCommentMsgService() service.article.likeService.GetIdByUsername err=", zap.Error(err))
		return nil, -1, err
	}

	// 获取所有评论及回复
	comments, err := mysql.QueryCommentRecordByUserArticles(uid, page, limit)
	if err != nil {
		zap.L().Error("GetCommentMsgService() service.article.likeService.QueryCommentRecordByUserArticles err=", zap.Error(err))
		return nil, -1, err
	}

	commentMsgs := make(nzx_model.CommentMsgs, 0)

	for _, comment := range comments {
		// 默认文章的评论
		commentType := 0
		if comment.Pid != 0 {
			// 如果是评论的回复
			commentType = 1
		}

		commentMsgs = append(commentMsgs, nzx_model.CommentMsg{
			Username:     comment.User.Username,
			Name:         comment.User.Name,
			Content:      comment.Content,
			UserHeadshot: comment.User.HeadShot,
			PostTime:     timeConverter.IntervalConversion(comment.CreatedAt),
			IsRead:       comment.IsRead,
			Type:         commentType,
			ArticleId:    comment.ArticleID,
			MsgId:        comment.ID,
		})
	}

	// 获取未读评论数
	num, err := mysql.QueryCommentRecordNumByUserId(uid)
	if err != nil {
		zap.L().Error("GetCommentMsgService() service.article.likeService.QueryCommentRecordNumByUserId err=", zap.Error(err))
		return nil, -1, err
	}

	return commentMsgs, num, nil
}

// AckInterMsgService 确认互动消息通知
func AckInterMsgService(msgId, msgType int) error {
	switch msgType {
	case constant.LikeMsgConstant:
		err := mysql.UpdateLikeRecordRead(msgId)
		if err != nil {
			zap.L().Error("AckInterMsgService() service.article.likeService.UpdateLikeRecordRead err=", zap.Error(err))
			return err
		}
	case constant.CommentMsgConstant:
		err := mysql.UpdateCommentRecordRead(msgId)
		if err != nil {
			zap.L().Error("AckInterMsgService() service.article.likeService.UpdateCommentRecordRead err=", zap.Error(err))
			return err
		}
	case constant.CollectMsgConstant:
		err := mysql.UpdateCollectRecordRead(msgId)
		if err != nil {
			zap.L().Error("AckInterMsgService() service.article.likeService.UpdateCollectRecordRead err=", zap.Error(err))
			return err
		}
	default:
		return myErr.DataFormatError()
	}
	return nil
}

// AckManagerMsgService 确认管理员消息
func AckManagerMsgService(username string) error {
	// 获取uid
	uid, err := mysql.GetIdByUsername(username)
	if err != nil {
		zap.L().Error("AckManagerMsgService() service.article.likeService.GetIdByUsername err=", zap.Error(err))
		return err
	}

	// 确认管理员消息
	err = mysql.UpdateManagerRecordRead(uid)
	if err != nil {
		zap.L().Error("AckManagerMsgService() service.article.likeService.UpdateManagerRecordRead err=", zap.Error(err))
		return err
	}
	return nil
}

// AckSystemMsgService 确认系统消息
func AckSystemMsgService(username string) error {
	// 获取uid
	uid, err := mysql.GetIdByUsername(username)
	if err != nil {
		zap.L().Error("AckSystemMsgService() service.article.likeService.GetIdByUsername err=", zap.Error(err))
		return err
	}

	// 确认管理员消息
	err = mysql.UpdateSystemRecordRead(uid)
	if err != nil {
		zap.L().Error("AckManagerMsgService() service.article.likeService.UpdateSystemRecordRead err=", zap.Error(err))
		return err
	}
	return nil
}

// PublishManagerMsgService 发布管理员通知
func PublishManagerMsgService(username, content, role string) error {
	// 权限验证
	if role != "college" {
		zap.L().Error("PublishManagerMsgService() service.article.likeService.role err=", zap.Error(myErr.OverstepCompetence))
		return myErr.OverstepCompetence
	}
	// 添加通知
	ids, err := mysql.QueryAllUserId()
	if err != nil {
		zap.L().Error("PublishManagerMsgService() service.article.likeService.QueryAllUserId err=", zap.Error(err))
		return err
	}
	for _, uid := range ids {
		err = mysql.AddManagerMsg(username, content, int(uid))
		if err != nil {
			zap.L().Error("PublishManagerMsgService() service.article.likeService.AddManagerMsg err=", zap.Error(err))
			return err
		}
	}

	notification, err := NotificationPush.BuildManagerNotification(username, content)
	if err != nil {
		zap.L().Error("PublishManagerMsgService() service.article.BuildManagerNotification.Transaction err=", zap.Error(err))
		return err
	}

	sse.SendSysNotification(*notification)

	return nil
}

// PublishSystemMsgService 发布系统通知
func PublishSystemMsgService(content, role, username string) error {
	// 权限验证
	if role != "superman" {
		return myErr.OverstepCompetence
	}

	// 添加通知

	//// 添加通知
	//ids, err := mysql.QueryAllUserId()
	//if err != nil {
	//	zap.L().Error("PublishSystemMsgService() service.article.likeService.QueryAllUserId err=", zap.Error(err))
	//	return err
	//}
	//err = mysql.DB.Transaction(func(tx *gorm.DB) error {
	//	for _, uid := range ids {
	//		err = mysql.AddSystemMsg(content, int(uid), tx, username)
	//		if err != nil {
	//			zap.L().Error("PublishSystemMsgService() service.article.likeService.AddSystemMsg err=", zap.Error(err))
	//			return err
	//		}
	//	}
	//	return nil
	//})
	//if err != nil {
	//	zap.L().Error("PublishSystemMsgService() service.article.likeService.Transaction err=", zap.Error(err))
	//	return err
	//}

	notification, err := NotificationPush.BuildSystemNotification(username, content)
	if err != nil {
		zap.L().Error("PublishSystemMsgService() service.article.BuildSystemNotification.Transaction err=", zap.Error(err))
		return err
	}

	sse.SendSysNotification(*notification)
	return nil
}
