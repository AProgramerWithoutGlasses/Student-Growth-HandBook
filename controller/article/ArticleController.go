package article

import (
	"fmt"
	"github.com/gin-gonic/gin"
	myErr "studentGrow/pkg/error"
	res "studentGrow/pkg/response"
	"studentGrow/service/article"
	readUtil "studentGrow/utils/readMessage"
	timeUtil "studentGrow/utils/timeConverter"
)

// GetArticleIdController article_id	获取文章详情
func GetArticleIdController(c *gin.Context) {
	//将数据解析到map中
	json, err := readUtil.GetJsonvalue(c)
	if err != nil {
		fmt.Println(" AnalyzeToMap err=", err)
		myErr.CheckErrors(err, c)
		return
	}
	// 获取文章详情
	err, user, acl := article.GetArticleService(json)
	//错误检查
	if err != nil {
		myErr.CheckErrors(err, c)
		return
	}
	// 无错误
	data := map[string]any{
		"article_id":          acl.ID,
		"username":            user.Username,
		"user_image":          user.HeadShot,
		"user_class":          user.Class,
		"article_post_time":   timeUtil.IntervalConversion(acl.CreatedAt),
		"article_content":     map[string]string{"article_text": acl.Content, "article_image": acl.Pic, "article_video": acl.Video},
		"topic_id":            acl.Topic,
		"article_collect_sum": acl.CollectAmount,
		"article_like_sum":    acl.LikeAmount,
		"article_comment_sum": acl.CommentAmount,
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
		fmt.Println("BannedArticle() controller.article.getArticle.GetJsonvalue err=", err)
		myErr.CheckErrors(err, c)
		return
	}

	// 对应帖子进行封禁或解封操作
	err = article.BannedArticleService(json)
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
	//获取前端发送的数据
	json, err := readUtil.GetJsonvalue(c)

	if err != nil {
		fmt.Println("DeleteArticleController() controller.article.getArticle.GetJsonvalue err=", err)
		myErr.CheckErrors(err, c)
		return
	}

	// 对文章进行删除操作
	err = article.DeleteArticleService(json)
	if err != nil {
		fmt.Println("DeleteArticleController() controller.article.getArticle.GetJsonvalue err=", err)
		myErr.CheckErrors(err, c)
		return
	}

	res.ResponseSuccess(c, nil)

}

//PublishArticle 发布文章
//func PublishArticle(c *gin.Context) {
//	//解析formdata
//	form, err := readUtil.GetFormData(c)
//	if err != nil {
//		return
//	}
//
//	//发布文章
//	//article.PublishArticleService(form)
//}
