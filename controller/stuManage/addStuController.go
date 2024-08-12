package stuManage

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"studentGrow/dao/mysql"
	"studentGrow/models/gorm_model"
	"studentGrow/pkg/response"
	"studentGrow/utils/readMessage"
)

// AddSingleStuContro 添加单个学生
func AddSingleStuContro(c *gin.Context) {
	// 接收请求数据
	stuMessage, err := readMessage.GetJsonvalue(c)
	if err != nil {
		fmt.Println("stuManage.AddSingleStuContro() readMessage.GetJsonvalue() err :", err)
	}

	// 获取请求信息中各个字段的值
	nameValue, err := stuMessage.GetString("name")
	if err != nil {
		fmt.Println("name GetString() err : ", err)
	}

	usernameValue, err := stuMessage.GetString("username")
	if err != nil {
		fmt.Println("username GetString() err : ", err)
	}

	passwordValue, err := stuMessage.GetString("password")
	if err != nil {
		fmt.Println("password GetString() err : ", err)
	}

	classValue, err := stuMessage.GetString("class")
	if err != nil {
		fmt.Println("class GetString() err : ", err)
	}

	// 将新增学生信息整合到结构体中
	user := gorm_model.User{
		Name:     nameValue,
		Username: usernameValue,
		Password: passwordValue,
		Class:    classValue,
		Identity: "学生",
	}

	// 在数据库中添加该学生信息
	mysql.AddSingleStudent(&user)

	// 成功响应
	response.ResponseSuccess(c, "添加成功！")
}
