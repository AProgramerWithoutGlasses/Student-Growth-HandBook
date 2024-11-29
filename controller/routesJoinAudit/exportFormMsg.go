package routesJoinAudit

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"studentGrow/dao/mysql"
	"studentGrow/pkg/response"
	token2 "studentGrow/utils/token"
)

type resList struct {
	ListName string                   `json:"list_name"`
	IsFinish bool                     `json:"is_finish"`
	List     []map[string]interface{} `json:"list"`
}

type rec struct {
	ActivityID int    `json:"activity_id"`
	CurMenu    string `json:"cur_menu"`
}

func ExportFormMsg(c *gin.Context) {
	token := token2.NewToken(c)
	_, exist := token.GetUser()
	if !exist {
		response.ResponseError(c, response.TokenError)
		zap.L().Error("token错误")
		return
	}
	var cr rec
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.ResponseErrorWithMsg(c, response.ParamFail, "json解析失败")
		return
	}
	userMsgList := mysql.UserListWithOrganizer(cr.ActivityID, cr.CurMenu)
	var isFinish = true
	if len(userMsgList) == 0 {
		isFinish = false
	}
	response.ResponseSuccess(c, resList{
		ListName: cr.CurMenu,
		IsFinish: isFinish,
		List:     userMsgList,
	})
}
