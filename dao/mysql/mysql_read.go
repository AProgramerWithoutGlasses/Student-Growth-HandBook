package mysql

import (
	"fmt"
	"studentGrow/models/gorm_model"
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
		fmt.Println("QueryReadListByUserId() dao.mysql.mysql_read")
		return nil, err
	}
	return readRecords, nil
}
