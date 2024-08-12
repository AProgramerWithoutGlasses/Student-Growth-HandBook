package service

import (
	"fmt"
	"sort"
	"strconv"
	"studentGrow/dao/mysql"
	"studentGrow/dao/redis"
	"time"
)

// ArticleData 帖子总数
func ArticleData(uidslice []int) (int64, error) {
	var allNumber int64
	//把每个用户发布的帖子个数都查询并把数据相加
	for _, id := range uidslice {
		number, err := mysql.SelArticleNum(id)
		if err != nil {
			fmt.Errorf("ArticleData mysql.SelArticleNum err", err)
		}
		allNumber += number
	}
	return allNumber, nil
}

// NarticleDataClass 今日新帖数
func NarticleDataClass(uidslice []int) (int64, error) {
	var allNumber int64
	data := time.Now()
	//2.把每个用户发布的帖子个数都查询并把数据相加
	for _, id := range uidslice {
		number, err := mysql.SelArticle(id, data)
		if err != nil {
			fmt.Errorf("ArticleData mysql.SelArticleNum err", err)
		}
		allNumber += number
	}
	return allNumber, nil
}

// 这个方法用来获取当天日期的前一天
func getYesterdayDate(date time.Time) time.Time {
	// 获取当前日期的年月日
	year, month, day := date.Date()

	// 如果是 1 号,则需要获取上个月的最后一天
	if day == 1 {
		// 减去 1 个月,然后获取那个月的最后一天
		prevMonth := month - 1
		prevYear := year
		if prevMonth == 0 {
			prevMonth = 12
			prevYear--
		}
		return time.Date(prevYear, prevMonth, 1, 0, 0, 0, 0, time.UTC).AddDate(0, 0, -1)
	} else {
		// 否则直接减去 1 天即可
		return time.Date(year, month, day, 0, 0, 0, 0, time.UTC).AddDate(0, 0, -1)
	}
}

// ArticleDataClassRate 昨日新帖数跟今日的相比
func ArticleDataClassRate(uidslice []int, nownumber int64) (float64, error) {
	data := time.Now()
	ydata := getYesterdayDate(data)
	var allNumber int64
	//遍历获取班级成员在指定日期的帖子数
	for _, id := range uidslice {
		number, err := mysql.SelArticle(id, ydata)
		if err != nil {
			fmt.Println("ArticleDataClassRate mysql.SelArticleNum err", err)
			return 0, err
		}
		allNumber += number
	}
	if nownumber == 0 || allNumber == 0 {
		return 0, nil
	}
	//计算比率
	rate, err := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(nownumber-allNumber)/float64(nownumber)), 64)
	return rate * 100, err
}

// TodayVictor 查询今日访客数
func TodayVictor(usernameslice []string) int {
	var allNumber int
	data := time.Now().Format("20060102")
	for _, username := range usernameslice {
		if redis.IfVictor(username, data) {
			allNumber += 1
		}
	}
	return allNumber
}

// VictorRate 今天和昨天的比率
func VictorRate(usernameslice []string, todayVictor int) (float64, error) {
	var yesdayVictor int
	ydata := getYesterdayDate(time.Now()).Format("20060102")
	for _, username := range usernameslice {
		if redis.IfVictor(username, ydata) {
			yesdayVictor += 1
		}
	}
	rate, err := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(todayVictor-yesdayVictor)/float64(todayVictor)), 64)
	return rate, err
}

// LikeAmount 获取帖子总赞数
func LikeAmount(uidslice []int) int64 {
	var allNumber int64
	for _, id := range uidslice {
		number, err := mysql.SelArticleLike(id)
		if err != nil {
			fmt.Println("LikeAmount SelArticleLike err", err)
		}
		allNumber += number
	}
	return allNumber
}

// ReadAmount 获取帖子总阅读数
func ReadAmount(uidslice []int) int64 {
	var allNumber int64
	for _, id := range uidslice {
		number, err := mysql.SelArticleRead(id)
		if err != nil {
			fmt.Println("LikeAmount SelArticleLike err", err)
		}
		allNumber += number
	}
	return allNumber
}

// 柱状图
func PillarData() ([]string, []int, error) {
	// 假设 mysql.SelTagArticle() 返回的是 []models.TagAmount 类型
	tagcount, err := mysql.SelTagArticle()
	if err != nil {
		fmt.Println("PillarData SelTagArticle err", err)
		return nil, nil, err
	}

	// 根据人数排序
	sort.Slice(tagcount, func(i, j int) bool {
		return tagcount[i].Count > tagcount[j].Count
	})

	// 只取前7个或更少的元素
	tagName := make([]string, 0, 7)
	count := make([]int, 0, 7)

	for i := 0; i < 7 && i < len(tagcount); i++ {
		tagName = append(tagName, tagcount[i].Tag)
		count = append(count, tagcount[i].Count)
	}

	return tagName, count, nil
}

// 特定日期的柱状图
func PillarDataTime(date string) ([]string, []int, error) {
	tagName := make([]string, 7)
	count := make([]int, 7)
	tagcount, err := mysql.SelTagArticleTime(date)
	if err != nil {
		fmt.Println("PillarData SelTagArticle err", err)
		return nil, nil, err
	}
	// 根据人数排序
	sort.Slice(tagcount, func(i, j int) bool {
		return tagcount[i].Count > tagcount[j].Count
	})
	for i := 0; i < len(tagcount) && i < 7; i++ {
		for _, tagmodel := range tagcount {
			tagName[i] = tagmodel.Tag
			count[i] = tagmodel.Count
		}
	}
	return tagName, count, nil
}
