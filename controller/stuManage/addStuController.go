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
	token2 "studentGrow/utils/token"
	"time"
)

// 添加单个学生
func AddSingleStuContro(c *gin.Context) {
	// 根据token拿到登陆者自己的信息
	token := c.GetHeader("token")
	username, err := token2.GetUsername(token)
	if err != nil {
		response.ResponseError(c, response.ParamFail)
		zap.L().Error(err.Error())
		return
	}
	fmt.Println(username)

	id, err := mysql.GetIdByUsername(username)
	if err != nil {
		response.ResponseError(c, response.ParamFail)
		zap.L().Error(err.Error())
		return
	}

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

	// 接收
	input := struct {
		Name     string `json:"name"`
		Username string `json:"username"`
		Password string `json:"password"`
		Class    string `json:"class"`
		Gender   string `json:"gender"`
	}{}

	err = c.ShouldBindJSON(&input)
	if err != nil {
		response.ResponseError(c, response.ParamFail)
		zap.L().Error(err.Error())
	}

	// 去除班级名称中的 ”班“ 字
	if len(input.Class) == 12 {
		input.Class = input.Class[:len(input.Class)-3]
	}

	// 使用正则表达式进行匹配
	pattern := `^[\p{Han}]{2}\d{3}$`
	match, _ := regexp.MatchString(pattern, input.Class)
	if !match {
		response.ResponseErrorWithMsg(c, response.ServerErrorCode, "请输入正确的班级格式")
		println("1111")
		return
	}

	// 计算出要添加的学生是大几的
	addStuNowGrade := service.CalculateNowGradeByClass(input.Class)

	// 导入班级权限判断
	if input.Class != class {
		if role == addStuNowGrade || role == "college" || role == "superman" {

		} else {
			response.ResponseErrorWithMsg(c, response.ServerErrorCode, "导入失败，您只能导入您所管班级的学生或所管年级的学生!")
			zap.L().Error("导入失败，您只能导入您所管班级的学生或所管年级的学生!")
			return
		}
	}

	// 根据班级获取入学时间
	re := regexp.MustCompile(`^\D*(\d{2})`)
	match1 := re.FindStringSubmatch(input.Class)

	yearEnd := match1[1]                   // 获取 "22"
	yearEndInt, _ := strconv.Atoi(yearEnd) // 将 "22" 转换为整数
	yearInt := yearEndInt + 2000           // 将整数转换为 "2022"

	now := time.Now()
	addStuPlusTime := time.Date(yearInt, 9, 1, 0, 0, 0, 0, now.Location())

	fmt.Println("plusTime:", addStuPlusTime)

	// 将新增学生信息整合到结构体中
	user := gorm_model.User{
		Name:     input.Name,
		Username: input.Username,
		Password: input.Password,
		Class:    input.Class,
		Gender:   input.Gender,
		Identity: "学生",
		PlusTime: addStuPlusTime,
		HeadShot: "https://student-grow.oss-cn-beijing.aliyuncs.com/image/user_headshot/user_headshot_5.png",
	}

	// 在数据库中添加该学生信息
	err = mysql.AddSingleStudent(&user)
	if err != nil {
		response.ResponseErrorWithMsg(c, 500, "添加失败, 该用户已存在")
		zap.L().Error("stuManage.AddMultipleStuControl() mysql.AddSingleStudent() failed: " + err.Error())
		return
	}

	// 添加学生记录
	addUserRecord := gorm_model.UserAddRecord{
		Username:    username,
		AddUsername: input.Username,
	}
	err = mysql.AddSingleStudentRecord(&addUserRecord)
	if err != nil {
		response.ResponseError(c, response.ServerErrorCode)
		zap.L().Error("stuManage.AddMultipleStuControl() mysql.AddSingleStudentRecord() failed: " + err.Error())
		return
	}

	// 成功响应
	response.ResponseSuccessWithMsg(c, input.Username+" 信息添加成功！", nil)
}
