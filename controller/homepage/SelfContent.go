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
	NewSelfContent string `json:"newSelfContent"`
}

// GetSelfContentContro 获取前端发送的用户id, 并将其在数据库中对应的用户自述响应给前端
func GetSelfContentContro(c *gin.Context) {
	// 接收
	token := c.GetHeader("token")
	username, err := token2.GetUsername(token)
	if err != nil {
		response.ResponseError(c, response.ParamFail)
		zap.L().Error(err.Error())
		return
	}

	// 接收前端发送的Username，将其存储到selfContentStruct结构体中
	var selfContentStruct SelfContentStruct
	if err := c.ShouldBindJSON(&selfContentStruct); err != nil {
		response.ResponseErrorWithMsg(c, response.ServerErrorCode, "获取用户自述失败")
		fmt.Println("controller.GetSelfContentContro() c.ShouldBindJSON() err : ", err.Error())
		return
	}

	// 根据学号获取id
	id, err := mysql.GetIdByUsername(username)
	if err != nil {
		fmt.Println("homepage.UpdateSelfContentContro() mysql.GetIdByUsername() err : ", err)
	}

	// 根据id，查找数据库中对应的selfContent
	selfContent, err := mysql.GetSelfContent(id)
	if err != nil {
		response.ResponseError(c, 401)
		fmt.Println("controller.GetSelfContentContro() mysql.GetSelfContent() err : ", err.Error())
		return
	}

	// 将selfContent发送给前端
	response.ResponseSuccess(c, "该用户自述为: "+selfContent)
}

// UpdateSelfContentContro 获取前端发送的学号和newSelfContent, 并将其在数据库中的旧selfContent更新
func UpdateSelfContentContro(c *gin.Context) {
	// 接收
	token := c.GetHeader("token")
	username, err := token2.GetUsername(token)
	if err != nil {
		response.ResponseError(c, response.ParamFail)
		zap.L().Error(err.Error())
		return
	}

	// 接收前端发送的学号和newSelfContent
	var selfContentStruct SelfContentStruct
	if err := c.ShouldBindJSON(&selfContentStruct); err != nil {
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
	response.ResponseSuccess(c, "")
}
