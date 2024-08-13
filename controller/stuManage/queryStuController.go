package stuManage

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"strings"
	"studentGrow/dao/mysql"
	"studentGrow/models/jrx_model"
	"studentGrow/pkg/response"
	"studentGrow/service"
	"studentGrow/utils/readMessage"
)

// { 响应
//
//		code:
//		msg:
//		data: {
//				year: [2021, 2022, 2023, 2024]
//				class: [计科221, 计科222, ]
//				stuInfo: [{name: "张四", userName: "20221544", password: "22mm", gender: true, }, {}, {}]
//				allStudentCount: 230
//			   }
//	}
//
// {	请求
//
//		year:
//		class:
//		gender:
//		isDisable:
//
//		searchSelect: [name, username, telephone]
//		searchMessage:
//	}

//// 用于入学年份下拉框
//type YearStruct struct {
//	Id_Year string `json:"value"`
//	Year    string `json:"label"`
//}
//
//// 用于班级下拉框
//type ClassStruct struct {
//	Id_class string `json:"value"`
//	Class    string `json:"label"`
//}
//
//type ResponseStruct struct {
//	Year            []YearStruct              `json:"year"`
//	Class           []ClassStruct             `json:"class"`
//	StuInfo         []gorm_model.StuMesStruct `json:"stuInfo"`
//	AllStudentCount int                       `json:"allStudentCount"`
//}

// 用于存储查询参数
var queryParmaStruct jrx_model.QueryParmaStruct
var querySql string
var queryAllStuNumber int

/*
func QueryStuContro(c *gin.Context) {
	// 接收参数
	var input struct {
		Year          int    `form:"year"`
		Class         string `form:"class"`
		Gender        int    `form:"gender"`
		IsDisable     bool   `form:"isDisable"`
		searchSelect  string `form:"searchSelect"`
		searchMessage string `form:"searchMessage"`
	}

	if err := c.Bind(&input); err != nil {
		response.ResponseError(c, response.ParamFail)
		return
	}
	// 参数校验

	// 调用server层方法

		data,err :=  server.QueryStu(input.Year,input.Class,input.Gender)
		if err != nil {
			log.logger("QueryStuContro failed:%v",err)
				response.ResponseError(c, response.ParamFail)
		        return
		}


	// 返回响应
	response.ResponseSuccess(c, data)
}
*/

// QueryStuContro 查询学生信息
func QueryStuContro(c *gin.Context) {
	// 接收请求数据
	stuMessage, err := readMessage.GetJsonvalue(c)
	if err != nil {
		fmt.Println("stuManage.QueryStuContro() readMessage.GetJsonvalue() err :", err)
	}

	// 将请求数据整理到结构体
	queryParmaStruct = service.GetReqMes(stuMessage)
	// 获取sql语句
	querySql = service.CreateQuerySql(stuMessage, queryParmaStruct)
	queryPageSql := querySql + " limit " + strconv.Itoa(8) + " offset " + strconv.Itoa(0)

	// 响应数据的获取
	stuInfo, err := mysql.GetStuMesList(querySql) // 所有学生数据
	if err != nil {
		zap.L().Error("mysql.GetStuMesList(querySql) failed", zap.Error(err))
		response.ResponseError(c, response.ServerErrorCode)
		return
	}
	stuPageInfo, err := mysql.GetStuMesList(queryPageSql) // 当页学生数据
	if err != nil {
		zap.L().Error("mysql.GetStuMesList(queryPageSql) failed", zap.Error(err))
		response.ResponseError(c, response.ServerErrorCode)
		return
	}
	yearStructSlice := service.GetYearStructSlice()
	classStructSlice := service.GetClassStructSlice()

	// 获取所有符合条件的学生数量
	queryAllStuNumber = len(stuInfo)

	// 响应结构体的初始化
	responseStruct := jrx_model.ResponseStruct{
		Year:            yearStructSlice,
		Class:           classStructSlice,
		StuInfo:         stuPageInfo,
		AllStudentCount: queryAllStuNumber,
	}

	// 响应数据
	response.ResponseSuccess(c, responseStruct)
}

// QueryPageStuContro 查询分页数据
func QueryPageStuContro(c *gin.Context) {
	// 接收请求数据
	stuMessage, err := readMessage.GetJsonvalue(c)
	if err != nil {
		fmt.Println("stuManage.QueryStuContro() readMessage.GetJsonvalue() err :", err)
	}

	offsetValue, err := stuMessage.GetInt("page")
	if err != nil {
		fmt.Println("page GetInt() err", err)
	}

	limitValue, err := stuMessage.GetInt("limit")
	if err != nil {
		fmt.Println("limit GetInt() err", err)
	}

	// 重置sql语句中的分页部分
	whereSqlIndex := strings.Index(querySql, "limit")
	if whereSqlIndex != -1 {
		afterWhere := querySql[:whereSqlIndex]
		querySql = afterWhere
	}

	// limit 分页查询语句的拼接
	querySql = querySql + " limit " + strconv.Itoa(limitValue) + " offset " + strconv.Itoa(offsetValue)

	// 响应数据的获取
	stuPageInfo, _ := mysql.GetStuMesList(querySql) // 当页学生数据
	yearStructSlice := service.GetYearStructSlice()
	classStructSlice := service.GetClassStructSlice()

	// 响应结构体的初始化
	responseStruct := jrx_model.ResponseStruct{
		Year:            yearStructSlice,
		Class:           classStructSlice,
		StuInfo:         stuPageInfo,
		AllStudentCount: queryAllStuNumber,
	}

	// 响应数据
	response.ResponseSuccess(c, responseStruct)

}
