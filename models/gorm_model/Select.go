package gorm_model

import "gorm.io/gorm"

type Select struct {
	gorm.Model
	UserID uint //收藏属于用户
	User   User //收藏属于用户
	IsRead bool `gorm:"default:false"`
}
