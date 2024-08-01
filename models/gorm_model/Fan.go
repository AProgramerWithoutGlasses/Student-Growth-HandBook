package gorm_model

import "gorm.io/gorm"

type Fan struct {
	gorm.Model
	Username string `gorm:"not null"`
	Status   bool   `gorm:"default:false"`
	UserId   int
}
