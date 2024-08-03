package gorm_model

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Content      string        `json:"content"`
	UpvoteAmount int           `gorm:"default:0"json:"UpvoteAmount"`
	IsRead       bool          `gorm:"default:false"json:"IsRead"`
	Pid          uint          //回复评论的ID
	ArticleID    uint          //评论属于文章
	Article      Article       //评论属于文章
	CommentLikes []CommentLike //评论拥有点赞
}
