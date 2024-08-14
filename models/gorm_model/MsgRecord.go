package gorm_model

import (
	"gorm.io/gorm"
	"time"
)

type MsgRecord struct {
	gorm.Model
	Content string    `gorm:"not null" json:"msg_content"`
	Type    int       `gorm:"not null" json:"msg_type"`
	Time    time.Time `gorm:"-" json:"msg_time"`
	IsRead  bool      `gorm:"default:false" json:"is_read"`
	UserID  uint
	User    User
}
