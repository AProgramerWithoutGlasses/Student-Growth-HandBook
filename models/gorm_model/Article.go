package gorm_model

import (
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
<<<<<<< HEAD
	Content       string `gorm:"size:350"`
	WordCount     int    `gorm:"not null"`
=======
	Content       string `gorm:"size:350"json:"content"`
	WordCount     int    `gorm:"not null"json:"wordCount"`
>>>>>>> origin/feature/xun
	Pic           string
	Video         string
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
}
