package gorm_model

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Content      string `json:"content"`
	LikeAmount   int    `gorm:"default:0"json:"UpvoteAmount"`
	IsRead       bool   `gorm:"default:false"json:"IsRead"`
	UserID       uint   `gorm:"not null"`
	User         User
	Pid          uint                    //回复评论的ID
	ArticleID    uint                    `gorm:"not null"` //评论属于文章
	Article      Article                 //评论属于文章
	CommentLikes []UserCommentLikeRecord //评论拥有点赞
	IsLike       bool                    `gorm:"-"`
}
