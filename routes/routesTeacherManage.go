package routes

import (
	"github.com/gin-gonic/gin"
	"studentGrow/controller/teacherManage"
)

// 后台老师管理路由
func routesTeacherManage(r *gin.Engine) {
	rt := r.Group("/teacherManage")

	rt.GET("/queryTeacher", teacherManage.QueryTeacherControl)
	rt.POST("/addSingleTeacher", teacherManage.AddSingleTeacherControl)
	rt.POST("/addMultipleTeacher", teacherManage.AddMultipleTeacherControl)
	//rt.POST("/queryStudent/class", stuManage.QueryStuContro)
	//rt.POST("/deleteStudent", stuManage.DeleteStuControl)
	//rt.POST("/setStudentManager", stuManage.SetStuManagerControl)
	//rt.POST("/editStudent", stuManage.EditStuControl)
	//rt.POST("/banStudent", stuManage.BanUserControl)
	//rt.POST("/outputMultipleStudent", stuManage.OutputMultipleStuControl)
}
