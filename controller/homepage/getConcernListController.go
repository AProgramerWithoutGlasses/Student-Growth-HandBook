package homepage

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"studentGrow/models/jrx_model"
	"studentGrow/pkg/response"
	"studentGrow/service"
)

func GetConcernListControl(c *gin.Context) {
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

	//token := c.GetHeader("token")
	//username, err := token2.GetUsername(token)
	//if err != nil {
	//	response.ResponseError(c, response.ParamFail)
	//	zap.L().Error(err.Error())
	//	return
	//}

	// 业务
	userConcern, err := service.GetConcernListService(input.Username)
	if err != nil {
		response.ResponseError(c, response.ServerErrorCode)
		zap.L().Error(err.Error())
		return
	}

	// 响应
	output := struct {
		UserConcern []jrx_model.HomepageFanStruct `json:"user_concern"`
	}{
		UserConcern: userConcern,
	}

	response.ResponseSuccess(c, output)

}
