package article

import (
	"fmt"
	"studentGrow/dao/mysql"
	"studentGrow/models/gorm_model"
	myErr "studentGrow/pkg/error"
)

// GetUnreadReportsForService 获取举报信息列表
func GetUnreadReportsForService(username string, role string) (reports []gorm_model.UserReportArticleRecord, err error) {

	// 选择管理员角色
	switch role {
	case "class":
		reports, err = mysql.GetUnreadReportsForClass(username)
	case "grade1":
		reports, err = mysql.GetUnreadReportsForGrade(1)
	case "grade2":
		reports, err = mysql.GetUnreadReportsForGrade(2)
	case "grade3":
		reports, err = mysql.GetUnreadReportsForGrade(3)
	case "grade4":
		reports, err = mysql.GetUnreadReportsForGrade(4)
	case "college":
		reports, err = mysql.GetUnreadReportsForSuperman()
	case "superman":
		reports, err = mysql.GetUnreadReportsForSuperman()
	default:
		return nil, myErr.NotFoundError()
	}

	if err != nil {
		fmt.Println("GetUnreadReportsForClassService() service.article.GetUnreadReports err=", err)
		return nil, err
	}

	return reports, nil
}

// AckUnreadReportsService 确认举报信息
func AckUnreadReportsService(reportId int, username string, role string) (err error) {
	// 选择管理员角色
	switch role {
	case "class":
		err = mysql.AckUnreadReportsForClass(reportId, username)
	case "grade1":
		err = mysql.AckUnreadReportsForGrade(reportId, 1)
	case "grade2":
		err = mysql.AckUnreadReportsForGrade(reportId, 2)
	case "grade3":
		err = mysql.AckUnreadReportsForGrade(reportId, 3)
	case "grade4":
		err = mysql.AckUnreadReportsForGrade(reportId, 4)
	case "college":
		err = mysql.AckUnreadReportsForSuperman(reportId)
	case "superman":
		err = mysql.AckUnreadReportsForSuperman(reportId)
	default:
		return myErr.NotFoundError()
	}

	if err != nil {
		fmt.Println("GetUnreadReportsForClassService() service.article.GetUnreadReports err=", err)
		return err
	}

	return nil
}
