package gorm_model

import (
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Content       string `gorm:"size:350"json:"content"`
	WordCount     int    `gorm:"not null"json:"wordCount"`
	Pic           string
	Video         string
	Topic         string    `json:"topic"`
	Tag           string    `json:"tag"`
	Status        bool      `gorm:"not null;default:true"json:"status"`
	ReadAmount    int       `gorm:"default:0"json:"readAmount"`
	UpvoteAmount  int       `gorm:"default:0"json:"upvoteAmount"`
	CollectAmount int       `gorm:"default:0"json:"collectAmount"`
	CommentAmount int       `gorm:"default:0"json:"commentAmount"`
	ReportAmount  int       `gorm:"default:0"json:"reportAmount"`
	Ban           bool      `gorm:"default:false"json:"-"`
	Del           bool      `gorm:"default:false"json:"-"`
	UserId        int       `gorm:"not null"json:"-"`
	User          User      `gorm:"foreignKey:UserId"json:"user"`
	Comment       []Comment `gorm:"foreignKey:Aid"json:"-"`
	Read          []Read    `gorm:"foreignKey:Aid"json:"-"`
	Upvote        []Upvote  `gorm:"foreignKey:Aid"json:"-"`
}
