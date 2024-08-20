package routes

import (
	"github.com/gin-gonic/gin"
	"studentGrow/controller/homepage"
	"studentGrow/utils/token"
)

// 前台个人主页
func routesHomepage(r *gin.Engine) {
	rh := r.Group("/user")

	rh.GET("/profiles_get", homepage.GetMesControl)
	rh.POST("/userHeadshot_update", token.AuthMiddleware(), homepage.UpdateHeadshotControl)
	rh.POST("/selfCotnent_get", homepage.GetSelfContentContro)
	rh.POST("/selfContent_update", homepage.UpdateSelfContentContro)
	rh.POST("/userMotto_update", homepage.UpdateHomepageMottoControl)
	rh.POST("/userPhone_update", homepage.UpdatePhoneNumberControl)
	rh.POST("/userEmail_update", homepage.UpdateEmailControl)
	rh.GET("/userData_get", homepage.GetUserDataControl)
	rh.GET("/fans_get", homepage.GetFansListControl)
	rh.GET("/concern_get", homepage.GetConcernListControl)
	rh.POST("/concern_change", homepage.ChangeConcernControl)
	rh.GET("/history_get", homepage.GetHistoryControl)
	// 我的足迹
	rh.GET("/star_get", homepage.GetStarControl)
	rh.GET("/class_get", homepage.GetClassControl)
	// 获取积分统计
	rh.GET("/article_get", homepage.GetArticleControl)
	rh.POST("/ban", homepage.BanUserControl)
	rh.POST("/unban", homepage.UnbanUserControl)

}
