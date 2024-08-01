package gorm_model

import (
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
<<<<<<< HEAD
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
	commentAmount int  `gorm:"default:0"`
	ReportAmount  int  `gorm:"default:0"`
	Ban           bool `gorm:"default:false"`
	Del           bool `gorm:"default:false"`
	UserId        int
	Comment       []Comment `gorm:"foreignKey:Aid"`
	Read          []Read    `gorm:"foreignKey:Aid"`
	Upvote        []Upvote  `gorm:"foreignKey:Aid"`
	UpvoteAmount  int       `gorm:"default:0"`
=======
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
	UserId        int       `json:"-"`
	Comment       []Comment `gorm:"foreignKey:Aid"json:"-"`
	Read          []Read    `gorm:"foreignKey:Aid"json:"-"`
	Upvote        []Upvote  `gorm:"foreignKey:Aid"json:"-"`
>>>>>>> bd64b59feb8245f5364f131e7324b0194666ecf9
}
