package stuManage

import (
	"fmt"
	"github.com/gin-gonic/gin"
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

// QueryStuContro 查询学生信息
func QueryStuContro(c *gin.Context) {
	// 接收请求数据
	stuMessage, err := readMessage.GetJsonvalue(c)
	if err != nil {
		fmt.Println("stuManage.QueryStuContro() readMessage.GetJsonvalue() err :", err)
	}

	// 获取请求信息中各个字段的值
	yearInt, err := stuMessage.GetInt("year")
	if err != nil {
		fmt.Println("year GetInt() err : ", err)
	}

	classValue, err := stuMessage.GetString("class")
	if err != nil {
		fmt.Println("class GetString() err : ", err)
	}

	genderValue, err := stuMessage.GetString("gender")
	if err != nil {
		fmt.Println("gender GetString() err : ", err)
	}

	isDisableValue, err := stuMessage.GetString("isDisable")
	if err != nil {
		fmt.Println("isDisable GetString() err : ", err)
	}

	searchSelectValue, err := stuMessage.GetString("searchSelect")
	if searchSelectValue == "telephone" {
		searchSelectValue = "phone_number"
	}
	if err != nil {
		fmt.Println("searchSelect GetString() err : ", err)
	}

	searchMessageValue, err := stuMessage.GetString("searchMessage")
	if err != nil {
		fmt.Println("searchMessage GetString() err : ", err)
	}

	// 将请求的数据转换成map
	stuMesMap := stuMessage.ForRangeObj()

	// 初始化查询学生信息的sql语句
	querySql := `Select name, username, password, class, plus_time, gender, phone_number, ban, is_manager from users`

	// temp标签用于在下方stuMesMap遍历中判断该字段是否为第一个有值的字段
	temp := 0

	// 对请求数据的map进行遍历，判断每个字段是否为空
	for k, v := range stuMesMap {
		switch k {
		case "year":
			if v.IsNull() || yearInt == 0 { //如果字段值为 null 或 零值 	// IsNull()只对值为null的起效，不对其余类型的空值起效
				fmt.Println("year null")
			} else { // 如果字段值有值
				fmt.Println("year")
				if temp == 0 { // 如果是第一个有值的字段
					querySql = querySql + " where YEAR(plus_time) = " + v.String()
					temp++
					break
				}
				querySql = querySql + " and YEAR(plus_time) = " + v.String() // 对sql语句加上该字段对应的限定条件
			}

		case "class":
			if v.IsNull() || classValue == "" {
				fmt.Println("class null")
			} else {
				fmt.Println("class")
				if temp == 0 {
					querySql = querySql + " where class = '" + v.String() + "'"
					temp++
					break
				}
				querySql = querySql + " and class = '" + v.String() + "'"
			}

		case "gender":
			if v.IsNull() || genderValue == "" {
				fmt.Println("gender null")
			} else {
				fmt.Println("gender")
				if temp == 0 {
					querySql = querySql + " where gender = '" + v.String() + "'"
					temp++
					break
				}
				querySql = querySql + " and gender = '" + v.String() + "'"
			}

		case "isDisable":
			if v.IsNull() || isDisableValue == "" {
				fmt.Println("isDisable null")
			} else {
				fmt.Println("isDisable")
				if temp == 0 {
					querySql = querySql + " where ban = " + v.String()
					temp++
					break
				}
				querySql = querySql + " and ban = " + v.String()
			}

		case "searchSelect":
			if v.IsNull() || searchSelectValue == "" {
				fmt.Println("searchSelect null")
			} else {
				fmt.Println("searchSelect")
				if temp == 0 {
					querySql = querySql + " where " + searchSelectValue + " like '%" + searchMessageValue + "%'"
					temp++
					break
				}
				querySql = querySql + " and " + searchSelectValue + " like '%" + searchMessageValue + "%'"
			}

		}
	}

	// 响应数据的获取
	stuInfo := mysql.GetStuMesList(querySql)
	yearStructSlice := service.GetYearStructSlice()
	classStructSlice := service.GetClassStructSlice()

	// 响应结构体的初始化
	responseStruct := jrx_model.ResponseStruct{
		Year:            yearStructSlice,
		Class:           classStructSlice,
		StuInfo:         stuInfo,
		AllStudentCount: len(stuInfo),
	}

	// 响应数据
	response.ResponseSuccess(c, responseStruct)
}
