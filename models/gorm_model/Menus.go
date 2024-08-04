package gorm_model

import (
	"gorm.io/gorm"
	"time"
)

type Menus struct {
	gorm.Model
	ParentId      int
	TreePath      string
	Name          string
	Type          int
	Path          string
	Component     string
	Perm          string
	Visible       int
	Sort          int
	Icon          string
	Redirect      string
	CreateTime    time.Time
	UpdateTime    time.Time
	Roles         string
	AlwaysShow    int
	KeepAlive     int
	DeletedAt     time.Time
	RequestUrl    string
	RequestMethod string
}
