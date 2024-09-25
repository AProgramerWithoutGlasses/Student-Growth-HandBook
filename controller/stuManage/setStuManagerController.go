package stuManage

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"studentGrow/pkg/response"
	"studentGrow/service"
	token2 "studentGrow/utils/token"
)

type InnerInput struct {
	Username string `json:"username"`
	Year     string `json:"year"`
}

type Input struct {
	Student     InnerInput `json:"student"`
	ManagerType string     `json:"managerType"`
}

// 设置用户为管理员
func SetStuManagerControl(c *gin.Context) {
	// 接收
	var input Input
	err := c.Bind(&input)
	if err != nil {
		zap.L().Error("stuManage.SetStuManagerControl() c.Bind() failed", zap.Error(err))
		response.ResponseError(c, response.ParamFail)
		return
	}
	token := c.GetHeader("token")
	username, err := token2.GetUsername(token) // class, grade(1-4), collge, superman
	if err != nil {
		response.ResponseError(c, response.ParamFail)
		zap.L().Error(err.Error())
		return
	}

	// 业务
	err = service.SetStuManagerService(input.Student.Username, username, input.ManagerType, input.Student.Year)
	if err != nil {
		response.ResponseErrorWithMsg(c, response.ServerErrorCode, err.Error())
		zap.L().Error("stuManager.SetStuManagerControl() service.SetStuManagerService() failed : ", zap.Error(err))
		return
	}

	// 响应
	response.ResponseSuccess(c, "已将用户 "+input.Student.Username+" 设置为 "+input.ManagerType)

}
