package gorm_model

import (
	"gorm.io/gorm"
)

type ArticleTag struct {
	gorm.Model
	TagName   string  `json:"TagName"`
	ArticleID uint    //标签属于文章
	Article   Article //标签属于文章
}
