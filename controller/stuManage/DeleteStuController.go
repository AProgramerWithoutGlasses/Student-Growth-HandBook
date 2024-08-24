package stuManage

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"studentGrow/dao/mysql"
	"studentGrow/models/jrx_model"
	"studentGrow/pkg/response"
)

// 删除选中用户
func DeleteStuControl(c *gin.Context) {
	// 接收请求
	var input struct {
		Selected_students []jrx_model.StuMesStruct `json:"selected_students"`
	}
	err := c.Bind(&input)
	if err != nil {
		response.ResponseErrorWithMsg(c, 500, "stuManager.DeleteStuControl() c.Bind() failed : "+err.Error())
		zap.L().Error("stuManager.DeleteStuControl() c.Bind() failed : ", zap.Error(err))
		return
	}

	// 删除选中的用户
	for i, _ := range input.Selected_students {

		id, err := mysql.GetIdByUsername(input.Selected_students[i].Username)
		if err != nil {
			response.ResponseErrorWithMsg(c, 500, "stuManager.DeleteStuControl() mysql.GetIdByUsername() failed : "+err.Error())
			zap.L().Error("stuManager.DeleteStuControl() mysql.GetIdByUsername() failed : ", zap.Error(err))
			return
		}

		err = mysql.DeleteSingleUser(id)
		if err != nil {
			response.ResponseErrorWithMsg(c, 400, err.Error())
			zap.L().Error("stuManager.DeleteStuControl() mysql.DeleteSingleStudent() err : ", zap.Error(err))
			return
		}

		isManager, err := mysql.GetIsManagerByUsername(input.Selected_students[i].Username)
		if err != nil {
			response.ResponseErrorWithMsg(c, response.ServerErrorCode, err.Error())
			zap.L().Error(err.Error())
			return
		}

		if isManager {
			mysql.DeleteSingleUserManager(input.Selected_students[i].Username)
		}

	}

	// 响应成功
	response.ResponseSuccess(c, "删除成功!")
}
