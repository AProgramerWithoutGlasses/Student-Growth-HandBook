package gorm_model

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Content      string `json:"content"`
	UpvoteAmount int    `gorm:"default:0"json:"UpvoteAmount"`
	IsRead       bool   `gorm:"default:false"json:"IsRead"`
	Del          bool   `gorm:"default:false"`
	Uid          int
	Pid          int
	Aid          int
	Upvote       []Upvote `gorm:"foreignKey:Cid"`
}
