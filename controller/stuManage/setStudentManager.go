package stuManage

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"studentGrow/dao/mysql"
	"studentGrow/models/jrx_model"

	"studentGrow/pkg/response"
)

// 设定用户为管理员
func SetStuManagerControl(c *gin.Context) {
	// 接收请求
	var input struct {
		Selected_students []jrx_model.StuMesStruct
	}
	err := c.Bind(&input)
	if err != nil {
		response.ResponseErrorWithMsg(c, 500, "stuManager.SetStuManagerControl() c.Bind() failed : "+err.Error())
		zap.L().Error("stuManager.SetStuManagerControl() c.Bind() failed : ", zap.Error(err))
		return
	}

	// 设为管理员
	for _, v := range input.Selected_students {
		id, err := mysql.GetIdByUsername(v.Username)
		if err != nil {
			response.ResponseErrorWithMsg(c, 500, "stuManager.SetStuManagerControl() mysql.GetIdByUsername() failed : "+err.Error())
			zap.L().Error("stuManager.SetStuManagerControl() mysql.GetIdByUsername() failed : ", zap.Error(err))
			return
		}

		err = mysql.SetStuManager(id)
		if err != nil {
			response.ResponseErrorWithMsg(c, 500, "stuManager.SetStuManagerControl() mysql.SetStuManager() failed : "+err.Error())
			zap.L().Error("stuManager.SetStuManagerControl() mysql.SetStuManager() failed : ", zap.Error(err))
			return
		}
	}

	// 响应成功
	response.ResponseSuccess(c, "操作成功")

}
