package service

import (
	"fmt"
	jsonvalue "github.com/Andrew-M-C/go.jsonvalue"
	"reflect"
	"strconv"
	"studentGrow/dao/mysql"
	"studentGrow/models/jrx_model"
	"time"
)

// 判断前端发来的结构体中非空字段的内容
func GetNotEmptyFields(v interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Struct {
		return result
	}

	for i := 0; i < val.NumField(); i++ {
		fieldVal := val.Field(i)
		fieldType := val.Type().Field(i)
		if !isValueEmpty(fieldVal) {
			result[fieldType.Name] = fieldVal.Interface()
		}
	}

	return result
}

func isValueEmpty(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.String, reflect.Array, reflect.Slice, reflect.Map, reflect.Ptr, reflect.Interface:
		return v.IsZero()
	default:
		return false
	}
}

func NowYearChange(num int) string {
	nowYearInt, err := strconv.Atoi(time.Now().Format("2006"))
	if err != nil {
		fmt.Println("strconv.Atoi(time.Now().Format(\"2006\")) err : ", err)
	}
	changedYear := nowYearInt + num
	changedYearStr := strconv.Itoa(changedYear)
	return changedYearStr
}

func GetYearStructSlice() []jrx_model.YearStruct {
	var yearStructSlice = []jrx_model.YearStruct{
		{
			Id_Year: NowYearChange(-3),
			Year:    NowYearChange(-3),
		},
		{
			Id_Year: NowYearChange(-2),
			Year:    NowYearChange(-2),
		},
		{
			Id_Year: NowYearChange(-1),
			Year:    NowYearChange(-1),
		},
		{
			Id_Year: NowYearChange(0),
			Year:    NowYearChange(0),
		},
	}
	return yearStructSlice
}

// 获取返回给前端的class结构体切片
func GetClassStructSlice() []jrx_model.ClassStruct {
	diffClassSlice := mysql.GetDiffClass() // 从mysql中获取不同的class
	classStructSlice := make([]jrx_model.ClassStruct, len(diffClassSlice))
	for i, class := range diffClassSlice {
		classStructSlice[i] = jrx_model.ClassStruct{
			Id_class: class,
			Class:    class,
		}
	}
	return classStructSlice
}

// 根据搜索条件，创建sql语句
func CreateQuerySql(stuMessage *jsonvalue.V, queryParmaStruct jrx_model.QueryParmaStruct) string {
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
			if v.IsNull() || queryParmaStruct.Year == 0 { //如果字段值为 null 或 零值 	// IsNull()只对值为null的起效，不对其余类型的空值起效
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
			if v.IsNull() || queryParmaStruct.Class == "" {
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
			if v.IsNull() || queryParmaStruct.Gender == "" {
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
			if v.IsNull() {
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
			if v.IsNull() || queryParmaStruct.SearchSelect == "" {
				fmt.Println("searchSelect null")
			} else {
				fmt.Println("searchSelect")
				if temp == 0 {
					querySql = querySql + " where " + queryParmaStruct.SearchSelect + " like '%" + queryParmaStruct.SearchMessage + "%'"
					temp++
					break
				}
				querySql = querySql + " and " + queryParmaStruct.SearchSelect + " like '%" + queryParmaStruct.SearchMessage + "%'"
			}

		}

	}

	return querySql
}

// GetReqMes 将请求信息整理到结构体
func GetReqMes(stuMessage *jsonvalue.V) jrx_model.QueryParmaStruct {
	// 获取请求信息中各个字段的值
	yearValue, err := stuMessage.GetInt("year")
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

	isDisableValue, err := stuMessage.GetBool("isDisable")
	if err != nil {
		fmt.Println("isDisable GetBool() err : ", err)
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

	queryParmaStruct := jrx_model.QueryParmaStruct{
		Year:          yearValue,
		Class:         classValue,
		Gender:        genderValue,
		IsDisable:     isDisableValue,
		SearchSelect:  searchSelectValue,
		SearchMessage: searchMessageValue,
	}

	return queryParmaStruct
}
