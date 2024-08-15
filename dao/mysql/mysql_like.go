package mysql

import (
	"fmt"
	"go.uber.org/zap"
	"studentGrow/models/gorm_model"
	myErr "studentGrow/pkg/error"
)

// UpdateLikeNum 修改点赞数量
func UpdateLikeNum(objId, likeType, likeNum int) error {
	// 修改文章点赞
	switch likeType {
	case 0:
		if err := DB.Model(gorm_model.Article{}).Where("id = ?", objId).Update("like_amount", likeNum).Error; err != nil {
			fmt.Println("UpdateLikeNum() dao.mysql.mysql_like")
			return err
		}
	case 1:
		if err := DB.Model(gorm_model.Comment{}).Where("id = ?", objId).Update("like_amount", likeNum).Error; err != nil {
			fmt.Println("UpdateLikeNum() dao.mysql.mysql_like")
			return err
		}
	}
	return nil
}

// QueryLikeNum 查询点赞数量
func QueryLikeNum(objId, likeType int) (int, error) {
	article := gorm_model.Article{}
	comment := gorm_model.Comment{}
	switch likeType {
	case 0:
		if err := DB.Model(gorm_model.Article{}).Where("id = ?", objId).First(&article).Error; err != nil {
			fmt.Println("QueryLikeNum() dao.mysql.mysql_like")
			return -1, err
		}
		return article.LikeAmount, nil
	case 1:
		if err := DB.Model(gorm_model.Comment{}).Where("id = ?", objId).First(&comment).Error; err != nil {
			fmt.Println("QueryLikeNum() dao.mysql.mysql_like")
			return -1, err
		}
		return comment.LikeAmount, nil
	default:
		return -1, myErr.DataFormatError()
	}
}

// InsertLikeRecord 插入点赞记录
func InsertLikeRecord(objId, likeType int, uid int) error {

	switch likeType {
	case 0:
		articleLike := gorm_model.UserArticleLikeRecord{ArticleID: uint(objId), UserID: uint(uid)}
		if err := DB.Model(gorm_model.UserArticleLikeRecord{}).Create(&articleLike).Error; err != nil {
			fmt.Println("InsertLikeRecord() dao.mysql.mysql_like")
			return err
		}
	case 1:
		commentLike := gorm_model.UserCommentLikeRecord{CommentID: uint(objId), UserID: uint(uid)}
		if err := DB.Model(gorm_model.UserCommentLikeRecord{}).Create(&commentLike).Error; err != nil {
			fmt.Println("InsertLikeRecord() dao.mysql.mysql_like")
			return err
		}
	default:
		return myErr.DataFormatError()
	}
	return nil
}

// DeleteLikeRecord 删除点赞记录
func DeleteLikeRecord(objId, likeType, uid int) error {
	switch likeType {
	case 0:
		if err := DB.Where("article_id = ? and user_id = ?", objId, uid).Delete(&gorm_model.UserArticleLikeRecord{}).Error; err != nil {
			fmt.Println("DeleteLikeRecord() dao.mysql.mysql_like")
			return err
		}
	case 1:
		if err := DB.Where("article_id = ? and user_id = ?", objId, uid).Delete(&gorm_model.UserCommentLikeRecord{}).Error; err != nil {
			fmt.Println("DeleteLikeRecord() dao.mysql.mysql_like")
			return err
		}
	default:
		return myErr.DataFormatError()
	}
	return nil
}

// QueryLikeRecordByUserArticle 通过uid分页查询其文章的点赞记录
func QueryLikeRecordByUserArticle(uid, page, limit int) ([]gorm_model.Article, error) {

	var articles []gorm_model.Article

	if err := DB.Preload("User", "id = ?", uid).Preload("ArticleLikes").Preload("ArticleLikes.User").
		Where("user_id = ?", uid).
		Limit(limit).Offset((page - 1) * limit).Order("created_at desc").
		Find(&articles).Error; err != nil {
		zap.L().Error("QueryLikeRecordByUserArticle() dao.mysql.mysql_like.Find err=", zap.Error(err))
		return nil, err
	}

	if len(articles) == 0 {
		zap.L().Error("QueryLikeRecordByUserArticle() dao.mysql.mysql_like err=", zap.Error(myErr.NotFoundError()))
		return nil, myErr.NotFoundError()
	}

	return articles, nil
}

// QueryLikeRecordNumByUserArticle 通过uid查询其文章的未读点赞记录数量
func QueryLikeRecordNumByUserArticle(uid int) (int, error) {
	var count int64

	if err := DB.Preload("ArticleLikes", "is_read = ?", false).Model(gorm_model.Article{}).Where("user_id = ?", uid).Count(&count).Error; err != nil {
		zap.L().Error("QueryLikeRecordNumByUserArticle() dao.mysql.mysql_like.Count err=", zap.Error(err))
		return -1, err
	}
	return int(count), nil
}

// QueryLikeRecordByUserComment 通过uid分页查询其评论的点赞记录
func QueryLikeRecordByUserComment(uid, page, limit int) ([]gorm_model.Comment, error) {

	var comments []gorm_model.Comment

	if err := DB.Preload("User", "id = ?", uid).Preload("CommentLikes").Preload("CommentLikes.User").
		Where("user_id = ?", uid).
		Limit(limit).Offset((page - 1) * limit).Order("created_at desc").
		Find(&comments).Error; err != nil {
		zap.L().Error("QueryLikeRecordByUserComment() dao.mysql.mysql_like.Find err=", zap.Error(err))
		return nil, err
	}

	if len(comments) == 0 {
		zap.L().Error("QueryLikeRecordByUserComment() dao.mysql.mysql_like err=", zap.Error(myErr.NotFoundError()))
		return nil, myErr.NotFoundError()
	}

	return comments, nil
}

// QueryLikeRecordNumByUserComment 通过uid查询其评论的未读点赞记录数量
func QueryLikeRecordNumByUserComment(uid int) (int, error) {
	var count int64

	if err := DB.Preload("CommentLikes", "is_read = ?", false).Model(gorm_model.Comment{}).Where("user_id = ?", uid).Count(&count).Error; err != nil {
		zap.L().Error("QueryLikeRecordNumByUserComment() dao.mysql.mysql_like.Count err=", zap.Error(err))
		return -1, err
	}
	return int(count), nil
}
