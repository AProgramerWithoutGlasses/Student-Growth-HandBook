package mysql

import (
	"fmt"
	"gorm.io/gorm"
	_ "gorm.io/gorm"
	"sort"
	model "studentGrow/models/gorm_model"
	myErr "studentGrow/pkg/error"
	"studentGrow/utils/timeConverter"
	"time"
)

// SelectUserById 查询数据库是否存在该用户
func SelectUserById(uid int) (err error, user *model.User) {
	//select * from users where id = uid
	// 查询用户
	if err := DB.Where("id = ?", uid).First(&user).Error; err != nil {
		return err, nil
	} else {
		return nil, user
	}
}

// SelectUserByUsername 通过username查找uid
func SelectUserByUsername(username string) (uid int, err error) {
	//select id from users where username = username
	var user model.User
	if err := DB.Model(model.User{}).Select("id").Where("username = ?", username).First(&user).Error; err != nil {
		return int(user.ID), err
	} else {
		return int(user.ID), nil
	}
}

// SelectArticleById 通过id查找文章
func SelectArticleById(aid int) (err error, article *model.Article) {
	//查询用户 select * from articles where id = aid
	// First自动检查记录是否存在
	if err := DB.Where("id = ?", aid).First(&article).Error; err != nil {
		// 处理查询错误
		fmt.Println("Error:", err)
		return err, nil
	} else {
		return nil, article
	}
}

// InsertIntoCommentsForArticle 向数据库插入评论数据(回复文章)

func InsertIntoCommentsForArticle(content string, aid int, uid int) (err error) {
	//content;id;username
	//comment := model.Comment{
	//	Model:        gorm.Model{},
	//	Content:      content,
	//	UpvoteAmount: 0,
	//	IsRead:       false,
	//	Del:          false,
	//	Uid:          uid,
	//	Pid:          0,
	//	Aid:          aid,
	//	Upvote:       nil,
	//}
	//DB.Create(&comment)

	return nil
}

// InsertIntoCommentsForComment 向数据库插入评论数据(回复评论)
func InsertIntoCommentsForComment(content string, uid int, pid int) (err error) {
	//content;id;username
	//comment := model.Comment{
	//	Model:        gorm.Model{},
	//	Content:      content,
	//	UpvoteAmount: 0,
	//	IsRead:       false,
	//	Del:          false,
	//	Uid:          uid,
	//	Pid:          pid,
	//	Aid:          0,
	//	Upvote:       nil,
	//}
	//
	//DB.Create(&comment)
	return nil
}

// SelectArticleAndUserListByPage 后台分页查询文章及用户列表并模糊查询
func SelectArticleAndUserListByPage(page, limit int, sort, order, startAt, endAt, topic, keyWords, name string, isBan bool) (result []model.Article, err error) {
	//SELECT articles.*, users.*
	//FROM articles
	//JOIN users ON articles.user_id = users.id
	//WHERE articles.created_at > (
	//    SELECT created_at
	//    FROM articles
	//    ORDER BY created_at DESC
	//    LIMIT ?, 1
	//)
	//LIMIT ?;
	var articles []model.Article
	var query *gorm.DB

	// 时间区间为空检查
	if startAt != "" && endAt != "" {
		query = DB.Where(fmt.Sprintf("articles.%s between ? and ? and topic like ? and content like ? and articles.ban = ?", sort),
			startAt, endAt, fmt.Sprintf("%%%s%%", topic), fmt.Sprintf("%%%s%%", keyWords), isBan)
	} else if startAt == "" && endAt != "" {
		query = DB.Where(fmt.Sprintf("articles.%s < ? and topic like ? and content like ? and articles.ban = ?", sort),
			endAt, fmt.Sprintf("%%%s%%", topic), fmt.Sprintf("%%%s%%", keyWords), isBan)
	} else if startAt != "" && endAt == "" {
		query = DB.Where(fmt.Sprintf("articles.%s > ? and topic like ? and content like ? and articles.ban = ?", sort),
			startAt, fmt.Sprintf("%%%s%%", topic), fmt.Sprintf("%%%s%%", keyWords), isBan)
	} else if startAt == "" && endAt == "" {
		query = DB.Where("topic like ? and content like ? and articles.ban = ?",
			fmt.Sprintf("%%%s%%", topic), fmt.Sprintf("%%%s%%", keyWords), isBan)
	}

	if err := query.InnerJoins("User").Where("name like ?", fmt.Sprintf("%%%s%%", name)).
		Order(fmt.Sprintf("%s %s", sort, order)).Limit(limit).Offset((page - 1) * limit).Find(&articles).Error; err != nil {
		fmt.Println()
		return nil, err
	}

	// 检查是否存在用户列表记录
	if len(articles) <= 0 {
		return nil, myErr.NotFoundError()
	}

	return articles, nil
}

// SelectArticleAndUserListByPageFirstPage 前台模糊查询文章列表
func SelectArticleAndUserListByPageFirstPage(keyWords, topic, articleSort string, limit, page int) (result []model.Article, err error) {
	var articles model.Articles
	if err = DB.Preload("User").Preload("ArticleTags").
		Where("topic = ? and key_words like ?", topic, fmt.Sprintf("%%%s%%", keyWords)).
		Order("created_at desc").
		Limit(limit).
		Offset((page - 1) * limit).Find(&articles).Error; err != nil {
		fmt.Println("SelectArticleAndUserListByPageFirstPage() dao.mysql.sql_nzx")
		return nil, err
	}

	if articleSort == "hot" {
		sort.Sort(articles)
	}

	return articles, nil
}

// BannedArticleByIdForClass 通过文章id对文章进行封禁或解封 - 班级
func BannedArticleByIdForClass(articleId int, isBan bool, username string) error {
	// 查询班级管理员信息
	user := model.User{}
	if err := DB.Where("username = ?", username).First(&user).Error; err != nil {
		fmt.Println("BannedArticleByIdForClass() dao.mysql.sql_nzx")
		return err
	}

	// 查询待封禁的文章;若查询不到，则返回
	article := model.Article{}
	if err := DB.Preload("User", "class = ?", user.Class).Where("id = ?", articleId).First(&article).Error; err != nil {
		fmt.Println("BannedArticleByIdForClass() dao.mysql.sql_nzx")
		return myErr.OverstepCompetence()
	}

	// 修改文章状态
	if err := DB.Model(&model.Article{}).Where("id = ?", articleId).Updates(model.Article{Ban: true}).Error; err != nil {
		fmt.Println("BannedArticleByIdForClass() dao.mysql.sql_nzx")
		return err
	}

	return nil
}

// BannedArticleByIdForGrade 通过文章id对文章进行封禁或解封 - 年级
func BannedArticleByIdForGrade(articleId int, grade int) error {
	// GetUnreadReportsForGrade
	year, err := timeConverter.GetEnrollmentYear(grade)
	if err != nil {
		fmt.Println("BannedArticleByIdForGrade() dao.mysql.sql_nzx")
		return err
	}

	// 获取需要被封禁的文章；若找不到则返回
	article := model.Article{}
	if err = DB.Preload("User", "plus_time between ? and ?",
		fmt.Sprintf("%s-01-01", year.Year()), fmt.Sprintf("%s-12-31", year.Year())).
		Where("id = ?", articleId).First(&article).Error; err != nil {
		fmt.Println("BannedArticleByIdForGrade() dao.mysql.sql_nzx")
		return myErr.OverstepCompetence()
	}

	// 修改文章状态
	if err := DB.Model(&model.Article{}).Where("id = ?", articleId).Updates(model.Article{Ban: true}).Error; err != nil {
		fmt.Println("BannedArticleByIdForGrade() dao.mysql.sql_nzx")
		return err
	}
	return nil
}

// BannedArticleByIdForSuperman 通过文章id对文章进行封禁或解封 - 院级(超级)
func BannedArticleByIdForSuperman(articleId int) error {
	// 修改文章状态
	if err := DB.Model(&model.Article{}).Where("id = ?", articleId).Updates(model.Article{Ban: true}).Error; err != nil {
		fmt.Println("BannedArticleByIdForSuperman() dao.mysql.sql_nzx")
		return err
	}
	return nil
}

// DeleteArticleByIdForClass 通过文章id删除文章 - 班级
func DeleteArticleByIdForClass(articleId int, username string) error {
	// 查询班级管理员信息
	user := model.User{}
	if err := DB.Where("username = ?", username).First(&user).Error; err != nil {
		fmt.Println("DeleteArticleByIdForClass() dao.mysql.sql_nzx")
		return err
	}

	// 查询待删除的文章
	article := model.Article{}
	if err := DB.Preload("User", "class = ?", user.Class).Where("id = ?", articleId).First(&article).Error; err != nil {
		fmt.Println("DeleteArticleByIdForClass() dao.mysql.sql_nzx")
		return err
	}

	if err := DB.Delete(&model.Article{}, article.ID).Error; err != nil {
		fmt.Println("DeleteArticleByIdForClass() dao.mysql.sql_nzx")
		return err
	}

	return nil
}

// DeleteArticleByIdForGrade 通过文章id删除文章 - 年级
func DeleteArticleByIdForGrade(articleId int, grade int) error {
	// 将年级转化为入学年份
	year, err := timeConverter.GetEnrollmentYear(grade)
	if err != nil {
		fmt.Println("DeleteArticleByIdForGrade() dao.mysql.sql_nzx")
		return err
	}

	// 获取需要被删除的文章
	article := model.Article{}
	if err = DB.Preload("User", "plus_time between ? and ?",
		fmt.Sprintf("%d-01-01", year.Year()), fmt.Sprintf("%d-12-31", year.Year())).
		Where("id = ?", articleId).First(&article).Error; err != nil {
		fmt.Println("DeleteArticleByIdForGrade() dao.mysql.sql_nzx")
		return err
	}

	// 删除文章
	if err = DB.Delete(&model.Article{}, article.ID).Error; err != nil {
		fmt.Println("DeleteArticleByIdForGrade() dao.mysql.sql_nzx")
		return err
	}
	return nil
}

// DeleteArticleByIdForSuperman 通过id删除文章 - 院级(超级)
func DeleteArticleByIdForSuperman(articleId int) error {
	article := model.Article{}
	if err := DB.Where("id = ?", articleId).First(&article).Error; err != nil {
		fmt.Println("DeleteArticleByIdForSuperman() dao.mysql.sql_nzx")
		return err
	}
	if err := DB.Delete(&model.Article{}, article.ID).Error; err != nil {
		fmt.Println("DeleteArticleByIdForSuperman() dao.mysql.sql_nzx")
		return err
	}

	return nil
}

// ReportArticleById 举报文章
func ReportArticleById(aid int, uid int, msg string) error {
	//由于举报逻辑需要先自增文章的举报字段，然后添加举报信息到记录表。
	//需要开启事务，若出现错误，则回滚
	bg := DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("ReportArticleById() panic rollback()", r)
			bg.Rollback()
		}
	}()

	// 获取被举报文章举报量，并对举报量+1操作
	article := model.Article{}
	if err := DB.Where("id = ?", uint(aid)).First(&article).Error; err != nil {
		fmt.Println("ReportArticleById() dao.mysql.sql_nzx")
		return err
	}
	article.ReportAmount += 1
	result := DB.Model(model.Article{}).Select("report_amount").Where("id = ?", aid).Save(&article)

	if result.Error != nil {
		bg.Rollback()
		panic("ss")
		return result.Error
	}
	// 查询更新结果
	if result.RowsAffected <= 0 {
		return myErr.NotFoundError()
	}

	// 检查举报记录：不允许重复举报
	var report []model.UserReportArticleRecord
	if err := DB.Where("user_id = ? and article_id = ?", uid, aid).Find(&report).Error; err != nil {
		fmt.Println("ReportArticleById() dao.mysql.sql_nzx")
		bg.Rollback()
		return err
	}

	//如果数据库有重复记录，则拒绝重复提交
	if len(report) > 0 {
		fmt.Println("ReportArticleById() dao.mysql.sql_nzx")
		bg.Rollback()
		return myErr.RejectRepeatSubmission()
	}

	// 写入举报记录
	reportRecord := model.UserReportArticleRecord{
		UserID:    uint(uid),
		ArticleID: uint(aid),
		Msg:       msg,
	}

	if err := DB.Create(&reportRecord).Error; err != nil {
		fmt.Println("ReportArticleById() dao.mysql.sql_nzx")
		bg.Rollback()
		return err
	}

	// 提交
	if err := bg.Commit().Error; err != nil {
		fmt.Println("ReportArticleById() dao.mysql.sql_nzx")
		bg.Rollback()
		return err
	}
	return nil
}

// SearchHotArticlesOfDay 查找今日热门文章
func SearchHotArticlesOfDay(startOfDay time.Time, endOfDay time.Time) (model.Articles, error) {
	var articles model.Articles
	if err := DB.Where("created_at >= ? and created_at < ?", startOfDay, endOfDay).
		Find(&articles).Error; err != nil {
		fmt.Println("SearchHotArticlesOfDay() dao.mysql.sql_nzx")
		return nil, err
	}

	if len(articles) <= 0 {
		fmt.Println("SearchHotArticlesOfDay() dao.mysql.sql_nzx")
		return nil, myErr.NotFoundError()
	}
	return articles, nil
}

// InsertIntoArticle 插入文章信息
//func InsertIntoArticle(username, content, topic string, tags []string, file[]) {
//
//}
