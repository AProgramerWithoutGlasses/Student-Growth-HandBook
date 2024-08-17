package teacherManage

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"studentGrow/pkg/response"
	"studentGrow/service"
)

// 设置老师管理员
func SetTeacherManagerControl(c *gin.Context) {
	// 响应
	var intput struct {
		Username    string `json:"username"`
		ManagerType string `json:"manager_type"`
	}
	err := c.Bind(&intput)
	if err != nil {
		response.ResponseError(c, response.ParamFail)
		zap.L().Error("teacherManage.SetTeacherManagerControl() c.Bind() failed : ", zap.Error(err))
		return
	}

	// 业务
	err = service.SetTeacherManagerService(intput.Username, intput.ManagerType)
	if err != nil {
		response.ResponseErrorWithMsg(c, response.ServerErrorCode, err.Error())
		zap.L().Error("teacherManage.SetTeacherManagerControl() service.SetTeacherManagerService() failed : ", zap.Error(err))
		return
	}

	// 响应
	response.ResponseSuccess(c, nil)
}
