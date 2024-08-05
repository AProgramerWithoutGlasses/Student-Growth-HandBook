package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"studentGrow/controller/login"
	"studentGrow/dao/mysql"
	"studentGrow/models/casbinModels"
	"studentGrow/utils/middleWare"
)

func RoutesXue(router *gin.Engine) {
	casbinService, err := casbinModels.NewCasbinService(mysql.DB)
	if err != nil {
		fmt.Println("Setup models.NewCasbinService()  err")
	}

	user := router.Group("userService")
	{
		//1.像前端返回验证码
		user.POST("/code", login.RidCode)
		//2.接收数据查询是否登录成功
		user.Use(middleWare.NewCasbinAuth(casbinService))

		user.POST("/hlogin", login.HLogin)
	}

}
