package routesJoinAudit

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"studentGrow/dao/mysql"
	"studentGrow/models/gorm_model"
	"studentGrow/pkg/response"
	token2 "studentGrow/utils/token"
)

// 组织部获取列表
func ActivityOrganizerList(c *gin.Context) {
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
	type file struct {
		ID       uint
		Username string `json:"username"`
		FilePath string `json:"file_path"`
		FileNote string `json:"file_note"`
	}
	var resList []struct {
		List gorm_model.JoinAudit `json:"msg"`
		File []file               `json:"file_list"`
	}

	for _, v := range list {
		fileList, _ := mysql.FilesList(v.Username, v.JoinAuditDutyID)
		var files file
		var filesList []file
		filesList = make([]file, 0)
		for _, val := range fileList {
			files.ID = val.ID
			files.Username = v.Username
			files.FilePath = val.FilePath
			files.FileNote = val.Note
			filesList = append(filesList, files)
		}
		resList = append(resList, struct {
			List gorm_model.JoinAudit `json:"msg"`
			File []file               `json:"file_list"`
		}{List: v, File: filesList})
	}

	response.ResponseSuccess(c, ListResponse{
		resList,
		count,
	})
}

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
			updatedJoinAudit := mysql.IsPass(id, "organizer_material_is_pass", "true")
			resMsg.NowStatus = updatedJoinAudit.OrganizerMaterialIsPass
			resList = append(resList, resMsg)
		}
	}
	if len(cr.Fail) != 0 {
		for _, id := range cr.Fail {
			var resMsg ResList
			resMsg.ID = id
			updatedJoinAudit := mysql.IsPass(id, "organizer_material_is_pass", "false")
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
			updatedJoinAudit := mysql.IsPass(id, "organizer_train_is_pass", "true")

			resMsg.NowStatus = updatedJoinAudit.OrganizerTrainIsPass
			resList = append(resList, resMsg)
		}
	}
	if len(cr.Fail) != 0 {
		for _, id := range cr.Fail {
			var resMsg ResList
			resMsg.ID = id
			updatedJoinAudit := mysql.IsPass(id, "organizer_train_is_pass", "false")
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
