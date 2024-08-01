package gorm_model

import (
	"gorm.io/gorm"
)

type ArticleTag struct {
	gorm.Model
	Tag     string `json:"tag"`
	TopicId int    `json:"-"`
}
