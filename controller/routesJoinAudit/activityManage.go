package routesJoinAudit

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"studentGrow/dao/mysql"
	"studentGrow/models/gorm_model"
	"studentGrow/pkg/response"
	token2 "studentGrow/utils/token"
	"time"
)

type ResDelMsg struct {
	ID        int    `json:"ID"`
	IsSuccess bool   `json:"is_success"`
	Msg       string `json:"msg"`
}
type DelActivityIDList struct {
	ID []int `from:"ID"`
}
type ListResponse struct {
	List  any   `json:"list"`
	Count int64 `json:"count"`
}
type ReceiveActivityMsg struct {
	ID                    uint   `json:"ID"`
	ActivityName          string `json:"activity_name"`
	StartTime             string `json:"start_time"`
	StopTime              string `json:"stop_time"`
	RulerName             string `json:"ruler_name"`
	OrganizerMaterialName string `json:"organizer_material_name"`
	OrganizerTrainName    string `json:"organizer_train_name"`
	IsShow                bool   `json:"is_show"`
	Note                  string `json:"note"`
}

func GetActivityList(c *gin.Context) {
	token := token2.NewToken(c)
	_, exist := token.GetUser()
	if !exist {
		response.ResponseError(c, response.TokenError)
		zap.L().Error("token错误")
		return
	}
	var pagMsg mysql.Pagination
	err := c.ShouldBindJSON(&pagMsg)
	fmt.Println(pagMsg)
	if err != nil {
		response.ResponseErrorWithMsg(c, response.ParamFail, "Query解析失败")
		return
	}
	pagMsg.Label = "ActivityList"
	ActivityMsgList, count, err := mysql.ComList(gorm_model.JoinAuditDuty{}, pagMsg)
	if err != nil {
		response.ResponseErrorWithMsg(c, response.ParamFail, "查询列表出现错误")
		return
	}
	response.ResponseSuccess(c, ListResponse{
		ActivityMsgList,
		count,
	})
}

// SaveActivityMsg 保存和更新活动
func SaveActivityMsg(c *gin.Context) {
	token := token2.NewToken(c)
	_, exist := token.GetUser()
	if !exist {
		response.ResponseError(c, response.TokenError)
		zap.L().Error("token错误")
		return
	}
	var cr ReceiveActivityMsg
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.ResponseErrorWithMsg(c, response.ParamFail, "json解析失败")
		return
	}
	startTime, err := time.Parse("2006-01-02 15:04:05", cr.StartTime)
	if err != nil {
		response.ResponseErrorWithMsg(c, response.ParamFail, "开始时间解析失败")
		return
	}
	stopTime, err := time.Parse("2006-01-02 15:04:05", cr.StopTime)
	if err != nil {
		response.ResponseErrorWithMsg(c, response.ParamFail, "结束时间解析失败")
		return
	}
	if cr.ID == 0 {
		err = mysql.DB.Create(&gorm_model.JoinAuditDuty{
			ActivityName:          cr.ActivityName,
			StartTime:             startTime,
			StopTime:              stopTime,
			RulerName:             cr.RulerName,
			OrganizerMaterialName: cr.OrganizerMaterialName,
			OrganizerTrainName:    cr.OrganizerTrainName,
			IsShow:                false,
			Note:                  cr.Note,
		}).Error
		if err != nil {
			fmt.Println(err.Error())
			response.ResponseErrorWithMsg(c, response.ParamFail, "活动创建失败")
			return
		} else {
			response.ResponseSuccessWithMsg(c, "活动创建成功", struct{}{})
			return
		}
	}
	var count int64
	mysql.DB.Model(&gorm_model.JoinAuditDuty{}).Where("id = ?", cr.ID).Count(&count)
	if count != 1 {
		response.ResponseErrorWithMsg(c, response.ParamFail, "查询活动异常")
		return
	}
	err = mysql.DB.Where("id =? ", cr.ID).Updates(&gorm_model.JoinAuditDuty{
		ActivityName:          cr.ActivityName,
		StartTime:             startTime,
		StopTime:              stopTime,
		RulerName:             cr.RulerName,
		OrganizerMaterialName: cr.OrganizerMaterialName,
		OrganizerTrainName:    cr.OrganizerTrainName,
		IsShow:                false,
		Note:                  cr.Note,
	}).Error
	if err != nil {
		fmt.Println(err.Error())
		response.ResponseErrorWithMsg(c, response.ParamFail, "活动更新失败")
		return
	}
	response.ResponseSuccessWithMsg(c, "活动更新成功", struct{}{})
}

func DelActivityMsg(c *gin.Context) {
	token := token2.NewToken(c)
	_, exist := token.GetUser()
	if !exist {
		response.ResponseError(c, response.TokenError)
		zap.L().Error("token错误")
		return
	}
	var cr DelActivityIDList
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		response.ResponseErrorWithMsg(c, response.ParamFail, "query数据解析失败")
		return
	}
	if len(cr.ID) == 0 {
		response.ResponseErrorWithMsg(c, response.ParamFail, "没有需要处理的数据")
		return
	}
	var resList []ResDelMsg
	var resDelMsg ResDelMsg
	for _, ActivityID := range cr.ID {
		resDelMsg.IsSuccess = false
		resDelMsg.ID = ActivityID
		var count int64
		mysql.DB.Model(&gorm_model.JoinAuditDuty{}).Where("id = ?", ActivityID).Count(&count)
		if count != 1 {
			resDelMsg.Msg = "活动查询异常"
			resList = append(resList, resDelMsg)
			continue
		}
		err = mysql.DB.Delete(&gorm_model.JoinAuditDuty{}, "id =? ", ActivityID).Error
		if err != nil {
			resDelMsg.Msg = "活动删除失败"
			resList = append(resList, resDelMsg)
			continue
		}
		resDelMsg.Msg = "活动删除成功"
		resDelMsg.IsSuccess = true
		resList = append(resList, resDelMsg)
	}
	response.ResponseSuccess(c, resList)
	return
}
