package routes

import (
	"github.com/gin-gonic/gin"
	"studentGrow/controller/stuManage"
	"studentGrow/controller/student"
	"studentGrow/logger"
	"studentGrow/utils/middleWare"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.Use(middleWare.CORSMiddleware())
	// 勋
	r.POST("/student/getSelfCotnent", student.GetSelfContentContro)
	r.POST("/student/updateSelfContent", student.UpdateSelfContentContro)
	r.POST("/stuManage/queryStudent/class", stuManage.QueryStuContro)
	r.GET("/stuManage/queryPageStudent", stuManage.QueryPageStuContro)
	r.POST("/stuManage/addSingleStudent", stuManage.AddSingleStuContro)
	r.POST("/stuManage/addMultipleStudent", stuManage.AddMultipleStuContro)
	r.POST("/stuManage/deleteStudent", stuManage.DeleteStuControl)
	r.POST("/stuManage/setStudentManager", stuManage.StuManagerControl)
	/*	r.POST("/stuManage/banStudent", stuManage.BanStuContro)
		r.POST("/stuManage/editStudent", stuManage.EditStuControl)

		r.POST("/stuManage/outputMultipleStudent", stuManage.outputMultipleStuContro)
	*/

	return r
}
