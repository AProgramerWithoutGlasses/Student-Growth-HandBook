package routesJoinAudit

import (
	"github.com/gin-gonic/gin"
	"studentGrow/dao/mysql"
	res "studentGrow/pkg/response"
	"time"
)

type ResOpenActivityMsg struct {
	IsOpen       bool `json:"is_open"`
	ID           uint
	ActivityName string    `json:"activity_name"`
	StartTime    time.Time `json:"start_time"`
	StopTime     time.Time `json:"stop_time"`
}

// OpenMsg 判断当前入团申请是否开放
func OpenMsg(c *gin.Context) {
	ActivityIsOpen, Msg, ActivityMsg := mysql.OpenActivityStates()
	Response := ResOpenActivityMsg{
		IsOpen:       ActivityIsOpen,
		ID:           ActivityMsg.ID,
		ActivityName: ActivityMsg.ActivityName,
		StartTime:    ActivityMsg.StartTime,
		StopTime:     ActivityMsg.StopTime,
	}
	res.ResponseSuccessWithMsg(c, Msg, Response)
	return
}
