package message

import (
	"fmt"
	"github.com/gin-gonic/gin"
	error2 "studentGrow/pkg/error"
	res "studentGrow/pkg/response"
	"studentGrow/service/article"
	"studentGrow/utils/token"
)

// GetUnreadReportsController 获取未读举报信息
func GetUnreadReportsController(c *gin.Context) {

	// 通过token获取username
	username, err := token.GetUsername(c.GetHeader("token"))

	//if err != nil {
	//	fmt.Println("GetUnreadReportsController() controller.message.GetUsername")
	//}

	// 获取未读举报列表
	reports, err := article.GetUnreadReportsService(username)
	if err != nil {
		fmt.Println("GetUnreadReportsController() controller.message.GetUnreadReportsService err=", err)
		error2.CheckErrors(err, c)
		return
	}

	// 返回响应
	res.ResponseSuccess(c, map[string]any{
		"reports":    reports,
		"report_num": len(reports),
	})

}

// 确认未读举报信息
