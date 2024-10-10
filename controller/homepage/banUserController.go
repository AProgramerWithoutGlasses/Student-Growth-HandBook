package homepage

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"studentGrow/models"
	"studentGrow/pkg/response"
	"studentGrow/service"
)

func BanUserControl(c *gin.Context) {
	// 接收数据
	input := struct {
		BanTime     int    `json:"ban_time"`
		BanUsername string `json:"ban_username"`
	}{}
	err := c.BindJSON(&input)
	if err != nil {
		fmt.Println("stuManage.BanStuControl() c.Bind() err : ", err)
		response.ResponseErrorWithMsg(c, 500, "stuManage.BanStuControl() c.Bind() failed : "+err.Error())
		return
	}
	//
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

	err = service.BanHomepageUserService(input.BanUsername, input.BanTime, username)
	if err != nil {
		response.ResponseErrorWithMsg(c, 500, "stuManage.BanStuControl() service.BanUserService() failed : "+err.Error())
		return
	}

	response.ResponseSuccess(c, struct{}{})
}
