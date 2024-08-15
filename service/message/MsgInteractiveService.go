package message

import (
	"go.uber.org/zap"
	"sort"
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
func GetArticleAndCommentLikedMsgService(username string, page, limit int) (nzx_model.Outs, int, error) {
	// 获取uid
	uid, err := mysql.GetIdByUsername(username)
	if err != nil {
		zap.L().Error("GetArticleAndCommentLikedMsgService() service.article.likeService.GetIdByUsername err=", zap.Error(err))
		return nil, -1, err
	}

	// 获取文章点赞列表
	articles, err := mysql.QueryLikeRecordByUserArticle(uid, page, limit)
	if err != nil {
		zap.L().Error("GetArticleAndCommentLikedMsgService() service.article.likeService.QueryLikeRecordByUserArticle err=", zap.Error(err))
		return nil, -1, err
	}

	// 获取评论点赞列表
	comments, err := mysql.QueryLikeRecordByUserComment(uid, page, limit)
	if err != nil {
		zap.L().Error("GetArticleAndCommentLikedMsgService() service.article.likeService.QueryLikeRecordByUserComment err=", zap.Error(err))
		return nil, -1, err
	}

	// 获取文章点赞未读消息总数
	articleNum, err := mysql.QueryLikeRecordNumByUserArticle(uid)
	if err != nil {
		zap.L().Error("GetArticleAndCommentLikedMsgService() service.article.likeService.QueryLikeRecordNumByUserArticle err=", zap.Error(err))
		return nil, -1, err
	}
	// 获取评论点赞未读消息总数
	commentNum, err := mysql.QueryLikeRecordNumByUserComment(uid)
	if err != nil {
		zap.L().Error("GetArticleAndCommentLikedMsgService() service.article.likeService.QueryLikeRecordNumByUserArticle err=", zap.Error(err))
		return nil, -1, err
	}

	var list nzx_model.Outs

	for _, article := range articles {
		for _, like := range article.ArticleLikes {
			list = append(list, nzx_model.Out{
				Username:     like.User.Username,
				Name:         like.User.Name,
				Content:      article.Content,
				UserHeadshot: like.User.HeadShot,
				PostTime:     timeConverter.IntervalConversion(like.CreatedAt),
				IsRead:       like.IsRead,
				CreatedAt:    like.CreatedAt,
			})
		}
	}

	for _, comment := range comments {
		for _, like := range comment.CommentLikes {
			list = append(list, nzx_model.Out{
				Username:     like.User.Username,
				Name:         like.User.Name,
				Content:      comment.Content,
				UserHeadshot: like.User.HeadShot,
				PostTime:     timeConverter.IntervalConversion(like.CreatedAt),
				IsRead:       like.IsRead,
				CreatedAt:    like.CreatedAt,
			})
		}
	}

	sort.Sort(list)

	return list, articleNum + commentNum, nil
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
	articles, err := mysql.QueryCollectRecordByUserArticles(uid, page, limit)
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
	for _, article := range articles {
		for _, collect := range article.Collects {
			list = append(list, map[string]any{
				"username":        collect.User.Username,
				"name":            collect.User.Name,
				"article_content": article.Content,
				"user_headshot":   collect.User.HeadShot,
				"post_time":       timeConverter.IntervalConversion(collect.CreatedAt),
				"is_read":         collect.IsRead,
			})
		}
	}

	return list, collectNum, nil

}
