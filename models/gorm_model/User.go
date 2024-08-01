package gorm_model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
<<<<<<< HEAD
	Username    string       `gorm:"not null;unique"`
	Password    string       `gorm:"not null"`
	Name        string       `gorm:"size:100"`
	HeadShot    string       `gorm:"size:300"`
	Gender      string       `gorm:"size:10"`
	College     string       `gorm:"size:100"`
	Class       string       `gorm:"size:100"`
	PhoneNumber string       `gorm:"size:100"`
	MailBox     string       `gorm:"size:100"`
	PlusTime    string       `gorm:"type:date"`
	Identity    string       `gorm:"not null;size:100"`
	Point       int          `gorm:"default:0"`
	Ban         bool         `gorm:"type:boolean;default:false"`
	Del         bool         `gorm:"type:boolean;default:false"`
	FanList     []Fan        `gorm:"foreignKey:UserId"`
	Follow      []Follow     `gorm:"foreignKey:UserId"`
	Article     []Article    `gorm:"foreignKey:UserId"`
	Comment     []Comment    `gorm:"foreignKey:Uid"`
	Read        []Read       `gorm:"foreignKey:Uid"`
	Upvote      []Upvote     `gorm:"foreignKey:Uid"`
	CasbinRule  []CasbinRule `gorm:"many2many:casbin_rules_users;"`
	SelfContent string       `gorm:"size:1000"`
	Motto       string       `gorm:"size:50"`
	Exper       int          `gorm:"default:0"`
=======
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
	Article     []Article    `gorm:"foreignKey:UserId"json:"-"`
	Comment     []Comment    `gorm:"foreignKey:Uid"json:"-"`
	Read        []Read       `gorm:"foreignKey:Uid"json:"-"`
	Upvote      []Upvote     `gorm:"foreignKey:Uid"json:"-"`
	CasbinRule  []CasbinRule `gorm:"many2many:casbin_rules_users;"json:"-"`
>>>>>>> bd64b59feb8245f5364f131e7324b0194666ecf9
}
