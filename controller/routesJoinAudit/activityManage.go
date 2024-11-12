package routesJoinAudit

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"studentGrow/dao/mysql"
	"studentGrow/models/gorm_model"
	"studentGrow/pkg/response"
	token2 "studentGrow/utils/token"
)

type ListResponse struct {
	List  any   `json:"list"`
	Count int64 `json:"count"`
}

func GetActivityMsg(c *gin.Context) {
	token := token2.NewToken(c)
	_, exist := token.GetUser()
	if !exist {
		response.ResponseError(c, response.TokenError)
		zap.L().Error("token错误")
		return
	}
	var pagMsg mysql.Pagination
	err := c.ShouldBindQuery(&pagMsg)
	if err != nil {
		response.ResponseErrorWithMsg(c, response.ParamFail, "Query解析失败")
	}
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
