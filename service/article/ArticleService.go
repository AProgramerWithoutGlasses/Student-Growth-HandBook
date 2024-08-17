package article

import (
	"fmt"
	jsonvalue "github.com/Andrew-M-C/go.jsonvalue"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"mime/multipart"
	"sort"
	"strconv"
	"studentGrow/dao/mysql"
	"studentGrow/dao/redis"
	"studentGrow/models/constant"
	model "studentGrow/models/gorm_model"
	myErr "studentGrow/pkg/error"
	"studentGrow/utils/fileProcess"
	"studentGrow/utils/timeConverter"
	"time"
)

// GetArticleService 获取文章详情
func GetArticleService(j *jsonvalue.V) (*model.Article, error) {
	//获取文章id
	aid, err := j.GetInt("article_id")
	if err != nil {
		fmt.Println("GetArticleService() dao.mysql.sqp_nzx.GetInt err=", err)
		return nil, err
	}

	// 获取用户名
	username, err := j.GetString("username")
	if err != nil {
		fmt.Println("GetArticleService() dao.mysql.sqp_nzx.GetInt err=", err)
		return nil, err
	}

	//查找文章信息
	err, article := mysql.SelectArticleById(aid)
	if err != nil {
		fmt.Println("GetArticleService() dao.mysql.sqp_nzx.SelectArticleById err=", err)
		return nil, err
	}

	// 该文章阅读量+1
	err = UpdateArticleReadNumService(aid, 1)
	if err != nil {
		zap.L().Error("GetArticleService() dao.mysql.sql_article.UpdateArticleReadNumService", zap.Error(err))
		return nil, err
	}

	// 查询是否点赞或收藏
	liked, err := redis.IsUserLiked(strconv.Itoa(aid), username, 0)
	if err != nil {
		fmt.Println("GetArticleService() dao.mysql.sqp_nzx.IsUserLiked err=", err)
		return nil, err
	}
	article.IsLike = liked
	selected, err := redis.IsUserCollected(username, strconv.Itoa(aid))
	if err != nil {
		fmt.Println("GetArticleService() dao.mysql.sqp_nzx.IsUserSelected err=", err)
		return nil, err
	}
	article.IsCollect = selected

	// 计算发布时间
	article.PostTime = timeConverter.IntervalConversion(article.CreatedAt)

	// 存储到浏览记录
	uid, err := mysql.GetIdByUsername(username)
	if err != nil {
		fmt.Println("GetArticleService() dao.mysql.sqp_nzx.GetIdByUsername err=", err)
		return nil, err
	}
	err = mysql.InsertReadRecord(uid, aid)
	if err != nil {
		fmt.Println("GetArticleService() dao.mysql.sqp_nzx.InsertReadRecord err=", err)
		return nil, err
	}

	return article, err
}

// GetArticleListService 后台获取文章列表
func GetArticleListService(page, limit int, sortType, order, startAt, endAt, topic, keyWords, name string, isBan bool) ([]model.Article, error) {
	//执行查询文章列表语句
	result, err := mysql.SelectArticleAndUserListByPage(page, limit, sortType, order, startAt, endAt, topic, keyWords, name, isBan)
	if err != nil {
		fmt.Println("GetArticleListService() service.article.GetString err=", err)
		return nil, myErr.NotFoundError()
	}

	return result, nil
}

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
func GetAllTopicsService() ([]model.Topic, error) {
	// 获取所有话题
	topics, err := mysql.QueryAllTopics()
	if err != nil {
		zap.L().Error("GetAllTopicsService() service.article.QueryAllTopics err=", zap.Error(err))
		return nil, err
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
func GetTagsByTopicService(topicId int) ([]map[string]any, error) {
	//获取想要添加标签的对应话题
	tags, err := mysql.QueryTagsByTopic(topicId)
	if err != nil {
		zap.L().Error("GetTagsByTopicService() service.article.QueryTagsByTopic err=", zap.Error(err))
		return nil, err
	}

	var list []map[string]any
	for i, v := range tags {
		list = append(list, map[string]any{
			"id":   i,
			"name": v.TagName,
		})
	}
	return list, nil
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
		zap.L().Error("BannedArticleService() service.article.GetIdByUsername err=", zap.Error(myErr.DataFormatError()))
		return err
	}

	/*
		回滚分数
	*/
	point := constant.PointConstant
	if isBan {
		point = -point
	}
	err = UpdatePointByUsernamePointAid(username, point, id)
	if err != nil {
		zap.L().Error("BannedArticleService() service.article.UpdatePointByUsernamePointAid err=", zap.Error(myErr.DataFormatError()))
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

	/*
		回滚分数
	*/
	err = UpdatePointByUsernamePointAid(username, -constant.PointConstant, id)
	if err != nil {
		zap.L().Error("DeleteArticleService() service.article.UpdatePointByUsernamePointAid err=", zap.Error(myErr.DataFormatError()))
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
		/*
			删除失败回滚分数
		*/
		err = UpdatePointByUsernamePointAid(username, constant.PointConstant, id)
		if err != nil {
			zap.L().Error("DeleteArticleService() service.article.UpdatePointByUsernamePointAid err=", zap.Error(myErr.DataFormatError()))
			return err
		}
		zap.L().Error("DeleteArticleService() service.article.DeleteArticleByIdForClass err=", zap.Error(myErr.DataFormatError()))
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
	for i := 0; i < count && i < len(articles); i++ {
		list = append(list, articles[i])
	}

	return list, nil
}

// SelectArticleAndUserListByPageFirstPageService 前台首页模糊查询文章列表
func SelectArticleAndUserListByPageFirstPageService(username, keyWords, topic, SortWay string, limit, page int) ([]model.Article, error) {

	// 查询符合模糊搜索的文章集合
	articles, err := mysql.SelectArticleAndUserListByPageFirstPage(keyWords, topic, limit, page)
	if err != nil {
		zap.L().Error("SelectArticleAndUserListByPageFirstPageService() service.article.SelectArticleAndUserListByPageFirstPage err=", zap.Error(err))
		return nil, err
	}

	if SortWay == "hot" {
		sort.Sort(articles)
	}

	// 遍历文章集合并判断当前用户是否点赞或收藏该文章
	for i := 0; i < len(articles); i++ {
		okSelect, err := redis.IsUserCollected(strconv.Itoa(int(articles[i].ID)), username)
		okLike, err := redis.IsUserLiked(strconv.Itoa(articles[i].LikeAmount), username, 0)
		if err != nil {
			zap.L().Error("SelectArticleAndUserListByPageFirstPageService() service.article.IsUserLiked err=", zap.Error(err))
			return nil, err
		}
		articles[i].IsCollect = okSelect
		articles[i].IsLike = okLike

		// 计算发布时间
		articles[i].PostTime = timeConverter.IntervalConversion(articles[i].CreatedAt)
	}

	return articles, nil
}

// PublishArticleService 发布文章
func PublishArticleService(username, content, topic string, wordCount int, tags []string, pics []*multipart.FileHeader, video []*multipart.FileHeader, status bool) error {
	// 检查文本内容字数
	if len(content) < 30 || len(content) > 300 {
		zap.L().Error("PublishArticleService() service.article.ArticleService err=", zap.Error(myErr.DataFormatError()))
		return myErr.DataFormatError()
	}

	// 检查标签
	if len(tags) <= 0 {
		zap.L().Error("PublishArticleService() service.article.ArticleService err=", zap.Error(myErr.DataFormatError()))
		return myErr.DataFormatError()
	}

	//  将图片上传至oss
	var picPath []string
	if len(pics) > 0 {
		for _, pic := range pics {
			url, err := fileProcess.UploadFile("image", pic)
			fmt.Println(url)
			if err != nil {
				zap.L().Error("PublishArticleService() service.article.UploadFile err=", zap.Error(myErr.DataFormatError()))
				return err
			}
			picPath = append(picPath, url)
		}
	}

	//  检查视频数量
	if len(video) > 1 {
		zap.L().Error("PublishArticleService() service.article.ArticleService err=", zap.Error(myErr.DataFormatError()))
		return myErr.DataFormatError()
	}

	// 将视频上传至oss
	var videoPath string
	if len(video) > 0 {
		url, err := fileProcess.UploadFile("video", video[0])
		if err != nil {
			zap.L().Error("PublishArticleService() service.article.UploadFile err=", zap.Error(myErr.DataFormatError()))
			return err
		}
		videoPath = url
	}
	uid, err := mysql.GetIdByUsername(username)
	if err != nil {
		zap.L().Error("PublishArticleService() service.article.GetIdByUsername err=", zap.Error(myErr.DataFormatError()))
		return err
	}

	// 插入新文章
	aid, err := mysql.InsertArticleContent(content, topic, uid, wordCount, tags, picPath, videoPath, status)
	if err != nil {
		zap.L().Error("PublishArticleService() service.article.InsertArticleContent err=", zap.Error(myErr.DataFormatError()))
		return err
	}

	// 上传标签
	err = mysql.InsertArticleTags(tags, aid)
	if err != nil {
		zap.L().Error("PublishArticleService() service.article.InsertArticleTags err=", zap.Error(myErr.DataFormatError()))
		return err
	}
	topicId, err := mysql.QueryTagIdByTagName(topic)
	if err != nil {
		zap.L().Error("PublishArticleService() service.article.QueryTagIdByTagName err=", zap.Error(myErr.DataFormatError()))
		return err
	}

	// 增加分数
	err = UpdatePointService(uid, constant.PointConstant, topicId)
	if err != nil {
		zap.L().Error("PublishArticleService() service.article.UpdatePointService err=", zap.Error(myErr.DataFormatError()))
		return err
	}

	// 将文章更新到redis点赞、收藏
	redis.RDB.HSet("article", strconv.Itoa(aid), 0)
	redis.RDB.HSet("collect", strconv.Itoa(aid), 0)
	fmt.Println("collect", strconv.Itoa(aid))
	return nil
}

// GetArticlesByClassService 班级分类查询文章
func GetArticlesByClassService(keyWords, username, sortWay string, limit, page, classId int) ([]model.Article, error) {
	// 获取class
	class, err := mysql.QueryClassByClassId(classId)
	if err != nil {
		zap.L().Error("GetArticlesByClassService() service.article.QueryClassByClassId err=", zap.Error(err))
		return nil, err
	}

	articles, err := mysql.QueryArticleByClass(limit, page, class, keyWords)
	if err != nil {
		zap.L().Error("GetArticlesByClassService() service.article.QueryArticleByClass err=", zap.Error(err))
		return nil, err
	}

	if sortWay == "hot" {
		sort.Sort(articles)
	}

	// 遍历文章集合并判断当前用户是否点赞或收藏该文章
	for i := 0; i < len(articles); i++ {
		okSelect, err := redis.IsUserCollected(strconv.Itoa(int(articles[i].ID)), username)
		okLike, err := redis.IsUserLiked(strconv.Itoa(articles[i].LikeAmount), username, 0)
		if err != nil {
			fmt.Println("SelectArticleAndUserListByPageFirstPageService() service.article.IsUserSelectedService err=", err)
			return nil, err
		}
		articles[i].IsCollect = okSelect
		articles[i].IsLike = okLike

		// 计算发布时间
		articles[i].PostTime = timeConverter.IntervalConversion(articles[i].CreatedAt)
	}
	return articles, nil
}
