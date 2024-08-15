package gorm_model

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	Topic   string `json:"topic"`
	TagName string `json:"tag_name"`
}
