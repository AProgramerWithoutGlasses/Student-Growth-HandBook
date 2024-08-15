package mysql

import (
	"fmt"
	"go.uber.org/zap"
	"studentGrow/models/gorm_model"
	myErr "studentGrow/pkg/error"
)

// UpdateCollectNum 修改收藏数量
func UpdateCollectNum(aid, collectNum int) error {
	if err := DB.Model(&gorm_model.Article{}).Where("id = ?", aid).Update("collect_amount", collectNum).Error; err != nil {
		fmt.Println("UpdateCollectNum() dao.mysql.mysql_collect")
		return err
	}

	return nil
}

// QueryCollectNum 查询收藏数量
func QueryCollectNum(aid int) (int, error) {
	article := gorm_model.Article{}
	if err := DB.Model(&gorm_model.Article{}).Where("id = ?", aid).First(&article).Error; err != nil {
		fmt.Println("QueryCollectNum() dao.mysql.mysql_collect")
		return -1, err
	}

	return article.CollectAmount, nil
}

// InsertCollectRecord 插入收藏记录
func InsertCollectRecord(aid, uid int) error {
	if err := DB.Model(&gorm_model.UserCollectRecord{}).Create(&gorm_model.UserCollectRecord{UserID: uint(uid), ArticleID: uint(aid)}).Error; err != nil {
		fmt.Println("QueryCollectNum() dao.mysql.mysql_collect")
		return err
	}
	return nil
}

// DeleteCollectRecord 删除收藏记录
func DeleteCollectRecord(aid, uid int) error {
	if err := DB.Where("article_id = ? and user_id = ?", aid, uid).Delete(&gorm_model.UserCollectRecord{}).Error; err != nil {
		fmt.Println("QueryCollectNum() dao.mysql.mysql_collect")
		return err
	}
	return nil
}

// QueryCollectRecordByUserArticles 通过用户的所有文章查找其收藏记录(该用户的文章被谁收藏了记录)
func QueryCollectRecordByUserArticles(uid, page, limit int) ([]gorm_model.Article, error) {
	var articles []gorm_model.Article

	if err := DB.Preload("User", "id = ?", uid).Preload("Collects").Preload("Collects.User").
		Where("user_id = ?", uid).
		Limit(limit).Offset((page - 1) * limit).Order("created_at desc").
		Find(&articles).Error; err != nil {
		zap.L().Error("QueryCollectRecordByUserArticles() dao.mysql.mysql_like.Find err=", zap.Error(err))
		return nil, err
	}

	if len(articles) == 0 {
		zap.L().Error("QueryCollectRecordByUserArticles() dao.mysql.mysql_like err=", zap.Error(myErr.NotFoundError()))
		return nil, myErr.NotFoundError()
	}
	return articles, nil
}

// QueryCollectRecordNumByUserArticle 通过uid查询其文章的未读收藏记录数量
func QueryCollectRecordNumByUserArticle(uid int) (int, error) {
	var count int64

	if err := DB.Model(gorm_model.Article{}).Preload("Collects", "is_read = ?", false).Where("user_id = ?", uid).Count(&count).Error; err != nil {
		zap.L().Error("QueryCollectRecordNumByUserArticle() dao.mysql.mysql_like.Count err=", zap.Error(err))
		return -1, err
	}
	return int(count), nil
}
