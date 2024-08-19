package homepage

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"studentGrow/pkg/response"
	"studentGrow/service"
	token2 "studentGrow/utils/token"
)

func GetMesControl(c *gin.Context) {
	// 接收
	token := c.GetHeader("token")
	username, err := token2.GetUsername(token)
	if err != nil {
		response.ResponseError(c, response.ParamFail)
		zap.L().Error(err.Error())
		return
	}

	// 校验
	if username == "" {
		response.ResponseErrorWithMsg(c, response.ParamFail, "请求参数为空")
		return
	}

	homepageMesStruct, err := service.GetHomepageMesService(input.Username)
	if err != nil {
		response.ResponseError(c, response.ServerErrorCode)
		zap.L().Error("homepage.GetHomepageMesContro() service.GetHomepageMesService() failed : ", zap.Error(err))
		return
	}

	fmt.Println("honemes is ： ", *homepageMesStruct)

	// 响应
	response.ResponseSuccess(c, &homepageMesStruct)

}
