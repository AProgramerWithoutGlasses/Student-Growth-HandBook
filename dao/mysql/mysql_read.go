package mysql

import (
	"fmt"
	"go.uber.org/zap"
	"studentGrow/models/gorm_model"
	myErr "studentGrow/pkg/error"
)

// InsertReadRecord 插入浏览记录
func InsertReadRecord(uid, aid int) error {
	readRecord := gorm_model.UserReadRecord{
		UserID:    uint(uid),
		ArticleID: uint(aid),
	}

	if err := DB.Create(&readRecord).Error; err != nil {
		fmt.Println("InsertReadRecord() dao.mysql.mysql_read")
		return err
	}
	return nil
}

// QueryReadListByUserId 查询浏览记录列表
func QueryReadListByUserId(uid int) (readRecords []gorm_model.UserReadRecord, err error) {
	if err = DB.Where("user_id = ?", uid).Find(&readRecords).Error; err != nil {
		zap.L().Error("QueryReadListByUserId() dao.mysql.sql_article", zap.Error(err))
		return nil, err
	}

	if len(readRecords) == 0 {
		return nil, myErr.NotFoundError()
	}
	return readRecords, nil
}

// QueryArticleReadNumById 通过文章id查询文章的浏览量
func QueryArticleReadNumById(aid int) (int, error) {
	var article gorm_model.Article
	if err := DB.Where("id = ?", aid).First(&article).Error; err != nil {
		zap.L().Error("QueryArticleReadNum() dao.mysql.sql_article", zap.Error(err))
		return -1, err
	}
	return article.ReadAmount, nil
}

// UpdateArticleReadNumById 通过文章id修改文章浏览量
func UpdateArticleReadNumById(aid, num int) error {
	if err := DB.Model(&gorm_model.Article{}).Where("id = ?", aid).Update("read_amount", num).Error; err != nil {
		zap.L().Error("UpdateArticleReadNumById() dao.mysql.sql_article", zap.Error(err))
		return err
	}
	return nil
}
