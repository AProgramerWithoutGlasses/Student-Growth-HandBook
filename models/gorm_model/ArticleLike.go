package gorm_model

import "gorm.io/gorm"

type ArticleLike struct {
	gorm.Model
	ArticleID uint    //点赞属于用户
	Article   Article //点赞属于用户
	IsRead    bool    `gorm:"default:false"` //文章发布者是否已读
}
