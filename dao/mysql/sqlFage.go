package mysql

import (
	"studentGrow/models"
	"time"
)

// SelGradeId 获取大一or大二or大三or大四的用户id 和usernameSlice
func SelGradeId(data time.Time, year int) ([]int, []string, error) {
	var uidslice []int
	var usernameslice []string
	//计算时间的间隔的右端
	// 计算时间间隔的左端
	CurrentYear := data.AddDate(year+1, 0, 0)
	YearAgo := data.AddDate(year, 0, 0)
	err := DB.Model("&User{}").Select("id").Where("plus_time >= ?", YearAgo).Where("plus_time <= ?", CurrentYear).Scan(&uidslice).Error
	err = DB.Model("&User{}").Select("username").Where("plus_time >= ?", YearAgo).Where("plus_time <= ?", CurrentYear).Scan(&usernameslice).Error
	if err != nil {
		return nil, nil, err
	}
	return uidslice, usernameslice, nil
}

// SelUid 根据班级查询班成员的id
func SelUid(class string) ([]int, error) {
	//班级成员的id切片
	var uidSlice []int
	err := DB.Table("users").Select("id").Where("class = ?", class).Scan(&uidSlice).Error
	// 检查并返回错误
	if err != nil {
		return nil, err
	}
	return uidSlice, nil
}

// SelCollegeId 查询所有人的uid和username
func SelCollegeId() ([]int, []string, error) {
	var uidslice []int
	var usernameslice []string
	err := DB.Table("users").Select("id").Scan(&uidslice).Error
	err = DB.Table("users").Select("username").Scan(&usernameslice).Error
	if err != nil {
		return nil, nil, err
	}
	return uidslice, usernameslice, nil
}

// SelArticleNum 根据id查询每个用户的帖子数
func SelArticleNum(id int) (int64, error) {
	var number int64
	err := DB.Table("articles").Where("user_id = ?", id).Count(&number).Error
	// 检查并返回错误
	if err != nil {
		return 0, err
	}
	return number, nil
}

// SelArticle 根据id查询每个用户目标天数的贴子数
func SelArticle(id int, date time.Time) (int64, error) {
	var number int64

	// 获取当天的开始时间（00:00:00）
	from := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC)

	// 获取第二天的开始时间（00:00:00），用于查询截止到当天结束的时间范围
	to := from.Add(24 * time.Hour)

	// 使用 BETWEEN 查询当天的记录数
	err := DB.Table("articles").Where("user_id = ? ", id).
		Where("created_at BETWEEN ? AND ?", from, to).
		Count(&number).Error
	// 检查并返回错误
	if err != nil {
		return 0, err
	}
	return number, nil
}

// SelUsername 根据班级查询班成员的username
func SelUsername(class string) ([]string, error) {
	//班级成员的id切片
	var usernameSlice []string
	err := DB.Table("users").Select("username").Where("class = ?", class).Scan(&usernameSlice).Error
	// 检查并返回错误
	if err != nil {
		return nil, err
	}
	return usernameSlice, nil
}

// SelStudent 查询所有学生人数
func SelStudent() (int64, error) {
	var number int64
	err := DB.Table("users").Where("identity = ?", "student").Count(&number).Error
	// 检查并返回错误
	if err != nil {
		return 0, err
	}
	return number, nil
}

// SelTeacher 查询所有教师人数
func SelTeacher() (int64, error) {
	var number int64
	err := DB.Table("users").Where("identity = ?", "teacher").Count(&number).Error
	// 检查并返回错误
	if err != nil {
		return 0, err
	}
	return number, nil
}

// SelArticleLike 查询帖子总赞数
func SelArticleLike(id int) (int64, error) {
	var number int64
	err := DB.Table("articles").Select("like_amount").Where("user_id = ?", id).Scan(&number).Error
	// 检查并返回错误
	if err != nil {
		return 0, err
	}
	return number, nil
}

// SelArticleRead 查询帖子总赞数
func SelArticleRead(id int) (int64, error) {
	var number int64
	err := DB.Table("articles").Select("read_amount").Where("user_id = ?", id).Scan(&number).Error
	// 检查并返回错误
	if err != nil {
		return 0, err
	}
	return number, nil
}

// 查询不同tag下的文章的大小

func SelTagArticle() ([]models.TagAmount, error) {
	var tagcount []models.TagAmount
	err := DB.Table("article_tags").Select("tag_name,COUNT(*)as count").Group("tag_name").Scan(&tagcount).Error
	if err != nil {
		return nil, err
	}
	return tagcount, nil
}

// 查询不同tag不同时间下的文章的大小
func SelTagArticleTime(date string) ([]models.TagAmount, error) {
	nowdate, err := time.Parse("2006-01-02", date)
	var tagcount []models.TagAmount
	err = DB.Table("article_tags").Where("created_at = ?", nowdate).Select("tag_name,COUNT(*)as count").Group("tag_name").Scan(&tagcount).Error
	if err != nil {
		return nil, err
	}
	return tagcount, nil
}
