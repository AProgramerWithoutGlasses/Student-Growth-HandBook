package article

import (
	"go.uber.org/zap"
	"strconv"
	"studentGrow/dao/mysql"
	"studentGrow/dao/redis"
	"studentGrow/models/nzx_model"
)

/*
redis
*/

// Like 点赞
func Like(objId, username string, likeType int) error {
	err := redis.AddUserToLikeSet(redis.List[likeType]+objId, username, likeType)
	if err != nil {
		zap.L().Error("Like() service.article.likeService.AddUserToLikeSet err=", zap.Error(err))
		return err
	}
	likes, err := redis.GetObjLikes(objId, likeType)
	if err != nil {
		zap.L().Error("Like() service.article.likeService.GetObjLikes err=", zap.Error(err))
		return err
	}
	if likes >= 0 {
		err = redis.SetObjLikes(objId, likes+1, likeType)
		if err != nil {
			zap.L().Error("Like() service.article.likeService.SetObjLikes err=", zap.Error(err))
			return err
		}
	}

	id, err := strconv.Atoi(objId)

	// 增加文章或评论的点赞数量
	switch likeType {
	case 0:
		num, err := mysql.QueryArticleLikeNum(id)
		if err != nil {
			zap.L().Error("Like() service.article.likeService.QueryArticleLikeNum err=", zap.Error(err))
			return err
		}
		err = mysql.UpdateArticleLikeNum(id, num+1)
		if err != nil {
			zap.L().Error("Like() service.article.likeService.UpdateArticleLikeNum err=", zap.Error(err))
			return err
		}
	case 1:
		num, err := mysql.QueryCommentLikeNum(id)
		if err != nil {
			zap.L().Error("Like() service.article.likeService.v err=", zap.Error(err))
			return err
		}
		err = mysql.UpdateCommentLikeNum(id, num+1)
		if err != nil {
			zap.L().Error("Like() service.article.likeService.UpdateCommentLikeNum err=", zap.Error(err))
			return err
		}
	}
	// 写入通道

	switch likeType {
	case 0:
		ArticleLikeChan <- nzx_model.RedisLikeArticleData{Aid: id, Username: username, Operator: "like"}
	case 1:
		CommentLikeChan <- nzx_model.RedisLikeCommentData{Cid: id, Username: username, Operator: "like"}
	}

	return nil
}

// CancelLike 取消点赞
func CancelLike(objId, username string, likeType int) error {

	ok, err := redis.IsUserLiked(objId, username, likeType)
	if err != nil {
		zap.L().Error("CancelLike() service.article.likeService.IsUserLiked err=", zap.Error(err))
		return err
	}
	if ok {
		err = redis.RemoveUserFromLikeSet(objId, username, likeType)
		if err != nil {
			zap.L().Error("CancelLike() service.article.likeService.RemoveUserFromLikeSet err=", zap.Error(err))
			return err
		}
		likes, err := redis.GetObjLikes(objId, likeType)
		if err != nil {
			zap.L().Error("CancelLike() service.article.likeService.SetObjLikes err=", zap.Error(err))
			return err
		}
		if likes > 0 {
			err = redis.SetObjLikes(objId, likes-1, likeType)
			if err != nil {
				zap.L().Error("CancelLike() service.article.likeService.SetObjLikes err=", zap.Error(err))
				return err
			}
		}
	}

	id, err := strconv.Atoi(objId)

	// 写入通道
	switch likeType {
	case 0:
		ArticleLikeChan <- nzx_model.RedisLikeArticleData{Aid: id, Username: username, Operator: "cancel_like"}
	case 1:
		CommentLikeChan <- nzx_model.RedisLikeCommentData{Cid: id, Username: username, Operator: "cancel_like"}
	}

	return nil
}

// LikeObjOrNot 检查是否点赞并点赞
func LikeObjOrNot(objId, username string, likeType int) error {
	//获取当前点赞文章列表
	slice, err := redis.GetObjLikedUsers(objId, likeType)
	if err != nil {
		zap.L().Error("LikeObjOrNot() service.article.likeService.GetObjLikedUsers err=", zap.Error(err))
		return err
	}
	likeUsers := make(map[string]struct{})
	for _, s := range slice {
		likeUsers[s] = struct{}{}
	}
	//若存在该用户，则取消点赞
	_, ok := likeUsers[username]
	if len(likeUsers) > 0 && ok {
		err = CancelLike(objId, username, likeType)
		if err != nil {
			zap.L().Error("LikeObjOrNot() service.article.likeService.CancelLike err=", zap.Error(err))
			return err
		}
	} else {
		//反之，点赞
		err = Like(objId, username, likeType)
		if err != nil {
			zap.L().Error("LikeObjOrNot() service.article.likeService.Like err=", zap.Error(err))
			return err
		}
	}
	return nil
}

/*
mysql
*/

// LikeToMysql 点赞
func LikeToMysql(objId, likeType int, username string) error {

	userId, err := mysql.GetIdByUsername(username)
	if err != nil {
		zap.L().Error("CancelLikeToMysql() service.article.likeService.GetIdByUsername err=", zap.Error(err))
		return err
	}
	// 插入点赞记录
	err = mysql.InsertLikeRecord(objId, likeType, userId)
	if err != nil {
		zap.L().Error("CancelLikeToMysql() service.article.likeService.InsertLikeRecord err=", zap.Error(err))
		return err
	}

	// 更新点赞数
	num, err := mysql.QueryLikeNum(objId, likeType)
	if err != nil {
		zap.L().Error("CancelLikeToMysql() service.article.likeService.QueryLikeNum err=", zap.Error(err))
		return err
	}
	err = mysql.UpdateLikeNum(objId, likeType, num+1)
	if err != nil {
		zap.L().Error("CancelLikeToMysql() service.article.likeService.UpdateLikeNum err=", zap.Error(err))
		return err
	}

	// 增加文章或评论的点赞数量
	switch likeType {
	case 0:
		num, err := mysql.QueryArticleLikeNum(objId)
		if err != nil {
			zap.L().Error("CancelLike() service.article.likeService.QueryArticleLikeNum err=", zap.Error(err))
			return err
		}
		err = mysql.UpdateArticleLikeNum(objId, num-1)
		if err != nil {
			zap.L().Error("CancelLike() service.article.likeService.UpdateArticleLikeNum err=", zap.Error(err))
			return err
		}
	case 1:
		num, err := mysql.QueryCommentLikeNum(objId)
		if err != nil {
			zap.L().Error("CancelLike() service.article.likeService.QueryCommentLikeNum err=", zap.Error(err))
			return err
		}
		err = mysql.UpdateCommentLikeNum(objId, num-1)
		if err != nil {
			zap.L().Error("CancelLike() service.article.likeService.UpdateCommentLikeNum err=", zap.Error(err))
			return err
		}
	}

	return nil
}

// CancelLikeToMysql 取消点赞
func CancelLikeToMysql(objId, likeType int, username string) error {
	userId, err := mysql.GetIdByUsername(username)
	if err != nil {
		zap.L().Error("CancelLikeToMysql() service.article.likeService.GetIdByUsername err=", zap.Error(err))
		return err
	}

	//删除点赞记录
	err = mysql.DeleteLikeRecord(objId, likeType, userId)
	if err != nil {
		zap.L().Error("CancelLikeToMysql() service.article.likeService.DeleteLikeRecord err=", zap.Error(err))
		return err
	}

	// 更新点赞数
	num, err := mysql.QueryLikeNum(objId, likeType)
	if err != nil {
		zap.L().Error("CancelLikeToMysql() service.article.likeService.QueryLikeNum err=", zap.Error(err))
		return err
	}
	err = mysql.UpdateLikeNum(objId, likeType, num-1)
	if err != nil {
		zap.L().Error("CancelLikeToMysql() service.article.likeService.UpdateLikeNum err=", zap.Error(err))
		return err
	}

	// 减少文章或评论的点赞数量

	switch likeType {
	case 0:
		num, err := mysql.QueryArticleLikeNum(objId)
		if err != nil {
			zap.L().Error("CancelLike() service.article.likeService.QueryArticleLikeNum err=", zap.Error(err))
			return err
		}
		err = mysql.UpdateArticleLikeNum(objId, num-1)
		if err != nil {
			zap.L().Error("CancelLike() service.article.likeService.UpdateArticleLikeNum err=", zap.Error(err))
			return err
		}
	case 1:
		num, err := mysql.QueryCommentLikeNum(objId)
		if err != nil {
			zap.L().Error("CancelLike() service.article.likeService.QueryCommentLikeNum err=", zap.Error(err))
			return err
		}
		err = mysql.UpdateCommentLikeNum(objId, num-1)
		if err != nil {
			zap.L().Error("CancelLike() service.article.likeService.UpdateCommentLikeNum err=", zap.Error(err))
			return err
		}
	}

	return nil
}
