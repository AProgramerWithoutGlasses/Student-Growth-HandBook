package comment

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"sort"
	"strconv"
	"studentGrow/dao/mysql"
	"studentGrow/dao/redis"
	"studentGrow/models/gorm_model"
	myErr "studentGrow/pkg/error"
	"studentGrow/service/article"
	timeUtil "studentGrow/utils/timeConverter"
)

// PostComment 发布评论
func PostComment(commentType, username, content string, id int) error {
	//类型comment_type:‘article’or‘comment’;id;comment_content;comment_username

	//获取用户id
	uid, err := mysql.SelectUserByUsername(username)
	fmt.Println(uid)
	if err != nil {
		zap.L().Error("PostComment() service.article.SelectUserByUsername err=", zap.Error(err))
		return err
	}

	var cid int
	//判断评论类型

	err = mysql.DB.Transaction(func(tx *gorm.DB) error {
		switch commentType {
		//给文章评论
		case "0":
			//向数据库插入评论数据
			cid, err = mysql.InsertIntoCommentsForArticle(content, id, uid, tx)
			if err != nil {
				zap.L().Error("PostComment() service.article.InsertIntoCommentsForArticle err=", zap.Error(err))
				return err
			}

		case "1":
			//向数据库插入评论数据
			cid, err = mysql.InsertIntoCommentsForComment(content, uid, id, tx)
			if err != nil {
				zap.L().Error("PostComment() service.article.InsertIntoCommentsForComment err=", zap.Error(err))
				return err
			}
		default:
			return myErr.DataFormatError()
		}
		return nil
	})
	if err != nil {
		zap.L().Error("PostComment() service.article.Transaction err=", zap.Error(err))
		return err
	}

	// 获取评论的文章id
	aid, err := mysql.QueryArticleIdByCommentId(cid)
	if err != nil {
		zap.L().Error("PostComment() service.article.QueryArticleIdByCommentId err=", zap.Error(err))
		return err
	}
	// 增加文章评论数
	num, err := mysql.QueryArticleCommentNum(aid)
	if err != nil {
		zap.L().Error("PostComment() service.article.QueryArticleCommentNum err=", zap.Error(err))
		return err
	}
	err = mysql.UpdateArticleCommentNum(aid, num+1, mysql.DB)
	if err != nil {
		zap.L().Error("PostComment() service.article.UpdateArticleCommentNum err=", zap.Error(err))
		return err
	}

	// 将评论数据加入redis
	redis.RDB.HSet("comment", strconv.Itoa(cid), 0)

	return nil
}

// GetLel1CommentsService 获取一级评论详情列表
func GetLel1CommentsService(aid, limit, page int, username, sortWay string) (gorm_model.Comments, error) {
	// 判断该用户是否被封禁或私密
	bl, err := article.QueryArticleStatusAndBanById(aid)
	if err != nil {
		zap.L().Error("GetLel1CommentsService() service.article.QueryArticleStatusAndBanById err=", zap.Error(err))
		return nil, err
	}
	if !bl {
		return nil, myErr.ErrNotFoundError
	}

	// 分页查询评论
	comments, err := mysql.QueryLevelOneComments(aid, limit, page)
	if err != nil {
		zap.L().Error("GetLel1CommentsService() service.article.QueryLevelOneComments err=", zap.Error(err))
		return nil, err
	}
	// 排序
	if sortWay == "hot" {
		sort.Sort(comments)
	}
	// 判断是否点赞, 并计算其子评论数量, 计算评论时间
	for i := 0; i < len(comments); i++ {
		liked, err := redis.IsUserLiked(strconv.Itoa(int(comments[i].ID)), username, 1)
		if err != nil {
			zap.L().Error("GetLel1CommentsService() service.article.IsUserLiked err=", zap.Error(err))
			return nil, err
		}
		comments[i].IsLike = liked

		num, err := mysql.QuerySonCommentNum(int(comments[i].ID))
		if err != nil {
			zap.L().Error("GetLel1CommentsService() service.article.QuerySonCommentNum err=", zap.Error(err))
			return nil, err
		}
		comments[i].ReplyCount = num

		comments[i].Time = timeUtil.IntervalConversion(comments[i].CreatedAt)
	}
	return comments, nil
}

// GetLelSonCommentListService 获取子评论列表
func GetLelSonCommentListService(cid, limit, page int, username string) ([]gorm_model.Comment, error) {
	// 获取文章对应的评论
	comments, err := mysql.QueryLevelSonComments(cid, limit, page)
	if err != nil {
		zap.L().Error("GetLelSonCommentListService() service.article.QueryLevelSonComments err=", zap.Error(err))
		return nil, err
	}

	// 该用户是否点赞, 计算评论时间
	for i := 0; i < len(comments); i++ {
		liked, err := redis.IsUserLiked(strconv.Itoa(int(comments[i].ID)), username, 1)
		if err != nil {
			zap.L().Error("GetLelSonCommentListService() service.article.IsUserLiked err=", zap.Error(err))
			return nil, err
		}
		comments[i].IsLike = liked

		comments[i].Time = timeUtil.IntervalConversion(comments[i].CreatedAt)
	}

	return comments, nil
}
