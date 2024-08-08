package mysql

import (
	"fmt"
	model "studentGrow/models/gorm_model"
	"studentGrow/models/jrx_model"
	"time"
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

// 查询下拉框字段的值
func GetPullValue() {

}

// 根据搜索框内容查询学生信息列表
func GetStuMesList() []jrx_model.StuMesStruct {
	// 从mysql中获取数据到user表中
	var userSlice []model.User
	DB.Unscoped().Select("name", "username", "password", "class", "PlusTime", "gender", "PhoneNumber", "ban").Find(&userSlice)

	// 从user表中获取数据到StuMesStruct中
	stuMesSlice := make([]jrx_model.StuMesStruct, len(userSlice))
	for i := 0; i < len(userSlice); i++ {
		stuMesSlice[i].Name = userSlice[i].Name
		stuMesSlice[i].Username = userSlice[i].Username
		stuMesSlice[i].Password = userSlice[i].Password
		stuMesSlice[i].Class = userSlice[i].Class

		// 解析 PlusTime 字符串并格式化
		t, err := time.Parse(time.RFC3339, userSlice[i].PlusTime)
		if err == nil {
			stuMesSlice[i].Year = t.Format("2006") // 格式化为 YYYY-MM-DD
		} else {
			stuMesSlice[i].Year = userSlice[i].PlusTime // 处理解析错误
		}

		stuMesSlice[i].Gender = userSlice[i].Gender
		stuMesSlice[i].Telephone = userSlice[i].PhoneNumber
		stuMesSlice[i].Ban = userSlice[i].Ban
		// stuMes[i].IsManager = users[i].is_manager
	}

	for k, user := range stuMesSlice {
		fmt.Println(k, user)
	}

	return stuMesSlice
}

// 获取不同的班级
func GetDiffClass() []string {
	var diffClassSlice []string
	DB.Table("users").Select("class").Distinct("class").Order("class ASC").Scan(&diffClassSlice)
	return diffClassSlice
}
