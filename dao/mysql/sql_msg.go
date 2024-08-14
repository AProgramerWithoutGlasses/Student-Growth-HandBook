package mysql

import (
	"fmt"
	"studentGrow/models/gorm_model"
	myErr "studentGrow/pkg/error"
	time "studentGrow/utils/timeConverter"
)

// GetUserByUsername 通过username获取user对象
func GetUserByUsername(username string) (*gorm_model.User, error) {
	var user gorm_model.User
	if err := DB.Where("username = ?", username).First(&user).Error; err != nil {
		fmt.Println("GetClassByUsername() dao.mysql.sql_msg")
		return nil, err
	}
	return &user, nil
}

// GetUnreadReportsForClass 获取未读举报信息-班级
func GetUnreadReportsForClass(username string, limit, page int) ([]gorm_model.UserReportArticleRecord, error) {
	// 通过username获取管理员
	user, err := GetUserByUsername(username)
	if err != nil {
		return nil, err
	}

	// 按举报时间逆序查询
	//  通过文章id查找到到对应的用户
	var reports []gorm_model.UserReportArticleRecord
	if err := DB.Preload("User", "class = ?", user.Class).Preload("Article", "ban = ?", false).
		Where("is_read = ?", false).Order("created_at DESC, article_id ASC").
		Limit(limit).Offset((page - 1) * limit).
		Find(&reports).Error; err != nil {
		fmt.Println("GetUnreadReportsForClass() dao.mysql.sql_msg")
		return nil, err
	}

	if len(reports) == 0 {
		return nil, myErr.NotFoundError()
	}

	return reports, nil
}

// GetUnreadReportsForGrade 获取未读举报信息-年级
func GetUnreadReportsForGrade(grade int, limit, page int) ([]gorm_model.UserReportArticleRecord, error) {
	// 通过年级获取入学年份
	year, err := time.GetEnrollmentYear(grade)
	if err != nil {
		fmt.Println("GetUnreadReportsForGrade() dao.mysql.sql_msg.GetEnrollmentYear")
		return nil, err
	}

	// 按举报时间逆序查询
	//  通过文章id查找到到对应的用户
	var reports []gorm_model.UserReportArticleRecord
	if err := DB.Preload("User", "plus_time between ? and ?",
		fmt.Sprintf("%d-01-01", year.Year()), fmt.Sprintf("%d-12-31", year.Year())).
		Preload("Article", "ban = ?", false).
		Where("is_read = ?", false).Order("created_at DESC, article_id ASC").
		Limit(limit).Offset((page - 1) * limit).
		Find(&reports).Error; err != nil {
		fmt.Println("GetUnreadReportForGrade() dao.mysql.sql_msg")
		return nil, err
	}

	if len(reports) == 0 {
		return nil, myErr.NotFoundError()
	}

	return reports, nil
}

// GetUnreadReportsForSuperman 获取未读举报信息 - 超级(院级)
func GetUnreadReportsForSuperman(limit, page int) ([]gorm_model.UserReportArticleRecord, error) {

	var reports []gorm_model.UserReportArticleRecord
	if err := DB.Preload("User").
		Preload("Article", "ban = ?", false).
		Where("is_read = ?", false).Order("created_at DESC, article_id ASC").
		Limit(limit).Offset((page - 1) * limit).
		Find(&reports).Error; err != nil {
		fmt.Println("GetUnreadReportForCollege() dao.mysql.sql_msg")
		return nil, err
	}

	if len(reports) == 0 {
		return nil, myErr.NotFoundError()
	}

	return reports, nil

}

// AckUnreadReportsForClass 确认未读举报信息 - 班级
func AckUnreadReportsForClass(reportsId int, username string) error {
	// 通过username获取管理员
	user, err := GetUserByUsername(username)
	if err != nil {
		return err
	}

	// 修改举报信息读取状态为已读
	result := DB.Preload("User", "class = ?", user.Class).
		Where("article_id = ?", reportsId).
		Updates(gorm_model.UserReportArticleRecord{IsRead: true})

	if result.Error != nil {
		fmt.Println("AckUnreadReportsById() dao.mysql.sql_msg")
		return result.Error
	}

	return nil
}

// AckUnreadReportsForGrade 确认未读举报信息 - 年级
func AckUnreadReportsForGrade(reportsId int, grade int) error {
	// 通过年级获取入学年份
	year, err := time.GetEnrollmentYear(grade)
	if err != nil {
		fmt.Println("GetUnreadReportsForGrade() dao.mysql.sql_msg.GetEnrollmentYear")
		return err
	}

	// 修改举报信息读取状态为已读
	result := DB.Preload("User", "plus_time between ? and ?",
		fmt.Sprintf("%s-01-01", year), fmt.Sprintf("%s-12-31", year)).
		Where("article_id = ?", reportsId).
		Updates(gorm_model.UserReportArticleRecord{IsRead: true})

	if result.Error != nil {
		fmt.Println("AckUnreadReportsById() dao.mysql.sql_msg")
		return result.Error
	}

	return nil
}

// AckUnreadReportsForSuperman 确认未读举报信息 - 超级(院级)
func AckUnreadReportsForSuperman(reportsId int) error {
	// 修改举报信息读取状态为已读
	result := DB.Preload("User").
		Where("article_id = ?", reportsId).
		Updates(gorm_model.UserReportArticleRecord{IsRead: true})

	if result.Error != nil {
		fmt.Println("AckUnreadReportsById() dao.mysql.sql_msg")
		return result.Error
	}

	return nil
}

//
