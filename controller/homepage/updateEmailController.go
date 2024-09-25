package homepage

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"studentGrow/pkg/response"
	"studentGrow/service"
	token2 "studentGrow/utils/token"
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
	token := c.GetHeader("token")
	username, err := token2.GetUsername(token) // class, grade(1-4), collge, superman
	if err != nil {
		response.ResponseError(c, response.ServerErrorCode)
		zap.L().Error(err.Error())
		return
	}

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
