package gorm_model

import "gorm.io/gorm"

type UserCasbinRules struct {
	gorm.Model
	CUsername string
	CasbinCid string
	Status    bool `gorm:"default:false"`
}
