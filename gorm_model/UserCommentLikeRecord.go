package gorm_model

import (
	"gorm.io/gorm"
)

type UserCommentLikeRecord struct {
	gorm.Model
	CommentID uint    //点赞属于评论
	Comment   Comment //点赞属于评论
	UserID    uint
	User      User
	IsRead    bool `gorm:"default:false"`
}
