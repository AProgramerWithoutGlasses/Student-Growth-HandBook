package homepage

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"studentGrow/dao/mysql"
	"studentGrow/pkg/response"
	token2 "studentGrow/utils/token"
)

// 用户自述
type SelfContentStruct struct {
	NewSelfContent string `json:"self_content"`
}

// GetSelfContentContro 获取前端发送的用户id, 并将其在数据库中对应的用户自述响应给前端
func GetSelfContentContro(c *gin.Context) {
	input := struct {
		Username string `json:"username"`
	}{}
	err := c.Bind(&input)
	if err != nil {
		response.ResponseError(c, response.ParamFail)
		zap.L().Error(err.Error())
		return
	}

	//hh

	// 根据学号获取id
	id, err := mysql.GetIdByUsername(input.Username)
	if err != nil {
		fmt.Println("homepage.UpdateSelfContentContro() mysql.GetIdByUsername() err : ", err)
	}

	// 根据id，查找数据库中对应的selfContent
	selfContent, err := mysql.GetSelfContent(id)
	if err != nil {
		response.ResponseError(c, response.ServerErrorCode)
		fmt.Println("controller.GetSelfContentContro() mysql.GetSelfContent() err : ", err.Error())
		return
	}

	output := struct {
		SelfContent string `json:"selfContent"`
	}{
		SelfContent: selfContent,
	}
	// 将selfContent发送给前端
	response.ResponseSuccess(c, output)
}

// UpdateSelfContentContro 获取前端发送的学号和newSelfContent, 并将其在数据库中的旧selfContent更新
func UpdateSelfContentContro(c *gin.Context) {
	token := c.GetHeader("token")
	username, err := token2.GetUsername(token)
	if err != nil {
		response.ResponseError(c, response.ParamFail)
		zap.L().Error(err.Error())
		return
	}

	// 接收前端发送的学号和newSelfContent
	var selfContentStruct SelfContentStruct
	if err = c.ShouldBindJSON(&selfContentStruct); err != nil {
		response.ResponseErrorWithMsg(c, response.ServerErrorCode, "获取用户自述失败")
		fmt.Println("selfContent.UpdateSelfContentContro() c.ShouldBindJSON() err : ", err)
		return
	}

	// 根据学号获取id
	id, err := mysql.GetIdByUsername(username)
	if err != nil {
		fmt.Println("homepage.UpdateSelfContentContro() mysql.GetIdByUsername() err : ", err)
	}

	// 在mysql中根据id查询到旧selfContent，用newSelfContent将其替换。
	err = mysql.UpdateSelfContent(id, selfContentStruct.NewSelfContent)
	if err != nil {
		response.ResponseError(c, response.ServerErrorCode)
		fmt.Println("UpdateSelfContentContro() mysql.UpdateSelfContent() err : ", err)
		return
	}

	// 响应成功信息
	response.ResponseSuccess(c, struct{}{})
}
