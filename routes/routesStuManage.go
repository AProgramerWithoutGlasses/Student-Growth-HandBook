package routes

import (
	"github.com/gin-gonic/gin"
	"studentGrow/controller/stuManage"
)

// 后台学生管理路由
func routesStudentManage(r *gin.Engine) {
	rs := r.Group("/stuManage")

	// rs.POST("/queryPageStudent", stuManage.QueryPageStuContro)
	rs.POST("/addSingleStudent", stuManage.AddSingleStuContro)
	rs.POST("/addMultipleStudent", stuManage.AddMultipleStuControl)
	rs.POST("/queryStudent/class", stuManage.QueryStuContro)
	rs.POST("/deleteStudent", stuManage.DeleteStuControl)
	rs.POST("/setStudentManager", stuManage.SetStuManagerControl)
	rs.POST("/editStudent", stuManage.EditStuControl)
	rs.POST("/banStudent", stuManage.BanUserControl)
	rs.POST("/outputMultipleStudent", stuManage.OutputMultipleStuControl)
}
