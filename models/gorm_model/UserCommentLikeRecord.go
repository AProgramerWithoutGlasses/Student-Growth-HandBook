package gorm_model

import (
	"gorm.io/gorm"
)

type UserCommentLikeRecord struct {
	gorm.Model
	CommentID uint    `gorm:"not null"` //点赞属于评论
	Comment   Comment //点赞属于评论
	UserID    uint    `gorm:"not null"`
	User      User
	IsRead    bool `gorm:"default:false"`
}
