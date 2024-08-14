package routes

import (
	"github.com/gin-gonic/gin"
	"studentGrow/controller/stu_manage"
)

// 后台学生管理路由
func routesStudentManage(r *gin.Engine) {
	rs := r.Group("/stu_manage")

	rs.POST("/queryPageStudent", stu_manage.QueryPageStuContro)
	rs.POST("/addSingleStudent", stu_manage.AddSingleStuContro)
	rs.POST("/addMultipleStudent", stu_manage.AddMultipleStuControl)
	rs.POST("/queryStudent/class", stu_manage.QueryStuContro)
	rs.POST("/deleteStudent", stu_manage.DeleteStuControl)
	rs.POST("/setStudentManager", stu_manage.SetStuManagerControl)
	rs.POST("/editStudent", stu_manage.EditStuControl)
	rs.POST("/banStudent", stu_manage.BanStuControl)
	rs.POST("/outputMultipleStudent", stu_manage.OutputMultipleStuControl)
}
