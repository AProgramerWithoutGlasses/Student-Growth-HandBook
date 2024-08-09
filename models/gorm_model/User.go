package gorm_model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Username              string                  `gorm:"not null;unique"json:"username"`
	Password              string                  `gorm:"not null"json:"password"`
	Name                  string                  `gorm:"size:100"json:"name"`
	HeadShot              string                  `gorm:"size:300"json:"headShot"`
	Gender                string                  `gorm:"size:10"json:"gender"`
	College               string                  `gorm:"size:100"json:"college"`
	Class                 string                  `gorm:"size:100"json:"class"`
	PhoneNumber           string                  `gorm:"size:100"json:"phoneNumber"`
	MailBox               string                  `gorm:"size:100"json:"mailBox"`
	PlusTime              time.Time               `gorm:"type:date"json:"plusTime"`
	Identity              string                  `gorm:"not null;size:100"json:"identity"`
	Point                 int                     `gorm:"default:0"json:"point"`
	SelfContent           string                  `gorm:"size:1000"json:"selfContent"`
	Motto                 string                  `gorm:"size:50"json:"motto"`
	Exper                 int                     `gorm:"default:0"json:"exper"`
	Ban                   bool                    `gorm:"type:boolean;default:false"json:"ban"`
	IsManager             bool                    `gorm:"default:false"json:"isManager"`
	UserPublisherRecordID uint                    // 用户属于用户添加者
	UserPublisherRecord   UserPublisherRecord     // 用户属于用户添加者
	Followers             []User                  `gorm:"many2many:user_followers"` //用户和用户之间的关注关系
	Articles              []Article               //用户拥有的文章列表
	ReadRecords           []UserReadRecord        //用户浏览记录
	Selects               []UserSelectRecord      //用户拥有收藏
	UserLoginRecords      []UserLoginRecord       // 用户拥有登录记录
	ArticleLikes          []UserArticleLikeRecord //用户拥有文章点赞记录
	CommentLikes          []UserCommentLikeRecord // 用户拥有评论点赞记录
}
