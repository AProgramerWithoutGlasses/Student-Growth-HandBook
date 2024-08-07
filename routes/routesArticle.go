package routes

import (
	"github.com/gin-gonic/gin"
	"studentGrow/controller/article"
	"studentGrow/utils/middleWare"
)

func routesArticle(r *gin.Engine) {
	r.POST("/article/content", middleWare.CORSMiddleware(), article.GetArticleIdController) // 获取文章详情
	//r.POST("/article/publish", article.PublishArticle) //发布文章
	r.POST("/article/list", article.GetArticleListController) // 获取文章列表
	r.POST("/article/comment", article.PostCom)               // 对文章进行评论
	r.POST("/article/publish/get_tags", middleWare.CORSMiddleware(), article.SendTopicTagsController)
	r.POST("/article/like")                                         //文章点赞
	r.POST("/article/cancel_like")                                  //取消文章点赞
	r.POST("/article/like_nums")                                    //获取文章点赞数量
	r.POST("/article/list/ban", article.BannedArticleController)    //封禁文章
	r.POST("/article/list/delete", article.DeleteArticleController) //删除文章
	r.POST("/article/report", article.ReportArticle)                //举报文章
}
