package homepage

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"studentGrow/pkg/response"
	"studentGrow/service"
)

func GetTopicPointsControl(c *gin.Context) {
	// 接收
	input := struct {
		Username string `json:"username"`
	}{}
	err := c.BindJSON(&input)
	if err != nil {
		response.ResponseError(c, response.ParamFail)
		zap.L().Error(err.Error())
		return
	} //HH

	//token := c.GetHeader("token")
	//username, err := token2.GetUsername(token)
	//if err != nil {
	//	response.ResponseError(c, response.ServerErrorCode)
	//	zap.L().Error(err.Error())
	//	return
	//}

	// 业务
	topicPointStruct, err := service.GetTopicPointsService(input.Username)
	if err != nil {
		response.ResponseErrorWithMsg(c, response.ServerErrorCode, err.Error())
		zap.L().Error(err.Error())
		return
	}

	// 响应
	response.ResponseSuccess(c, topicPointStruct)
}
