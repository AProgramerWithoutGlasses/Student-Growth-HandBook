package stuManage

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"studentGrow/dao/mysql"
	"studentGrow/pkg/response"
	"studentGrow/utils/readMessage"
)

func DeleteStuControl(c *gin.Context) {
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
		fmt.Println("stuManage.DeleteStuControl() mysql.GetIdByUsername() err : ", err)
		response.ResponseErrorWithMsg(c, 400, err.Error())
		return
	}

	// mysql中删除该学生
	err = mysql.DeleteSingleStudent(id)
	if err != nil {
		fmt.Println("stuManage.DeleteStuControl() mysql.DeleteSingleStudent() err : ", err)
		response.ResponseErrorWithMsg(c, 400, err.Error())
		zap.L().Error("mysql.DeleteSingleStudent() err : ", zap.Error(err))
		return
	}
	response.ResponseSuccess(c, 200)

}
