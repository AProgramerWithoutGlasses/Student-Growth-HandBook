package stuManage

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"studentGrow/dao/mysql"
	"studentGrow/models/jrx_model"
	"studentGrow/pkg/response"
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
	id, err := mysql.GetIdByUsername(user.Username)
	if err != nil {
		response.ResponseErrorWithMsg(c, 500, err.Error())
		zap.L().Error("mysql.GetIdByUsername failed", zap.Error(err))
		return
	}

	if err := mysql.ChangeStudentMessage(id, user); err != nil {
		response.ResponseErrorWithMsg(c, 500, err.Error())
		zap.L().Error("mysql.ChangeStudentMessage failed", zap.Error(err))
		return
	}

	// 响应
	response.ResponseSuccess(c, "")

}
