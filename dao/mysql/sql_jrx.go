package mysql

import (
	"fmt"
	model "studentGrow/models/gorm_model"
	"studentGrow/models/jrx_model"
)

// 将新的用户自述在mysql中进行更行
func UpdateSelfContent(id int, newSelfContent string) error {
	return DB.Table("users").Where("id = ?", id).Update("self_content", newSelfContent).Error
}

// 获取mysql中的用户自述
func GetSelfContent(id int) (string, error) {
	var users model.User
	err := DB.Unscoped().Where("id = ?", id).First(&users).Error // Unscoped()用于解除搜索时会自动加上deleted_at字段的限制
	return users.SelfContent, err
}

// 根据学号获取id
func GetIdByUsername(username string) (int, error) {
	var users model.User
	err := DB.Unscoped().Where("username = ?", username).First(&users).Error // Unscoped()用于解除搜索时会自动加上deleted_at字段的限制
	return int(users.ID), err
}

// 获取不同的班级
func GetDiffClass() []string {
	var diffClassSlice []string
	DB.Table("users").Select("class").Distinct("class").Order("class ASC").Scan(&diffClassSlice)
	return diffClassSlice
}

// GetStuMesList 根据搜索框内容查询学生信息列表
func GetStuMesList(querySql string) []jrx_model.StuMesStruct {

	// 从mysql中获取数据到user表中
	var userSlice []model.User

	//DB.Select("name", "username", "password", "class", "plus_time", "gender", "phone_number", "ban", "is_manager").Where("YEAR(plus_time) = ?  and class IS NULL OR class = ? and gender = ? and ban = ?", parmaStruct.Year, parmaStruct.Class, parmaStruct.Gender, parmaStruct.IsDisable).Find(&userSlice)
	DB.Raw(querySql).Find(&userSlice)

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
		stuMesSlice[i].IsManager = userSlice[i].IsManager
	}

	for k, user := range stuMesSlice {
		fmt.Println("转化成功", k, user)
	}

	return stuMesSlice
}

// 添加单个学生
func AddSingleStudent(users *model.User) {
	DB.Select("name", "username", "password", "class", "identity").Create(users)
}

// 删除单个学生
func DeleteSingleStudent(id string) {
	// DB.Delete()
}
