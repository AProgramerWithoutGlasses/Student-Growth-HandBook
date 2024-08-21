package homepage

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"studentGrow/models/jrx_model"
	"studentGrow/pkg/response"
	"studentGrow/service"
	token2 "studentGrow/utils/token"
)

func GetFansListControl(c *gin.Context) {
	// 接收
	input := struct {
		Username string `json:"username"`
	}{}
	err := c.BindJSON(&input)
	if err != nil {
		response.ResponseError(c, response.ParamFail)
		zap.L().Error(err.Error())
		return
	}

	/////
	token := c.GetHeader("token")
	tokenUsername, err := token2.GetUsername(token)
	if err != nil {
		response.ResponseError(c, response.ParamFail)
		zap.L().Error(err.Error())
		return
	}

	// 业务
	userfans, err := service.GetFansListService(input.Username, tokenUsername)
	if err != nil {
		response.ResponseError(c, response.ServerErrorCode)
		zap.L().Error(err.Error())
		return
	}

	// 响应
	output := struct {
		Userfans []jrx_model.HomepageFanStruct `json:"user_fans"`
	}{
		Userfans: userfans,
	}

	response.ResponseSuccess(c, output)

}
