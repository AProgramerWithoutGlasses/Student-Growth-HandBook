package gorm_model

import (
	"gorm.io/gorm"
)

type CommentLike struct {
	gorm.Model
	CommentID uint    //点赞属于评论
	Comment   Comment //点赞属于评论
	IsRead    bool    `gorm:"default:false"`
}
