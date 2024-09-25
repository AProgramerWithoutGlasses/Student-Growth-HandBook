package gorm_model

import (
	"gorm.io/gorm"
)

type UserEditRecord struct {
	gorm.Model
	Username       string `json:"username"`
	EditUsername   string `json:"add_username"`
	OldClass       string `json:"old_class"`
	NewClass       string `json:"new_class"`
	OldPhoneNumber string `json:"old_phone_number"`
	NewPhoneNumber string `json:"new_phone_number"`
	OldPassword    string `json:"old_password"`
	NewPassword    string `json:"new_password"`
}
