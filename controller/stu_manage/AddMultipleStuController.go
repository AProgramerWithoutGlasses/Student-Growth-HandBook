package stu_manage

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"studentGrow/dao/mysql"
	"studentGrow/models/gorm_model"
	"studentGrow/pkg/response"
)

// 批量添加学生
func AddMultipleStuControl(c *gin.Context) {
	// 获取上传的Excel文件
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		response.ResponseErrorWithMsg(c, 500, "stu_manage.AddMultipleStuControl() c.Request.FormFile() failed : "+err.Error())
		zap.L().Error("stu_manage.AddMultipleStuControl() c.Request.FormFile() failed : " + err.Error())
		return
	}
	defer file.Close()

	// 解析Excel文件并获取数据
	f, err := excelize.OpenReader(file)
	if err != nil {
		response.ResponseErrorWithMsg(c, 500, "stu_manage.AddMultipleStuControl() excelize.OpenReader() failed : "+err.Error())
		zap.L().Error("stu_manage.AddMultipleStuControl() excelize.OpenReader() failed : " + err.Error())
		return
	}

	rows := f.GetRows("Sheet1")

	// 批量导入数据
	for _, row := range rows[1:] { // 忽略表头行
		user := gorm_model.User{
			Class:    row[0],
			Name:     row[1],
			Username: row[2],
			Password: row[3],
			Identity: "学生",
		}
		err = mysql.AddSingleStudent(&user)
		if err != nil {
			response.ResponseErrorWithMsg(c, 500, "stu_manage.AddMultipleStuControl() excelize.OpenReader() failed : "+err.Error())
			zap.L().Error("stu_manage.AddMultipleStuControl() excelize.OpenReader() failed : " + err.Error())
			return
		}
	}

	response.ResponseSuccess(c, 200)

}
