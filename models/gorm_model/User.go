package gorm_model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
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
}
