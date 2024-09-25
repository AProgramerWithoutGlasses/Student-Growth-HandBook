package stuManage

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"studentGrow/models/jrx_model"
	"studentGrow/pkg/response"
	"studentGrow/service"
)

func EditStuControl(c *gin.Context) {
	var user jrx_model.ChangeStuMesStruct

	// 接收数据
	if err := c.Bind(&user); err != nil {
		response.ResponseError(c, response.ParamFail)
		zap.L().Error("stuManage.EditStuControl() c.Bind() err : ", zap.Error(err))
		return
	}

	// 业务处理
	err := service.EditStuService(user)
	if err != nil {
		response.ResponseError(c, response.ServerErrorCode)
		zap.L().Error(err.Error())
		return
	}

	// 响应
	response.ResponseSuccessWithMsg(c, "信息修改成功", nil)

}
