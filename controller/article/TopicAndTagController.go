package article

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	myErr "studentGrow/pkg/error"
	res "studentGrow/pkg/response"
	"studentGrow/service/article"
	readUtil "studentGrow/utils/readMessage"
)

// AddTopicsController 添加话题
func AddTopicsController(c *gin.Context) {
	//获取前端发送的数据
	json, err := readUtil.GetJsonvalue(c)

	if err != nil {
		fmt.Println("AddTopicsController() controller.article.getArticle.GetJsonvalue err=", err)
		myErr.CheckErrors(err, c)
		return
	}
	err = article.AddTopicsService(json)
	if err != nil {
		fmt.Println("AddTopicsController() controller.article.getArticle.AddTopicsService err=", err)
		myErr.CheckErrors(err, c)
		return
	}

	res.ResponseSuccess(c, nil)

}

// GetAllTopicsController 获取所有话题
func GetAllTopicsController(c *gin.Context) {

	result, err := article.GetAllTopicsService()
	if err != nil {
		fmt.Println("GetAllTopicsController() controller.article.getArticle.GetAllTopicsService err=", err)
		if err != nil {
			myErr.CheckErrors(err, c)
			return
		}
	}

	res.ResponseSuccess(c, result)

}

// AddTagsByTopicController 添加标签
func AddTagsByTopicController(c *gin.Context) {
	//获取前端发送的数据
	json, err := readUtil.GetJsonvalue(c)
	if err != nil {
		fmt.Println("AddArticleTagsController() controller.article.getArticle.GetJsonvalue err=", err)
		myErr.CheckErrors(err, c)
		return
	}

	err = article.AddTagsByTopicService(json)
	if err != nil {
		fmt.Println("AddArticleTagsController() controller.article.getArticle.AddTagsByTopicService err=", err)
		if err != nil {
			myErr.CheckErrors(err, c)
			return
		}
		res.ResponseError(c, res.ServerErrorCode)
		return
	}

	res.ResponseSuccess(c, nil)
}

// GetTagsByTopicController 获取标签
func GetTagsByTopicController(c *gin.Context) {
	//获取前端发送的数据
	json, err := readUtil.GetJsonvalue(c)
	if err != nil {
		fmt.Println("AddArticleTagsController() controller.article.getArticle.GetJsonvalue err=", err)
		myErr.CheckErrors(err, c)
		return
	}

	result, err := article.GetTagsByTopicService(json)
	if err != nil {
		fmt.Println("AddArticleTagsController() controller.article.getArticle.GetTagsByTopicService err=", err)
		if errors.Is(err, myErr.NotFoundError()) {
			myErr.CheckErrors(err, c)
			return
		}
		res.ResponseError(c, res.ServerErrorCode)
		return
	}
	res.ResponseSuccess(c, result)
}

// SendTopicTagsController 发送话题标签数据
func SendTopicTagsController(c *gin.Context) {
	//获取前端发送的数据
	json, err := readUtil.GetJsonvalue(c)

	if err != nil {
		fmt.Println("SendTopicTagsController() controller.article.getArticle.GetJsonvalue err=")
		myErr.CheckErrors(err, c)
		return
	}

	//获取到查询的标签
	result, err := article.GetTagsByTopicService(json)
	if err != nil {
		fmt.Println("SendTopicTagsController() controller.article.getArticle.GetTagsByTopicService err=")
		myErr.CheckErrors(err, c)
		return
	}

	//返回响应
	res.ResponseSuccess(c, result)
}
