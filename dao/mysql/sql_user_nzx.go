package mysql

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm"
	model "studentGrow/models/gorm_model"
	"studentGrow/utils/timeConverter"
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

// QueryUserByAdvancedFilter 高级筛选用户(年级、班级、姓名)
func QueryUserByAdvancedFilter(grade int, class []string, name string) (*gorm.DB, error) {

	year, err := timeConverter.GetEnrollmentYear(grade)
	if err != nil {
		zap.L().Error("QueryClassByUsername() dao.mysql.sql_user_nzx err=", zap.Error(err))
		return nil, err
	}
	query := DB.Model(&model.User{}).
		Where("plus_time BETWEEN ? AND ? AND class IN ? AND name LIKE ?",
			fmt.Sprintf("%s-01-01", year.Year()), fmt.Sprintf("%s-12-31", year.Year()), class, fmt.Sprintf("%%%s%%", name))
	if query.Error != nil {
		zap.L().Error("QueryArticleByAdvancedFilter() service.article.Error err=", zap.Error(err))
		return nil, err
	}
	return query, nil
}
