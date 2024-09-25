package article

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	myErr "studentGrow/pkg/error"
	res "studentGrow/pkg/response"
	"studentGrow/service/article"
	"studentGrow/utils/readMessage"
	"studentGrow/utils/token"
)

// CollectArticleController 收藏文章
func CollectArticleController(c *gin.Context) {
	// 获取文章id
	json, err := readMessage.GetJsonvalue(c)
	if err != nil {
		zap.L().Error("CollectArticleController() controller.article.CollectController.GetJsonvalue err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}
	aid, err := json.GetInt("article_id")
	if err != nil {
		zap.L().Error("CollectArticleController() controller.article.CollectController.GetInt err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}

	// 通过token获取username
	username, err := token.GetUsername(c.GetHeader("token"))
	if err != nil {
		zap.L().Error("CollectArticleController() controller.article.CollectController.GetUsername err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}

	// 收藏
	err = article.CollectOrNotService(strconv.Itoa(aid), username)
	if err != nil {
		zap.L().Error("CollectArticleController() controller.article.CollectController.CollectOrNotService err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}

	res.ResponseSuccess(c, struct{}{})

}
