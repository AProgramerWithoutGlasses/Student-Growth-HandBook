package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"studentGrow/controller/RoleController"
	"studentGrow/controller/growth"
	"studentGrow/controller/login"
	"studentGrow/controller/menuController"
	"studentGrow/dao/mysql"
	"studentGrow/models/casbinModels"
	"studentGrow/utils/middleWare"
	"studentGrow/utils/token"
)

func RoutesXue(router *gin.Engine) {
	casbinService, err := casbinModels.NewCasbinService(mysql.DB)
	if err != nil {
		fmt.Println("Setup models.NewCasbinService()  err")
	}
	router.Use(middleWare.CORSMiddleware())
	user := router.Group("user")
	{
		//1.像前端返回验证码
		user.POST("/code", login.RidCode)
		//2.接收数据查询是否登录成功id
		user.POST("/hlogin", login.HLogin)
		//前台登录
		user.POST("/qlogin", login.QLogin)
	}
	userMessage := router.Group("user")
	{
		//获取注册天数
		userMessage.POST("/register/day", login.RegisterDay)
	}

	//casbin鉴权
	userLoginAfter := router.Group("user")
	userLoginAfter.Use(middleWare.NewCasbinAuth(casbinService))
	{
		//班级管理员首页
		userLoginAfter.POST("/fpage/class", login.FPageClass)
		//年级管理员首页
		userLoginAfter.POST("/fpage/grade", login.FPageGrade)
		//学院管理员首页
		userLoginAfter.POST("/fpage/college", login.FPageCollege)
		//超级管理员首页
		userLoginAfter.POST("/fpage/superman", login.FPageCollege)
		//首页柱状图
		userLoginAfter.GET("/fpage/pillar", login.Pillar)
		//获取登陆者的全部信息
		userLoginAfter.GET("/message", menuController.HeadRoute)
	}

	//暂时不添加casbin中间件
	elected := router.Group("star")
	elected.Use(token.AuthMiddleware())
	{
		//成长之星退选时展示的表格
		elected.GET("/select", growth.Search)
		//班级管理员推选
		elected.POST("/elected/class", growth.ElectClass)
		//年级管理员推选
		elected.POST("/elected/grade", growth.ElectGrade)
		//学院管理员推选
		elected.POST("/elected/college", growth.ElectCollege)
		//院级管理员公布
		elected.POST("/public/college", growth.PublicStar)
		//搜索第几届成长之星的接口
		elected.GET("/termStar", growth.StarPub)
		elected.GET("/class_star", growth.BackStarClass)
		elected.GET("/grade_star", growth.BackStarGrade)
		elected.GET("/college_star", growth.BackStarCollege)
		elected.POST("/change_disabled", growth.ChangeStatus)
	}
	//前端侧边栏
	router.GET("/sidebar/message", menuController.MenuSide)
	//菜单管理 casbin鉴权
	menu := router.Group("menuManage")
	//menu.Use(middleWare.NewCasbinAuth(casbinService))
	{
		//菜单初始化
		menu.GET("/init", menuController.MenuMangerInit)
		//添加菜单
		menu.POST("/newelyBuilt", menuController.AddMenu)
		//删除菜单
		menu.POST("/delete", menuController.DeleteMenu)
		//搜索菜单
		menu.GET("/selectInfo", menuController.SearchMenu)
		//编辑菜单
		menu.POST("/edit", menuController.UpdateMenu)
	}
	//角色管理
	role := router.Group("role")
	{
		role.GET("/list", RoleController.RoleList)
	}
}
