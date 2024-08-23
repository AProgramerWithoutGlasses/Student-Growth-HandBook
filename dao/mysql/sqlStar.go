package mysql

import (
	"studentGrow/models/gorm_model"
	"time"
)

// SelName 根据学号查名字，等一系列数据
func SelName(username string) (string, error) {
	var name string
	err := DB.Model(&gorm_model.User{}).Select("name").Where("username = ?", username).Scan(&name).Error
	if err != nil {
		return "", err
	}
	return name, nil
}

// SelId 根据账号查找id
func SelId(username string) (int, error) {
	var id int
	err := DB.Model(&gorm_model.User{}).Select("id").Where("username = ?", username).Scan(&id).Error
	if err != nil {
		return 0, err
	}
	return id, nil
}

// Selfans 根据 id查询粉丝
func Selfans(id int) (int64, error) {
	var fans int64
	err := DB.Table("user_followers").Where("user_id = ?", id).Count(&fans).Error
	if err != nil {
		return 0, err
	}
	return fans, nil
}

// Score 查询积分
func Score(username string) (int, error) {
	var score int
	err := DB.Model(&gorm_model.User{}).Select("point").Where("username = ?", username).Error
	if err != nil {
		return 0, err
	}
	return score, nil
}

// Frequency 被推举次数
func Frequency(username string) (int64, error) {
	var number int64
	err := DB.Model(&gorm_model.Star{}).Where("username = ?", username).Count(&number).Error
	if err != nil {
		return 0, err
	}
	return number, nil
}

// SelHot 查询热度
func SelHot(id int) (int, error) {
	var like int
	var collect int
	err := DB.Model(&gorm_model.Article{}).Select("like_amount").Where("user_id =?", id).Scan(&like).Error
	err = DB.Model(&gorm_model.Article{}).Select("collect_amount").Where("user_id = ?", id).Scan(&collect).Error
	if err != nil {
		return 0, err
	}
	return collect + like, nil
}

// SelStarUser 查询未公布的学号合集
func SelStarUser() ([]string, error) {
	var alluser []string
	err := DB.Table("stars").Where("session = ?", 0).Select("username").Scan(&alluser).Error
	if err != nil {
		return nil, err
	}
	return alluser, nil
}

// SelSearchGrade 查询未公布的学号合集
func SelSearchGrade(name string) ([]string, error) {
	var alluser []string
	err := DB.Model(&gorm_model.Star{}).Where("name LIKE ?", name).Where("session = ?", 0).Select("username").Scan(&alluser).Error
	if err != nil {
		return nil, err
	}
	return alluser, nil
}

// SelPlus 查询入学时间
func SelPlus(username string) (time.Time, error) {
	var plus time.Time
	err := DB.Model(&gorm_model.User{}).Select("plus_time").Where("username = ?", username).Scan(&plus).Error
	if err != nil {
		return time.Time{}, err
	}
	return plus, nil
}

// SelStarColl 院级查询表里学号合集
func SelStarColl() ([]string, error) {
	var alluser []string
	err := DB.Model(&gorm_model.Star{}).Where("type = ?", 2).Where("session = ?", 0).Select("username").Scan(&alluser).Error
	if err != nil {
		return nil, err
	}
	return alluser, nil
}

// SelSearchUser 根据名字班级查找学号--班级管理员搜索
func SelSearchUser(name string, class string) ([]string, error) {
	var username []string
	err := DB.Model(&gorm_model.User{}).Where("class = ?", class).Where("name LIKE ?", "%"+name+"%").Select("username").Scan(&username).Error
	if err != nil {
		return nil, err
	}
	return username, nil
}

// SelSearchColl 院级管理员搜索
func SelSearchColl(name string) ([]string, error) {
	var usernamesli []string
	err := DB.Model(&gorm_model.Star{}).Where("name LIKE ?", "%"+name+"%").Where("type = ?", 2).Where("session = ?", 0).Select("username").Scan(&usernamesli).Error
	if err != nil {
		return nil, err
	}
	return usernamesli, nil
}

// CreatClass 班级管理员推选班级之星
func CreatClass(username string, name string) error {
	stars := gorm_model.Star{
		Username: username,
		Name:     name,
		Type:     1,
		Session:  0,
	}
	err := DB.Create(&stars).Error
	if err != nil {
		return err
	}
	return nil
}

// UpdateGrade 年级管理员推选更新数据
func UpdateGrade(username string) error {
	var star gorm_model.Star
	err := DB.Model(gorm_model.Star{}).Where("username = ?", username).Where("session = ?", 0).Where("type = ?", 1).First(&star).Error
	if err != nil {
		return err
	}
	star.Type = 2
	err = DB.Save(&star).Error
	if err != nil {
		return err
	}
	return nil
}

// UpdateCollege 院级管理员推选更新数据
func UpdateCollege(username string) error {
	var star gorm_model.Star
	err := DB.Model(gorm_model.Star{}).Where("username = ?", username).Where("session = ?", 0).Where("type = ?", 2).First(&star).Error
	if err != nil {
		return err
	}
	star.Type = 3
	err = DB.Save(&star).Error
	if err != nil {
		return err
	}
	return nil
}

// SelMax 查询session字段最大值
func SelMax() (int, error) {
	var maxnum int
	err := DB.Model(&gorm_model.Star{}).Select("MAX(session)").Scan(&maxnum).Error
	if err != nil {
		return 0, err
	}
	return maxnum, nil
}

// UpdateSession 更新字段
func UpdateSession(session int) error {
	err := DB.Model(&gorm_model.Star{}).Where("session = ? ", 0).Updates(map[string]interface{}{"session": session}).Error
	if err != nil {
		return err
	}
	return nil
}

// SelStar 查找指定届数的班级之星
func SelStar(session int, starType int) ([]string, error) {
	var username []string
	err := DB.Model(&gorm_model.Star{}).Where("session = ?", session).Where("type = ?", starType).Select("username").Scan(&username).Error
	if err != nil {
		return nil, err
	}
	return username, nil
}

// Selstarexit 查找这条数据是否存在数据库中
func Selstarexit(username string, SType int) (int64, error) {
	var number int64
	err := DB.Model(&gorm_model.Star{}).Where("username = ?", username).Where("type = ?", SType).Where("session = ?", 0).Count(&number).Error
	if err != nil {
		return 0, err
	}
	return number, nil
}

// SelStatus 查询管理员是否可以添加数据
func SelStatus(username string) (bool, error) {
	var status bool
	err := DB.Model(&gorm_model.UserCasbinRules{}).Where("c_username = ?", username).Select("status").Scan(&status).Error
	if err != nil {
		return false, err
	}
	return status, nil
}

// UpdateStatus 批量更新管理员的状态字段
func UpdateStatus() error {
	ok := true
	err := DB.Model(&gorm_model.UserCasbinRules{}).Where("status = ? ", ok).Updates(map[string]interface{}{"status": !ok}).Error
	if err != nil {
		return err
	}
	return nil
}

// UpdateOne 更新一个管理员的字段
func UpdateOne(username string) error {
	var user gorm_model.UserCasbinRules
	err := DB.Where("c_username = ?", username).First(&user).Error
	if err != nil {
		return err
	}
	user.Status = true
	err = DB.Save(&user).Error
	if err != nil {
		return err
	}
	return nil
}

// SelNotClass 查询被推选为年级院级成长之星的学号合集
func SelNotClass() ([]string, error) {
	var username []string
	err := DB.Model(&gorm_model.Star{}).Where("type <> ? ", 1).Select("username").Scan(&username).Error
	if err != nil {
		return nil, err
	}
	return username, nil
}
