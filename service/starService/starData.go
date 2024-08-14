package starService

import (
	"fmt"
	"studentGrow/dao/mysql"
	"studentGrow/models"
	"studentGrow/utils/timeConverter"
)

// StarGridClass 查询表格所有数据
func StarGrid(usernameslice []string) ([]models.StarBack, error) {
	var stakback []models.StarBack
	for _, username := range usernameslice {
		//结构体对象存放数据
		//var star models.StarBack
		//查询 name ，id，userfans，score，hot
		//查询名字
		name, err := mysql.SelName(username)
		if err != nil {
			fmt.Println("StarGridClass err", err)
			return nil, err
		}

		//查询 id
		id, err := mysql.SelId(username)
		if err != nil {
			fmt.Println("StarGridClass err", err)
			return nil, err
		}

		//查询粉丝数
		userfans, err := mysql.Selfans(id)
		if err != nil {
			fmt.Println("StarGridClass Selfans err", err)
			return nil, err
		}

		//查询积分
		score, err := mysql.Score(username)
		if err != nil {
			fmt.Println("StarGridClass Score err", err)
			return nil, err
		}

		//查询被推举次数
		frequency, err := mysql.Frequency(username)
		if err != nil {
			fmt.Println("StarGridClass Frequency err", err)
			return nil, err
		}

		//查询文章数
		article, err := mysql.SelArticleNum(id)
		if err != nil {
			fmt.Println("StarGridClass SelArticleNum err", err)
			return nil, err
		}

		//查询热度
		hot, err := mysql.SelHot(id)
		if err != nil {
			fmt.Println("StarGridClass SelHot err", err)
			return nil, err
		}

		//查询状态
		status, err := mysql.SelStatus(username)

		star := models.StarBack{
			Username:           username,
			Frequency:          frequency,
			Name:               name,
			User_article_total: article,
			Userfans:           userfans,
			Score:              score,
			Hot:                hot,
			Status:             status,
		}
		stakback = append(stakback, star)
	}
	return stakback, nil
}

// PageQuery 实现分页查询
func PageQuery(starback []models.StarBack, page, limit int) []models.StarBack {
	length := len(starback)
	left := (page - 1) * limit
	right := page * limit
	if left >= length {
		return nil
	}
	if right > length+1 {
		return starback[left:]
	}
	return starback[left:right]
}

// StarGrade 把表中数据以年级分开
func StarGuidGrade(usernamesli []string, year int) ([]string, error) {
	var gradeusersli []string
	for _, username := range usernamesli {
		plus, err := mysql.SelPlus(username)
		if err != nil {
			fmt.Println("StarGrade SelPlus err", err)
			return nil, err
		}
		grade := timeConverter.GetUserGrade(plus)
		if grade == year {
			gradeusersli = append(gradeusersli, username)
		}
	}
	return gradeusersli, nil
}
