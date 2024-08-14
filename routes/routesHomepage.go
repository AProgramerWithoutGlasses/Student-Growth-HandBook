package routes

import (
	"github.com/gin-gonic/gin"
	"studentGrow/controller/homepage"
)

// 前台个人主页
func routesHomepage(r *gin.Engine) {
	rt := r.Group("/user")

	rt.POST("/getSelfCotnent", homepage.GetSelfContentContro)
	rt.POST("/updateSelfContent", homepage.UpdateSelfContentContro)
}
