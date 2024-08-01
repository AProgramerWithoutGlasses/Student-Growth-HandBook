package gorm_model

import (
	"gorm.io/gorm"
)

type Upvote struct {
	gorm.Model
	Type   string `gorm:"not null"`
	IsRead bool   `gorm:"default:false"`
	Uid    int
	Aid    int
	Cid    int
	Del    bool `form:"default:false"`
}
