package stuManage

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"studentGrow/dao/mysql"
	"studentGrow/pkg/response"
	"studentGrow/utils/readMessage"
)

// 设置用户为管理员
func SetStuManagerControl(c *gin.Context) {
	// 接收请求信息
	stuMessage, err := readMessage.GetJsonvalue(c)
	if err != nil {
		fmt.Println("readMessage.GetJsonvalue() err : ", err)
	}

	usernameValue, err := stuMessage.GetString("username")
	if err != nil {
		fmt.Println("username GetString() err : ", err)
	}

	// 根据学号获取id
	id, err := mysql.GetIdByUsername(usernameValue)
	if err != nil {
		response.ResponseErrorWithMsg(c, 500, "stuManager.SetStuManagerControl() mysql.GetIdByUsername() failed : "+err.Error())
		zap.L().Error("stuManager.SetStuManagerControl() mysql.GetIdByUsername() failed : ", zap.Error(err))
		return
	}

	// 设置管理员
	err = mysql.SetStuManager(id)
	if err != nil {
		response.ResponseErrorWithMsg(c, 500, "stuManager.SetStuManagerControl() mysql.SetStuManager() failed : "+err.Error())
		zap.L().Error("stuManager.SetStuManagerControl() mysql.SetStuManager() failed : ", zap.Error(err))
		return
	}

	// 响应成功
	response.ResponseSuccess(c, 200)

}
