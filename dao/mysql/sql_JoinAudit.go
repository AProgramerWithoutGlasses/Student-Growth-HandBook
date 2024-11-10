package mysql

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"studentGrow/models/gorm_model"
)

// OpenActivityMsg 查询活动的开放信息
func OpenActivityMsg() (isOpen bool, Msg string, ActivityMsg gorm_model.JoinAuditDuty) {
	count := DB.Select("is_show = ?", true).RowsAffected
	if count < 1 {
		isOpen = false
		Msg = "不存在开放活动"
		return
	}
	if count > 1 {
		isOpen = false
		Msg = "开放活动数量异常"
		return
	}
	//currentTime := time.Now()
	count = DB.Where("is_show = ?", true).Find(&ActivityMsg).RowsAffected
	if count != 1 {
		isOpen = false
		Msg = "活动信息查询异常"
		return
	}
	fmt.Println(ActivityMsg.StartTime)
	fmt.Println(ActivityMsg.StopTime)
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

// 申请人信息保存

//查询纪权部所需信息

//查询组织部所需信息

//查询各期负责人
