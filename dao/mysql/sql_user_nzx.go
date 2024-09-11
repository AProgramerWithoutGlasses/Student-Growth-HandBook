package mysql

import (
	"go.uber.org/zap"
	model "studentGrow/models/gorm_model"
)

// QueryClassByUsername 通过用户名查找班级
func QueryClassByUsername(username string) (string, error) {
	var class string
	if err := DB.Model(&model.User{}).Select("class").Where("username = ?", username).First(&class).Error; err != nil {
		zap.L().Error("QueryClassByUsername() dao.mysql.sql_user_nzx err=", zap.Error(err))
		return "", err
	}
	return class, nil
}
