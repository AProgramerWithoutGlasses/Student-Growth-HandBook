package article

import (
	"fmt"
	"studentGrow/dao/mysql"
	"studentGrow/models/gorm_model"
)

// GetUnreadReportsService 获取举报信息列表
func GetUnreadReportsService(username string) ([]gorm_model.UserReportArticleRecord, error) {
	// 查询举报信息列表
	reports, err := mysql.GetUnreadReportsForClass(username)
	if err != nil {
		fmt.Println("GetUnreadReportsService() service.article.GetUnreadReports err=", err)
		return nil, err
	}

	return reports, nil
}
