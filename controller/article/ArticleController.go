package article

import (
	"fmt"
	"github.com/gin-gonic/gin"
	res "studentGrow/pkg/response"
	"studentGrow/service/article"
	readUtil "studentGrow/utils/readMessage"
	timeUtil "studentGrow/utils/timeConverter"
)

// GetArticleId article_id
func GetArticleId(c *gin.Context) {
	//将数据解析到map中
	mp, e := readUtil.AnalyzeDataToMap(c)
	if e != nil {
		fmt.Println(" AnalyzeToMap err,", e)
	}
	//查询结果
	err, user, acl := article.GetArticleService(mp)

	//若在数据库中找不到该文章或用户
	if err != nil {
		fmt.Println("GetArticleId() controller.article.getArticle.GetArticleService err=", err)
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
		"article_like_sum":    acl.UpvoteAmount,
		"article_comment_sum": acl.CommentAmount,
	}
	//若可以找到
	res.ResponseSuccess(c, data)
}

// PublishArticle 发布文章
func PublishArticle(c *gin.Context) {
	//将数据解析到map中
	mp, e := readUtil.AnalyzeDataToMap(c)
	if e != nil {
		fmt.Println(" AnalyzeToMap err,", e)
	}
	fmt.Println(mp)

}
