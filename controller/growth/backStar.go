package growth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"studentGrow/dao/mysql"
	"studentGrow/pkg/response"
	"studentGrow/service/starService"
	token2 "studentGrow/utils/token"
)

// Student 定义接收前端数据结构体
type Student struct {
	Username         string `json:"username"`
	Name             string `json:"name"`
	Userarticletotal int    `json:"user_article_total"`
	Userfans         int    `json:"userfans"`
	Score            int    `json:"score"`
	Hot              int    `json:"hot"`
	Frequency        int    `json:"frequency"`
}

// Search 搜索表格数据
func Search(c *gin.Context) {
	//返回前端限制人数
	var peopleLimit int
	var usernamesli []string
	//获取前端传来的数据
	var datas struct {
		Name  string `form:"search"`
		Page  int    `form:"page"`
		Limit int    `form:"pageCapacity"`
	}
	err := c.Bind(&datas)
	if err != nil {
		zap.L().Error("Search Bind err", zap.Error(err))
		response.ResponseError(c, response.ServerErrorCode)
		return
	}

	//得到登录者的角色和账号
	token := c.GetHeader("token")
	role, err := token2.GetRole(token)
	username, err := token2.GetUsername(token)
	if err != nil {
		zap.L().Error("Search err", zap.Error(err))
		response.ResponseError(c, response.ServerErrorCode)
		return
	}

	//查找stars库中所有的数据
	//查找权限下的数据
	alluser, err := mysql.SelStarUser()
	if err != nil {
		fmt.Println("StarGrade err", err)
		response.ResponseErrorWithMsg(c, 400, "获取表格数据失败")
	}

	//根据角色分类
	switch role {
	case "class":
		//查找班级
		class, err := mysql.SelClass(username)
		if err != nil {
			zap.L().Error("Search SelClass err", zap.Error(err))
			response.ResponseError(c, response.ServerErrorCode)
			return
		}
		if datas.Name == "" {
			usernamesli, err = mysql.SelUsername(class)
		} else {
			usernamesli, err = mysql.SelSearchUser(datas.Name, class)
		}
		peopleLimit = 3
		if err != nil {
			zap.L().Error("Search SelSearchUser err", zap.Error(err))
			response.ResponseError(c, response.ServerErrorCode)
			return
		}
	case "grade1":
		if datas.Name == "" {
			usernamesli, err = starService.StarGuidGrade(alluser, 1)
		} else {
			usernamesli, err = starService.SearchGrade(datas.Name, 1)
		}
		peopleLimit = 5
		if err != nil {
			zap.L().Error("Search GetEnrollmentYear err", zap.Error(err))
			response.ResponseError(c, response.ServerErrorCode)
			return
		}
	case "grade2":
		if datas.Name == "" {
			usernamesli, err = starService.StarGuidGrade(alluser, 2)
		} else {
			usernamesli, err = starService.SearchGrade(datas.Name, 2)
		}
		peopleLimit = 5
		if err != nil {
			zap.L().Error("Search GetEnrollmentYear err", zap.Error(err))
			response.ResponseError(c, response.ServerErrorCode)
			return
		}
	case "grade3":
		if datas.Name == "" {
			usernamesli, err = starService.StarGuidGrade(alluser, 3)
		} else {
			usernamesli, err = starService.SearchGrade(datas.Name, 3)
		}
		peopleLimit = 5
		if err != nil {
			zap.L().Error("Search GetEnrollmentYear err", zap.Error(err))
			response.ResponseError(c, response.ServerErrorCode)
			return
		}
	case "grade4":
		if datas.Name == "" {
			usernamesli, err = starService.StarGuidGrade(alluser, 4)
		} else {
			usernamesli, err = starService.SearchGrade(datas.Name, 4)
		}
		peopleLimit = 5
		if err != nil {
			zap.L().Error("Search GetEnrollmentYear err", zap.Error(err))
			response.ResponseError(c, response.ServerErrorCode)
			return
		}
	case "college":
		if datas.Name == "" {
			usernamesli, err = mysql.SelStarColl()
		} else {
			usernamesli, err = mysql.SelSearchColl(datas.Name)
		}
		peopleLimit = 10
		if err != nil {
			zap.L().Error("Search SelSearchColl err", zap.Error(err))
			response.ResponseError(c, response.ServerErrorCode)
			return
		}
	case "superman":
		if datas.Name == "" {
			usernamesli, err = mysql.SelStarColl()
		} else {
			usernamesli, err = mysql.SelSearchColl(datas.Name)
		}
		peopleLimit = 0
		if err != nil {
			zap.L().Error("Search SelSearchColl err", zap.Error(err))
			response.ResponseError(c, response.ServerErrorCode)
			return
		}
	}

	//表格所需所有数据
	starback, err := starService.StarGrid(usernamesli)
	if err != nil {
		fmt.Println("StarGrade starback err", err)
		return
	}
	total := len(starback)
	//实现分页
	tableData := starService.PageQuery(starback, datas.Page, datas.Limit)
	//查询状态
	status, err := mysql.SelStatus(username)
	if err != nil {
		response.ResponseError(c, 400)
		return
	}
	data := map[string]any{
		"tableData":   tableData,
		"total":       total,
		"peopleLimit": peopleLimit,
		"isDisabled":  status,
	}
	response.ResponseSuccess(c, data)
}

// ElectClass  班级管理员推选数据
func ElectClass(c *gin.Context) {
	var Responsedata struct {
		ElectedArr []Student `json:"electedArr"`
	}
	err := c.Bind(&Responsedata)
	if err != nil {
		zap.L().Error("Search Bind err", zap.Error(err))
		response.ResponseErrorWithMsg(c, response.ServerErrorCode, "为获取到数据")
		return
	}
	for _, student := range Responsedata.ElectedArr {
		username := student.Username
		name := student.Name
		//防止有重复数据
		number, err := mysql.Selstarexit(username)
		if err != nil || number != 0 {
			response.ResponseErrorWithMsg(c, 400, "数据已存在")
			return
		}
		//添加数据
		err = mysql.CreatClass(username, name)
		if err != nil {
			response.ResponseErrorWithMsg(c, 400, "推选失败")
			return
		}
	}
}

// ElectGrade 年级管理员推选数据
func ElectGrade(c *gin.Context) {
	var Responsedata struct {
		ElectedArr []Student `json:"electedArr"`
	}
	err := c.Bind(&Responsedata)
	if err != nil {
		zap.L().Error("Search Bind err", zap.Error(err))
		response.ResponseErrorWithMsg(c, response.ServerErrorCode, "未获取到数据")
		return
	}
	for _, user := range Responsedata.ElectedArr {
		username := user.Username
		//更新数据
		err := mysql.UpdateGrade(username)
		if err != nil {
			response.ResponseErrorWithMsg(c, 400, "推选失败")
			return
		}
	}
}

// ElectCollege 院级管理员推选
func ElectCollege(c *gin.Context) {
	var Responsedata struct {
		ElectedArr []Student `json:"electedArr"`
	}
	err := c.Bind(&Responsedata)
	if err != nil {
		zap.L().Error("Search Bind err", zap.Error(err))
		response.ResponseErrorWithMsg(c, response.ServerErrorCode, "未获取到数据")
		return
	}
	for _, user := range Responsedata.ElectedArr {
		username := user.Username
		err := mysql.UpdateCollege(username)
		if err != nil {
			response.ResponseErrorWithMsg(c, 400, "推选失败")
			return
		}
	}
}

// PublicStar 公布成长之星
func PublicStar(c *gin.Context) {
	//获取session字段的最大值
	nowSession, err := mysql.SelMax()
	if err != nil {
		response.ResponseErrorWithMsg(c, 400, "获取值失败")
		return
	}
	session := nowSession + 1
	//更新字段
	err = mysql.UpdateSession(session)
	if err != nil {
		response.ResponseErrorWithMsg(c, 400, "公布失败")
		return
	}

	//更新所有管理员状态字段
	err = mysql.UpdateStatus()
	if err != nil {
		response.ResponseError(c, 400)
		return
	}

	//展示最新一期
	session, err = mysql.SelMax()
	if err != nil {
		response.ResponseErrorWithMsg(c, 400, "获取最新数据失败")
		return
	}

	//返回数据
	//班级成长之星
	classData, err := starService.StarClass(session)
	if err != nil {
		response.ResponseErrorWithMsg(c, 400, "班级之星查找失败")
		return
	}

	//年级之星
	gradeData, err := starService.StarGrade(session)
	if err != nil {
		response.ResponseErrorWithMsg(c, 400, "年级之星查找失败")
		return
	}

	//院级之星
	hospitalData, err := starService.StarCollege(session)
	if err != nil {
		response.ResponseErrorWithMsg(c, 400, "院级之星查找失败")
		return
	}

	data := map[string]any{
		"classData":    classData,
		"gradeData":    gradeData,
		"hospitalData": hospitalData,
	}
	response.ResponseSuccess(c, data)
}

// StarPub 搜索第几届成长之星
func StarPub(c *gin.Context) {
	//定义届数
	var session int
	//接收前端数据
	var term struct {
		TermNumber int `form:"termNumber"`
	}
	err := c.Bind(&term)
	if err != nil {
		zap.L().Error("Search Bind err", zap.Error(err))
		response.ResponseError(c, response.ServerErrorCode)
		return
	}
	//设置session的值
	if term.TermNumber == 0 {
		//找到最大的session最新一期进行展示
		session, err = mysql.SelMax()
		if session == 0 {
			response.ResponseSuccess(c, "")
		}
		if err != nil {
			response.ResponseErrorWithMsg(c, 400, "获取最新数据失败")
			return
		}
	} else {
		session = term.TermNumber
	}
	//班级成长之星
	classData, err := starService.StarClass(session)
	if err != nil {
		response.ResponseErrorWithMsg(c, 400, "班级之星查找失败")
		return
	}

	//年级之星
	gradeData, err := starService.StarGrade(session)
	if err != nil {
		response.ResponseErrorWithMsg(c, 400, "年级之星查找失败")
		return
	}

	//院级之星
	hospitalData, err := starService.StarCollege(session)
	if err != nil {
		response.ResponseErrorWithMsg(c, 400, "院级之星查找失败")
		return
	}

	data := map[string]any{
		"classData":    classData,
		"gradeData":    gradeData,
		"hospitalData": hospitalData,
	}
	response.ResponseSuccess(c, data)
}

// BackStarClass 返回前台班级成长之星
func BackStarClass(c *gin.Context) {
	starlist, err := starService.QStarClass(1)
	if err != nil {
		response.ResponseErrorWithMsg(c, 400, "未找到班级之星")
		return
	}
	data := map[string]any{
		"starlist": starlist,
	}
	response.ResponseSuccess(c, data)
}

// BackStarGrade 返回前台年级成长之星
func BackStarGrade(c *gin.Context) {
	starlist, err := starService.QStarClass(2)
	if err != nil {
		response.ResponseErrorWithMsg(c, 400, "未找到班级之星")
	}
	data := map[string]any{
		"starlist": starlist,
	}
	response.ResponseSuccess(c, data)
}

// BackStarCollege 返回前台年级成长之星
func BackStarCollege(c *gin.Context) {
	starlist, err := starService.QStarClass(3)
	if err != nil {
		response.ResponseErrorWithMsg(c, 400, "未找到班级之星")
	}
	data := map[string]any{
		"starlist": starlist,
	}
	response.ResponseSuccess(c, data)
}

// ChangeStatus 修改是否可以再次推选
func ChangeStatus(c *gin.Context) {
	token := c.GetHeader("token")
	username, err := token2.GetUsername(token)
	err = mysql.UpdateOne(username)
	if err != nil {
		response.ResponseErrorWithMsg(c, 400, "修改状态失败")
	}
}
