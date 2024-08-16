package teacherManage

//import (
//	"github.com/gin-gonic/gin"
//	"go.uber.org/zap"
//	"studentGrow/models/jrx_model"
//	"studentGrow/pkg/response"
//	"studentGrow/service"
//)
//
//func QueryTeacher(c *gin.Context) {
//	var queryParama jrx_model.QueryTeacherParamStruct
//	err := c.Bind(&queryParama)
//	if err != nil {
//		response.ResponseError(c, response.ParamFail)
//		zap.L().Error("teacherManage.QueryTeacher() c.Bind() err : ", zap.Error(err))
//		return
//	}
//
//	// 业务
//	teacherList, allTeacherCount, err := service.QueryTeacher(queryParama)
//	if err != nil {
//		response.ResponseError(c, 500)
//		zap.L().Error("teacherManage.QueryTeacher() service.QueryTeacher() err : ", zap.Error(err))
//		return
//	}
//
//	response.ResponseSuccess(c)
//
//}
