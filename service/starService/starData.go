package starService

import (
	"fmt"
	"studentGrow/dao/mysql"
	"studentGrow/models"
	"studentGrow/utils/timeConverter"
)

// StarGrid 查询表格所有数据
func StarGrid(usernameslice []string) ([]models.StarBack, error) {
	var starBack []models.StarBack
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
		userFans, err := mysql.Selfans(id)
		if err != nil {
			fmt.Println("StarGridClass Selfans err", err)
			return nil, err
		}
		var score int
		//查询积分
		allScore, err := mysql.Score(id)
		for _, thisScore := range allScore {
			score += thisScore
		}
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

		var hot int
		//查询热度
		likes, collects, err := mysql.SelHot(id)
		for _, like := range likes {
			hot += like
		}
		for _, collect := range collects {
			hot += collect
		}
		if err != nil {
			fmt.Println("StarGridClass SelHot err", err)
			return nil, err
		}

		star := models.StarBack{
			Username:           username,
			Frequency:          frequency,
			Name:               name,
			User_article_total: article,
			Userfans:           userFans,
			Score:              score,
			Hot:                hot,
			Status:             false,
		}
		starBack = append(starBack, star)
	}
	return starBack, nil
}

// PageQuery 实现分页查询
func PageQuery(starback []models.StarBack, page, limit int) []models.StarBack {
	if limit <= 0 {
		// 返回空切片，limit 应该大于0
		return []models.StarBack{}
	}
	length := len(starback)
	left := (page - 1) * limit
	right := left + limit

	// 修正right的计算，确保不会超出slice的长度
	if right > length {
		right = length
	}

	// 如果left超出length，返回空切片
	if left >= length {
		return []models.StarBack{}
	}

	return starback[left:right]
}

// StarGuidGrade StarGrade 把表中数据以年级分开
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

// SearchGrade 年级管理员搜索
func SearchGrade(name string, year int) ([]string, error) {
	//查询表格未公布与名字相匹配的所有数据
	usernamesli, err := mysql.SelSearchGrade(name)
	if err != nil {
		return nil, err
	}
	//根据管理员权限查找数据
	gUsername, err := StarGuidGrade(usernamesli, year)
	return gUsername, nil
}

// StarClass 返回成长之星的班级之星
func StarClass(session int) ([]models.StarClass, error) {
	var StarClasssli []models.StarClass
	//查询出所有班级之星
	usernamesli, err := mysql.SelStar(session, 1)
	if err != nil {
		return nil, err
	}
	//通过班级进行分组
	starMap, err := GroupByClass(usernamesli)
	if err != nil {
		return nil, err
	}
	//对结构体赋值
	for class, name := range starMap {
		starclass := models.StarClass{
			ClassName: class,
			ClassStar: name,
		}
		StarClasssli = append(StarClasssli, starclass)
	}
	return StarClasssli, nil
}

// GroupByClass 通过班级进行分组
func GroupByClass(usernamesli []string) (map[string][]string, error) {
	starmap := make(map[string][]string)
	for _, username := range usernamesli {
		class, err := mysql.SelClass(username)
		name, err := mysql.SelName(username)
		if err != nil {
			return nil, err
		}
		if _, exists := starmap[class]; !exists {
			starmap[class] = []string{name}
		} else {
			starmap[class] = append(starmap[class], name)
		}
	}
	return starmap, nil
}

// StarGrade 返回年级之星
func StarGrade(session int) ([]models.StarGrade, error) {
	var starGrade []models.StarGrade
	usernamesli, err := mysql.SelStar(session, 2)
	if err != nil {
		return nil, err
	}
	for _, username := range usernamesli {
		name, _ := mysql.SelName(username)
		class, _ := mysql.SelClass(username)
		//赋值结构体
		stargrade := models.StarGrade{
			GradeName:  name,
			GradeClass: class,
		}
		//加入切片
		starGrade = append(starGrade, stargrade)
	}
	return starGrade, nil
}

// StarCollege 返回院级之星
func StarCollege(session int) ([]models.StarGrade, error) {
	var starGrade []models.StarGrade
	usernamesli, err := mysql.SelStar(session, 3)
	if err != nil {
		return nil, err
	}
	for _, username := range usernamesli {
		name, _ := mysql.SelName(username)
		class, _ := mysql.SelClass(username)
		//赋值结构体
		stargrade := models.StarGrade{
			GradeName:  name,
			GradeClass: class,
		}
		//加入切片
		starGrade = append(starGrade, stargrade)
	}
	return starGrade, nil
}

// QStarClass 返回前台成长之星
func QStarClass(starType int) ([]models.StarStu, error) {
	var starlist []models.StarStu
	session, err := mysql.SelMax()
	if err != nil {
		return nil, err
	}
	usernameslic, err := mysql.SelStar(session, starType)
	if err != nil {
		return nil, err
	}
	for _, username := range usernameslic {
		name, err := mysql.SelName(username)
		headshot, err := mysql.SelHead(username)
		starstu := models.StarStu{
			Username:     username,
			Name:         name,
			UserHeadshot: headshot,
		}
		if err != nil {
			return nil, err
		}
		starlist = append(starlist, starstu)
	}
	return starlist, nil
}

// SelNumClass 查询表中跟管理员一个班的有多少人
func SelNumClass(class string) (int, error) {
	var classnum int
	//1.查询表中目前的人员
	usernamesli, err := mysql.SelStarUser()
	if err != nil {
		return 0, err
	}
	//2.匹配
	for _, username := range usernamesli {
		thisclass, err := mysql.SelClass(username)
		if err != nil {
			return 0, err
		}
		if thisclass == class {
			classnum += 1
		}
	}
	return classnum, nil
}

// SelNumGrade 查询年级管理员已推选的人数
func SelNumGrade(year int) (int, error) {
	var classNum int
	usernamesli, err := mysql.SelNotClass()
	if err != nil {
		return 0, err
	}
	for _, username := range usernamesli {
		plus, err := mysql.SelPlus(username)
		if err != nil {
			return 0, err
		}
		grade := timeConverter.GetUserGrade(plus)
		if grade == year {
			classNum += 1
		}
	}
	return classNum, nil
}
