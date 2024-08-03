package article

import (
	"fmt"
	jsonvalue "github.com/Andrew-M-C/go.jsonvalue"
	"github.com/pkg/errors"
	"studentGrow/dao/mysql"
	model "studentGrow/models/gorm_model"
)

// GetArticleService 获取文章详情
func GetArticleService(j *jsonvalue.V) (err error, user *model.User, article *model.Article) {
	//获取文章id
	aid, _ := j.GetInt("article_id")

	//查找文章信息
	err, article = mysql.SelectArticleById(aid)
	if err != nil {
		fmt.Println("GetArticleService() dao.mysql.sqp_nzx.SelectArticleById err=", err)
		return err, nil, nil
	}
	//查找用户信息
	err, user = mysql.SelectUserById(article.UserId)
	fmt.Println(article.UserId)
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
	//执行查询文章列表语句
	result, err := mysql.SelectArticleAndUserListByPage(page, limit, sort, order)
	if err != nil {
		fmt.Println("GetArticleListService() service.article.GetString err=", err)
		return nil, errors.New("no records")
	}

	return result, nil
}

// PublishArticleService 发布文章
//func PublishArticleService() map[string]any{
//
//}

// GetTopicTagsService GetTopicTags 根据话题获得对应的标签
func GetTopicTagsService(j *jsonvalue.V) []map[string]any {
	//解析获得话题
	topic, err := j.GetString("topic")
	fmt.Println(topic)
	if err != nil {
		fmt.Println("GetTopicTags() service.article.GetString err=", err)
		return nil
	}

	//查询标签
	return mysql.SelectTagsByTopic(topic)

}

//GetAllTopicsService 获取所有话题
//func GetAllTopicsService(j *jsonvalue.V) []map[string]any {
//
//}
