package stuManage

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"studentGrow/dao/mysql"
	"studentGrow/models/gorm_model"
	"studentGrow/pkg/response"
)

func AddMultipleStuControl(c *gin.Context) {
	// 获取上传的Excel文件
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		response.ResponseErrorWithMsg(c, 500, "stuManage.AddMultipleStuControl() c.Request.FormFile() failed: "+err.Error())
		zap.L().Error("stuManage.AddMultipleStuControl() c.Request.FormFile() failed: " + err.Error())
		return
	}
	defer file.Close()

	// 解析Excel文件并获取数据
	f, err := excelize.OpenReader(file)
	if err != nil {
		response.ResponseErrorWithMsg(c, 500, "stuManage.AddMultipleStuControl() excelize.OpenReader() failed: "+err.Error())
		zap.L().Error("stuManage.AddMultipleStuControl() excelize.OpenReader() failed: " + err.Error())
		return
	}

	rows := f.GetRows("Sheet1")

	duplicatedUser := make([]string, 0)

	// 检查数据库中是否已经存在该用户
	for _, row := range rows[1:] { // 忽略表头行
		err = mysql.ExistedUsername(row[2])
		if err != nil {
			if err == gorm.ErrRecordNotFound {

			} else {
				response.ResponseErrorWithMsg(c, 500, "stuManage.AddMultipleStuControl() mysql.ExistedUsername() failed: "+err.Error())
				zap.L().Error("stuManage.AddMultipleStuControl() mysql.ExistedUsername() failed: " + err.Error())
				return
			}

		} else { // 用户存在
			duplicatedUser = append(duplicatedUser, row[2])
		}
	}

	var duplicatedUserStr string
	if len(duplicatedUser) > 0 {
		for _, v := range duplicatedUser {
			duplicatedUserStr = duplicatedUserStr + v + ", "
		}
		duplicatedUserStr = duplicatedUserStr[:len(duplicatedUserStr)-2]
		response.ResponseErrorWithMsg(c, 500, "导入失败，请不要导入已存在的学生学号: "+duplicatedUserStr)
		zap.L().Error("导入失败，请不要导入已存在的学生学号: " + duplicatedUserStr)
		return
	}

	// 创建新的用户
	for _, row := range rows[1:] {
		user := gorm_model.User{
			Class:    row[0],
			Name:     row[1],
			Username: row[2],
			Password: row[3],
			Gender:   row[4],
			Identity: "学生",
		}
		err = mysql.AddSingleStudent(&user)
		if err != nil {
			response.ResponseErrorWithMsg(c, 500, "stuManage.AddMultipleStuControl() mysql.AddSingleStudent failed: "+err.Error())
			zap.L().Error("stuManage.AddMultipleStuControl() mysql.AddSingleStudent failed: " + err.Error())
			return
		}
	}

	response.ResponseSuccess(c, "导入成功!")

}
