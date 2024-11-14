package routesJoinAudit

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"studentGrow/dao/mysql"
	"studentGrow/models/gorm_model"
	"studentGrow/pkg/response"
	"studentGrow/service/JoinAudit"
	token2 "studentGrow/utils/token"
)

type StuFileMsg struct {
	ID       uint
	FileName string `json:"file_name"`
	FilePath string `json:"file_path"`
	Note     string `json:"note"`
}
type StuFileDelMsg struct {
	ID        int
	IsSuccess bool `json:"is_success"`
}

func GetStuFile(c *gin.Context) {
	token := token2.NewToken(c)
	user, exist := token.GetUser()
	if !exist {
		response.ResponseError(c, response.TokenError)
		zap.L().Error("token错误")
		return
	}
	var resMsg StuFileMsg
	var resList []StuFileMsg
	ActivityIsOpen, Msg, ActivityMsg := mysql.OpenActivityMsg()
	if !ActivityIsOpen {
		response.ResponseErrorWithMsg(c, response.ParamFail, Msg)
		return
	}
	var stuFromMsg gorm_model.JoinAudit
	mysql.DB.Where("username = ? AND join_audit_duty_id = ?", user.Username, ActivityMsg.ID).Take(&stuFromMsg)
	if stuFromMsg.ClassIsPass != "pass" {
		response.ResponseErrorWithMsg(c, response.ParamFail, "班级审核未通过")
		return
	}
	if stuFromMsg.RulerIsPass == "fail" {
		response.ResponseErrorWithMsg(c, response.ParamFail, "综测成绩审核未通过")
		return
	}
	if stuFromMsg.OrganizerMaterialIsPass == "pass" {
		response.ResponseErrorWithMsg(c, response.ParamFail, "材料审核已通过")
		return
	}
	var fileList []gorm_model.JoinAuditFile
	mysql.DB.Where("username = ? AND join_audit_duty_id = ?", user.Username, ActivityMsg.ID).Find(&fileList)
	if len(fileList) == 0 {
		response.ResponseSuccessWithMsg(c, "用户文件不存在", resList)
		return
	}
	for _, file := range fileList {
		resMsg.ID = file.ID
		resMsg.FileName = file.FileName
		resMsg.FilePath = file.FilePath
		resMsg.Note = file.Note
		resList = append(resList, resMsg)
	}
	response.ResponseSuccess(c, resList)
	return
}

// SaveStuFile 保存提交文件
func SaveStuFile(c *gin.Context) {
	token := token2.NewToken(c)
	user, exist := token.GetUser()
	if !exist {
		response.ResponseError(c, response.TokenError)
		zap.L().Error("token错误")
		return
	}
	ActivityIsOpen, Msg, ActivityMsg := mysql.OpenActivityMsg()
	if !ActivityIsOpen {
		response.ResponseErrorWithMsg(c, response.ParamFail, Msg)
		return
	}
	form, err := c.MultipartForm()
	if err != nil {
		response.ResponseErrorWithMsg(c, response.ServerErrorCode, err.Error())
	}
	var resList []JoinAudit.StuFileUpload
	if form.File["material"] == nil || form.File["application"] == nil {
		response.ResponseErrorWithMsg(c, response.ParamFail, "无文件需要处理")
		return
	}
	//删除之前存在的文件
	var imageList []gorm_model.JoinAuditFile
	count := mysql.DB.Find(&imageList, "username = ? AND join_audit_duty_id = ?", user.Username, ActivityMsg.ID).RowsAffected
	if count != 0 {
		//for _, file := range imageList {
		//	err, getHeader := fileProcess.DelOssFile(file.FilePath)
		//	if err != nil {
		//		response.ResponseErrorWithMsg(c, response.ParamFail, "阿里云文件删除失败")
		//		return
		//	}
		//	fmt.Println(getHeader)
		//}
		err = mysql.DB.Delete(&imageList).Error
		if err != nil {
			response.ResponseErrorWithMsg(c, response.ParamFail, "material 旧文件删除失败")
			return
		}
	}
	//获取传入的文件
	fileList, ok := form.File["material"]
	if ok {
		for _, file := range fileList {
			resMsg := JoinAudit.FileUpload(file, user, ActivityMsg, "material")
			resList = append(resList, resMsg)
		}
	}
	fileList, ok = form.File["application"]
	if ok {
		for _, file := range fileList {
			resMsg := JoinAudit.FileUpload(file, user, ActivityMsg, "application")
			resList = append(resList, resMsg)
		}
	}
	if len(resList) == 0 {
		response.ResponseErrorWithMsg(c, response.ParamFail, "文件获取失败")
		return
	}
	mysql.DB.Model(&gorm_model.JoinAudit{}).Where("username = ? AND join_audit_duty_id = ?", user.Username, ActivityMsg.ID).Update("organizer_material_is_pass", "")
	response.ResponseSuccess(c, resList)
}

// DelStuFile 文件删除
func DelStuFile(c *gin.Context) {
	type DelFileList struct {
		ID []int
	}
	token := token2.NewToken(c)
	user, exist := token.GetUser()
	if !exist {
		response.ResponseError(c, response.TokenError)
		zap.L().Error("token错误")
		return
	}
	var cr DelFileList
	if err := c.ShouldBindQuery(&cr); err != nil {
		response.ResponseErrorWithMsg(c, response.ParamFail, "客户端数据解析失败")
		return
	}
	//判断删除时间是否合法
	ActivityIsOpen, Msg, ActivityMsg := mysql.OpenActivityMsg()
	if !ActivityIsOpen {
		response.ResponseErrorWithMsg(c, response.ParamFail, Msg)
		return
	}
	var resFileDelMsgList []StuFileDelMsg
	for _, fileID := range cr.ID {
		var FileDelMsg StuFileDelMsg
		FileDelMsg.ID = fileID
		count := mysql.DB.Delete(&gorm_model.JoinAuditFile{}, "id = ? AND username = ? AND join_audit_duty_id = ?", fileID, user.Username, ActivityMsg.ID).RowsAffected
		if count != 1 {
			FileDelMsg.IsSuccess = false
			resFileDelMsgList = append(resFileDelMsgList, FileDelMsg)
			continue
		}
		FileDelMsg.IsSuccess = true
		resFileDelMsgList = append(resFileDelMsgList, FileDelMsg)
	}
	response.ResponseSuccess(c, resFileDelMsgList)
}
