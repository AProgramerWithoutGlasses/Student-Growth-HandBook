package routesJoinAudit

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"studentGrow/dao/mysql"
	"studentGrow/models/gorm_model"
	"studentGrow/pkg/response"
	token2 "studentGrow/utils/token"
)

// 组织部材料审核
func ActivityOrganizerMaterialManager(c *gin.Context) {
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
			mysql.DB.Model(&gorm_model.JoinAudit{}).Where("id = ?", id).Update("organizer_material_is_pass", "true")
			var updatedJoinAudit gorm_model.JoinAudit
			mysql.DB.Select("organizer_material_is_pass").Where("id = ?", id).First(&updatedJoinAudit)
			resMsg.NowStatus = updatedJoinAudit.OrganizerMaterialIsPass
			resList = append(resList, resMsg)
		}
	}
	if len(cr.Fail) != 0 {
		for _, id := range cr.Fail {
			var resMsg ResList
			resMsg.ID = id
			mysql.DB.Model(&gorm_model.JoinAudit{}).Where("id = ?", id).Update("organizer_material_is_pass", "false")
			var updatedJoinAudit gorm_model.JoinAudit
			mysql.DB.Select("organizer_material_is_pass").Where("id = ?", id).First(&updatedJoinAudit)
			resMsg.NowStatus = updatedJoinAudit.OrganizerMaterialIsPass
			resList = append(resList, resMsg)
		}
	}
	response.ResponseSuccess(c, resList)
}

// 组织部获取成绩列表
func ActivityOrganizerTrainList(c *gin.Context) {
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
		response.ResponseErrorWithMsg(c, response.ParamFail, "json数据解析失败")
		return
	}
	cr.Label = "ActivityOrganizerTrainList"
	list, count, err := mysql.ComList(gorm_model.JoinAudit{}, cr)
	if err != nil {
		response.ResponseErrorWithMsg(c, response.ParamFail, "列表查询失败")
		return
	}

	response.ResponseSuccess(c, ListResponse{
		list,
		count,
	})
}

// 组织部考核成绩审核
func ActivityOrganizerTrainManager(c *gin.Context) {
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
			mysql.DB.Model(&gorm_model.JoinAudit{}).Where("id = ?", id).Update("organizer_train_is_pass", "true")
			var updatedJoinAudit gorm_model.JoinAudit
			mysql.DB.Select("organizer_train_is_pass").Where("id = ?", id).First(&updatedJoinAudit)
			resMsg.NowStatus = updatedJoinAudit.OrganizerTrainIsPass
			resList = append(resList, resMsg)
		}
	}
	if len(cr.Fail) != 0 {
		for _, id := range cr.Fail {
			var resMsg ResList
			resMsg.ID = id
			mysql.DB.Model(&gorm_model.JoinAudit{}).Where("id = ?", id).Update("organizer_train_is_pass", "true")
			var updatedJoinAudit gorm_model.JoinAudit
			mysql.DB.Select("organizer_train_is_pass").Where("id = ?", id).First(&updatedJoinAudit)
			resMsg.NowStatus = updatedJoinAudit.OrganizerTrainIsPass
			resList = append(resList, resMsg)
		}
	}
	response.ResponseSuccess(c, resList)
}

// 组织部更新分数
func SaveTrainScore(c *gin.Context) {
	token := token2.NewToken(c)
	_, exist := token.GetUser()
	if !exist {
		response.ResponseError(c, response.TokenError)
		zap.L().Error("token错误")
		return
	}

	type score struct {
		ID    int
		Score float64 `json:"score"`
	}

	var cr []score
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.ResponseErrorWithMsg(c, response.ParamFail, "json数据解析失败")
		return
	}
	type scoreMsg struct {
		ID       int
		NowScore int
	}
	var resList []scoreMsg
	for _, v := range cr {
		var resMsg scoreMsg
		resMsg.ID = v.ID
		var trainScore gorm_model.JoinAudit
		mysql.DB.Model(&gorm_model.JoinAudit{}).Where("id = ?", v.ID).Update("train_score", v.Score)
		mysql.DB.Select("train_score").Find(&trainScore)
		resMsg.NowScore = trainScore.TrainScore
		resList = append(resList, resMsg)
	}
	response.ResponseSuccess(c, resList)
}
