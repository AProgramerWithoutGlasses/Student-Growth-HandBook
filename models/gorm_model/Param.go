package gorm_model

import "gorm.io/gorm"

type Param struct {
	gorm.Model
	ParamsKey   string `json:"paramsKey"`
	ParamsValue string `json:"paramsValue"`
	MenuId      int    `json:"menuId"`
}
