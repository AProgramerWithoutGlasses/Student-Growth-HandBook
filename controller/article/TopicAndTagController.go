package article

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"studentGrow/pkg/error"
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
		res.ResponseError(c, res.ServerErrorCode)
		return
	}
	err = article.AddTopicsService(json)
	if err != nil {
		if errors.Is(err, error.HasExistError()) {
			res.ResponseErrorWithMsg(c, res.ServerErrorCode, error.HasExistError().Error())
			return
		}
		res.ResponseError(c, res.ServerErrorCode)
		return
	}

	res.ResponseSuccess(c, nil)

}

// GetAllTopicsController 获取所有话题
func GetAllTopicsController(c *gin.Context) {

	result, err := article.GetAllTopicsService()
	if err != nil {
		fmt.Println("GetAllTopicsController() controller.article.getArticle.AnalyzeDataToMyData err=", err)
		if err != nil {
			if errors.Is(err, error.NotFoundError()) {
				res.ResponseErrorWithMsg(c, res.ServerErrorCode, error.NotFoundError().Error())
				return
			}
			res.ResponseErrorWithMsg(c, res.ServerErrorCode, error.NotFoundError().Error())
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
		res.ResponseError(c, res.ServerErrorCode)
		return
	}

	err = article.AddTagsByTopicService(json)
	if err != nil {
		fmt.Println("AddArticleTagsController() controller.article.getArticle.AddTagsByTopicService err=", err)
		if err != nil {
			res.ResponseErrorWithMsg(c, res.ServerErrorCode, error.HasExistError().Error())
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
		res.ResponseError(c, res.ServerErrorCode)
		return
	}

	result, err := article.GetTagsByTopicService(json)
	if err != nil {
		fmt.Println("AddArticleTagsController() controller.article.getArticle.GetTagsByTopicService err=", err)
		if errors.Is(err, error.NotFoundError()) {
			res.ResponseErrorWithMsg(c, res.ServerErrorCode, error.NotFoundError().Error())
			return
		}
		res.ResponseError(c, res.ServerErrorCode)
		return
	}
	res.ResponseSuccess(c, result)
}
