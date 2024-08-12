package routes

import (
	"github.com/gin-gonic/gin"
	"studentGrow/controller/article"
	"studentGrow/utils/token"
)

func routesArticle(r *gin.Engine) {
	at := r.Group("/article")
	// 获取文章内容
	at.POST("/content", article.GetArticleIdController)
	// 获取文章列表
	at.POST("/list", article.GetArticleListController)
	// 对文章进行评论
	at.POST("/comment", article.PostCom)
	//发布文章
	//r.POST("/article/publish", article.PublishArticle)
	// 获取文章标签
	at.POST("/publish/get_tags", article.SendTopicTagsController)
	//文章点赞
	at.POST("/like")
	//取消文章点赞
	at.POST("/cancel_like")
	//获取文章点赞数量
	at.POST("/like_nums")
	//封禁文章
	at.POST("/ban", article.BannedArticleController)
	//删除文章
	at.POST("/delete", token.AuthMiddleware(), article.DeleteArticleController)
	//举报文章
	at.POST("/report", article.ReportArticle)
	// 获取今日热帖
	at.POST("/hot_articles", article.GetHotArticlesOfDayController)
	// 首页模糊搜索
	at.POST("/search_first", article.SelectArticleAndUserListByPageFirstPageController)
	// 收藏
	at.POST("/collect", article.CollectArticleController)
	// 取消收藏
	at.POST("/cancel_collect", article.CancelCollectArticleController)
	// 查看收藏列表
	at.POST("/get_collects", article.GetArticleListForSelectController)
}
