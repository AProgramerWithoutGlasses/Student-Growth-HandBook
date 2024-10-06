package gorm_model

import (
	"gorm.io/gorm"
)

type Notification struct {
	gorm.Model
	TarUserId  uint // 目标用户ID
	OwnUserId  uint // 拥有者ID
	NoticeType int  // 消息类型
	SuperType  int  // 父级消息类型
	SuperId    int  // 父级消息ID
	IsRead     bool `gorm:"default:false"`
	Status     int  `gorm:"default:true"`
}
