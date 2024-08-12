package gorm_model

import (
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Content       string `gorm:"size:350"json:"content"`
	WordCount     int    `gorm:"default:0"json:"wordCount"`
	Pic           string
	Video         string
	Topic         string                  `json:"topic"`
	Status        bool                    `gorm:"not null;default:true"json:"status"`
	ReadAmount    int                     `gorm:"default:0"json:"readAmount"`
	LikeAmount    int                     `gorm:"default:0"json:"LikeAmount"`
	CollectAmount int                     `gorm:"default:0"json:"collectAmount"`
	CommentAmount int                     `gorm:"default:0"json:"commentAmount"`
	ReportAmount  int                     `gorm:"default:0"json:"reportAmount"`
	Ban           bool                    `gorm:"default:false"json:"-"`
	UserID        uint                    `gorm:"not null"` //文章属于用户
	User          User                    `json:"user"`     //文章属于用户
	Comments      []Comment               //文章拥有评论
	ArticleLikes  []UserArticleLikeRecord //文章拥有点赞
	ArticleTags   []ArticleTag            //文章拥有标签
	Selects       []UserSelectRecord      //文章拥有收藏
}

// Articles 自定义加权排序
type Articles []Article

func (a Articles) Len() int      { return len(a) }
func (a Articles) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a Articles) Less(i, j int) bool {
	return float64(a[i].LikeAmount)*0.5+float64(a[i].CollectAmount)*0.2+float64(a[i].CommentAmount)*0.3 > float64(a[j].LikeAmount)*0.5+float64(a[j].CollectAmount)*0.2+float64(a[j].CommentAmount)*0.3 // 降序
}
