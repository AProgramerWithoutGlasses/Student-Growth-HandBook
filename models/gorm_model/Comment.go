package gorm_model

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
<<<<<<< HEAD
	Content      string
	UpvoteAmount int  `gorm:"default:0"`
	IsRead       bool `gorm:"default:false"`
	Del          bool `gorm:"default:false"`
=======
	Content      string `json:"content"`
	UpvoteAmount int    `gorm:"default:0"json:"UpvoteAmount"`
	IsRead       bool   `gorm:"default:false"json:"IsRead"`
	Del          bool   `gorm:"default:false"`
>>>>>>> bd64b59feb8245f5364f131e7324b0194666ecf9
	Uid          int
	Pid          int
	Aid          int
	Upvote       []Upvote `gorm:"foreignKey:Cid"`
}
