package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
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

	user := router.Group("user")
	user.Use(middleWare.CORSMiddleware())
	{
		//1.像前端返回验证码
		user.POST("/code", login.RidCode)
		//2.接收数据查询是否登录成功id
		user.POST("/hlogin", login.HLogin)
		//前台登录
		user.POST("/qlogin", login.QLogin)
	}

	userLoginAfter := router.Group("user")
	userLoginAfter.Use(middleWare.CORSMiddleware(), token.AuthMiddleware(), middleWare.NewCasbinAuth(casbinService))
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
	elected.Use(middleWare.CORSMiddleware(), token.AuthMiddleware())
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
	router.GET("/sidebar/message", menuController.MenuSide)
}
