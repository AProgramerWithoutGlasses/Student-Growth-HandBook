package mysql

import (
	"fmt"
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
		articleLike := gorm_model.UserLikeRecord{ArticleID: uint(objId), UserID: uint(uid), Type: likeType}
		if err := DB.Model(gorm_model.UserLikeRecord{}).Create(&articleLike).Error; err != nil {
			fmt.Println("InsertLikeRecord() dao.mysql.mysql_like")
			return err
		}
	case 1:
		commentLike := gorm_model.UserLikeRecord{CommentID: uint(objId), UserID: uint(uid), Type: likeType}
		if err := DB.Model(gorm_model.UserLikeRecord{}).Create(&commentLike).Error; err != nil {
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
		if err := DB.Where("article_id = ? and user_id = ?", objId, uid).Delete(&gorm_model.UserLikeRecord{}).Error; err != nil {
			fmt.Println("DeleteLikeRecord() dao.mysql.mysql_like")
			return err
		}
	case 1:
		if err := DB.Where("comment_id = ? and user_id = ?", objId, uid).Delete(&gorm_model.UserLikeRecord{}).Error; err != nil {
			fmt.Println("DeleteLikeRecord() dao.mysql.mysql_like")
			return err
		}
	default:
		return myErr.DataFormatError()
	}
	return nil
}
