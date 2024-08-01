package gorm_model

import "gorm.io/gorm"

type Select struct {
	gorm.Model
	IsRead bool `gorm:"default:false"`
	Uid    int
	Aid    int
	Del    bool `form:"default:false"`
}
