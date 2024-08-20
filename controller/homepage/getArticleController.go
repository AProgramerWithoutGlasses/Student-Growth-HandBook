package homepage

import (
	"github.com/gin-gonic/gin"
)

func GetArticleControl(c *gin.Context) {
	// 接收
	//input := struct {
	//	Page  int `json:"page"`
	//	Limit int `json:"limit"`
	//}{}
	//err := c.BindJSON(&input)
	//if err != nil {
	//	response.ResponseError(c, response.ParamFail)
	//	zap.L().Error(err.Error())
	//	return
	//}
	//
	//token := c.GetHeader("token")
	//username, err := token2.GetUsername(token)
	//if err != nil {
	//	response.ResponseError(c, response.ParamFail)
	//	zap.L().Error(err.Error())
	//	return
	//}

	//// 业务
	//articleList, err := service.GetArticleService(input.Page, input.Limit, username)
	//if err != nil {
	//	response.ResponseError(c, response.ServerErrorCode)
	//	zap.L().Error(err.Error())
	//	return
	//}
	//
	//// 响应
	//output := struct {
	//	History []jrx_model.HomepageArticleHistoryStruct `json:"star"`
	//}{
	//	History: homepageStarList,
	//}
	//
	//response.ResponseSuccess(c, output)
}
