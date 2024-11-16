package routes

import (
	"github.com/gin-gonic/gin"
	"studentGrow/controller/routesJoinAudit"
	"studentGrow/utils/token"
)

func RoutsJoinAudit(router *gin.Engine) {

	r := router.Group("/routesJoinAudit")
	r.Use(token.AuthMiddleware())
	r.POST("/isOpen", routesJoinAudit.OpenMsg)
	r.GET("/StuForm", routesJoinAudit.GetStudForm)
	r.POST("/StuForm", routesJoinAudit.SaveStudForm)
	r.GET("/StudFile", routesJoinAudit.GetStuFile)
	r.POST("/StudFile", routesJoinAudit.SaveStuFile)
	r.POST("/DelStudFile", routesJoinAudit.DelStuFile)
	r.GET("/activity", routesJoinAudit.GetActivityList)
	r.POST("/activity", routesJoinAudit.SaveActivityMsg)
	r.POST("/delActivity", routesJoinAudit.DelActivityMsg)
	r.GET("/activityClass", routesJoinAudit.ClassApplicationList)
	r.POST("/activityClass", routesJoinAudit.ClassApplicationManager)
	r.GET("/activityRuler", routesJoinAudit.ActivityRulerList)
	r.POST("/activityRuler", routesJoinAudit.ActivityRulerManager)
	r.POST("/activityMaterial", routesJoinAudit.ActivityOrganizerMaterialManager)
	r.GET("/activityTrain", routesJoinAudit.ActivityOrganizerTrainList)
	r.POST("/activityTrain", routesJoinAudit.ActivityOrganizerTrainManager)
	r.POST("/saveTrainScore", routesJoinAudit.SaveTrainScore)
}
