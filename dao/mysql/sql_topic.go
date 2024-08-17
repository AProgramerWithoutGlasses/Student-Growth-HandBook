package mysql

import (
	"go.uber.org/zap"
	"studentGrow/models/gorm_model"
	myErr "studentGrow/pkg/error"
)

// QueryAllTopics 获取所有话题
func QueryAllTopics() ([]gorm_model.Topic, error) {
	var topics []gorm_model.Topic
	if err := DB.Find(&topics).Error; err != nil {
		zap.L().Error("QueryAllTopics() dao.mysql.sql_topic err=", zap.Error(err))
		return nil, err
	}
	if len(topics) == 0 {
		zap.L().Error("QueryAllTopics() dao.mysql.sql_topic err=", zap.Error(myErr.NotFoundError()))
		return nil, myErr.NotFoundError()
	}
	return topics, nil
}

// QueryTagsByTopic 获取话题对应的标签
func QueryTagsByTopic(id int) ([]gorm_model.Tag, error) {
	var tags []gorm_model.Tag
	if err := DB.Where("topic_id = ?", id).Find(&tags).Error; err != nil {
		zap.L().Error("QueryAllTopics() dao.mysql.sql_topic err=", zap.Error(err))
		return nil, err
	}

	if len(tags) == 0 {
		zap.L().Error("QueryTagsByTopic() dao.mysql.sql_topic err=", zap.Error(myErr.NotFoundError()))
		return nil, myErr.NotFoundError()
	}

	return tags, nil
}

// QueryTagIdByTagName 根据标签名字查找标签ID
func QueryTagIdByTagName(name string) (int, error) {
	var tag gorm_model.Tag
	if err := DB.Where("tag_name = ?", name).First(&tag).Error; err != nil {
		zap.L().Error("QueryTagIdByTagName() dao.mysql.sql_topic err=", zap.Error(myErr.NotFoundError()))
		return -1, myErr.NotFoundError()
	}
	return int(tag.ID), nil
}

// InsertArticleTags 添加文章标签
func InsertArticleTags(tags []string, articleId int) error {
	for _, tag := range tags {
		tagId, err := QueryTagIdByTagName(tag)
		if err != nil {
			zap.L().Error("InsertArticleTags() dao.mysql.sql_topic.QueryTagIdByTagName err=", zap.Error(err))
			return err
		}
		articleTag := gorm_model.ArticleTag{
			ArticleID: uint(articleId),
			TagID:     uint(tagId),
		}
		if err = DB.Create(&articleTag).Error; err != nil {
			zap.L().Error("InsertArticleTags() dao.mysql.sql_topic.Create err=", zap.Error(err))
			return err
		}
	}
	return nil
}

// QueryTopicByArticleId 通过文章id查询文章的话题
func QueryTopicByArticleId(aid int) (string, error) {
	var article gorm_model.Article
	if err := DB.Where("id = ?", aid).First(&article).Error; err != nil {
		zap.L().Error("QueryTopicByArticleId() dao.mysql.sql_topic.First err=", zap.Error(err))
		return "", err
	}
	return article.Topic, nil
}

// QueryTopicIdByTopicName 通过topicName查询topicId
func QueryTopicIdByTopicName(topicName string) (int, error) {
	var topic gorm_model.Topic
	if err := DB.Where("topic_name = ?", topicName).First(&topic).Error; err != nil {
		zap.L().Error("QueryTopicByArticleId() dao.mysql.sql_topic.First err=", zap.Error(err))
		return -1, err
	}
	return int(topic.ID), nil
}
