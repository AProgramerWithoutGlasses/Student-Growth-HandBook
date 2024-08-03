package gorm_model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username    string       `gorm:"not null;unique"json:"username"`
	Password    string       `gorm:"not null"json:"password"`
	Name        string       `gorm:"size:100"json:"name"`
	HeadShot    string       `gorm:"size:300"json:"headShot"`
	Gender      string       `gorm:"size:10"json:"gender"`
	College     string       `gorm:"size:100"json:"college"`
	Class       string       `gorm:"size:100"json:"class"`
	PhoneNumber string       `gorm:"size:100"json:"phoneNumber"`
	MailBox     string       `gorm:"size:100"json:"mailBox"`
	PlusTime    string       `gorm:"type:date"json:"plusTime"`
	Identity    string       `gorm:"not null;size:100"json:"identity"`
	Point       int          `gorm:"default:0"json:"point"`
	SelfContent string       `gorm:"size:1000"json:"selfContent"`
	Motto       string       `gorm:"size:50"json:"motto"`
	Exper       int          `gorm:"default:0"json:"exper"`
	Ban         bool         `gorm:"type:boolean;default:false"json:"-"`
	Del         bool         `gorm:"type:boolean;default:false"json:"-"`
	FanList     []Fan        `gorm:"foreignKey:UserId"json:"-"`
	Follow      []Follow     `gorm:"foreignKey:UserId"json:"-"`
	Articles    []Article    `gorm:"foreignKey:UserId"json:"-"`
	Comment     []Comment    `gorm:"foreignKey:Uid"json:"-"`
	Read        []Read       `gorm:"foreignKey:Uid"json:"-"`
	Upvote      []Upvote     `gorm:"foreignKey:Uid"json:"-"`
	CasbinRule  []CasbinRule `gorm:"many2many:casbin_rules_users;"json:"-"`
}
