package homepage

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"studentGrow/models/jrx_model"
	"studentGrow/pkg/response"
	"studentGrow/service"
)

func GetHistoryControl(c *gin.Context) {
	// 接收
	input := struct {
		Page     int    `json:"page"`
		Limit    int    `json:"limit"`
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
	homepageArticleHistoryList, err := service.GetHistoryService(input.Page, input.Limit, input.Username)
	if err != nil {
		response.ResponseError(c, response.ServerErrorCode)
		zap.L().Error(err.Error())
		return
	}

	// 响应
	output := struct {
		History []jrx_model.HomepageArticleHistoryStruct `json:"history"`
	}{
		History: homepageArticleHistoryList,
	}

	response.ResponseSuccess(c, output)
}
