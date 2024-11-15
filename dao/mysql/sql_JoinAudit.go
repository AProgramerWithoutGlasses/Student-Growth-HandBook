package mysql

import (
	"errors"
	"gorm.io/gorm"
	"studentGrow/models/gorm_model"
	"time"
)

type Pagination struct {
	Page           int    `json:"page"`
	Label          string `json:"Label"`
	Limit          int    `json:"limit"`
	Sort           string `json:"sort"`
	PersonInCharge string `json:"person_in_charge"`
	ActivityName   string `json:"activity_name"`
	IsShow         bool   `json:"is_show"`
	StartTime      string `json:"start_time"`
	StopTime       string `json:"stop_time"`
	UserClass      string `json:"user_class"`
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

// ComList分页查询
func ComList[T any](model T, pagMsg Pagination) (list []T, count int64, err error) {
	if pagMsg.Sort == "asc" {
		pagMsg.Sort = "created_at asc"
	} else {
		pagMsg.Sort = "created_at desc"
	}
	db := DB
	//判断来源，匹配特殊的选项
	switch pagMsg.Label {
	case "ActivityList":
		db = db.Where("is_show = ?", pagMsg.IsShow)
		if pagMsg.ActivityName != "" {
			db = db.Where("activity_name like", "%"+pagMsg.ActivityName+"%")
		}
		if pagMsg.PersonInCharge != "" {
			db = db.Where("person_in_charge like", "%"+pagMsg.PersonInCharge+"%")
		}
		//判断时间区间
		if pagMsg.StartTime != "" && pagMsg.StopTime != "" {
			db = db.Where("start_time >= ? AND stop_time <= ?", pagMsg.StartTime, pagMsg.StopTime)
		} else if pagMsg.StartTime != "" && pagMsg.StopTime == "" {
			db = db.Where("start_time >= ?", pagMsg.StartTime)
		} else if pagMsg.StartTime == "" && pagMsg.StopTime != "" {
			db = db.Where("stop_time <= ?", pagMsg.StartTime, pagMsg.StopTime)
		}
	case "ClassApplicationList":
		db = db.Preload("JoinAuditDuty").Where("user_class = ? ", pagMsg.UserClass)
	case "ActivityRulerList":
		db = db.Preload("JoinAuditDuty").Where("class_is_pass = ?", "pass")
	case "ActivityOrganizerTrainList":
		db = db.Preload("JoinAuditDuty").Where("class_is_pass = ? AND ruler_is_pass = ? AND organizer_material_is_pass = ?", "pass", "pass", "pass")
	}

	offset := (pagMsg.Page - 1) * pagMsg.Limit
	if offset < 0 {
		offset = 0
	}
	err = db.Limit(pagMsg.Limit).Offset(offset).Order(pagMsg.Sort).Find(&list).Count(&count).Error
	return
}

//申请人信息保存

//查询纪权部所需信息

//查询组织部所需信息

//查询各期负责人
