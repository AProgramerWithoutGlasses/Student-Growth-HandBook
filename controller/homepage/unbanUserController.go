package homepage

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"studentGrow/pkg/response"
	"studentGrow/service"
	token2 "studentGrow/utils/token"
)

func UnbanUserControl(c *gin.Context) {
	// 接收数据
	input := struct {
		UnbanUsername string `json:"unban_username"`
	}{}
	err := c.Bind(&input)
	if err != nil {
		response.ResponseError(c, response.ParamFail)
		zap.L().Error(err.Error())
		return
	}

	token := c.GetHeader("token")
	username, err := token2.GetUsername(token)
	if err != nil {
		response.ResponseError(c, response.ParamFail)
		zap.L().Error(err.Error())
		return
	}

	// 业务
	err = service.UnbanHomepageUserService(input.UnbanUsername, username)
	if err != nil {
		response.ResponseError(c, response.ServerErrorCode)
		return
	}

	// 响应
	response.ResponseSuccess(c, nil)
}
