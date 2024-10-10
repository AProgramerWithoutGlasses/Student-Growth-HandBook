package homepage

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"studentGrow/models"
	"studentGrow/pkg/response"
	"studentGrow/service"
)

func UpdateEmailControl(c *gin.Context) {
	// 接收请求
	input := struct {
		UserEmail string `json:"user_email"`
	}{}
	err := c.BindJSON(&input)
	if err != nil {
		response.ResponseError(c, response.ParamFail)
		zap.L().Error(err.Error())
		return
	}

	// 获取角色
	claim, exist := c.Get("claim")
	if !exist {
		response.ResponseError(c, response.TokenError)
		zap.L().Error("token错误")
		return
	}
	username := claim.(*models.Claims).Username

	// 业务
	err = service.UpdateHomepageEmailService(username, input.UserEmail)
	if err != nil {
		response.ResponseError(c, response.ServerErrorCode)
		zap.L().Error(err.Error())
		return
	}

	// 响应
	response.ResponseSuccess(c, struct{}{})
}
