package stuManage

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"studentGrow/dao/mysql"
	model "studentGrow/models/gorm_model"
	"studentGrow/pkg/response"
)

// 设定用户为管理员
func SetStuManagerControl(c *gin.Context) {
	// 接收数据
	var users []model.User
	err := c.Bind(&users)
	if err != nil {
		response.ResponseErrorWithMsg(c, response.ParamFail, "stuManage.SetStuManagerControl() c.Bind() failed : "+err.Error())
		zap.L().Error("stuManage.SetStuManagerControl() c.Bind() failed", zap.Error(err))
		return
	}

	// 将选中的学生们设置成管理员
	for _, v := range users {
		id, err := mysql.GetIdByUsername(v.Username)
		if err != nil {
			response.ResponseErrorWithMsg(c, 500, "stuManage.SetStuManagerControl() mysql.GetIdByUsername() failed : "+err.Error())
			zap.L().Error("stuManage.SetStuManagerControl() mysql.GetIdByUsername() failed", zap.Error(err))
			return
		}

		err = mysql.SetStuManager(id)
		if err != nil {
			response.ResponseErrorWithMsg(c, 500, "stuManage.SetStuManagerControl() mysql.SetStuManager() failed : "+err.Error())
			zap.L().Error("stuManage.SetStuManagerControl() mysql.SetStuManager() failed", zap.Error(err))
			return
		}
	}

	// 响应成功
	response.ResponseSuccess(c, "操作成功")

}
