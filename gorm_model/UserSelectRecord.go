package gorm_model

import "gorm.io/gorm"

type UserSelectRecord struct {
	gorm.Model
	UserID    uint    //收藏属于用户
	User      User    //收藏属于用户
	ArticleID uint    //收藏属于文章
	Article   Article //收藏属于文章
	IsRead    bool    `gorm:"default:false"`
}
