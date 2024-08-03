package gorm_model

import "gorm.io/gorm"

type UserCasbinRules struct {
	gorm.Model
	UserId   int
	CasbinId string
	part     string
}
