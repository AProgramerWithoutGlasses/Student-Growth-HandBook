package gorm_model

import (
	"gorm.io/gorm"
)

type UserReadRecord struct {
	gorm.Model
	UserID uint //属于
	Aid    int
}
