package routes

import (
	"github.com/gin-gonic/gin"
	"studentGrow/controller/article"
	"studentGrow/utils/middleWare"
)

func routesArticle(r *gin.Engine) {
	r.POST("/article/content", middleWare.CORSMiddleware(), article.GetArticleIdController)
	//r.POST("/article/publish", article.PublishArticle) //发布文章
	r.POST("/article/list", article.GetArticleListController)
	r.POST("/article/comment", article.PostCom)
	r.POST("/article/publish/get_tags", middleWare.CORSMiddleware(), article.SendTopicTagsController)
	r.POST("/article/like")        //文章点赞
	r.POST("/article/cancel_like") //取消文章点赞
	r.POST("/article/like_nums")   //获取文章点赞数量
}
