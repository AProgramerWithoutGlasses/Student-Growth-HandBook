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
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		response.ResponseErrorWithMsg(c, response.ParamFail, "query解析失败")
		return
	}
	list, count, err := mysql.ActivityRulerList(gorm_model.JoinAudit{}, cr)
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
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		response.ResponseErrorWithMsg(c, response.ParamFail, "query数据解析失败")
		return
	}
	var resList []ResList
	if len(cr.Pass) != 0 {
		for _, id := range cr.Pass {
			var resMsg ResList
			resMsg.ID = id
			mysql.DB.Model(&gorm_model.JoinAudit{}).Where("id = ?", id).Update("ruler_is_pass", "pass")
			var updatedJoinAudit gorm_model.JoinAudit
			mysql.DB.Select("ruler_is_pass").Where("id = ?", id).First(&updatedJoinAudit)
			resMsg.NowStatus = updatedJoinAudit.RulerIsPass
			resList = append(resList, resMsg)
		}
	}
	if len(cr.Fail) != 0 {
		for _, id := range cr.Fail {
			var resMsg ResList
			resMsg.ID = id
			mysql.DB.Model(&gorm_model.JoinAudit{}).Where("id = ?", id).Update("ruler_is_pass", "fail")
			var updatedJoinAudit gorm_model.JoinAudit
			mysql.DB.Select("ruler_is_pass").Where("id = ?", id).First(&updatedJoinAudit)
			resMsg.NowStatus = updatedJoinAudit.RulerIsPass
			resList = append(resList, resMsg)
		}
	}
	response.ResponseSuccess(c, resList)
}
