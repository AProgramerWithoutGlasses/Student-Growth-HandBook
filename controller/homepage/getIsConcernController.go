package homepage

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"studentGrow/pkg/response"
	"studentGrow/service"
	token2 "studentGrow/utils/token"
)

func GetIsConcernControl(c *gin.Context) {
	input := struct {
		//Username      string `json:"username"`
		OtherUsername string `json:"other_username"`
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

	isConcern, err := service.GetIsConcernService(username, input.OtherUsername)
	if err != nil {
		response.ResponseError(c, response.ServerErrorCode)
		zap.L().Error(err.Error())
		return
	}

	output := struct {
		IsConcern bool `json:"is_concern"`
	}{
		IsConcern: isConcern,
	}
	response.ResponseSuccess(c, output)

}
