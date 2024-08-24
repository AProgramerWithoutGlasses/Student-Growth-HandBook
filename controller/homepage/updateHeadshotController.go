package homepage

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"studentGrow/pkg/response"
	"studentGrow/service"
	token2 "studentGrow/utils/token"
)

func UpdateHeadshotControl(c *gin.Context) {
	//input := struct {
	//	UserHeadshot string `form:"user_headshot"`
	//}{}
	//err := c.Bind(&input)
	//if err != nil {
	//	response.ResponseError(c, response.ParamFail)
	//	zap.L().Error(err.Error())
	//	return
	//}

	// 获取上传的文件
	file, err := c.FormFile("file")
	if err != nil {
		response.ResponseError(c, response.ParamFail)
		zap.L().Error(err.Error())
		return
	}
	fmt.Println("成功接收到file")

	token := c.GetHeader("token")
	username, err := token2.GetUsername(token)
	if err != nil {
		response.ResponseError(c, response.ParamFail)
		zap.L().Error(err.Error())
		return
	}
	fmt.Println("成功接收到token")

	// 业务
	err = service.UpdateHeadshotService(file, username)
	if err != nil {
		response.ResponseError(c, response.ServerErrorCode)
		zap.L().Error(err.Error())
		return
	}

	fmt.Println("成功执行完业务")

	// 响应
	response.ResponseSuccess(c, struct{}{})
}
