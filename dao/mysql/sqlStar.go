package mysql

import (
	"errors"
	"gorm.io/gorm"
	"studentGrow/models/gorm_model"
	"time"
)

// SelName 根据学号查名字，等一系列数据
func SelName(username string) (string, error) {
	var name string
	err := DB.Table("users").Select("name").Where("username = ?", username).Scan(&name).Error
	if err != nil {
		return "", err
	}
	return name, nil
}

// SelId 根据账号查找id
func SelId(username string) (int, error) {
	var id int
	err := DB.Table("users").Select("id").Where("username = ?", username).Scan(&id).Error
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
	err := DB.Table("users").Select("point").Where("username = ?", username).Error
	if err != nil {
		return 0, err
	}
	return score, nil
}

// Frequency 被推举次数
func Frequency(username string) (int64, error) {
	var number int64
	err := DB.Table("stars").Where("username = ?", username).Count(&number).Error
	if err != nil {
		return 0, err
	}
	return number, nil
}

// SelHot 查询热度
func SelHot(id int) (int, error) {
	var like int
	var collect int
	err := DB.Table("articles").Select("like_amount").Where("user_id =?", id).Scan(&like).Error
	err = DB.Table("articles").Select("collect_amount").Where("user_id = ?", id).Scan(&collect).Error
	if err != nil {
		return 0, err
	}
	return collect + like, nil
}

// SelStatus 查询状态
func SelStatus(username string) (bool, error) {
	var star []gorm_model.Star
	err := DB.Where("username = ?", username).Where("session = ?", 0).Find(&star).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		} else {
			return false, err
		}
	}
	if len(star) == 0 {
		return false, nil
	}
	return true, nil
}

// SelStarUser SelGrade 查询未公布的学号合集
func SelStarUser() ([]string, error) {
	var alluser []string
	err := DB.Table("stars").Where("session = ?", 0).Select("username").Scan(&alluser).Error
	if err != nil {
		return nil, err
	}
	return alluser, nil
}

// SelPlus 查询入学时间
func SelPlus(username string) (time.Time, error) {
	var plus time.Time
	err := DB.Table("users").Select("plus_time").Where("username = ?", username).Scan(&plus).Error
	if err != nil {
		return time.Time{}, err
	}
	return plus, nil
}

// SelStarColl 院级查询表里学号合集
func SelStarColl() ([]string, error) {
	var alluser []string
	err := DB.Table("stars").Where("type = ?", 2).Where("session = ?", 0).Select("username").Scan(&alluser).Error
	if err != nil {
		return nil, err
	}
	return alluser, nil
}

// SelSearchUser 根据名字班级查找学号--班级管理员搜索
func SelSearchUser(name string, class string) ([]string, error) {
	var username []string
	err := DB.Table("users").Where("class = ?", class).Where("name = ?", name).Select("username").Scan(&username).Error
	if err != nil {
		return nil, err
	}
	return username, nil
}

// SelSearchColl 院级管理员搜索
func SelSearchColl(name string) ([]string, error) {
	var usernamesli []string
	err := DB.Table("stars").Where("name = ?", name).Where("type = ?", 2).Where("session = ?", 0).Select("username").Scan(&usernamesli).Error
	if err != nil {
		return nil, err
	}
	return usernamesli, nil
}
