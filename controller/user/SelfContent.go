package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"studentGrow/dao/mysql"
	response2 "studentGrow/pkg/response"
)

// 用户自述
type SelfContentStruct struct {
	Username       int    `json:"username"`
	NewSelfContent string `json:"newSelfContent"`
}

// 获取前端发送的用户id, 并将其在数据库中对应的用户自述响应给前端
func GetSelfContentContro(c *gin.Context) {
	// 接收前端发送的Username，将其存储到selfContentStruct结构体中
	var selfContentStruct SelfContentStruct
	if err := c.ShouldBindJSON(&selfContentStruct); err != nil {
		response2.ResponseErrorWithMsg(c, response2.ServerErrorCode, "获取用户自述失败")
		fmt.Println("controller.GetSelfContentContro() c.ShouldBindJSON() err : ", err.Error())
		return
	}

	// 根据username，查找数据库中对应的selfContent
	selfContent, err := mysql.GetSelfContent(selfContentStruct.Username)
	if err != nil {
		response2.ResponseError(c, 401)
		fmt.Println("controller.GetSelfContentContro() mysql.GetSelfContent() err : ", err.Error())
		return
	}

	// 将selfContent发送给前端
	response2.ResponseSuccess(c, selfContent)
}

// 获取前端发送的用户id和newSelfContent, 并将其在数据库中的旧selfContent更新
func UpdateSelfContentContro(c *gin.Context) {
	// 接收前端发送的用户id和newSelfContent
	var selfContentStruct SelfContentStruct
	if err := c.ShouldBindJSON(&selfContentStruct); err != nil {
		response2.ResponseErrorWithMsg(c, response2.ServerErrorCode, "获取用户自述失败")
		fmt.Println("UpdateSelfContentContro() c.ShouldBindJSON() err : ", err)
		return
	}

	// 在mysql中根据id查询到旧selfContent，用newSelfContent将其替换。
	err := mysql.UpdateSelfContent(selfContentStruct.Username, selfContentStruct.NewSelfContent)
	if err != nil {
		response2.ResponseError(c, response2.ServerErrorCode)
		fmt.Println("UpdateSelfContentContro() mysql.UpdateSelfContent() err : ", err)
		return
	}

	// 响应成功信息
	response2.ResponseSuccess(c, "")
}
