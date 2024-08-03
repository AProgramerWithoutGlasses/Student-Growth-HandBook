package routes

import (
	"github.com/gin-gonic/gin"
	"studentGrow/controller/student"
	"studentGrow/logger"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

<<<<<<< HEAD
	// 勋
	r.POST("/student/getSelfCotnent", student.GetSelfContentContro)
	r.POST("/student/updateSelfContent", student.UpdateSelfContentContro)
	/*	r.POST("/stuManage/queryStudent", stuManage.QueryStuContro)
		r.POST("/stuManage/addSingleStudent", stuManage.AddSingleStuContro)
		r.POST("/stuManage/addMultipleStudent", stuManage.AddMultipleStuContro)
		r.POST("/stuManage/deleteStudent", stuManage.DeleteStuContro)
		r.POST("/stuManage/banStudent", stuManage.BanStuContro)
		r.POST("/stuManage/editStudent", stuManage.EditStuContro)
		r.POST("/stuManage/setStudentManager", stuManage.setStuManagerContro)
		r.POST("/stuManage/outputMultipleStudent", stuManage.outputMultipleStuContro)


	*/
=======
	r.POST("/article/content", middleWare.CORSMiddleware(), article.GetArticleId)
	//r.POST("/article/publish", article.PublishArticle) //发布文章
	r.POST("/article/list", article.GetArticleList)
	r.POST("/article/comment", article.PostCom)
	r.POST("/article/publish/get_tags", middleWare.CORSMiddleware(), article.SendTopicTags)
	r.POST("/article/like")        //文章点赞
	r.POST("/article/cancel_like") //取消文章点赞
	r.POST("/article/like_nums")   //获取文章点赞数量
>>>>>>> bbd4d3eaa3b86d0900cb2b387d9481155a2f2743

	RoutesXue(r)

	return r
}
