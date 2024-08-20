package routes

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"studentGrow/controller/stuManage"
	"studentGrow/dao/mysql"
	"studentGrow/models/casbinModels"
	"studentGrow/utils/middleWare"
)

// 后台学生管理路由
func routesStudentManage(r *gin.Engine) {
	rs := r.Group("/stuManage")

	casbinService, err := casbinModels.NewCasbinService(mysql.DB)
	if err != nil {
		zap.L().Error("routesArticle() routes.routesArticle.NewCasbinService err=", zap.Error(err))
		return
	}

	// rs.POST("/queryPageStudent", stuManage.QueryPageStuContro)
	rs.POST("/addSingleStudent", stuManage.AddSingleStuContro)
	rs.POST("/addMultipleStudent", stuManage.AddMultipleStuControl)
	rs.POST("/queryStudent/class", middleWare.NewCasbinAuth(casbinService), stuManage.QueryStuContro)

	rs.POST("/deleteStudent", stuManage.DeleteStuControl)
	rs.POST("/setStudentManager", stuManage.SetStuManagerControl)
	rs.POST("/editStudent", stuManage.EditStuControl)
	rs.POST("/banStudent", stuManage.BanUserControl)
	rs.POST("/outputMultipleStudent", stuManage.OutputMultipleStuControl)
}
