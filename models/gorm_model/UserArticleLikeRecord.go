package gorm_model

import "gorm.io/gorm"

type UserArticleLikeRecord struct {
	gorm.Model
	ArticleID uint    `gorm:"not null"` //点赞属于文章
	Article   Article //点赞属于文章
	UserID    uint    `gorm:"not null"` // 点赞者用户ID
	User      User
	IsRead    bool `gorm:"default:false"` //文章发布者是否已读
}
