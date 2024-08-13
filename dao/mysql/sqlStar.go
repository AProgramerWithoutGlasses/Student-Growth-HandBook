package mysql

import (
	"errors"
	"gorm.io/gorm"
	"studentGrow/models/gorm_model"
)

// 根据id 查询文章id 切片

// SelGuid 根据学号查名字，等一系列数据
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
	err := DB.Model("&User{}").Select("id").Where("username = ?", username).Scan(&id).Error
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
	err := DB.Model("&User{}").Select("point").Where("username = ?", username).Error
	if err != nil {
		return 0, err
	}
	return score, nil
}

// Frequency 被推举次数
func Frequency(username string) (int64, error) {
	var number int64
	err := DB.Model("&Star{}").Where("username = ?", username).Count(&number).Error
	if err != nil {
		return 0, err
	}
	return number, nil
}

// SelHot 查询热度
func SelHot(id int) (int, error) {
	var like int
	var collect int
	err := DB.Model("&Article{}").Select("like_amount").Where("user_id =?", id).Scan(&like).Error
	err = DB.Model("&Article{}").Select("collect_amount").Where("user_id = ?", id).Scan(&collect).Error
	if err != nil {
		return 0, err
	}
	return collect + like, nil
}

// 查询状态
func SelStatus(username string) (bool, error) {
	var star gorm_model.Star
	result := DB.Find(&star).Where("username = ?", username).Where("session = ?", 0)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false, nil
	} else if result.Error == nil {
		return true, nil
	} else {
		return false, result.Error
	}
}
