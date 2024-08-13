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
	HeadShot              string                  `gorm:"size:300"json:"user_headShot"`
	Gender                string                  `gorm:"size:10"json:"user_gender"`
	College               string                  `gorm:"size:100"json:"user_college"`
	Class                 string                  `gorm:"size:100"json:"user_class"`
	PhoneNumber           string                  `gorm:"size:100"json:"phone_number"`
	MailBox               string                  `gorm:"size:100"json:"user_mail"`
	PlusTime              time.Time               `gorm:"type:date"json:"plus_time"`
	Identity              string                  `gorm:"not null;size:100"json:"user_identity"`
	Point                 int                     `gorm:"default:0"json:"user_point"`
	SelfContent           string                  `gorm:"size:1000"json:"self_content"`
	Motto                 string                  `gorm:"size:50"json:"user_motto"`
	Exper                 int                     `gorm:"default:0"json:"user_exper"`
	Ban                   bool                    `gorm:"type:boolean;default:false"json:"user_ban"`
	IsManager             bool                    `gorm:"default:false"json:"user_is_manager"`
	UserPublisherRecordID uint                    // 用户属于用户添加者
	UserPublisherRecord   UserPublisherRecord     // 用户属于用户添加者
	Followers             []User                  `gorm:"many2many:user_followers"` //用户和用户之间的关注关系
	Articles              []Article               //用户拥有的文章列表
	ReadRecords           []UserReadRecord        //用户浏览记录
	Collect               []UserCollectRecord     //用户拥有收藏
	UserLoginRecords      []UserLoginRecord       // 用户拥有登录记录
	ArticleLikes          []UserArticleLikeRecord //用户拥有文章点赞记录
	CommentLikes          []UserCommentLikeRecord // 用户拥有评论点赞记录
	Comments              []Comment               //用户拥有评论
}
