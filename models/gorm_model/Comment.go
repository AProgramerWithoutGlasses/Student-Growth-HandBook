package gorm_model

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Content      string
	UpvoteAmount int  `gorm:"default:0"`
	IsRead       bool `gorm:"default:false"`
	Del          bool `gorm:"default:false"`
	Uid          int
	Pid          int
	Aid          int
	Upvote       []Upvote `gorm:"foreignKey:Cid"`
}
