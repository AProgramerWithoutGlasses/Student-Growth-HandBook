package gorm_model

import (
	"gorm.io/gorm"
)

type Read struct {
	gorm.Model
	Uid int
	Aid int
}
