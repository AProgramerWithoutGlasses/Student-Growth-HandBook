package gorm_model

import (
	"gorm.io/gorm"
	"time"
)

// JoinAuditDuty 各期责任人表
type JoinAuditDuty struct {
	gorm.Model
	ActivityName              string          `json:"activity_name"`                                //期数
	StartTime                 time.Time       `json:"start_time"`                                   //开始时间
	StopTime                  time.Time       `json:"stop_time"`                                    //结束时间
	RulerName                 string          `json:"ruler_name"`                                   //纪权部综测审核人
	RulerUsername             string          `json:"ruler_username"`                               //纪权部综测审核人账号
	OrganizerMaterialName     string          `json:"organizer_material_name"`                      //组织部材料审核人
	OrganizerMaterialUsername string          `json:"organizer_material_username"`                  //组织部材料审核人账号
	OrganizerTrainName        string          `json:"organizer_train_name"`                         //组织部培训审核人
	OrganizerTrainUsername    string          `json:"organizer_train_username"`                     //组织部培训审核人账号
	IsShow                    bool            `json:"is_show" gorm:"default:false"`                 //是否展示
	JoinAudit                 []JoinAudit     `json:"join_audit" gorm:"foreignKey:JoinAuditDutyID"` //学生提交的表单信息
	JoinAuditFIle             []JoinAuditFile `json:"images" gorm:"foreignKey:JoinAuditDutyID"`     //每一期对应的照片
	Note                      string          `json:"note"`                                         //备注
}
