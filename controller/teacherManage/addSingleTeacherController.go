package teacherManage

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strings"
	"studentGrow/models/gorm_model"
	"studentGrow/pkg/response"
	"studentGrow/service"
)

func AddSingleTeacherControl(c *gin.Context) {
	var addSingleTeacherReqStruct gorm_model.User
	err := c.ShouldBindJSON(&addSingleTeacherReqStruct)
	if err != nil {
		response.ResponseError(c, response.ParamFail)
		zap.L().Error(err.Error())
		return
	}

	err = service.AddSingleTeacherService(addSingleTeacherReqStruct)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			response.ResponseErrorWithMsg(c, 500, addSingleTeacherReqStruct.Name+" 该用户已存在")
			zap.L().Error(err.Error())
			return
		} else {
			response.ResponseErrorWithMsg(c, 500, err.Error())
			zap.L().Error(err.Error())
			return
		}

	}

	response.ResponseSuccess(c, addSingleTeacherReqStruct.Name+" 添加成功")

}
