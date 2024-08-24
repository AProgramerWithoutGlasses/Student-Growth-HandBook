package homepage

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"studentGrow/models/jrx_model"
	"studentGrow/pkg/response"
	"studentGrow/service"
	token2 "studentGrow/utils/token"
)

func GetStarControl(c *gin.Context) {
	// 接收
	input := struct {
		Page  int `json:"page"`
		Limit int `json:"limit"`
	}{}
	err := c.BindJSON(&input)
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
	homepageStarList, err := service.GetStarService(input.Page, input.Limit, username)
	if err != nil {
		response.ResponseError(c, response.ServerErrorCode)
		zap.L().Error(err.Error())
		return
	}

	// 响应
	output := struct {
		History []jrx_model.HomepageArticleHistoryStruct `json:"star"`
	}{
		History: homepageStarList,
	}

	if output.History == nil || len(output.History) == 0 {
		response.ResponseSuccess(c, "")
		return
	}

	response.ResponseSuccess(c, output)
}
