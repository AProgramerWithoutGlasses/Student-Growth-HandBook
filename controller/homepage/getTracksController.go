package homepage

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"studentGrow/models"
	"studentGrow/models/jrx_model"
	"studentGrow/pkg/response"
	"studentGrow/service"
)

func GetTracksControl(c *gin.Context) {
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

	claim, exist := c.Get("claim")
	if !exist {
		response.ResponseError(c, response.TokenError)
		zap.L().Error("token错误")
		return
	}
	username := claim.(*models.Claims).Username

	// 业务
	Tracks, err := service.GetTracksService(input.Page, input.Limit, username)
	if err != nil {
		response.ResponseError(c, response.ServerErrorCode)
		zap.L().Error(err.Error())
		return
	}

	// 响应
	output := struct {
		InterInfo []jrx_model.HomepageTrack `json:"inter_info"`
	}{
		InterInfo: Tracks,
	}

	response.ResponseSuccess(c, output)
}
