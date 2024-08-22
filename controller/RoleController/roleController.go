package RoleController

import (
	"github.com/gin-gonic/gin"
	"studentGrow/dao/mysql"
	"studentGrow/models/casbinModels"
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

// UpdateRoleMenu 修改权限
func UpdateRoleMenu(c *gin.Context) {
	var casbinService *casbinModels.CasbinService
	//1.接收前端传来的新的权限组
	var NewMenuList struct {
		Role     string   `json:"role"`
		MenuList []string `json:"menuList"`
	}
	err := c.Bind(&NewMenuList)
	if err != nil {
		response.ResponseError(c, 400)
		return
	}
	//2.查询老的权限组
	menuId, err := mysql.SelOldRole(NewMenuList.Role)
	//--1.删除老的权限
	for _, v1 := range menuId {
		rolelist := casbinModels.RolePolicy{
			RoleName: NewMenuList.Role,
			MenuId:   v1,
		}
		err := casbinService.DeleteRolePolicy(rolelist)
		if err != nil {
			return
		}
	}
	//--2.创建新权限
	for _, newMenuId := range NewMenuList.MenuList {
		roleList := casbinModels.RolePolicy{
			RoleName: NewMenuList.Role,
			MenuId:   newMenuId,
		}
		err := casbinService.CreateRolePolicy(roleList)
		if err != nil {
			return
		}
	}
	response.ResponseSuccess(c, "操作成功")
}
