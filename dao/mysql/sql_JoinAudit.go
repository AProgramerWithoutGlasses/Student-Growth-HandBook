package mysql

import (
	"errors"
	"gorm.io/gorm"
	"studentGrow/models/gorm_model"
	"time"
)

type Pagination struct {
	Page  int    `form:"page"`
	Limit int    `form:"limit"`
	Sort  string `form:"sort"`
}

// OpenActivityMsg 查询活动的开放信息
func OpenActivityMsg() (bool, string, gorm_model.JoinAuditDuty) {
	var count int64
	var ActivityMsg gorm_model.JoinAuditDuty
	var FailActivityMsg gorm_model.JoinAuditDuty
	DB.Where("is_show = ?", true).Find(&ActivityMsg).Count(&count)
	if count < 1 {
		return false, "不存在开放活动", FailActivityMsg
	}
	if count > 1 {
		return false, "开放活动数量异常", FailActivityMsg
	}
	format := "2006-01-02 15:04:05"
	now := time.Now().Format(format)
	startTime := ActivityMsg.StartTime.Format(format)
	stopTime := ActivityMsg.StopTime.Format(format)
	if now <= startTime {
		return false, "活动未到开放时间", ActivityMsg
	}
	if now >= stopTime {
		return false, "活动已结束", ActivityMsg
	}
	return true, "活动已开放", ActivityMsg
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

// 分页查询
//func ComList[T any](model T, pagMsg Pagination) (list []T) {
//	if pagMsg.Sort == "" || pagMsg.Sort == "desc" {
//		pagMsg.Sort = "created_at desc"
//	}
//
//}

//申请人信息保存

//查询纪权部所需信息

//查询组织部所需信息

//查询各期负责人
