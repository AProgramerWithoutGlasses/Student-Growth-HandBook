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
<<<<<<< HEAD
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
=======
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
>>>>>>> bbd4d3eaa3b86d0900cb2b387d9481155a2f2743
}
