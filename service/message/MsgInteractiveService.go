package message

import (
	"go.uber.org/zap"
	"studentGrow/dao/mysql"
	"studentGrow/models/gorm_model"
	"studentGrow/models/nzx_model"
	"studentGrow/utils/timeConverter"
)

// GetSystemMsgService 获取系统消息通知
func GetSystemMsgService(limit, page int, username string) ([]gorm_model.MsgRecord, int, error) {
	msg, err := mysql.QuerySystemMsg(page, limit, username)
	if err != nil {
		zap.L().Error("GetSystemMsgService() service.message.QuerySystemMsg", zap.Error(err))
		return nil, -1, err
	}

	// 查询未读消息条数
	count, err := mysql.QueryUnreadSystemMsg(username)
	if err != nil {
		zap.L().Error("GetSystemMsgService() service.message.QueryUnreadSystemMsg", zap.Error(err))
		return nil, -1, err
	}

	return msg, count, nil
}

// GetManagerMsgService 获取管理员消息
func GetManagerMsgService(limit, page int, username string) ([]gorm_model.MsgRecord, int, error) {
	msg, err := mysql.QueryManagerMsg(page, limit, username)
	if err != nil {
		zap.L().Error("GetManagerMsgService() service.message.QueryManagerMsg", zap.Error(err))
		return nil, -1, err
	}

	// 查询未读消息条数
	count, err := mysql.QueryUnreadManagerMsg(username)
	if err != nil {
		zap.L().Error("GetManagerMsgService() service.message.QueryUnreadManagerMsg", zap.Error(err))
		return nil, -1, err
	}

	return msg, count, nil
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

	var list []nzx_model.Out

	for _, like := range likes {
		// 判断文章点赞还是评论点赞
		usernameL := like.Article.User.Username
		name := like.Article.User.Name
		content := like.Article.Content
		userHeadshot := like.Article.User.HeadShot
		likeType := 0
		articleId := like.ArticleID
		if like.Type == 1 {
			username = like.Comment.User.Username
			name = like.Comment.User.Name
			content = like.Comment.Content
			userHeadshot = like.Comment.User.HeadShot
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

	var list []map[string]any

	for _, collect := range articleCollects {
		list = append(list, map[string]any{
			"username":        collect.User.Username,
			"name":            collect.User.Name,
			"article_content": collect.Article.Content,
			"user_headshot":   collect.User.HeadShot,
			"post_time":       timeConverter.IntervalConversion(collect.CreatedAt),
			"is_read":         collect.IsRead,
		})
	}

	return list, collectNum, nil

}

// GetCommentMsgService 获取评论消息
func GetCommentMsgService(username string, page, limit int) (nzx_model.CommentMsgs, error) {
	// 获取uid
	uid, err := mysql.GetIdByUsername(username)
	if err != nil {
		zap.L().Error("GetCommentMsgService() service.article.likeService.GetIdByUsername err=", zap.Error(err))
		return nil, err
	}

	// 获取所有评论及回复
	comments, err := mysql.QueryCommentRecordByUserArticles(uid, page, limit)
	if err != nil {
		zap.L().Error("GetCommentMsgService() service.article.likeService.QueryCommentRecordByUserArticles err=", zap.Error(err))
		return nil, err
	}

	var commentMsgs nzx_model.CommentMsgs

	for _, comment := range comments {
		// 判断其为文章评论还是评论回复
		content := comment.Article.Content
		commentType := 0
		if comment.Pid != 0 {
			content = comment.Content
			commentType = 1
		}

		commentMsgs = append(commentMsgs, nzx_model.CommentMsg{
			Username:     comment.Article.User.Username,
			Name:         comment.Article.User.Name,
			Content:      content,
			UserHeadshot: comment.Article.User.HeadShot,
			PostTime:     timeConverter.IntervalConversion(comment.CreatedAt),
			IsRead:       comment.IsRead,
			Type:         commentType,
			ArticleId:    comment.ArticleID,
		})
	}

	return commentMsgs, nil

	//commentMsgs := make([]nzx_model.CommentMsgs, page+5) // 每个元素为一页，一页limit条记录
	//itemCount := 0                                       // 记录项数
	//isLimit := false                                     // 是否超出记录限制

	//// 根据用户的一级评论获取评论回复
	//for _, lel1 := range lel1Comments {
	//	// 满足客户端传来的记录总数据，则退出循环
	//	if isLimit {
	//		break
	//	}
	//	comments, err := mysql.QueryCommentRecordByUserComments(int(lel1.ID))
	//	if err != nil {
	//		zap.L().Error("GetCommentMsgService() service.article.likeService.QueryCommentRecordByUserComments err=", zap.Error(err))
	//		return
	//	}
	//	// 遍历回复记录
	//	for _, comment := range comments {
	//		itemCount++
	//		// 若回复项数超出客户端传来的页数，则退出循环
	//		if itemCount/limit > page-1 {
	//			isLimit = true
	//			break
	//		}
	//		commentMsgs[itemCount/limit] = append(commentMsgs[itemCount/limit], nzx_model.CommentMsg{
	//			Username:     comment.User.Username,
	//			Name:         comment.User.Name,
	//			Content:      lel1.Content,
	//			UserHeadshot: comment.User.HeadShot,
	//			PostTime:     timeConverter.IntervalConversion(comment.CreatedAt),
	//			IsRead:       comment.IsRead,
	//			Type:         1,
	//			ArticleId:    lel1.ArticleID,
	//			CreatedAt:    comment.CreatedAt,
	//		})
	//	}
	//}
	//// 如果查询到的记录总数不满足客户端要求
	//if !isLimit {
	//	// 若第page页没有记录
	//	if len(commentMsgs[page-1]) == 0 {
	//		return nil, myErr.NotFoundError()
	//	} else {
	//		// 若page页还有记录
	//
	//	}
	//}

}
