package gorm_model

import "gorm.io/gorm"

type UserLoginRecord struct {
	gorm.Model
	Username string
	UserID   uint // 用户登录记录属于用户
	User     User // 用户登录记录属于用户
}
