package mysql

import (
	"errors"
	"gorm.io/gorm"
	"studentGrow/models/gorm_model"
	"time"
)

// OpenActivityMsg 查询活动的开放信息
func OpenActivityMsg() (isOpen bool, Msg string, ActivityMsg gorm_model.JoinAuditDuty) {
	isOpen = false
	var count int64
	DB.Where("is_show = ?", true).Find(&ActivityMsg).Count(&count)
	if count < 1 {
		Msg = "不存在开放活动"
		return
	}
	if count > 1 {
		Msg = "开放活动数量异常"
		return
	}
	format := "2006-01-02 15:04:05"
	now := time.Now().Format(format)
	startTime := ActivityMsg.StartTime.Format(format)
	stopTime := ActivityMsg.StopTime.Format(format)
	if now <= startTime {
		Msg = "活动未到开放时间"
		return
	}
	if now >= stopTime {
		Msg = "活动已结束"
		return
	}
	isOpen = true
	Msg = "活动已开放"
	return
}

// StuFormMsg 查询提交过的信息
func StuFormMsg(username string, activityID uint) (isExist bool, stuMsg gorm_model.JoinAudit) {
	err := DB.Where("username = ? AND join_audit_duty_id = ?", username, activityID).First(&stuMsg).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		isExist = false
		return
	}
	isExist = true
	return
}

//申请人信息保存

//查询纪权部所需信息

//查询组织部所需信息

//查询各期负责人
