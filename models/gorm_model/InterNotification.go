package gorm_model

import (
	"gorm.io/gorm"
)

type InterNotification struct {
	gorm.Model
	TarUserId  uint   // 目标用户ID
	OwnUserId  uint   // 拥有者ID
	NoticeType int    // 消息类型
	SuperType  int    // 父级消息类型
	SuperId    int    // 父级消息ID
	Content    string // 消息内容
	IsRead     bool   `gorm:"default:false"`
	Status     bool   `gorm:"default:true"`
}
