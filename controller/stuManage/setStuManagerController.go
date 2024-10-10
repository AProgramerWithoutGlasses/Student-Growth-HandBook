package stuManage

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"studentGrow/models"
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
	// 接收
	var input Input
	err := c.Bind(&input)
	if err != nil {
		zap.L().Error("stuManage.SetStuManagerControl() c.Bind() failed", zap.Error(err))
		response.ResponseError(c, response.ParamFail)
		return
	}

	claim, exist := c.Get("claim")
	if !exist {
		response.ResponseError(c, response.TokenError)
		zap.L().Error("token错误")
		return
	}
	username := claim.(*models.Claims).Username

	// 业务
	err = service.SetStuManagerService(input.Student.Username, username, input.ManagerType, input.Student.Year)
	if err != nil {
		response.ResponseErrorWithMsg(c, response.ServerErrorCode, err.Error())
		zap.L().Error("stuManager.SetStuManagerControl() service.SetStuManagerService() failed : ", zap.Error(err))
		return
	}

	// 响应
	response.ResponseSuccessWithMsg(c, "已将用户 "+input.Student.Username+" 设置为 "+input.ManagerType, "")

}
