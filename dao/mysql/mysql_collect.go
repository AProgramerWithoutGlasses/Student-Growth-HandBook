package mysql

import (
	"fmt"
	"gorm.io/gorm"
	"studentGrow/models/gorm_model"
)

// UpdateCollectNum 修改收藏数量
func UpdateCollectNum(aid, collectNum int, db *gorm.DB) error {
	if err := db.Model(&gorm_model.Article{}).Where("id = ?", aid).Update("collect_amount", collectNum).Error; err != nil {
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
func DeleteCollectRecord(aid, uid int, db *gorm.DB) error {
	if err := db.Where("article_id = ? and user_id = ?", aid, uid).Delete(&gorm_model.UserCollectRecord{}).Error; err != nil {
		fmt.Println("QueryCollectNum() dao.mysql.mysql_collect")
		return err
	}
	return nil
}
