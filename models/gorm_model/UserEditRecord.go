package gorm_model

import (
	"gorm.io/gorm"
)

type UserEditRecord struct {
	gorm.Model
	Username       string `json:"username"`
	EditedUsername string `json:"edited_username"`
	OldClass       string `json:"old_class"`
	NewClass       string `json:"new_class"`
	OldPhoneNumber string `json:"old_phone_number"`
	NewPhoneNumber string `json:"new_phone_number"`
}
