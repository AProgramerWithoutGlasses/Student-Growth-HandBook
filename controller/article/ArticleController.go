package article

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"studentGrow/pkg/error"
	res "studentGrow/pkg/response"
	"studentGrow/service/article"
	readUtil "studentGrow/utils/readMessage"
	timeUtil "studentGrow/utils/timeConverter"
)

// GetArticleIdController article_id	获取文章详情
func GetArticleIdController(c *gin.Context) {
	//将数据解析到map中
	json, e := readUtil.GetJsonvalue(c)
	if e != nil {
		fmt.Println(" AnalyzeToMap err=", e)
	}
	err, user, acl := article.GetArticleService(json)

	//若在数据库中找不到该文章或用户
	if err != nil {
		fmt.Println("GetArticleIdController() controller.article.getArticle.GetArticleService err=", err)
		res.ResponseErrorWithMsg(c, res.ServerErrorCode, "NOT FOUND")
		return
	}
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
	//若可以找到
	res.ResponseSuccess(c, data)
}

// SendTopicTagsController 发送话题标签数据
func SendTopicTagsController(c *gin.Context) {
	//获取前端发送的数据
	json, err := readUtil.GetJsonvalue(c)

	if err != nil {
		fmt.Println("SendTopicTagsController() controller.article.getArticle.GetJsonvalue err=")
		res.ResponseError(c, res.ServerErrorCode)
		return
	}

	//获取到查询的标签
	result, err := article.GetTagsByTopicService(json)
	if err != nil {
		fmt.Println("SendTopicTagsController() controller.article.getArticle.GetTagsByTopicService err=")
		res.ResponseErrorWithMsg(c, res.ServerErrorCode, error.HasExistError().Error())
		return
	}

	//返回响应
	res.ResponseSuccess(c, result)

}

// GetArticleListController 获取文章列表
func GetArticleListController(c *gin.Context) {
	//获取前端发送的数据
	json, err := readUtil.GetJsonvalue(c)

	if err != nil {
		fmt.Println("GetArticleListController() controller.article.getArticle.AnalyzeDataToMyData err=", err)
		res.ResponseError(c, res.ServerErrorCode)
		return
	}

	//查询文章列表
	result, err := article.GetArticleListService(json)

	if err != nil {
		fmt.Println("GetArticleList() controller.article.getArticle.AnalyzeDataToMyData err=", err)
		res.ResponseError(c, res.ServerErrorCode)
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

// PublishArticle 发布文章
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
