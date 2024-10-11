package stuManage

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"studentGrow/dao/mysql"
	"studentGrow/models"
	"studentGrow/models/gorm_model"
	"studentGrow/models/jrx_model"
	"studentGrow/pkg/response"
	"studentGrow/service"
)

// 添加单个学生
func AddSingleStuContro(c *gin.Context) {
	// 根据token拿到登陆者自己的信息
	var myMes jrx_model.MyTokenMes
	var err error

	claim, exist := c.Get("claim")
	if !exist {
		response.ResponseError(c, response.TokenError)
		zap.L().Error("token错误")
		return
	}
	myMes.MyUsername = claim.(*models.Claims).Username

	myMes.MyId, err = mysql.GetIdByUsername(myMes.MyUsername)
	if err != nil {
		response.ResponseError(c, response.ParamFail)
		zap.L().Error(err.Error())
		return
	}

	myMes.MyClass, err = mysql.GetClassById(myMes.MyId)
	if err != nil {
		response.ResponseError(c, response.ParamFail)
		zap.L().Error(err.Error())
		return
	}

	//myMes.MyRole, err = token2.GetRole(token)
	myMes.MyRole = claim.(*models.Claims).Role

	// 接收
	input := struct {
		gorm_model.User
		Class  string `json:"class" binding:"required, len=9"`
		Gender string `json:"gender" binding:"required, one of=男 女"`
	}{}
	err = c.ShouldBindJSON(&input)
	if err != nil {
		response.ResponseError(c, response.ParamFail)
		zap.L().Error(err.Error())
		return
	}

	// 业务
	err = service.AddStuService(input, myMes)
	if err != nil {
		response.ResponseErrorWithMsg(c, response.ServerErrorCode, err.Error())
		zap.L().Error(err.Error())
		return
	}

	// 成功响应
	response.ResponseSuccessWithMsg(c, input.Name+" 信息添加成功！", nil)
}
