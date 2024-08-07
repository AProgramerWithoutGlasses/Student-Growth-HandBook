package mysql

import (
	"fmt"
	"studentGrow/models/gorm_model"
	myErr "studentGrow/pkg/error"
)

// GetClassByUsername 通过username获取班级
func GetClassByUsername(username string) (string, error) {
	var user gorm_model.User
	if err := DB.Where("username = ?", username).First(&user).Error; err != nil {
		fmt.Println("GetClassByUsername() dao.mysql.sql_msg")
		return "", err
	}
	return user.Class, nil
}

// GetUnreadReportsForClass 获取未读举报信息-班级
func GetUnreadReportsForClass(username string) ([]gorm_model.UserReportArticleRecord, error) {
	// 通过username获取管理员班级
	class, err := GetClassByUsername(username)
	if err != nil {
		return nil, err
	}

	// 按举报时间逆序查询
	//  通过文章id查找到到对应的用户
	var reports []gorm_model.UserReportArticleRecord
	DB.Preload("User", "class = ?", class).Preload("Article").
		Where("is_read = ?", false).Order("created_at DESC").Find(&reports)

	if len(reports) == 0 {
		return nil, myErr.NotFoundError()
	}

	return reports, nil
}

// GetUnreadReportForGrade 获取未读举报信息-年级
func GetUnreadReportForGrade(username string) {
	// 通过username获取管理员年级

}

// AckUnreadReportsById 确认未读举报信息
func AckUnreadReportsById(reportsId []int) {

}
