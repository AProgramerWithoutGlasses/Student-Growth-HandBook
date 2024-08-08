package service

import (
	"fmt"
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
			Id_Year: "1",
			Year:    NowYearChange(-3),
		},
		{
			Id_Year: "2",
			Year:    NowYearChange(-2),
		},
		{
			Id_Year: "3",
			Year:    NowYearChange(-1),
		},
		{
			Id_Year: "4",
			Year:    NowYearChange(0),
		},
	}
	return yearStructSlice
}

// 获取返回给前端的class结构体切片
func GetClassStructSlice() []jrx_model.ClassStruct {
	diffClassSlice := mysql.GetDiffClass()	// 从mysql中获取不同的class
	classStructSlice := make([]jrx_model.ClassStruct, len(diffClassSlice))
	for i, class := range diffClassSlice {
		classStructSlice[i] = jrx_model.ClassStruct{
			Id_class: strconv.Itoa(i + 1),
			Class:    class,
		}
	}
	return classStructSlice
}
