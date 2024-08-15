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

// 用于存储查询参数
var queryParmaStruct jrx_model.QueryParmaStruct
var querySql string
var queryAllStuNumber int

//token := c.GetHeader("token")
//username, err := token2.GetUsername(token)
//
//role, err := token2.GetRole(token)
//switch role {
//case "grade1":
//
//}

// QueryStuContro 查询学生信息
func QueryStuContro(c *gin.Context) {

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

	// 将请求数据整理到结构体
	queryParmaStruct = service.GetReqMes(stuMessage)
	// 获取sql语句
	querySql = service.CreateQuerySql(stuMessage, queryParmaStruct)

	// 获取符合条件的所有学生，用于计算长度
	stuInfo, err := mysql.GetStuMesList(querySql) // 所有学生数据
	if err != nil {
		zap.L().Error("mysql.GetStuMesList(querySql) failed", zap.Error(err))
		response.ResponseError(c, response.ServerErrorCode)
		return
	}

	// 获取所有符合条件的学生数量
	queryAllStuNumber = len(stuInfo)

	// 重置sql语句中的分页部分
	whereSqlIndex := strings.Index(querySql, "limit")
	if whereSqlIndex != -1 {
		afterWhere := querySql[:whereSqlIndex]
		querySql = afterWhere
	}

	// limit 分页查询语句的拼接
	querySql = querySql + " ORDER BY class ASC" + " limit " + strconv.Itoa(limitValue) + " offset " + strconv.Itoa((offsetValue-1)*limitValue)

	// 获取符合条件的当页学生
	stuPageInfo, err := mysql.GetStuMesList(querySql) // 当页学生数据
	if err != nil {
		zap.L().Error("mysql.GetStuMesList(queryPageSql) failed", zap.Error(err))
		response.ResponseError(c, response.ServerErrorCode)
		return
	}
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

//// QueryPageStuContro 查询分页数据
//func QueryPageStuContro(c *gin.Context) {
//	// 接收请求数据
//	stuMessage, err := readMessage.GetJsonvalue(c)
//	if err != nil {
//		fmt.Println("stuManage.QueryStuContro() readMessage.GetJsonvalue() err :", err)
//	}
//
//	offsetValue, err := stuMessage.GetInt("page")
//	if err != nil {
//		fmt.Println("page GetInt() err", err)
//	}
//
//	limitValue, err := stuMessage.GetInt("limit")
//	if err != nil {
//		fmt.Println("limit GetInt() err", err)
//	}
//
//	// 重置sql语句中的分页部分
//	whereSqlIndex := strings.Index(querySql, "limit")
//	if whereSqlIndex != -1 {
//		afterWhere := querySql[:whereSqlIndex]
//		querySql = afterWhere
//	}
//
//	// limit 分页查询语句的拼接
//	querySql = querySql + " limit " + strconv.Itoa(limitValue) + " offset " + strconv.Itoa((offsetValue-1)*limitValue)
//
//	// 响应数据的获取
//	stuPageInfo, _ := mysql.GetStuMesList(querySql) // 当页学生数据
//	yearStructSlice := service.GetYearStructSlice()
//	classStructSlice := service.GetClassStructSlice()
//
//	// 响应结构体的初始化
//	responseStruct := jrx_model.ResponseStruct{
//		Year:            yearStructSlice,
//		Class:           classStructSlice,
//		StuInfo:         stuPageInfo,
//		AllStudentCount: queryAllStuNumber,
//	}
//
//	// 响应数据
//	response.ResponseSuccess(c, responseStruct)
//
//}
