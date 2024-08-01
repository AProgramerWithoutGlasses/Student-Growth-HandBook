package gorm_model

import "gorm.io/gorm"

type Follow struct {
	gorm.Model
	Username string `gorm:"size:100"`
	UserId   int
	del      bool `gorm:"default:false"`
}
