package gorm_model

import "gorm.io/gorm"

type JoinAuditFile struct {
	gorm.Model
	Username        string        `json:"username"` //照片归属
	FileName        string        `json:"name"`
	FileHash        string        `json:"hash"` //照片hash值
	FilePath        string        `json:"path"` //路径
	JoinAuditDutyID uint          `json:"join_audit_duty_id"`
	JoinAuditDuty   JoinAuditDuty `gorm:"foreignKey:JoinAuditDutyID"`
	Note            string        `json:"image_note"` //备注
}
