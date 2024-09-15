package homepage

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"studentGrow/models/jrx_model"
	"studentGrow/pkg/response"
	"studentGrow/service"
	token2 "studentGrow/utils/token"
)

func GetHistoryControl(c *gin.Context) {
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
	homepageArticleHistoryList, err := service.GetHistoryService(input.Page, input.Limit, username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

		} else {
			response.ResponseError(c, response.ServerErrorCode)
			zap.L().Error(err.Error())
			return
		}
	}

	// 响应
	output := struct {
		History []jrx_model.HomepageArticleHistoryStruct `json:"history"`
	}{
		History: homepageArticleHistoryList,
	}

	if homepageArticleHistoryList == nil || len(homepageArticleHistoryList) == 0 {
		output.History = []jrx_model.HomepageArticleHistoryStruct{}
	}

	response.ResponseSuccess(c, output)
}
