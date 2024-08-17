package article

import (
	"fmt"
	"github.com/gin-gonic/gin"
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
		fmt.Println("CollectArticleController() controller.article.CollectController.GetJsonvalue err=", err)
		myErr.CheckErrors(err, c)
		return
	}
	aid, err := json.GetInt("article_id")
	if err != nil {
		fmt.Println("CollectArticleController() controller.article.CollectController.GetInt err=", err)
		myErr.CheckErrors(err, c)
		return
	}

	// 通过token获取username
	username, err := token.GetUsername(c.GetHeader("token"))
	if err != nil {
		fmt.Println("CollectArticleController() controller.article.CollectController.GetUsername err=", err)
		myErr.CheckErrors(err, c)
		return
	}

	// 收藏
	err = article.CollectOrNotService(strconv.Itoa(aid), username)
	if err != nil {
		fmt.Println("CollectArticleController() controller.article.CollectController.SelectService err=", err)
		myErr.CheckErrors(err, c)
		return
	}

	res.ResponseSuccess(c, nil)

}

// CancelCollectArticleController 取消收藏
//func CancelCollectArticleController(c *gin.Context) {
//	// 获取文章id
//	json, err := readMessage.GetJsonvalue(c)
//	if err != nil {
//		fmt.Println("CollectArticleController() controller.article.CollectController.GetJsonvalue err=", err)
//		myErr.CheckErrors(err, c)
//		return
//	}
//	aid, err := json.GetInt("article_id")
//	if err != nil {
//		fmt.Println("CollectArticleController() controller.article.CollectController.GetInt err=", err)
//		myErr.CheckErrors(err, c)
//		return
//	}
//
//	// 通过token获取username
//	username, err := token.GetUsername(c.GetHeader("token"))
//	if err != nil {
//		fmt.Println("CollectArticleController() controller.article.CollectController.GetUsername err=", err)
//		myErr.CheckErrors(err, c)
//		return
//	}
//
//	// 取消收藏
//	err = article.CancelSelectService(strconv.Itoa(aid), username)
//	fmt.Println("CollectArticleController() controller.article.CollectController.GetUsername err=", err)
//	myErr.CheckErrors(err, c)
//	if err != nil {
//		return
//	}
//
//	res.ResponseSuccess(c, nil)
//}

// GetArticleListForSelectController 获取收藏文章列表
//func GetArticleListForSelectController(c *gin.Context) {
//	// 通过token获取username
//	username, err := token.GetUsername(c.GetHeader("token"))
//	if err != nil {
//		fmt.Println("GetArticleListForSelectController() controller.article.CollectController.GetUsername err=", err)
//		myErr.CheckErrors(err, c)
//		return
//	}
//
//	// 获取收藏文章列表
//	set, err := redis.GetUserSelectionSet(username)
//	if err != nil {
//		fmt.Println("GetArticleListForSelectController() controller.article.CollectController.GetUserSelectionSet err=", err)
//		myErr.CheckErrors(err, c)
//		return
//	}
//
//	res.ResponseSuccess(c, set)
//}
