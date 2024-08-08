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

// 查询学生信息
func QueryStuContro(c *gin.Context) {
	stuMessage, err := readMessage.GetJsonvalue(c)
	if err != nil {
		fmt.Println("stuManage.QueryStuContro() readMessage.GetJsonvalue() err :", err)
	}

	stuMesMap := stuMessage.ForRangeObj()

	for k, v := range stuMesMap {
		switch k {
		case "year":
			if v != nil {

			} else {

			}
		}
	}

	// 响应数据的获取
	stuInfo := mysql.GetStuMesList()
	yearStructSlice := service.GetYearStructSlice()
	classStructSlice := service.GetClassStructSlice()

	responseStruct := jrx_model.ResponseStruct{
		Year:            yearStructSlice,
		Class:           classStructSlice,
		StuInfo:         stuInfo,
		AllStudentCount: len(stuInfo),
	}

	// 响应数据
	response.ResponseSuccess(c, responseStruct)
}
