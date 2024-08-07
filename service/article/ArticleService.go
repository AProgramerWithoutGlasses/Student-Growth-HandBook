package article

import (
	"fmt"
	jsonvalue "github.com/Andrew-M-C/go.jsonvalue"
	"github.com/pkg/errors"
	"studentGrow/dao/mysql"
	"studentGrow/dao/redis"
	model "studentGrow/models/gorm_model"
	myErr "studentGrow/pkg/error"
)

// GetArticleService 获取文章详情
func GetArticleService(j *jsonvalue.V) (err error, user *model.User, article *model.Article) {
	//获取文章id
	aid, err := j.GetInt("article_id")
	if err != nil {
		fmt.Println("GetArticleService() dao.mysql.sqp_nzx.GetInt err=", err)
		return
	}
	//查找文章信息
	err, article = mysql.SelectArticleById(aid)
	if err != nil {
		fmt.Println("GetArticleService() dao.mysql.sqp_nzx.SelectArticleById err=", err)
		return err, nil, nil
	}
	//查找用户信息
	err, user = mysql.SelectUserById(int(article.UserID))

	if err != nil {
		fmt.Println("GetArticleService() dao.mysql.sqp_nzx.SelectUserById err=", err)
		return err, nil, nil
	}

	return nil, user, article
}

// GetArticleListService 获取文章列表
func GetArticleListService(j *jsonvalue.V) ([]model.Article, error) {
	//获取参数：page, limit, sort, username
	page, _ := j.GetInt("page")
	limit, _ := j.GetInt("limit")
	sort, _ := j.GetString("sort")
	order, _ := j.GetString("order")
	//获取发布时间、封禁状态、发布人、话题分类、关键词
	startAt, err := j.GetString("start_at")
	if err != nil {
		fmt.Println("SearchArticleService() service.article.GetString err=", err)
		return nil, err
	}
	endAt, err := j.GetString("end_at")
	if err != nil {
		fmt.Println("SearchArticleService() service.article.GetString err=", err)
		return nil, err
	}
	isBan, err := j.GetBool("article_ban")
	if err != nil {
		fmt.Println("SearchArticleService() service.article.GetBool err=", err)
		return nil, err
	}
	name, err := j.GetString("name")
	if err != nil {
		fmt.Println("SearchArticleService() service.article.GetString err=", err)
		return nil, err
	}
	topic, err := j.GetString("topic")
	if err != nil {
		fmt.Println("SearchArticleService() service.article.GetString err=", err)
		return nil, err
	}
	keyWords, err := j.GetString("key_words")
	if err != nil {
		fmt.Println("SearchArticleService() service.article.GetString err=", err)
		return nil, err
	}
	//执行查询文章列表语句
	result, err := mysql.SelectArticleAndUserListByPage(page, limit, sort, order, startAt, endAt, topic, keyWords, name, isBan)
	if err != nil {
		fmt.Println("GetArticleListService() service.article.GetString err=", err)
		return nil, myErr.NotFoundError()
	}

	return result, nil
}

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
func BannedArticleService(j *jsonvalue.V) error {
	// 获取文章id和封禁状态
	id, err := j.GetInt("article_id")
	if err != nil {
		fmt.Println("BannedArticleService() service.article.GetString err=", err)
		return err
	}
	isBan, err := j.GetBool("article_ban")

	// 修改文章封禁状态
	err = mysql.BannedArticleById(id, isBan)
	if err != nil {
		fmt.Println("BannedArticleService() service.article.BannedArticleById err=", err)
		return err
	}

	return nil
}

// DeleteArticleService 删除文章
func DeleteArticleService(j *jsonvalue.V) error {
	// 获取文章id
	id, err := j.GetInt("article_id")
	if err != nil {
		fmt.Println("DeleteArticleService() service.article.GetInt err=", err)
		return err
	}

	// 删除文章
	err = mysql.DeleteArticleById(id)
	if err != nil {
		fmt.Println("DeleteArticleService() service.article.DeleteArticleById err=", err)
		return err
	}
	return nil
}

// ReportArticleService 举报文章
func ReportArticleService(j *jsonvalue.V) error {
	// 获取文章id和举报者用户id
	aid, err := j.GetInt("article_id")
	if err != nil {
		fmt.Println("ReportArticleService() service.article.GetInt err=", err)
		return err
	}

	username, err := j.GetString("username")
	if err != nil {
		fmt.Println("ReportArticleService() service.article.GetString err=", err)
		return err
	}

	// 通过username查询id
	uid, err := mysql.SelectUserByUsername(username)
	if err != nil {
		fmt.Println("ReportArticleService() service.article.SelectUserByUsername err=", err)
		return err
	}

	err = mysql.ReportArticleById(aid, uid)
	if err != nil {
		fmt.Println("ReportArticleService() service.article.ReportArticleById err=", err)
		return err
	}
	return nil
}
