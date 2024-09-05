package gorm_model

import (
	"gorm.io/gorm"
)

type UserEditRecord struct {
	gorm.Model
	Username     string `json:"username"`
	EditUsername string `json:"add_username"`
	EditMessage  string `json:"edit_message"`
}
