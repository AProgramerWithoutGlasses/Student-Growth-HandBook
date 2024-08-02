package gorm_model

type CasbinRule struct {
	Id    int    `gorm:"primaryKey;autoIncrement"`
	PType string `gorm:"not null"`
	V0    string `gorm:"size:100"`
	V1    string `gorm:"size:100"`
	V2    string `gorm:"size:100"`
	V3    string `gorm:"size:100"`
	V4    string `gorm:"size:100"`
	V5    string `gorm:"size:100"`
}
