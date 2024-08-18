package mysql

import (
	"fmt"
	_ "gorm.io/gorm/clause"
	"studentGrow/models/gorm_model"
	"studentGrow/models/jrx_model"
)

// 将新的用户自述在mysql中进行更行
func UpdateSelfContent(id int, newSelfContent string) error {
	return DB.Table("users").Where("id = ?", id).Update("self_content", newSelfContent).Error
}

// 获取mysql中的用户自述
func GetSelfContent(id int) (string, error) {
	var users gorm_model.User
	err := DB.Where("id = ?", id).First(&users).Error // Unscoped()用于解除搜索时会自动加上deleted_at字段的限制
	return users.SelfContent, err
}

// 根据学号获取id
func GetIdByUsername(username string) (int, error) {
	var users gorm_model.User
	err := DB.Where("username = ?", username).First(&users).Error
	return int(users.ID), err
}

// 根据id获取姓名
func GetNameById(id int) (string, error) {
	var users gorm_model.User
	err := DB.Where("id = ?", id).First(&users).Error
	return users.Name, err
}

// 获取不同的班级
func GetDiffClass() ([]string, error) {
	var diffClassSlice []string
	err := DB.Table("users").Select("class").Distinct("class").Where("LENGTH(class) < 10").Order("class ASC").Scan(&diffClassSlice).Error
	return diffClassSlice, err
}

// 添加单个学生
func AddSingleStudent(users *gorm_model.User) error {
	err := DB.Select("name", "username", "password", "class", "gender", "identity", "plus_time").Create(users).Error
	return err
}

// 添加单个老师
func AddSingleTeacher(users *gorm_model.User) error {
	err := DB.Select("name", "username", "password", "gender", "identity").Create(users).Error
	return err
}

// 删除单个学生
func DeleteSingleUser(id int) error {
	err := DB.Table("users").Where("id = ?", id).Delete(nil).Error
	return err
}

// 封禁该用户
func BanStudent(id int) (int, error) {
	var temp int
	var users gorm_model.User
	DB.Take(&users, id)
	var err error
	if users.Ban == false {
		err = DB.Model(&gorm_model.User{}).Where("id = ?", id).Update("ban", 1).Error
		temp = 1
	} else if users.Ban == true {
		err = DB.Model(&gorm_model.User{}).Where("id = ?", id).Update("ban", 0).Error
		temp = 0
	}

	return temp, err
}

// 修改用户信息username
func ChangeStudentMessage(id int, users jrx_model.ChangeStuMesStruct) error {
	err := DB.Model(&gorm_model.User{}).Where("id = ?", id).Updates(users).Error
	return err
}

// 修改老师信息username
func ChangeTeacherMessage(id int, newTeacher jrx_model.ChangeTeacherMesStruct) error {
	err := DB.Model(&gorm_model.User{}).Where("id = ?", id).Select("name", "username", "password", "gender").Updates(&newTeacher).Error
	return err
}

// 将用户设置为管理员
func SetStuManager(username string, casbinCid string) error {
	// user表设置
	id, err := GetIdByUsername(username)
	if err != nil {
		return err
	}

	err = DB.Model(&gorm_model.User{}).Where("id = ?", id).Update("is_manager", 1).Error
	if err != nil {
		return err
	}

	// casbin_ruler表设置
	casbinUser := gorm_model.UserCasbinRules{
		CUsername: username,
		CasbinCid: casbinCid,
	}
	err = DB.Create(&casbinUser).Error
	fmt.Println("casbinCid:", casbinUser.CasbinCid)
	if err != nil {
		return err
	}

	return err
}

// 将用户修改管理员等级
func ChangeStuManager(username string, casbinCid string) error {
	// user表设置
	id, err := GetIdByUsername(username)
	if err != nil {
		return err
	}

	err = DB.Model(&gorm_model.User{}).Where("id = ?", id).Update("is_manager", 1).Error
	if err != nil {
		return err
	}

	// casbin_ruler表设置
	casbinUser := gorm_model.UserCasbinRules{
		CUsername: username,
		CasbinCid: casbinCid,
	}
	err = DB.Model(&casbinUser).Where("c_username = ?", username).Update("casbin_cid", casbinCid).Error
	fmt.Println("casbinCid:", casbinUser.CasbinCid)
	if err != nil {
		return err
	}

	return err
}

// 取消用户管理员身份
func CancelStuManager(username string, casbinCid string) error {
	// user表设置
	id, err := GetIdByUsername(username)
	if err != nil {
		return err
	}

	err = DB.Model(&gorm_model.User{}).Where("id = ?", id).Update("is_manager", 0).Error
	if err != nil {
		return err
	}

	// casbin_ruler表设置

	err = DB.Table("user_casbin_rules").Where("c_username = ?", username).Delete(nil).Error

	if err != nil {
		return err
	}

	return err
}

// 判断用户是否存在
func ExistedUsername(username string) error {
	err := DB.Where("username = ?", username).First(&gorm_model.User{}).Error
	return err
}

// 查询选中的用户
func QuerySelectedUser(usernameSlice []string) ([]gorm_model.User, error) {
	var users []gorm_model.User
	err := DB.Where("username IN (?)", usernameSlice).Find(&users).Error
	return users, err
}

func GetAllUserCount(identity string) (int64, error) {
	var user gorm_model.User
	var count int64
	err := DB.Model(&user).Where("identity = ?", identity).Count(&count).Error
	return count, err
}

// GetStuMesList 根据搜索框内容查询学生信息列表
func GetTeacherList(querySql string) ([]jrx_model.QueryTeacherResStruct, error) {
	// 从mysql中获取数据到user表中
	var userSlice []jrx_model.QueryTeacherResStruct

	err := DB.Raw(querySql).Find(&userSlice).Error
	if err != nil {
		return nil, err
	}

	return userSlice, nil
}

func GetManagerCId(username string) (string, error) {
	var casbinUser gorm_model.UserCasbinRules
	err := DB.Where("c_username = ?", username).First(&casbinUser).Error
	return casbinUser.CasbinCid, err
}

func GetUserListBySql(querySql string) ([]gorm_model.User, error) {
	var userSlice []gorm_model.User
	err := DB.Raw(querySql).Find(&userSlice).Error
	if err != nil {
		return nil, err
	}
	return userSlice, err
}

// 根据学号获取 managerType
func GetIsManagerByUsername(username string) (bool, error) {
	var users gorm_model.User
	err := DB.Where("username = ?", username).First(&users).Error
	return users.IsManager, err
}

func GetHomepageUserMesDao(id int) (*gorm_model.User, error) {
	var userMes gorm_model.User
	err := DB.Model(&gorm_model.User{}).Where("id = ?", id).First(&userMes).Error
	if err != nil {
		return nil, err
	}
	return &userMes, err
}

// 获取个人主页粉丝个数
func GetHomepageFansCountDao(id int) (int, error) {
	var count int64
	err := DB.Table("user_followers").Where("user_id = ?", id).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return int(count), err
}

// 获取个人主页关注个数
func GetHomepageConcernCountDao(id int) (int, error) {
	var count int64
	err := DB.Table("user_followers").Where("follower_id = ?", id).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return int(count), err
}

// 获取个人主页获赞个数(文章获赞)
func GetHomepageLikeCountDao(id int) (int, error) {
	var count int
	// Article表 满足 user_id = id 的所有行中 like_amount 列的值的总和
	err := DB.Model(&gorm_model.Article{}).Where("user_id = ?", id).Select("SUM(like_amount)").Scan(&count).Error
	if err != nil {
		return 0, err
	}
	return count, err
}

// 获取个人主页积分值
func GetHomepagePointDao(id int) (int, error) {
	var count int
	// Article表 满足 user_id = id 的所有行中 like_amount 列的值的总和
	err := DB.Model(&gorm_model.UserPoint{}).Where("user_id = ?", id).Select("SUM(point)").Scan(&count).Error
	if err != nil {
		return 0, err
	}
	return count, err
}
