package routes

import (
	"github.com/gin-gonic/gin"
	"studentGrow/controller/article"
	"studentGrow/logger"
	"studentGrow/utils/middleWare"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.POST("/article/content", middleWare.CORSMiddleware(), article.GetArticleId)
	//r.POST("/article/publish", article.PublishArticle) //发布文章
	r.POST("/article/list", article.GetArticleList)
	r.POST("/article/comment", article.PostCom)
	r.POST("/article/publish/get_tags", middleWare.CORSMiddleware(), article.SendTopicTags)
	r.POST("/article/like")        //文章点赞
	r.POST("/article/cancel_like") //取消文章点赞
	r.POST("/article/like_nums")   //获取文章点赞数量

	RoutesXue(r)

	return r
}
