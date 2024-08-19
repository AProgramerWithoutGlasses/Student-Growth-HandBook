package menuController

import (
	"github.com/gin-gonic/gin"
	"studentGrow/dao/mysql"
	"studentGrow/models"
	"studentGrow/models/gorm_model"
	"studentGrow/pkg/response"
	service "studentGrow/service/permission"
)

// MenuMangerInit 初始化菜单
func MenuMangerInit(c *gin.Context) {
	menu, err := service.BuildMenuTree(0)
	if err != nil {
		response.ResponseError(c, 400)
		return
	}
	response.ResponseSuccess(c, menu)
}

// AddMenu 新增菜单
func AddMenu(c *gin.Context) {
	var fId int
	var backMenu models.Menu
	err := c.Bind(&backMenu)
	if err != nil {
		response.ResponseErrorWithMsg(c, 400, "获取数据失败")
	}
	//1.查找父id
	if backMenu.FatherMenu == "" {
		fId = 0
	} else {
		fId, err = mysql.SelMenuFId(backMenu.FatherMenu)
	}
	//2.增添数据模型
	menu := gorm_model.Menus{
		ParentId:      fId,
		TreePath:      "",
		Name:          backMenu.Name,
		Type:          backMenu.Type,
		RouteName:     backMenu.RouteName,
		Path:          backMenu.Path,
		Component:     backMenu.Component,
		Perm:          backMenu.Perm,
		Visible:       backMenu.Visible,
		Sort:          backMenu.Sort,
		Icon:          backMenu.Icon,
		Redirect:      backMenu.Redirect,
		Roles:         "",
		RequestUrl:    backMenu.RequestUrl,
		RequestMethod: backMenu.RequestMethod,
	}
	//3.新增菜单
	err = mysql.AddMenu(menu)
	if err != nil {
		response.ResponseError(c, 400)
		return
	}
	response.ResponseSuccess(c, "")
}

// DeleteMenu 删除菜单
func DeleteMenu(c *gin.Context) {
	//接收前端数据
	var fromdata struct {
		MenuName string `json:"menuName"`
	}
	err := c.Bind(&fromdata)
	if err != nil {
		response.ResponseError(c, 400)
		return
	}
	//删除menus表中的数据
	//1.查询id
	id, err := mysql.SelMenuFId(fromdata.MenuName)
	//删除menu中的数据
	err = mysql.DeleteMenu(id)
	if err != nil {
		response.ResponseError(c, 400)
		return
	}
	//删除路由参数
	err = mysql.DeleteParam(id)
	if err != nil {
		response.ResponseError(c, 400)
		return
	}
	response.ResponseSuccess(c, "")
}

// SearchMenu 搜索菜单
func SearchMenu(c *gin.Context) {
	//接收前端数据
	var inputdata struct {
		Input string `form:"input"`
	}
	err := c.Bind(&inputdata)
	if err != nil {
		response.ResponseError(c, 401)
		return
	}
	//返回前端的数据
	menus, err := service.BuildMenu(inputdata.Input)
	if err != nil {
		response.ResponseError(c, 400)
		return
	}
	response.ResponseSuccess(c, menus)
}

// UpdateMenu 编辑菜单
func UpdateMenu(c *gin.Context) {
	//接收前端数据
	var menu models.Menu
	err := c.Bind(&menu)
	if err != nil {
		response.ResponseErrorWithMsg(c, 401, "没有接收到数据")
		return
	}
	err = service.UpdateMenuData(menu)
	if err != nil {
		response.ResponseError(c, 400)
		return
	}
	response.ResponseSuccess(c, "")
}
