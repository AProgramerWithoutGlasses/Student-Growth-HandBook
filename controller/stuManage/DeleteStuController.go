package stuManage

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"studentGrow/dao/mysql"
	"studentGrow/models/gorm_model"
	"studentGrow/models/jrx_model"
	"studentGrow/pkg/response"
	token2 "studentGrow/utils/token"
)

// 删除选中用户
func DeleteStuControl(c *gin.Context) {
	token := c.GetHeader("token")
	username, err := token2.GetUsername(token)
	if err != nil {
		response.ResponseError(c, response.ParamFail)
		zap.L().Error(err.Error())
		return
	}

	// 接收请求
	var input struct {
		Selected_students []jrx_model.StuMesStruct `json:"selected_students"`
	}
	err = c.Bind(&input)
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
			if errors.Is(err, gorm.ErrRecordNotFound) {
				err = nil
			} else {
				response.ResponseErrorWithMsg(c, response.ServerErrorCode, err.Error())
				zap.L().Error(err.Error())
				return
			}
		}

		if isManager {
			err := mysql.DeleteSingleUserManager(input.Selected_students[i].Username)
			if err != nil {
				return
			}
		}

		// 删除学生记录
		deleteUserRecord := gorm_model.UserDeleteRecord{
			Username:       username,
			DeleteUsername: input.Selected_students[i].Username,
		}
		err = mysql.DeleteStudentRecord(&deleteUserRecord)
		if err != nil {
			response.ResponseErrorWithMsg(c, 400, err.Error())
			zap.L().Error("stuManager.DeleteStuControl() mysql.DeleteStudentRecord() err : ", zap.Error(err))
			return
		}

	}

	// 响应成功
	response.ResponseSuccess(c, "删除成功!")
}
