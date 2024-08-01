package gorm_model

import "gorm.io/gorm"

type ArticleTopic struct {
	gorm.Model
	Topic      string
	ArticleTag []ArticleTag `gorm:"foreignKey:TopicId"`
}
