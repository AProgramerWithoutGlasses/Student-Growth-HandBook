package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"studentGrow/controller/login"
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
		userLoginAfter.POST("/fpage/class", login.FPageClass)
		userLoginAfter.POST("/fpage/grade", login.FPageGrade)
		userLoginAfter.POST("/fpage/college", login.FPageCollege)
		userLoginAfter.POST("/fpage/superman", login.FPageCollege)
		userLoginAfter.GET("/fpage/pillar", login.Pillar)
	}

}
