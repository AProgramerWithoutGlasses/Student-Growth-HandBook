package routes

import (
	"github.com/gin-gonic/gin"
	"studentGrow/controller/homepage"
)

// 前台个人主页
func routesHomepage(r *gin.Engine) {
	rh := r.Group("/user")

	rh.POST("/getSelfCotnent", homepage.GetSelfContentContro)
	rh.POST("/updateSelfContent", homepage.UpdateSelfContentContro)
	rh.GET("/profiles", homepage.GetHomepageMesContro)
}
