package article

import (
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
		zap.L().Error("GetArticleIdController() controller.article.GetJsonvalue err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}

	// 获取用户名
	username, err := json.GetString("username")
	if err != nil {
		zap.L().Error("GetArticleService() service.article.GetString err=", zap.Error(err))
		return
	}

	// 获取文章详情
	art, err := article.GetArticleService(json)
	if err != nil {
		zap.L().Error("GetArticleIdController() controller.article.GetArticleService err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}

	if err != nil {
		zap.L().Error("GetArticleIdController() controller.article.GetArticleService err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}

	var tags []string
	for _, tag := range art.ArticleTags {
		tags = append(tags, tag.Tag.TagName)
	}
	var pics []string
	for _, pic := range art.ArticlePics {
		pics = append(pics, pic.Pic)
	}

	articleContent := map[string]any{
		"article_image": pics,
		"article_text":  art.Content,
		"article_video": art.Video,
	}

	var data map[string]any
	if (art.Ban == true && art.User.Username != username) || art.Status == false {
		data = map[string]any{
			"ban":             art.Ban,
			"status":          art.Status,
			"user_headshot":   nil,
			"name":            nil,
			"username":        nil,
			"user_class":      nil,
			"article_tags":    nil,
			"post_time":       nil,
			"article_content": nil,
			"like_amount":     nil,
			"collect_amount":  nil,
			"comment_amount":  nil,
			"is_like":         nil,
			"is_collect":      nil,
		}
	} else {
		data = map[string]any{
			"ban":             art.Ban,
			"status":          art.Status,
			"user_headshot":   art.User.HeadShot,
			"name":            art.User.Name,
			"username":        art.User.Username,
			"user_class":      art.User.Class,
			"article_tags":    tags,
			"post_time":       art.PostTime,
			"article_content": articleContent,
			"like_amount":     art.LikeAmount,
			"collect_amount":  art.CollectAmount,
			"comment_amount":  art.CommentAmount,
			"is_like":         art.IsLike,
			"is_collect":      art.IsCollect,
		}
	}

	res.ResponseSuccess(c, data)
}

// GetArticleListController 获取文章列表
func GetArticleListController(c *gin.Context) {
	in := struct {
		Page     int    `json:"page"`
		Limit    int    `json:"limit"`
		SortType string `json:"sort"`
		Order    string `json:"order"`
		StartAt  string `json:"start_at"`
		EndAt    string `json:"end_at"`
		IsBan    bool   `json:"article_ban"`
		Name     string `json:"name"`
		Topic    string `json:"topic"`
		KeyWords string `json:"key_words"`
	}{}

	err := c.ShouldBindJSON(&in)
	if err != nil {
		zap.L().Error("GetArticleListController() controller.article.ShouldBindJSON err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}

	// 通过token获取身份
	role, err := token.GetRole(c.GetHeader("token"))
	if err != nil {
		zap.L().Error("GetArticleListController() controller.article.GetRole err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}

	// 通过token获取username
	username, err := token.GetUsername(c.GetHeader("token"))
	if err != nil {
		zap.L().Error("GetArticleListController() controller.article.GetRole err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}

	//查询文章列表
	result, articleAmount, err := article.GetArticleListService(in.Page, in.Limit, in.SortType, in.Order, in.StartAt, in.EndAt, in.Topic, in.KeyWords, in.Name, in.IsBan, role, username)
	if err != nil {
		zap.L().Error("GetArticleListController() controller.article.GetArticleListService err=", zap.Error(myErr.DataFormatError()))
		myErr.CheckErrors(err, c)
		return
	}

	list := make([]map[string]any, 0)
	for _, val := range result {
		list = append(list, map[string]any{
			"article_id":      val.ID,
			"article_content": val.Content,
			"user_headshot":   val.User.HeadShot,
			"article_ban":     val.Ban,
			"upvote_amount":   val.LikeAmount,
			"comment_amount":  val.CommentAmount,
			"username":        val.User.Username,
			"created_at":      val.CreatedAt,
			"name":            val.User.Name,
		})
	}

	res.ResponseSuccess(c, map[string]any{
		"list":           list,
		"article_amount": articleAmount,
	})
}

// BannedArticleController 封禁文章
func BannedArticleController(c *gin.Context) {
	//获取前端发送的数据
	json, err := readUtil.GetJsonvalue(c)

	if err != nil {
		zap.L().Error("BannedArticleController() controller.article.GetJsonvalue err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}

	// 获取管理员username
	username, err := token.GetUsername(c.GetHeader("token"))
	if err != nil {
		zap.L().Error("BannedArticleController() controller.article.GetUsername err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}

	// 获取管理员角色
	role, err := token.GetRole(c.GetHeader("token"))
	if err != nil {
		zap.L().Error("BannedArticleController() controller.article.GetRole err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}

	// 对应帖子进行封禁或解封操作
	err = article.BannedArticleService(json, role, username)
	// 检查错误
	if err != nil {
		zap.L().Error("BannedArticleController() controller.article.BannedArticleService err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}

	res.ResponseSuccess(c, struct{}{})

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
		zap.L().Error("DeleteArticleController() controller.article.GetJsonvalue err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}

	// 对文章进行删除操作
	err = article.DeleteArticleService(json, role, username)
	if err != nil {
		zap.L().Error("DeleteArticleController() controller.article.DeleteArticleService err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}

	res.ResponseSuccess(c, struct{}{})

}

// ReportArticle 举报文章
func ReportArticle(c *gin.Context) {
	//获取前端发送的数据
	json, err := readUtil.GetJsonvalue(c)
	if err != nil {
		zap.L().Error("ReportArticle() controller.article.GetJsonvalue err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}
	// 通过token获取username
	username, err := token.GetUsername(c.GetHeader("token"))
	if err != nil {
		zap.L().Error("ReportArticle() controller.article.GetUsername err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}

	// 对文章进行举报并记录
	err = article.ReportArticleService(json, username)

	if err != nil {
		zap.L().Error("ReportArticle() controller.article.ReportArticleService err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}

	res.ResponseSuccess(c, struct{}{})
}

// GetHotArticlesOfDayController 获取今日十条热帖
func GetHotArticlesOfDayController(c *gin.Context) {
	//获取前端发送的数据
	json, err := readUtil.GetJsonvalue(c)
	if err != nil {
		zap.L().Error("GetHotArticlesOfDayController() controller.article.GetJsonvalue err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}

	articles, err := article.SearchHotArticlesOfDayService(json)
	if err != nil {
		zap.L().Error("GetHotArticlesOfDayController() controller.article.SearchHotArticlesOfDayService err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}

	list := make([]map[string]any, 0)
	for _, a := range articles {
		list = append(list, map[string]any{
			"article_id":    a.ID,
			"article_title": a.Content,
		})
	}

	res.ResponseSuccess(c, map[string]any{
		"article_list": list,
	})
}

// SelectArticleAndUserListByPageFirstPageController 前台首页模糊搜索文章列表
func SelectArticleAndUserListByPageFirstPageController(c *gin.Context) {
	in := struct {
		Username string `json:"username"`
		KeyWords string `json:"key_word"`
		Topic    string `json:"topic_name"`
		SortWay  string `json:"article_sort"`
		Limit    int    `json:"article_count"`
		Page     int    `json:"article_page"`
	}{}

	err := c.ShouldBindJSON(&in)
	if err != nil {
		zap.L().Error("SelectArticleAndUserListByPageFirstPageController() controller.article.ShouldBindJSON err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}

	articles, err := article.SelectArticleAndUserListByPageFirstPageService(in.Username, in.KeyWords, in.Topic, in.SortWay, in.Limit, in.Page)
	if err != nil {
		zap.L().Error("SelectArticleAndUserListByPageFirstPageController() controller.article.SelectArticleAndUserListByPageFirstPageService err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}

	list := make([]map[string]any, 0)
	for _, a := range articles {
		var pics []string
		var tags []string
		for _, pic := range a.ArticlePics {
			pics = append(pics, pic.Pic)
		}
		for _, tag := range a.ArticleTags {
			tags = append(tags, tag.Tag.TagName)
		}

		list = append(list, map[string]any{
			"user_headshot":   a.User.HeadShot,
			"user_class":      a.User.Class,
			"name":            a.User.Name,
			"article_id":      a.ID,
			"like_amount":     a.LikeAmount,
			"collect_amount":  a.CollectAmount,
			"comment_amount":  a.CommentAmount,
			"article_content": a.Content,
			"article_pics":    pics,
			"article_video":   a.Video,
			"article_tags":    tags,
			"article_topic":   a.Topic,
			"is_like":         a.IsLike,
			"is_collect":      a.IsCollect,
			"post_time":       a.PostTime,
			"username":        a.User.Username,
		})
	}

	res.ResponseSuccess(c, map[string]any{
		"content": list,
	})
}

// GetArticleByClassController 班级分类获取文章列表
func GetArticleByClassController(c *gin.Context) {
	input := struct {
		Username string `json:"username"`
		KeyWords string `json:"key_word"`
		SortWay  string `json:"article_sort"`
		Class    string `json:"class_name"`
		Limit    int    `json:"article_count"`
		Page     int    ` json:"article_page"`
	}{}

	err := c.ShouldBindJSON(&input)
	if err != nil {
		zap.L().Error("GetArticleByClassController() controller.article.ShouldBindJSON err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}

	// 获取列表
	articles, err := article.GetArticlesByClassService(input.KeyWords, input.Username, input.SortWay, input.Limit, input.Page, input.Class)
	if err != nil {
		zap.L().Error("GetArticleByClassController() controller.article.GetArticlesByClassService err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}

	list := make([]map[string]any, 0)
	for _, a := range articles {
		var pics []string
		var tags []string
		for _, pic := range a.ArticlePics {
			pics = append(pics, pic.Pic)
		}
		for _, tag := range a.ArticleTags {
			tags = append(tags, tag.Tag.TagName)
		}

		list = append(list, map[string]any{
			"user_headshot":   a.User.HeadShot,
			"user_class":      a.User.Class,
			"name":            a.User.Name,
			"article_id":      a.ID,
			"like_amount":     a.LikeAmount,
			"collect_amount":  a.CollectAmount,
			"comment_amount":  a.CommentAmount,
			"article_content": a.Content,
			"article_pics":    pics,
			"article_video":   a.Video,
			"article_tags":    tags,
			"article_topic":   a.Topic,
			"is_like":         a.IsLike,
			"is_collect":      a.IsCollect,
			"post_time":       a.PostTime,
			"username":        a.User.Username,
		})
	}
	res.ResponseSuccess(c, map[string]any{
		"content": list,
	})
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
	countString := form.Value["word_count"][0]
	var wordCount int
	if countString != "" {
		wordCount, err = strconv.Atoi(countString)
		if err != nil {
			zap.L().Error("PublishArticleController() controller.article.getArticle.Atoi err=", zap.Error(err))
			myErr.CheckErrors(err, c)
			return
		}
	}

	tags := form.Value["article_tags"]
	topic := form.Value["article_topic"][0]

	statusString := form.Value["article_status"][0]
	var status bool

	if statusString != "" {
		status, err = strconv.ParseBool(statusString)
		if err != nil {
			zap.L().Error("PublishArticleController() controller.article.getArticle.ParseBool err=", zap.Error(err))
			myErr.CheckErrors(err, c)
			return
		}
	}

	// 获取图片和视频文件
	pics := form.File["pic"]
	video := form.File["video"]

	err = article.PublishArticleService(username, content, topic, wordCount, tags, pics, video, status)
	if err != nil {
		zap.L().Error("PublishArticleController() controller.article.getArticle.PublishArticleService err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}

	res.ResponseSuccess(c, struct{}{})

}

// ReviseArticleStatusController 修改文章私密状态
func ReviseArticleStatusController(c *gin.Context) {
	in := struct {
		ArticleId int  `json:"article_id"`
		Status    bool `json:"article_status"`
	}{}

	err := c.ShouldBindJSON(&in)
	if err != nil {
		zap.L().Error("ReviseArticleStatusController() controller.article.getArticle.ShouldBindJSON err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}

	err = article.ReviseArticleStatusService(in.ArticleId, in.Status)
	if err != nil {
		zap.L().Error("ReviseArticleStatusController() controller.article.getArticle.ReviseArticleStatusService err=", zap.Error(err))
		myErr.CheckErrors(err, c)
		return
	}

	res.ResponseSuccess(c, struct{}{})
}
