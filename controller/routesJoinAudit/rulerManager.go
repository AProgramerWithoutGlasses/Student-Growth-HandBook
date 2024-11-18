package routesJoinAudit

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"studentGrow/dao/mysql"
	"studentGrow/models/gorm_model"
	"studentGrow/pkg/response"
	token2 "studentGrow/utils/token"
)

func ActivityRulerList(c *gin.Context) {
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
		response.ResponseErrorWithMsg(c, response.ParamFail, "query解析失败")
		return
	}
	cr.Label = "ActivityRulerList"
	list, count, err := mysql.ComList(gorm_model.JoinAudit{}, cr)
	if err != nil {
		response.ResponseErrorWithMsg(c, response.ParamFail, "列表查询出错")
		return
	}
	response.ResponseSuccess(c, ListResponse{
		list,
		count,
	})
}
func ActivityRulerManager(c *gin.Context) {
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
			updatedJoinAudit := mysql.IsPass(id, "ruler_is_pass", "true")
			resMsg.NowStatus = updatedJoinAudit.RulerIsPass
			resList = append(resList, resMsg)
		}
	}
	if len(cr.False) != 0 {
		for _, id := range cr.False {
			var resMsg ResList
			resMsg.ID = id
			updatedJoinAudit := mysql.IsPass(id, "ruler_is_pass", "false")
			resMsg.NowStatus = updatedJoinAudit.RulerIsPass
			resList = append(resList, resMsg)
		}
	}
	response.ResponseSuccess(c, resList)
}
