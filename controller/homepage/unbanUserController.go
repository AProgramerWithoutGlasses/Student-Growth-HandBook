package homepage

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"studentGrow/models"
	"studentGrow/pkg/response"
	"studentGrow/service"
)

func UnbanUserControl(c *gin.Context) {
	// 接收数据
	input := struct {
		UnbanUsername string `json:"unban_username"`
	}{}
	err := c.BindJSON(&input)
	if err != nil {
		response.ResponseError(c, response.ParamFail)
		zap.L().Error(err.Error())
		return
	}

	claim, exist := c.Get("claim")
	if !exist {
		response.ResponseError(c, response.TokenError)
		zap.L().Error("token错误")
		return
	}
	username := claim.(*models.Claims).Username

	// 业务
	err = service.UnbanHomepageUserService(input.UnbanUsername, username)
	if err != nil {
		response.ResponseError(c, response.ServerErrorCode)
		return
	}

	// 响应
	response.ResponseSuccess(c, struct{}{})
}
