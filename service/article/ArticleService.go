package article

import (
	"fmt"
	jsonvalue "github.com/Andrew-M-C/go.jsonvalue"
	"github.com/pkg/errors"
	"sort"
	"strconv"
	"studentGrow/dao/mysql"
	"studentGrow/dao/redis"
	model "studentGrow/models/gorm_model"
	myErr "studentGrow/pkg/error"
	"studentGrow/utils/timeConverter"
	"time"
)

// GetArticleService 获取文章详情
func GetArticleService(j *jsonvalue.V) (error, map[string]any) {
	//获取文章id
	aid, err := j.GetInt("article_id")
	if err != nil {
		fmt.Println("GetArticleService() dao.mysql.sqp_nzx.GetInt err=", err)
		return err, nil
	}

	// 获取用户名
	username, err := j.GetString("username")
	if err != nil {
		fmt.Println("GetArticleService() dao.mysql.sqp_nzx.GetInt err=", err)
		return err, nil
	}
	// 获取uid
	uid, err := mysql.GetIdByUsername(username)
	if err != nil {
		fmt.Println("GetArticleService() dao.mysql.sqp_nzx.GetIdByUsername err=", err)
		return err, nil
	}

	//查找文章信息
	err, article := mysql.SelectArticleById(aid)
	if err != nil {
		fmt.Println("GetArticleService() dao.mysql.sqp_nzx.SelectArticleById err=", err)
		return err, nil
	}
	//查找用户信息
	err, user := mysql.SelectUserById(int(article.UserID))

	if err != nil {
		fmt.Println("GetArticleService() dao.mysql.sqp_nzx.SelectUserById err=", err)
		return err, nil
	}

	// 查询是否点赞或收藏
	liked, err := redis.IsUserLiked(strconv.Itoa(aid), username, 0)
	if err != nil {
		fmt.Println("GetArticleService() dao.mysql.sqp_nzx.IsUserLiked err=", err)
		return err, nil
	}
	selected, err := redis.IsUserCollected(username, strconv.Itoa(aid))
	if err != nil {
		fmt.Println("GetArticleService() dao.mysql.sqp_nzx.IsUserSelected err=", err)
		return err, nil
	}

	// 存储到浏览记录
	err = mysql.InsertReadRecord(uid, aid)
	if err != nil {
		fmt.Println("GetArticleService() dao.mysql.sqp_nzx.InsertReadRecord err=", err)
		return err, nil
	}

	data := map[string]any{
		"article_id":          article.ID,
		"username":            user.Username,
		"user_image":          user.HeadShot,
		"user_class":          user.Class,
		"article_post_time":   timeConverter.IntervalConversion(article.CreatedAt),
		"article_content":     map[string]string{"article_text": article.Content, "article_image": article.Pic, "article_video": article.Video},
		"topic_id":            article.Topic,
		"article_collect_sum": article.CollectAmount,
		"article_like_sum":    article.LikeAmount,
		"article_comment_sum": article.CommentAmount,
		"if_like":             liked,
		"if_collect":          selected,
	}

	return nil, data
}

// GetArticleListService 后台获取文章列表
func GetArticleListService(j *jsonvalue.V) ([]model.Article, error) {
	//获取参数：page, limit, sort, username
	page, e1 := j.GetInt("page")
	limit, e2 := j.GetInt("limit")
	sortType, e3 := j.GetString("sort")
	order, e4 := j.GetString("order")
	//获取发布时间、封禁状态、发布人、话题分类、关键词
	startAt, e5 := j.GetString("start_at")
	endAt, e6 := j.GetString("end_at")
	isBan, e7 := j.GetBool("article_ban")
	name, e8 := j.GetString("name")
	topic, e9 := j.GetString("topic")
	keyWords, e10 := j.GetString("key_words")

	if e1 != nil || e2 != nil || e3 != nil || e4 != nil || e5 != nil || e6 != nil || e7 != nil || e8 != nil || e9 != nil || e10 != nil {
		fmt.Println("GetArticleListService() service.article.GetString err=")
		return nil, myErr.DataFormatError()
	}

	//执行查询文章列表语句
	result, err := mysql.SelectArticleAndUserListByPage(page, limit, sortType, order, startAt, endAt, topic, keyWords, name, isBan)
	if err != nil {
		fmt.Println("GetArticleListService() service.article.GetString err=", err)
		return nil, myErr.NotFoundError()
	}

	return result, nil
}

// GetArticleListFirstPageService 前台首页模糊搜索文章列表
//func GetArticleListFirstPageService(j *jsonvalue.V) ([]model.Article, error) {
//	keyWords, e1 := j.GetString("key_word")
//	topic, e2 := j.GetString("topic_name")
//	sort, e3 := j.GetString("article_sort")
//	limit, e4 := j.GetInt("article_count")
//	page, e5 := j.GetInt("article_page")
//
//	if e1 != nil || e2 != nil || e3 != nil || e4 != nil || e5 != nil {
//		fmt.Println("GetArticleListFirstPageService() service.article.GetString err=")
//		return nil, myErr.DataFormatError()
//	}
//
//}

// PublishArticleService 发布文章
//func PublishArticleService() map[string]any{
//
//}

// GetTopicTagsService GetTopicTags 根据话题获得对应的标签mysql
//func GetTopicTagsService(j *jsonvalue.V) []map[string]any {
//	//解析获得话题
//	topic, err := j.GetString("topic")
//	fmt.Println(topic)
//	if err != nil {
//		fmt.Println("GetTopicTags() service.article.GetString err=", err)
//		return nil
//	}
//
//	//查询标签
//	return mysql.SelectTagsByTopic(topic)
//
//}

// AddTopicsService 添加话题
func AddTopicsService(j *jsonvalue.V) error {
	// 获取话题
	v, err := j.GetArray("topics")
	if err != nil {
		fmt.Println("AddTopicsService() service.article.GetArray err=", err)
		return err
	}
	//添加话题
	for _, v := range v.ForRangeArr() {
		if ok := redis.RDB.SIsMember("topics", v.String()).Val(); ok {
			return errors.New("has existed")
		}
		redis.RDB.SAdd("topics", v.String())
		fmt.Println(v.String())
	}
	return nil
}

// GetAllTopicsService 获取所有话题
func GetAllTopicsService() (topics []map[string]any, err error) {
	// 获取所有话题
	slice, err := redis.RDB.SMembers("topics").Result()

	if err != nil {
		fmt.Println("GetAllTopicsService() service.article.SMembers err=", err)
		return nil, err
	}
	// 检查是否存在记录
	if len(slice) == 0 {
		return nil, myErr.NotFoundError()
	}
	//生成数据包
	for i, v := range slice {
		topics = append(topics, map[string]any{
			"id":   i,
			"name": v,
		})
	}
	return topics, nil
}

// AddTagsByTopicService 添加话题标签
func AddTagsByTopicService(j *jsonvalue.V) error {
	//获取想要添加标签的对应话题
	topic, err := j.GetString("topic")

	if err != nil {
		fmt.Println("AddTagsByTopicService() service.article.GetString err=", err)
		return err
	}

	//获取想要添加的标签
	v, err := j.GetArray("tags")
	if err != nil {
		fmt.Println("AddTagsByTopicService() service.article.GetArray err=", err)
		return err
	}

	//添加标签
	for _, v := range v.ForRangeArr() {
		if ok := redis.RDB.SIsMember(topic, v.String()).Val(); ok {
			return errors.New("has existed")
		}
		redis.RDB.SAdd(topic, v.String())
	}
	return nil
}

// GetTagsByTopicService 获取话题对应的标签
func GetTagsByTopicService(j *jsonvalue.V) (tags []map[string]any, err error) {
	//获取想要添加标签的对应话题
	topic, err := j.GetString("topic")
	if err != nil {
		fmt.Println("GetTagsByTopicService() service.article.GetString err=", err)
		return nil, err
	}

	// 获取对应的标签
	slice, err := redis.RDB.SMembers(topic).Result()

	if err != nil {
		fmt.Println("GetTagsByTopicService() service.article.SMembers err=", err)
		return nil, err
	}

	if len(slice) == 0 {
		return nil, myErr.NotFoundError()
	}

	for i, v := range slice {
		tags = append(tags, map[string]any{
			"id":   i,
			"name": v,
		})
	}
	return tags, nil
}

// BannedArticleService 解封或封禁文章
func BannedArticleService(j *jsonvalue.V, role string, username string) error {
	// 获取文章id和封禁状态
	id, err := j.GetInt("article_id")
	if err != nil {
		fmt.Println("BannedArticleService() service.article.GetString err=", err)
		return err
	}
	isBan, err := j.GetBool("article_ban")

	switch role {
	case "class":
		err = mysql.BannedArticleByIdForClass(id, isBan, username)
	case "grade1":
		err = mysql.BannedArticleByIdForGrade(id, 1)
	case "grade2":
		err = mysql.BannedArticleByIdForGrade(id, 2)
	case "grade3":
		err = mysql.BannedArticleByIdForGrade(id, 3)
	case "grade4":
		err = mysql.BannedArticleByIdForGrade(id, 4)
	case "college":
		err = mysql.BannedArticleByIdForSuperman(id)
	case "superman":
		err = mysql.BannedArticleByIdForSuperman(id)
	default:
		return myErr.NotFoundError()
	}

	if err != nil {
		fmt.Println("BannedArticleService() service.article.BannedArticleById err=", err)
		return err
	}

	return nil
}

// DeleteArticleService 删除文章
func DeleteArticleService(j *jsonvalue.V, role string, username string) error {
	// 获取文章id
	id, err := j.GetInt("article_id")
	if err != nil {
		fmt.Println("DeleteArticleService() service.article.GetInt err=", err)
		return err
	}

	switch role {
	case "class":
		err = mysql.DeleteArticleByIdForClass(id, username)
	case "grade1":
		err = mysql.DeleteArticleByIdForGrade(id, 1)
	case "grade2":
		err = mysql.DeleteArticleByIdForGrade(id, 2)
	case "grade3":
		err = mysql.DeleteArticleByIdForGrade(id, 3)
	case "grade4":
		err = mysql.DeleteArticleByIdForGrade(id, 4)
	case "college":
		err = mysql.DeleteArticleByIdForSuperman(id)
	case "superman":
		err = mysql.DeleteArticleByIdForSuperman(id)
	default:
		return myErr.NotFoundError()
	}

	if err != nil {
		fmt.Println("DeleteArticleService() service.article.DeleteArticleById err=", err)
		return err
	}
	return nil
}

// ReportArticleService 举报文章
func ReportArticleService(j *jsonvalue.V, username string) error {
	// 获取文章id和举报者用户id,举报信息
	aid, err := j.GetInt("article_id")
	if err != nil {
		fmt.Println("ReportArticleService() service.article.GetInt err=", err)
		return err
	}

	reportMsg, err := j.GetString("report_msg")

	// 通过username查询id
	uid, err := mysql.SelectUserByUsername(username)
	if err != nil {
		fmt.Println("ReportArticleService() service.article.SelectUserByUsername err=", err)
		return err
	}

	err = mysql.ReportArticleById(aid, uid, reportMsg)
	if err != nil {
		fmt.Println("ReportArticleService() service.article.ReportArticleById err=", err)
		return err
	}
	return nil
}

// SearchHotArticlesOfDayService 获取今日十条热帖
// 简单加权:点赞0.5；评论0.3；收藏0.2
func SearchHotArticlesOfDayService(j *jsonvalue.V) (model.Articles, error) {

	// 获取热帖条数
	count, err := j.GetInt("article_count")
	if err != nil {
		fmt.Println("SearchHotArticlesOfDayService() service.article.GetInt err=", err)
		return nil, err
	}
	// 计算今日的始末时间
	startOfDay := time.Now().Truncate(24 * time.Hour) // 今天的开始时间
	endOfDay := startOfDay.Add(24 * time.Hour)        // 明天的开始时间

	articles, err := mysql.SearchHotArticlesOfDay(startOfDay, endOfDay)
	if err != nil {
		fmt.Println("SearchHotArticlesOfDayService() service.article.SearchHotArticlesOfDay err=", err)
		return nil, err
	}
	// 排序
	sort.Sort(articles)

	var list model.Articles

	// 获取热度前count条数据
	for i := 0; i < count; i++ {
		list = append(list, articles[i])
	}

	return list, nil
}

// SelectArticleAndUserListByPageFirstPageService 前台首页模糊查询文章列表
func SelectArticleAndUserListByPageFirstPageService(j *jsonvalue.V) ([]map[string]any, error) {
	keyWords, e1 := j.GetString("key_word")
	topic, e2 := j.GetString("topic_name")
	articleSort, e3 := j.GetString("article_sort")
	limit, e4 := j.GetInt("article_count")
	page, e5 := j.GetInt("article_page")
	username, e6 := j.GetString("username")

	if e1 != nil || e2 != nil || e3 != nil || e4 != nil || e5 != nil || e6 != nil {
		fmt.Println("SelectArticleAndUserListByPageFirstPageService() service.article.GetString err=")
		return nil, myErr.DataFormatError()
	}

	// 查询符合模糊搜索的文章集合
	articles, err := mysql.SelectArticleAndUserListByPageFirstPage(keyWords, topic, articleSort, limit, page)
	if err != nil {
		fmt.Println("SelectArticleAndUserListByPageFirstPageService() service.article.SelectArticleAndUserListByPageFirstPage err=", err)
		return nil, err
	}

	// 遍历文章集合并判断当前用户是否点赞或收藏该文章
	var list []map[string]any
	for _, article := range articles {
		okSelect, err := redis.IsUserCollected(strconv.Itoa(int(article.ID)), username)
		okLike, err := redis.IsUserLiked(strconv.Itoa(article.LikeAmount), username, 0)
		if err != nil {
			fmt.Println("SelectArticleAndUserListByPageFirstPageService() service.article.IsUserSelectedService err=", err)
			return nil, err
		}
		list = append(list, map[string]any{
			"user_headshot":   article.User.HeadShot,
			"user_class":      article.User.Class,
			"name":            article.User.Name,
			"article_id":      article.ID,
			"upvote_amount":   article.LikeAmount,
			"collect_amount":  article.CollectAmount,
			"comment_amount":  article.CommentAmount,
			"article_content": article.Content,
			"tag_name":        article.ArticleTags,
			"if_like":         okLike,
			"if_collect":      okSelect,
			"post_time":       timeConverter.IntervalConversion(article.CreatedAt),
		})

	}

	return list, nil
}
