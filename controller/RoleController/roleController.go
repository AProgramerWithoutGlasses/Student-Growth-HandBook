package RoleController

import (
	"github.com/gin-gonic/gin"
	"studentGrow/pkg/response"
	service "studentGrow/service/permission"
)

// RoleList 展示角色
func RoleList(c *gin.Context) {
	rolelist, err := service.RoleData()
	if err != nil {
		response.ResponseError(c, 400)
		return
	}
	response.ResponseSuccess(c, rolelist)
}

// ShowMenu 权限列表
func ShowMenu(c *gin.Context) {
	//接收前端数据
	var fromData struct {
		Role string `json:"role"`
	}
	err := c.Bind(&fromData)
	if err != nil {
		response.ResponseErrorWithMsg(c, 400, "获取数据失败")
		return
	}
	//定义返回前端的数据
	menuList, err := service.RoleMenuTree(fromData.Role, 0)
	if err != nil {
		response.ResponseError(c, 400)
		return
	}
	response.ResponseSuccess(c, menuList)
}
