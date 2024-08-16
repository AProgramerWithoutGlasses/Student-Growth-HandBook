package stuManage

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"studentGrow/pkg/response"
	"studentGrow/service"
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
	// 接收请求信息
	var input Input
	err := c.Bind(&input)
	if err != nil {
		zap.L().Error("stuManage.SetStuManagerControl() c.Bind() failed", zap.Error(err))
		response.ResponseError(c, response.ParamFail)
		return
	}

	// 设置管理员
	err = service.SetStuManagerService(input.Student.Username, input.ManagerType, input.Student.Year)
	if err != nil {
		response.ResponseErrorWithMsg(c, 500, err.Error())
		zap.L().Error("stuManager.SetStuManagerControl() mysql.SetStuManager() failed : ", zap.Error(err))
		return
	}

	// 响应成功
	response.ResponseSuccess(c, 200)

}
