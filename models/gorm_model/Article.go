package gorm_model

import (
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Content       string `gorm:"size:350"`
	WordCount     int    `gorm:"not null"`
	Pic           string
	Video         string
	Topic         string
	Tag           string
	Status        bool `gorm:"not null;default:true"`
	ReadAmount    int  `gorm:"default:0"`
	UpvoteCount   int  `gorm:"default:0"`
	CollectAmount int  `gorm:"default:0"`
	CommentAmount int  `gorm:"default:0"`
	ReportAmount  int  `gorm:"default:0"`
	Ban           bool `gorm:"default:false"`
	Del           bool `gorm:"default:false"`
	UserId        int
	Comment       []Comment `gorm:"foreignKey:Aid"`
	//Read          []Read    `gorm:"foreignKey:Aid"`
	//Upvote        []Upvote  `gorm:"foreignKey:Aid"`
	UpvoteAmount int `gorm:"default:0"`
}
