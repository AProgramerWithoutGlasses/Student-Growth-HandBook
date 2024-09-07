package service

import (
	"bytes"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	jsonvalue "github.com/Andrew-M-C/go.jsonvalue"
	"reflect"
	"strconv"
	"studentGrow/dao/mysql"
	"studentGrow/models/gorm_model"
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
func GetClassStructSlice() ([]jrx_model.ClassStruct, error) {
	diffClassSlice, err := mysql.GetDiffClass() // 从mysql中获取不同的class
	if err != nil {
		return nil, err
	}
	classStructSlice := make([]jrx_model.ClassStruct, len(diffClassSlice))
	for i, class := range diffClassSlice {
		classStructSlice[i] = jrx_model.ClassStruct{
			Id_class: class,
			Class:    class,
		}
	}
	return classStructSlice, err
}

// 根据搜索条件，创建sql语句
func CreateQuerySql(stuMessage *jsonvalue.V, queryParmaStruct jrx_model.QueryParmaStruct) string {
	// 将请求的数据转换成map
	stuMesMap := stuMessage.ForRangeObj()

	// 初始化查询学生信息的sql语句
	querySql := `Select name, username, password, class, plus_time, gender, phone_number, ban, is_manager from users where identity = '学生' and deleted_at is NULL`

	// temp标签用于在下方stuMesMap遍历中判断该字段是否为第一个有值的字段

	// 对请求数据的map进行遍历，判断每个字段是否为空
	for k, v := range stuMesMap {
		switch k {
		case "year":
			if v.IsNull() || v.String() == "" { //如果字段值为 null 或 零值 	// IsNull()只对值为null的起效，不对其余类型的空值起效
				fmt.Println("year null")
			} else { // 如果字段值有值
				fmt.Println("year")
				querySql = querySql + " and YEAR(plus_time) = " + v.String() // 对sql语句加上该字段对应的限定条件
			}

		case "class":
			if v.IsNull() || queryParmaStruct.Class == "" {
				fmt.Println("class null")
			} else {
				fmt.Println("class")
				querySql = querySql + " and class = '" + v.String() + "'"
			}

		case "gender":
			if v.IsNull() || queryParmaStruct.Gender == "" {
				fmt.Println("gender null")
			} else {
				fmt.Println("gender")
				querySql = querySql + " and gender = '" + v.String() + "'"
			}

		case "isDisable":
			if v.IsNull() || v.String() == "" {
				fmt.Println("isDisable null")
			} else {
				fmt.Println("isDisable")
				querySql = querySql + " and ban = " + v.String()
			}

		case "searchSelect":
			if v.IsNull() || queryParmaStruct.SearchSelect == "" {
				fmt.Println("searchSelect null")
			} else {
				fmt.Println("searchSelect")
				querySql = querySql + " and " + queryParmaStruct.SearchSelect + " like '%" + queryParmaStruct.SearchMessage + "%'"
			}

		}

	}

	return querySql
}

// GetReqMes 将请求信息整理到结构体
func GetReqMes(stuMessage *jsonvalue.V) jrx_model.QueryParmaStruct {
	// 获取请求信息中各个字段的值
	yearValue, err := stuMessage.GetString("year")
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

// 获取导出学生信息的 excel表格
func GetSelectedStuExcel(selectedStuMesStruct jrx_model.SelectedStuMesStruct) (*bytes.Buffer, error) {
	// 提取处学号数组
	usernameSlice := make([]string, len(selectedStuMesStruct.Selected_students))
	for i, v := range selectedStuMesStruct.Selected_students {
		usernameSlice[i] = v.Username
	}
	fmt.Println(usernameSlice)
	users, err := mysql.QuerySelectedUser(usernameSlice)
	if err != nil {
		return nil, err
	}

	// 创建 Excel 文件
	f := excelize.NewFile()

	// 设置表头
	f.SetCellValue("Sheet1", "A1", "入学年份")
	f.SetCellValue("Sheet1", "B1", "班级")
	f.SetCellValue("Sheet1", "C1", "姓名")
	f.SetCellValue("Sheet1", "D1", "学号")

	// 填充数据
	for i, user := range users {
		row := i + 2 // 从第二行开始填充数据
		f.SetCellValue("Sheet1", fmt.Sprintf("A%d", row), user.PlusTime.Format("2006"))
		f.SetCellValue("Sheet1", fmt.Sprintf("B%d", row), user.Class)
		f.SetCellValue("Sheet1", fmt.Sprintf("C%d", row), user.Name)
		f.SetCellValue("Sheet1", fmt.Sprintf("D%d", row), user.Username)
	}

	// 将 Excel 文件写入内存
	excelData, err := f.WriteToBuffer()
	if err != nil {
		return nil, err
	}

	return excelData, err
}

// banUserService
func BanUserService(username string) (name string, temp int, err error) {
	// 根据学号获取id
	id, err := mysql.GetIdByUsername(username)
	if err != nil {
		return name, temp, err
	}

	// 获取该学生姓名
	name, err = mysql.GetNameById(id)
	if err != nil {
		return name, temp, err
	}

	// mysql中封禁该学生
	temp, err = mysql.BanStudent(id)
	if err != nil {
		return name, temp, err
	}

	return name, temp, err
}

// GetStuMesList 根据搜索框内容查询学生信息列表
func GetStuMesList(querySql string) ([]jrx_model.StuMesStruct, error) {
	// 从mysql中获取数据到user表中
	userSlice, err := mysql.GetUserListBySql(querySql)
	if err != nil {
		return nil, err
	}

	// 从user表中获取数据到stuMesSlice中
	stuMesSlice := make([]jrx_model.StuMesStruct, len(userSlice))
	for i := 0; i < len(userSlice); i++ {
		stuMesSlice[i].Name = userSlice[i].Name
		stuMesSlice[i].Username = userSlice[i].Username
		stuMesSlice[i].Password = userSlice[i].Password
		stuMesSlice[i].Class = userSlice[i].Class
		stuMesSlice[i].Year = userSlice[i].PlusTime.Format("2006") // 日期只保留年份
		stuMesSlice[i].Gender = userSlice[i].Gender
		stuMesSlice[i].Telephone = userSlice[i].PhoneNumber
		stuMesSlice[i].Ban = userSlice[i].Ban

		// 获取管理员等级信息
		if userSlice[i].IsManager {
			managerType, err := GetManagerType(userSlice[i].Username)
			if err != nil {
				return nil, err
			}
			stuMesSlice[i].ManagerType = managerType
		} else {
			stuMesSlice[i].ManagerType = "无"
		}

	}

	for k, user := range stuMesSlice {
		fmt.Println("转化成功", k, user)
	}

	return stuMesSlice, err
}

func CalculateNowGrade(birthday time.Time) (grade string) {
	now := time.Now()

	// 计算出相差的年数
	age := now.Year() - birthday.Year()

	// 如果当前月份小于出生月份,或者当前月份等于出生月份但当前日期小于出生日期,则age减1
	if now.Month() < birthday.Month() || (now.Month() == birthday.Month() && now.Day() < birthday.Day()) {
		age--
	}

	switch age {
	case 0:
		grade = "grade1"

	case 1:
		grade = "grade2"

	case 2:
		grade = "grade3"

	case 3:
		grade = "grade4"
	}

	return grade
}

func CalculateNowGradeByClass(class string) (grade string) {
	classNumber, err := strconv.Atoi("20" + class[6:8])
	if err != nil {
		return grade
	}

	startGrade1 := time.Date(classNumber, 8, 1, 0, 0, 0, 0, time.UTC)
	endGrade1 := time.Date(classNumber+1, 8, 1, 0, 0, 0, 0, time.UTC)
	startGrade2 := time.Date(classNumber+1, 8, 1, 0, 0, 0, 0, time.UTC)
	endGrade2 := time.Date(classNumber+2, 8, 1, 0, 0, 0, 0, time.UTC)
	startGrade3 := time.Date(classNumber+2, 8, 1, 0, 0, 0, 0, time.UTC)
	endGrade3 := time.Date(classNumber+3, 8, 1, 0, 0, 0, 0, time.UTC)
	startGrade4 := time.Date(classNumber+3, 8, 1, 0, 0, 0, 0, time.UTC)
	endGrade4 := time.Date(classNumber+4, 8, 1, 0, 0, 0, 0, time.UTC)

	now := time.Now()
	if now.After(startGrade1) && now.Before(endGrade1) {
		grade = "grade1"
	} else if now.After(startGrade2) && now.Before(endGrade2) {
		grade = "grade2"
	} else if now.After(startGrade3) && now.Before(endGrade3) {
		grade = "grade3"
	} else if now.After(startGrade4) && now.Before(endGrade4) {
		grade = "grade4"
	}

	return grade
}

func EditStuService(user jrx_model.ChangeStuMesStruct) error {
	id, err := mysql.GetIdByUsername(user.Username)
	if err != nil {
		return err
	}

	oldUser, err := mysql.GetUser(id)
	if err != nil {
		return err
	}

	err = mysql.ChangeStudentMessage(id, user)
	if err != nil {
		return err
	}

	// 记录更改
	var userEditRecord gorm_model.UserEditRecord

	if oldUser.Class != user.Class {
		userEditRecord.OldClass = oldUser.Class
		userEditRecord.NewClass = user.Class
	}

	if oldUser.PhoneNumber != user.Phone_number {
		userEditRecord.OldPhoneNumber = oldUser.PhoneNumber
		userEditRecord.NewPhoneNumber = user.Phone_number
	}

	if oldUser.Password != user.Password {
		userEditRecord.OldPassword = oldUser.Password
		userEditRecord.NewPassword = user.Password
	}

	fmt.Println("userEditRecordStruct ： ", userEditRecord)

	err = mysql.EditUserRecord(userEditRecord)
	if err != nil {
		return err
	}

	return nil

}
