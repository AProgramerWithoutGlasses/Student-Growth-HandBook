package routes

import (
	"github.com/gin-gonic/gin"
	_ "studentGrow/controller/article"
	"studentGrow/controller/student"
	"studentGrow/logger"
	"studentGrow/utils/middleWare"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.Use(middleWare.CORSMiddleware()) // 跨域中间件

	routesArticle(r)
	routesTopic(r)
	routesMsg(r)

	RoutesXue(r)

	// 勋
	r.POST("/student/getSelfCotnent", student.GetSelfContentContro)
	r.POST("/student/updateSelfContent", student.UpdateSelfContentContro)
	/*	r.POST("/stuManage/queryStudent", stuManage.QueryStuContro)
		r.POST("/stuManage/addSingleStudent", stuManage.AddSingleStuContro)
		r.POST("/stuManage/addMultipleStudent", stuManage.AddMultipleStuContro)
		r.POST("/stuManage/deleteStudent", stuManage.DeleteStuContro)
		r.POST("/stuManage/banStudent", stuManage.BanStuContro)
		r.POST("/stuManage/editStudent", stuManage.EditStuContro)
		r.POST("/stuManage/setStudentManager", stuManage.setStuManagerContro)
		r.POST("/stuManage/outputMultipleStudent", stuManage.outputMultipleStuContro)
	*/

	return r
}
