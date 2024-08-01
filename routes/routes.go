package routes

import (
	"github.com/gin-gonic/gin"
<<<<<<< HEAD
	"studentGrow/controller/user"
	"studentGrow/logger"
=======
	"studentGrow/controller/article"
	"studentGrow/logger"
	"studentGrow/utils/middleWare"
>>>>>>> bd64b59feb8245f5364f131e7324b0194666ecf9
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

<<<<<<< HEAD
	//r.POST("/article/content", article.GetArticleId)
	r.POST("/user/getSelfCotnent", user.GetSelfContentContro)
	r.POST("/user/updateSelfContent", user.UpdateSelfContentContro)
=======
	r.POST("/article/content", middleWare.CORSMiddleware(), article.GetArticleId)
	r.POST("/article/publish", article.PublishArticle) //发布文章
	r.POST("/article/comment", article.PostCom)
	r.POST("/article/like")        //文章点赞
	r.POST("/article/cancel_like") //取消文章点赞
	r.POST("/article/like_nums")   //获取文章点赞数量
>>>>>>> bd64b59feb8245f5364f131e7324b0194666ecf9

	return r
}
