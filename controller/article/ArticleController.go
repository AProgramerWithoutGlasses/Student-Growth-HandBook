package article

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	myErr "studentGrow/pkg/error"
	res "studentGrow/pkg/response"
	"studentGrow/service/article"
	readUtil "studentGrow/utils/readMessage"
	"studentGrow/utils/token"
)

// GetArticleIdController article_id	获取文章详情
func GetArticleIdController(c *gin.Context) {
	//将数据解析到map中
	json, err := readUtil.GetJsonvalue(c)
	if err != nil {
		fmt.Println(" GetArticleIdController() controller.article.GetJsonvalue err=", err)
		myErr.CheckErrors(err, c)
		return
	}
	// 获取文章详情
	err, data := article.GetArticleService(json)
	//错误检查
	if err != nil {
		fmt.Println(" GetArticleIdController() controller.article.GetArticleService err=", err)
		myErr.CheckErrors(err, c)
		return
	}

	res.ResponseSuccess(c, data)
}

// GetArticleListController 获取文章列表
func GetArticleListController(c *gin.Context) {
	//获取前端发送的数据
	json, err := readUtil.GetJsonvalue(c)

	if err != nil {
		fmt.Println("GetArticleListController() controller.article.getArticle.AnalyzeDataToMyData err=", err)
		myErr.CheckErrors(err, c)
		return
	}

	//查询文章列表
	result, err := article.GetArticleListService(json)

	if err != nil {
		fmt.Println("GetArticleList() controller.article.getArticle.AnalyzeDataToMyData err=", err)
		myErr.CheckErrors(err, c)
		return
	}
	var list []map[string]any
	for _, val := range result {
		list = append(list, map[string]any{
			"article_id":      val.ID,
			"article_content": val.Content,
			"user_headshot":   val.User.HeadShot,
			"article_ban":     val.Ban,
			"upvote_amount":   val.LikeAmount,
			"comment_amount":  val.CommentAmount,
			"username":        val.User.Name,
			"created_at":      val.CreatedAt,
		})
	}
	res.ResponseSuccess(c, map[string][]map[string]any{
		"list": list,
	})
}

// BannedArticleController 封禁文章
func BannedArticleController(c *gin.Context) {
	//获取前端发送的数据
	json, err := readUtil.GetJsonvalue(c)

	if err != nil {
		fmt.Println("BannedArticleController() controller.article.getArticle.GetJsonvalue err=", err)
		myErr.CheckErrors(err, c)
		return
	}

	// 获取管理员username
	username, err := token.GetUsername(c.GetHeader("token"))
	if err != nil {
		fmt.Println("BannedArticleController() controller.article.getArticle.GetUsername err=", err)
		myErr.CheckErrors(err, c)
		return
	}

	// 获取管理员角色
	role, err := token.GetRole(c.GetHeader("token"))
	if err != nil {
		fmt.Println("BannedArticleController() controller.article.getArticle.GetRole err=", err)
		myErr.CheckErrors(err, c)
		return
	}

	// 对应帖子进行封禁或解封操作
	err = article.BannedArticleService(json, role, username)
	// 检查错误
	if err != nil {
		fmt.Println("BannedArticle() controller.article.getArticle.BannedArticleService err=", err)
		myErr.CheckErrors(err, c)
		return
	}

	res.ResponseSuccess(c, nil)

}

// DeleteArticleController 删除文章
func DeleteArticleController(c *gin.Context) {

	// 获取权限角色
	role, err := token.GetRole(c.GetHeader("token"))

	// 获取管理员username
	username, err := token.GetUsername(c.GetHeader("token"))

	//获取前端发送的数据
	json, err := readUtil.GetJsonvalue(c)

	if err != nil {
		fmt.Println("DeleteArticleController() controller.article.getArticle.GetJsonvalue err=", err)
		myErr.CheckErrors(err, c)
		return
	}

	// 对文章进行删除操作
	err = article.DeleteArticleService(json, role, username)
	if err != nil {
		fmt.Println("DeleteArticleController() controller.article.getArticle.GetJsonvalue err=", err)
		myErr.CheckErrors(err, c)
		return
	}

	res.ResponseSuccess(c, nil)

}

// ReportArticle 举报文章
func ReportArticle(c *gin.Context) {
	//获取前端发送的数据
	json, err := readUtil.GetJsonvalue(c)
	if err != nil {
		fmt.Println("ReportArticle() controller.article.getArticle.GetJsonvalue err=", err)
		myErr.CheckErrors(err, c)
		return
	}
	// 通过token获取username
	username, err := token.GetUsername(c.GetHeader("token"))
	if err != nil {
		fmt.Println("ReportArticle() controller.article.getArticle.GetUsername err=", err)
		myErr.CheckErrors(err, c)
		return
	}

	// 对文章进行举报并记录
	err = article.ReportArticleService(json, username)

	if err != nil {
		fmt.Println("ReportArticle() controller.article.getArticle.ReportArticleService err=", err)
		myErr.CheckErrors(err, c)
		return
	}

	res.ResponseSuccess(c, nil)
}

// GetHotArticlesOfDayController 获取今日十条热帖
func GetHotArticlesOfDayController(c *gin.Context) {
	//获取前端发送的数据
	json, err := readUtil.GetJsonvalue(c)
	if err != nil {
		fmt.Println("ReportArticle() controller.article.getArticle.GetJsonvalue err=", err)
		myErr.CheckErrors(err, c)
		return
	}

	articles, err := article.SearchHotArticlesOfDayService(json)
	if err != nil {
		fmt.Println("GetHotArticlesOfDayController() controller.article.getArticle.SearchHotArticlesOfDayService err=", err)
		myErr.CheckErrors(err, c)
		return
	}

	res.ResponseSuccess(c, articles)
}

// SelectArticleAndUserListByPageFirstPageController 前台首页模糊搜索文章列表
func SelectArticleAndUserListByPageFirstPageController(c *gin.Context) {
	//获取前端发送的数据
	json, err := readUtil.GetJsonvalue(c)
	if err != nil {
		fmt.Println("SelectArticleAndUserListByPageFirstPageController() controller.article.getArticle.GetJsonvalue err=", err)
		myErr.CheckErrors(err, c)
		return
	}

	list, err := article.SelectArticleAndUserListByPageFirstPageService(json)
	if err != nil {
		fmt.Println("SelectArticleAndUserListByPageFirstPageController() controller.article.getArticle.SelectArticleAndUserListByPageFirstPageService err=", err)
		myErr.CheckErrors(err, c)
		return
	}

	res.ResponseSuccess(c, list)
}

// PublishArticleController 发布文章
func PublishArticleController(c *gin.Context) {
	// 通过token获取username
	username, err := token.GetUsername(c.GetHeader("token"))
	if err != nil {
		zap.L().Error("PublishArticleController() controller.article.getArticle.GetUsername err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}

	err = c.Request.ParseMultipartForm(10 << 23) // 最大 80MB

	if err != nil {
		zap.L().Error("PublishArticleController() controller.article.getArticle.ParseMultipartForm err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}

	form, err := c.MultipartForm()
	if err != nil {
		zap.L().Error("PublishArticleController() controller.article.getArticle.MultipartForm err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}

	// 获取基本数据
	content := form.Value["article_content"][0]
	wordCount, err := strconv.Atoi(form.Value["word_count"][0])
	if err != nil {
		zap.L().Error("PublishArticleController() controller.article.getArticle.Atoi err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}
	tags := form.Value["article_tags"]
	topic := form.Value["article_topic"][0]
	// 获取图片和视频文件
	pics := form.File["pic"]
	video := form.File["video"]

	article.PublishArticleService(username, content, topic, wordCount, tags, pics, video)

}
