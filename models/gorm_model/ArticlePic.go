package gorm_model

import "gorm.io/gorm"

type ArticlePic struct {
	gorm.Model
	Pic       string `json:"pic"`
	ArticleID uint
	Article   Article
}
