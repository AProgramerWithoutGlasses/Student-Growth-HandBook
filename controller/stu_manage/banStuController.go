package stu_manage

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"studentGrow/dao/mysql"
	"studentGrow/pkg/response"
	"studentGrow/utils/readMessage"
)

func BanStuControl(c *gin.Context) {
	// 接收请求信息
	stuMessage, err := readMessage.GetJsonvalue(c)
	if err != nil {
		fmt.Println("stu_manage.BanStuControl() readMessage.GetJsonvalue() err : ", err)
		response.ResponseErrorWithMsg(c, 500, err.Error())
		return
	}

	usernameValue, err := stuMessage.GetString("username")
	if err != nil {
		fmt.Println("stu_manage.BanStuControl() username GetString() err : ", err)
		response.ResponseErrorWithMsg(c, 500, err.Error())
		return
	}

	// 根据学号获取id
	id, err := mysql.GetIdByUsername(usernameValue)
	if err != nil {
		fmt.Println("stu_manage.DeleteStuControl() mysql.GetIdByUsername() err : ", err)
		response.ResponseErrorWithMsg(c, 500, err.Error())
		return
	}

	// mysql中封禁该学生
	err = mysql.BanStudent(id)
	if err != nil {
		fmt.Println("stu_manage.DeleteStuControl() mysql.BanStudent() err : ", err)
		response.ResponseErrorWithMsg(c, 500, err.Error())
		return
	}
	response.ResponseSuccess(c, 200)

}
