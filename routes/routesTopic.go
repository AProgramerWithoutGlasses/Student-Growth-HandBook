package routes

import (
	"github.com/gin-gonic/gin"
	"studentGrow/controller/article"
)

func routesTopic(r *gin.Engine) {
	// 添加话题
	r.POST("/article/add_topic", article.AddTopicsController)
	// 获取话题
	r.POST("/article/get_topic", article.GetAllTopicsController)
	// 添加标签
	r.POST("/article/add_tags", article.AddTagsByTopicController)
	// 获取标签
	r.POST("/article/get_tags", article.GetTagsByTopicController)
}
