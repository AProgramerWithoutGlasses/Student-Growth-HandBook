package homepage

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"studentGrow/models"
	"studentGrow/models/jrx_model"
	"studentGrow/pkg/response"
	"studentGrow/service"
)

func GetStarControl(c *gin.Context) {
	// 接收
	input := struct {
		Page  int `json:"page" binding:"required"`
		Limit int `json:"limit" binding:"required"`
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

	claim, exist := c.Get("claim")
	if !exist {
		response.ResponseError(c, response.TokenError)
		zap.L().Error("token错误")
		return
	}
	username := claim.(*models.Claims).Username

	// 校验
	if input.Page <= 0 || input.Limit <= 0 {
		response.ResponseError(c, response.ParamFail)
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

	if homepageStarList == nil || len(homepageStarList) == 0 {
		output.History = []jrx_model.HomepageArticleHistoryStruct{}
	}

	response.ResponseSuccess(c, output)
}
