package homepage

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"studentGrow/pkg/response"
	"studentGrow/service"
	_ "studentGrow/utils/token"
)

func GetHomepageMesContro(c *gin.Context) {
	// 接收
	input := struct {
		Username string `json:"username"`
	}{}
	err := c.BindJSON(&input)
	if err != nil {
		response.ResponseError(c, response.ParamFail)
		return
	}

	// 校验
	if input.Username == "" {
		response.ResponseErrorWithMsg(c, response.ParamFail, "请求参数为空")
		return
	}

	// 业务
	// 获取角色
	/*	token := c.GetHeader("token")
		role,err := token2.GetRole(token)	// class, grade(1-4), collge, superman */

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
