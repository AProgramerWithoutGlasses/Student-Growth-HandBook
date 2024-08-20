package article

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"strconv"
	"studentGrow/dao/mysql"
	"studentGrow/dao/redis"
	"studentGrow/models/constant"
	"studentGrow/models/nzx_model"
)

/*
redis
*/

// Like 点赞
func Like(objId, username string, likeType int) error {
	// 将用户添加到点赞列表
	err := redis.AddUserToLikeSet(objId, username, likeType)
	if err != nil {
		zap.L().Error("Like() service.article.likeService.AddUserToLikeSet err=", zap.Error(err))
		return err
	}
	// 获取对象点赞数
	likes, err := redis.GetObjLikes(objId, likeType)
	if err != nil {
		zap.L().Error("Like() service.article.likeService.GetObjLikes err=", zap.Error(err))
		return err
	}
	// 设置对象点赞数
	if likes >= 0 {
		err = redis.SetObjLikes(objId, likes+1, likeType)
		if err != nil {
			zap.L().Error("Like() service.article.likeService.SetObjLikes err=", zap.Error(err))
			return err
		}
	}
	if err != nil {
		zap.L().Error("Like() service.article.TxPipelined.GetObjLikes err=", zap.Error(err))
		return err
	}

	id, err := strconv.Atoi(objId)
	if err != nil {
		zap.L().Error("Like() service.article.likeService.Atoi err=", zap.Error(err))
		return err
	}
	//
	//// 增加文章或评论的点赞数量
	//switch likeType {
	//case constant.ArticleInteractionConstant:
	//	num, err := redis.GetObjLikes(strconv.Itoa(id), constant.ArticleInteractionConstant)
	//	if err != nil {
	//		zap.L().Error("Like() service.article.likeService.QueryArticleLikeNum err=", zap.Error(err))
	//		return err
	//	}
	//	err = redis.SetObjLikes(strconv.Itoa(id), num+1, constant.ArticleInteractionConstant)
	//	if err != nil {
	//		zap.L().Error("Like() service.article.likeService.UpdateArticleLikeNum err=", zap.Error(err))
	//		return err
	//	}
	//case 1:
	//	num, err := redis.GetObjLikes(strconv.Itoa(id), constant.CommentInteractionConstant)
	//	if err != nil {
	//		zap.L().Error("Like() service.article.likeService.v err=", zap.Error(err))
	//		return err
	//	}
	//	err = redis.SetObjLikes(strconv.Itoa(id), num+1, constant.CommentInteractionConstant)
	//	if err != nil {
	//		zap.L().Error("Like() service.article.likeService.UpdateCommentLikeNum err=", zap.Error(err))
	//		return err
	//	}
	//}
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
	// 判断是否存在点赞
	ok, err := redis.IsUserLiked(objId, username, likeType)
	if err != nil {
		zap.L().Error("CancelLike() service.article.likeService.IsUserLiked err=", zap.Error(err))
		return err
	}
	if ok {
		// 移除点赞列表
		err = redis.RemoveUserFromLikeSet(objId, username, likeType)
		if err != nil {
			zap.L().Error("CancelLike() service.article.likeService.RemoveUserFromLikeSet err=", zap.Error(err))
			return err
		}
		// 获取对象点赞数
		likes, err := redis.GetObjLikes(objId, likeType)
		if err != nil {
			zap.L().Error("CancelLike() service.article.likeService.SetObjLikes err=", zap.Error(err))
			return err
		}
		if likes > 0 {
			// 设置对象点赞数
			err = redis.SetObjLikes(objId, likes-1, likeType)
			if err != nil {
				zap.L().Error("CancelLike() service.article.likeService.SetObjLikes err=", zap.Error(err))
				return err
			}
		}
		if err != nil {
			zap.L().Error("CancelLike() service.article.likeService.TxPipelined err=", zap.Error(err))
			return err
		}
	}

	id, err := strconv.Atoi(objId)

	// 减少文章或评论的点赞数量
	//switch likeType {
	//case 0:
	//	num, err := redis.GetObjLikes(strconv.Itoa(id), constant.ArticleInteractionConstant)
	//	if err != nil {
	//		zap.L().Error("CancelLike() service.article.likeService.GetObjLikes err=", zap.Error(err))
	//		return err
	//	}
	//	err = redis.SetObjLikes(strconv.Itoa(id), num-1, constant.ArticleInteractionConstant)
	//	if err != nil {
	//		zap.L().Error("CancelLike() service.article.likeService.SetObjLikes err=", zap.Error(err))
	//		return err
	//	}
	//case 1:
	//	num, err := redis.GetObjLikes(strconv.Itoa(id), constant.CommentInteractionConstant)
	//	if err != nil {
	//		zap.L().Error("CancelLike() service.article.likeService.GetObjLikes err=", zap.Error(err))
	//		return err
	//	}
	//	err = redis.SetObjLikes(strconv.Itoa(id), num-1, constant.CommentInteractionConstant)
	//	if err != nil {
	//		zap.L().Error("CancelLike() service.article.likeService.SetObjLikes err=", zap.Error(err))
	//		return err
	//	}
	//}

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
	//slice, err := redis.GetObjLikedUsers(objId, likeType)
	//if err != nil {
	//	zap.L().Error("LikeObjOrNot() service.article.likeService.GetObjLikedUsers err=", zap.Error(err))
	//	return err
	//}
	//likeUsers := make(map[string]struct{})
	//for _, s := range slice {
	//	likeUsers[s] = struct{}{}
	//}
	//likes, err := redis.GetObjLikes(objId, 0)
	//if err != nil {
	//	return err
	//}
	//fmt.Println("likenum:", likes)
	//_, ok := likeUsers[username]
	//若存在该用户，则取消点赞

	ok, err := redis.IsUserLiked(objId, username, constant.ArticleInteractionConstant)
	if err != nil {
		return err
	}
	fmt.Println(ok)
	if ok {
		err = CancelLike(objId, username, likeType)
		fmt.Println("cancel")
		if err != nil {
			zap.L().Error("LikeObjOrNot() service.article.likeService.CancelLike err=", zap.Error(err))
			return err
		}
	} else {
		//反之，点赞
		err = Like(objId, username, likeType)
		if err != nil {
			fmt.Println("like")
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
	err = mysql.DB.Transaction(func(tx *gorm.DB) error {
		// 插入点赞记录
		err = mysql.InsertLikeRecord(objId, likeType, userId, tx)
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
		err = mysql.UpdateLikeNum(objId, likeType, num+1, tx)
		if err != nil {
			zap.L().Error("CancelLikeToMysql() service.article.likeService.UpdateLikeNum err=", zap.Error(err))
			return err
		}

		// 增加文章或评论的点赞数量
		//switch likeType {
		//case 0:
		//	num, err := mysql.QueryArticleLikeNum(objId)
		//	if err != nil {
		//		zap.L().Error("CancelLike() service.article.likeService.QueryArticleLikeNum err=", zap.Error(err))
		//		return err
		//	}
		//	err = mysql.UpdateArticleLikeNum(objId, num+1, tx)
		//	if err != nil {
		//		zap.L().Error("CancelLike() service.article.likeService.UpdateArticleLikeNum err=", zap.Error(err))
		//		return err
		//	}
		//case 1:
		//	num, err := mysql.QueryCommentLikeNum(objId)
		//	if err != nil {
		//		zap.L().Error("CancelLike() service.article.likeService.QueryCommentLikeNum err=", zap.Error(err))
		//		return err
		//	}
		//	err = mysql.UpdateCommentLikeNum(objId, num-1, tx)
		//	if err != nil {
		//		zap.L().Error("CancelLike() service.article.likeService.UpdateCommentLikeNum err=", zap.Error(err))
		//		return err
		//	}
		//}
		return nil
	})
	if err != nil {
		zap.L().Error("CancelLikeToMysql() service.article.likeService.Transaction err=", zap.Error(err))
		return err
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

	err = mysql.DB.Transaction(func(db *gorm.DB) error {
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
		err = mysql.UpdateLikeNum(objId, likeType, num-1, db)
		if err != nil {
			zap.L().Error("CancelLikeToMysql() service.article.likeService.UpdateLikeNum err=", zap.Error(err))
			return err
		}

		// 减少文章或评论的点赞数量

		//switch likeType {
		//case 0:
		//	num, err := mysql.QueryArticleLikeNum(objId)
		//	if err != nil {
		//		zap.L().Error("CancelLike() service.article.likeService.QueryArticleLikeNum err=", zap.Error(err))
		//		return err
		//	}
		//	err = mysql.UpdateArticleLikeNum(objId, num-1, db)
		//	if err != nil {
		//		zap.L().Error("CancelLike() service.article.likeService.UpdateArticleLikeNum err=", zap.Error(err))
		//		return err
		//	}
		//case 1:
		//	num, err := mysql.QueryCommentLikeNum(objId)
		//	if err != nil {
		//		zap.L().Error("CancelLike() service.article.likeService.QueryCommentLikeNum err=", zap.Error(err))
		//		return err
		//	}
		//	err = mysql.UpdateCommentLikeNum(objId, num-1, db)
		//	if err != nil {
		//		zap.L().Error("CancelLike() service.article.likeService.UpdateCommentLikeNum err=", zap.Error(err))
		//		return err
		//	}
		//}
		return nil
	})
	if err != nil {
		zap.L().Error("CancelLike() service.article.likeService.Transaction err=", zap.Error(err))
		return err
	}

	return nil
}
