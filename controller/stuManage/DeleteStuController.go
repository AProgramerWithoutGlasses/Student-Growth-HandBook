package stuManage

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"studentGrow/dao/mysql"
	"studentGrow/pkg/response"
)

type name struct {
	Name      string `json:"name"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Class     string `json:"class"`
	Year      int    `json:"year"`
	Gender    string `json:"gender"`
	Telephone string `json:"telephone"`
	Ban       bool   `json:"ban"`
	IsManager bool   `json:"isManager"`
}

// 删除选中用户
func DeleteStuControl(c *gin.Context) {
	// 接收请求
	var input struct {
		Selected_students []name
	}
	err := c.Bind(&input)
	if err != nil {
		response.ResponseErrorWithMsg(c, 500, "stuManager.DeleteStuControl() c.Bind() failed : "+err.Error())
		zap.L().Error("stuManager.DeleteStuControl() c.Bind() failed : ", zap.Error(err))
		return
	}

	// 删除选中的用户
	for _, v := range input.Selected_students {

		id, err := mysql.GetIdByUsername(v.Username)
		if err != nil {
			response.ResponseErrorWithMsg(c, 500, "stuManager.DeleteStuControl() mysql.GetIdByUsername() failed : "+err.Error())
			zap.L().Error("stuManager.DeleteStuControl() mysql.GetIdByUsername() failed : ", zap.Error(err))
			return
		}

		err = mysql.DeleteSingleStudent(id)
		if err != nil {
			response.ResponseErrorWithMsg(c, 400, err.Error())
			zap.L().Error("stuManager.DeleteStuControl() mysql.DeleteSingleStudent() err : ", zap.Error(err))
			return
		}
	}

	// 响应成功
	response.ResponseSuccess(c, "删除成功!")
}
