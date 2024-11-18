package routesJoinAudit

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"studentGrow/dao/mysql"
	"studentGrow/models/gorm_model"
	"studentGrow/pkg/response"
	token2 "studentGrow/utils/token"
)

type RecList struct {
	True  []int `form:"true"`
	False []int `form:"false"`
}
type ResList struct {
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
	msgList, count, err := mysql.ComList(gorm_model.JoinAudit{}, cr)
	if err != nil {
		response.ResponseErrorWithMsg(c, response.ParamFail, "列表查询出错")
		return
	}
	fmt.Println(count)
	response.ResponseSuccess(c, ListResponse{
		msgList,
		count,
	})
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
	var resList []ResList
	if len(cr.True) != 0 {
		for _, id := range cr.True {
			var resMsg ResList
			resMsg.ID = id
			updatedJoinAudit := mysql.IsPass(id, "class_is_pass", "true")
			resMsg.NowStatus = updatedJoinAudit.ClassIsPass
			resList = append(resList, resMsg)
		}
	}
	if len(cr.False) != 0 {
		for _, id := range cr.False {
			var resMsg ResList
			resMsg.ID = id
			updatedJoinAudit := mysql.IsPass(id, "class_is_pass", "false")
			resMsg.NowStatus = updatedJoinAudit.ClassIsPass
			resList = append(resList, resMsg)
		}
	}
	response.ResponseSuccess(c, resList)
}
