package stuManage

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"regexp"
	"strconv"
	"studentGrow/dao/mysql"
	"studentGrow/models/gorm_model"
	"studentGrow/pkg/response"
	"studentGrow/service"
	"studentGrow/utils/readMessage"
	token2 "studentGrow/utils/token"
	"time"
)

// AddSingleStuContro 添加单个学生
func AddSingleStuContro(c *gin.Context) {
	token := c.GetHeader("token")

	username, err := token2.GetUsername(token)
	if err != nil {
		response.ResponseError(c, response.ParamFail)
		zap.L().Error(err.Error())
		return
	}

	id, err := mysql.GetIdByUsername(username)
	if err != nil {
		response.ResponseError(c, response.ParamFail)
		zap.L().Error(err.Error())
		return
	}

	// 拿到登陆者自己的class
	class, err := mysql.GetClassById(id)
	if err != nil {
		response.ResponseError(c, response.ParamFail)
		zap.L().Error(err.Error())
		return
	}

	role, err := token2.GetRole(token)
	if err != nil {
		response.ResponseError(c, response.ParamFail)
		zap.L().Error(err.Error())
		return
	}

	// 接收请求数据
	stuMessage, err := readMessage.GetJsonvalue(c)
	if err != nil {
		fmt.Println("stuManage.AddSingleStuContro() readMessage.GetJsonvalue() err :", err)
	}

	// 获取请求信息中各个字段的值
	nameValue, err := stuMessage.GetString("name")
	if err != nil {
		fmt.Println("name GetString() err : ", err)
	}

	usernameValue, err := stuMessage.GetString("username")
	if err != nil {
		fmt.Println("username GetString() err : ", err)
	}

	passwordValue, err := stuMessage.GetString("password")
	if err != nil {
		fmt.Println("password GetString() err : ", err)
	}

	classValue, err := stuMessage.GetString("class")
	if err != nil {
		fmt.Println("class GetString() err : ", err)
	}

	// 拿到添加学生的grade
	addStuId, err := mysql.GetIdByUsername(usernameValue)
	if err != nil {
		response.ResponseError(c, response.ServerErrorCode)
		zap.L().Error(err.Error())
		return
	}

	addStuPlusTime, err := mysql.GetPlusTimeById(addStuId)
	if err != nil {
		response.ResponseError(c, response.ServerErrorCode)
		zap.L().Error(err.Error())
		return
	}

	// 计算出要添加的学生是大几的
	addStuNowGrade := service.CalculateNowGradeByClass(classValue)

	// 去除班级名称中的 ”班“ 字
	if len(classValue) == 12 {
		classValue = classValue[:len(classValue)-3]
	}

	// 导入班级权限判断
	if classValue != class {
		if role == addStuNowGrade || role == "college" {

		} else {
			response.ResponseErrorWithMsg(c, response.ServerErrorCode, "导入失败，您只能导入您所管班级的学生或所管年级的学生!")
			zap.L().Error("导入失败，您只能导入您所管班级的学生或所管年级的学生!")
			return
		}
	}

	genderValue, err := stuMessage.GetString("gender")
	if err != nil {
		fmt.Println("gender GetString() err : ", err)
	}

	// 根据班级获取入学时间
	re := regexp.MustCompile(`^\D*(\d{2})`)
	match := re.FindStringSubmatch(classValue)

	yearEnd := match[1]                    // 获取 "22"
	yearEndInt, _ := strconv.Atoi(yearEnd) // 将 "22" 转换为整数
	yearInt := yearEndInt + 2000           // 将整数转换为 "2022"

	now := time.Now()
	addStuPlusTime = time.Date(yearInt, 9, 1, 0, 0, 0, 0, now.Location())

	fmt.Println("plusTime:", addStuPlusTime)

	// 将新增学生信息整合到结构体中
	user := gorm_model.User{
		Name:     nameValue,
		Username: usernameValue,
		Password: passwordValue,
		Class:    classValue,
		Gender:   genderValue,
		Identity: "学生",
		PlusTime: addStuPlusTime,
	}

	// 在数据库中添加该学生信息
	err = mysql.AddSingleStudent(&user)
	if err != nil {
		response.ResponseErrorWithMsg(c, 500, "添加失败, 该用户已存在")
		zap.L().Error("stuManage.AddMultipleStuControl() mysql.AddSingleStudent() failed: " + err.Error())
		return
	}

	// 成功响应
	response.ResponseSuccess(c, nameValue+" 信息添加成功！")
}
