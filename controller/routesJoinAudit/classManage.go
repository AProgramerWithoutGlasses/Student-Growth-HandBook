package routesJoinAudit

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"studentGrow/dao/mysql"
	"studentGrow/pkg/response"
	"studentGrow/service/JoinAudit"
	token2 "studentGrow/utils/token"
)

type RecList struct {
	True  []int `form:"true"`
	False []int `form:"false"`
}
type ResListWithIsPass struct {
	ID        int
	NowStatus string
}

// 班长获取对应的申请列表
func ClassApplicationList(c *gin.Context) {
	token := token2.NewToken(c)
	_, exist := token.GetUser()
	if !exist {
		response.ResponseError(c, response.TokenError)
		zap.L().Error("token错误")
		return
	}
	var cr mysql.Pagination
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.ResponseErrorWithMsg(c, response.ParamFail, "json解析出错")
	}
	cr.Label = "ClassApplicationList"
	//cr.UserClass = user.Class
	cr.UserClass = "计科211"
	//判断返回全部还是只有开启的活动
	var ResAllMsgList = make([]JoinAudit.ResList, 0)
	ResAllMsgList, err = JoinAudit.ResListWithJSON(cr)
	if err != nil {
		response.ResponseErrorWithMsg(c, response.ParamFail, err.Error())
	}
	response.ResponseSuccess(c, ResAllMsgList)
}

func ClassApplicationManager(c *gin.Context) {
	token := token2.NewToken(c)
	_, exist := token.GetUser()
	if !exist {
		response.ResponseError(c, response.TokenError)
		zap.L().Error("token错误")
		return
	}
	var cr RecList
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.ResponseErrorWithMsg(c, response.ParamFail, "json数据解析失败")
		return
	}
	var resList []ResListWithIsPass
	if len(cr.True) != 0 {
		for _, id := range cr.True {
			var resMsg ResListWithIsPass
			resMsg.ID = id
			updatedJoinAudit := mysql.IsPass(id, "class_is_pass", "true")
			resMsg.NowStatus = updatedJoinAudit.ClassIsPass
			resList = append(resList, resMsg)
		}
	}
	if len(cr.False) != 0 {
		for _, id := range cr.False {
			var resMsg ResListWithIsPass
			resMsg.ID = id
			updatedJoinAudit := mysql.IsPass(id, "class_is_pass", "false")
			resMsg.NowStatus = updatedJoinAudit.ClassIsPass
			resList = append(resList, resMsg)
		}
	}
	response.ResponseSuccess(c, resList)
}
