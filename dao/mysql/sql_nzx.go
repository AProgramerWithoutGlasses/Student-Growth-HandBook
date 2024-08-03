package mysql

import (
	"errors"
	"fmt"
	_ "gorm.io/gorm"
	model "studentGrow/models/gorm_model"
	"time"
)

// SelectUserById 查询数据库是否存在该用户
func SelectUserById(uid int) (err error, user *model.User) {
	//select * from users where id = uid
	// 查询用户
	if err := DB.Where("id = ?", uid).First(&user).Error; err != nil {
		fmt.Println("SelectUserById() dao.mysql err=", err)
		return err, nil
	} else {
		return nil, user
	}
}

// SelectUserByUsername 通过username查找uid
func SelectUserByUsername(username string) (uid int, err error) {
	//select id from users where username = username
	var user model.User
	if err := DB.Model(model.User{}).Select("id").Where("username = ?", username).Find(&user).Error; err != nil {
		fmt.Println("Error:", err)
		return int(user.ID), err
	} else {
		fmt.Println("user.ID", int(user.ID))
		return int(user.ID), nil
	}
}

// SelectArticleById 通过id查找文章
func SelectArticleById(aid int) (err error, article *model.Article) {
	//查询用户 select * from articles where id = aid
	fmt.Println("id:", aid)
	if err := DB.Where("id = ?", aid).First(&article).Error; err != nil {
		// 处理查询错误
		fmt.Println("Error:", err)
		return err, nil
	} else {
		return nil, article
	}
}

// InsertIntoCommentsForArticle 向数据库插入评论数据(回复文章)

func InsertIntoCommentsForArticle(content string, aid int, uid int) (err error) {
	//content;id;username
	//comment := model.Comment{
	//	Model:        gorm.Model{},
	//	Content:      content,
	//	UpvoteAmount: 0,
	//	IsRead:       false,
	//	Del:          false,
	//	Uid:          uid,
	//	Pid:          0,
	//	Aid:          aid,
	//	Upvote:       nil,
	//}
	//DB.Create(&comment)

	return nil
}

// InsertIntoCommentsForComment 向数据库插入评论数据(回复评论)
func InsertIntoCommentsForComment(content string, uid int, pid int) (err error) {
	//content;id;username
	//comment := model.Comment{
	//	Model:        gorm.Model{},
	//	Content:      content,
	//	UpvoteAmount: 0,
	//	IsRead:       false,
	//	Del:          false,
	//	Uid:          uid,
	//	Pid:          pid,
	//	Aid:          0,
	//	Upvote:       nil,
	//}
	//
	//DB.Create(&comment)
	return nil
}

// SelectTagsByTopic 通过话题查询该话题对应的标签列表
func SelectTagsByTopic(topic string) (tags []map[string]any) {
	//id = select id from article_topics where topic = topic
	//tags = select tag from article_tags where article_topic_id = id

	//var articleTopic model.ArticleTopic
	//var articleTags []model.ArticleTag
	////查询对应的话题ID
	//if err := DB.Where("topic = ?", topic).First(&articleTopic).Error; err != nil {
	//	fmt.Println("SelectTagsByTopic() dao.mysql.sql_nzx.Error=", err)
	//}
	//if err := DB.Where("article_topic_id = ?", articleTopic.ID).Find(&articleTags).Error; err != nil {
	//	fmt.Println("SelectTagsByTopic() dao.mysql.sql_nzx.Error=", err)
	//}
	//
	//for index, tag := range articleTags {
	//	tags = append(tags, map[string]any{
	//		"id":   index,
	//		"name": tag.Tag,
	//	})
	//}
	//
	return tags
}

// SelectArticleAndUserListByPage 分页查询文章及用户列表
func SelectArticleAndUserListByPage(page, limit int, sort, order string) (result []model.Article, err error) {
	//SELECT articles.*, users.*
	//FROM articles
	//JOIN users ON articles.user_id = users.id
	//WHERE articles.created_at > (
	//    SELECT created_at
	//    FROM articles
	//    ORDER BY created_at DESC
	//    LIMIT ?, 1
	//)
	//LIMIT ?;

	var articles []model.Article
	if order == "asc" {
		var createdAt time.Time
		_ = DB.Model(&model.Article{}).
			Select("created_at").
			Order("created_at ASC").
			Limit(limit).
			Offset((page - 1) * limit).
			Scan(&createdAt).Error

		DB.InnerJoins("User").Where("articles.created_at > ?", createdAt).Find(&articles)
	} else {
		var createdAt time.Time
		_ = DB.Model(&model.Article{}).
			Select("created_at").
			Order("created_at DESC").
			Limit(limit).
			Offset((page - 1) * limit).
			Scan(&createdAt).Error
		DB.InnerJoins("User").Where("articles.created_at < ?", createdAt).Find(&articles)
	}

	if len(articles) <= 0 {
		return nil, errors.New("no records")
	}

	return articles, nil
}

// SelectAllTopics 查找所有的话题
//func SelectAllTopics() (topics []map[string]any) {
//	//select topic from article_topic
//
//	var articleTopic model.ArticleTopic
//
//	if err := DB.Find(&articleTopic).Error; err != nil {
//		fmt.Println("SelectAllTopics() dao.mysql.sql_nzx.Error=", err)
//	}
//
//}
