package gorm_model

import "gorm.io/gorm"

type SysNotification struct {
	gorm.Model
	UserID     uint
	NoticeType int
	Content    string
	IsRead     bool `gorm:"default:false"`
}
