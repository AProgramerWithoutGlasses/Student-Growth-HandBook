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
	Topic         string        `json:"topic"`
	Status        bool          `gorm:"not null;default:true"json:"status"`
	ReadAmount    int           `gorm:"default:0"json:"readAmount"`
	LikeAmount    int           `gorm:"default:0"json:"LikeAmount"`
	CollectAmount int           `gorm:"default:0"json:"collectAmount"`
	CommentAmount int           `gorm:"default:0"json:"commentAmount"`
	ReportAmount  int           `gorm:"default:0"json:"reportAmount"`
	Ban           bool          `gorm:"default:false"json:"-"`
	UserID        uint          //文章属于用户
	User          User          `json:"user"` //文章属于用户
	Comments      []Comment     //文章拥有评论
	ArticleLikes  []ArticleLike //文章拥有点赞
	ArticleTags   []ArticleTag  //文章拥有标签
>>>>>>> 6820bb9dec9c9fbede6712769c244eca04b27ff7
}
