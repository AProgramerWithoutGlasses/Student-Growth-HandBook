package teacherManage

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"studentGrow/models/jrx_model"
	"studentGrow/pkg/response"
	"studentGrow/service"
)

type ResponseStruct struct {
	TeacherInfo     []jrx_model.QueryTeacherResStruct `json:"teacherInfo"`
	AllTeacherCount int                               `json:"allTeacherCount"`
}

func QueryTeacherControl(c *gin.Context) {
	// 接收
	var queryParama jrx_model.QueryTeacherParamStruct
	err := c.BindJSON(&queryParama)
	if err != nil {
		zap.L().Error("teacherManage.QueryTeacher() c.Bind() err : ", zap.Error(err))
		response.ResponseError(c, response.ServerErrorCode)
	}

	fmt.Printf("queryParama: %+v\n", queryParama)

	// 业务
	teacherList, allTeacherCount, err := service.QueryTeacher(queryParama)
	if err != nil {
		response.ResponseError(c, 500)
		zap.L().Error("teacherManage.QueryTeacher() service.QueryTeacher() err : ", zap.Error(err))
		return
	}

	// 响应
	responseStruct := ResponseStruct{
		TeacherInfo:     teacherList,
		AllTeacherCount: allTeacherCount,
	}

	response.ResponseSuccess(c, responseStruct)

}
