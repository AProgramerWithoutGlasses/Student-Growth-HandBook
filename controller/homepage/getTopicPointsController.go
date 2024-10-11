package homepage

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"studentGrow/models"
	"studentGrow/pkg/response"
	"studentGrow/service"
	_ "studentGrow/utils/token"
)

func GetTopicPointsControl(c *gin.Context) {
	//// 接收
	//input := struct {
	//	Username string `json:"username"`
	//}{}
	//err := c.BindJSON(&input)
	//if err != nil {
	//	response.ResponseError(c, response.ParamFail)
	//	zap.L().Error(err.Error())
	//	return
	//} //HH

	//token := c.GetHeader("token")
	//username, err := token2.GetUsername(token)
	//if err != nil {
	//	response.ResponseError(c, response.ServerErrorCode)
	//	zap.L().Error(err.Error())
	//	return
	//}
	claim, exist := c.Get("claim")
	if !exist {
		response.ResponseError(c, response.TokenError)
		zap.L().Error("token错误")
		return
	}
	username := claim.(*models.Claims).Username

	// 业务
	topicPointStruct, err := service.GetTopicPointsService(username)
	if err != nil {
		response.ResponseErrorWithMsg(c, response.ServerErrorCode, err.Error())
		zap.L().Error(err.Error())
		return
	}

	// 响应
	response.ResponseSuccess(c, topicPointStruct)
}
