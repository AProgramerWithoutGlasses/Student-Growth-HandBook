package mysql

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm"
	_ "gorm.io/gorm"
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
	if err := DB.Preload("ArticlePics").Preload("ArticleTags").Preload("User").
		Where("id = ?", aid).First(&article).Error; err != nil {
		return err, nil
	} else {
		return nil, article
	}
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
func SelectArticleAndUserListByPageFirstPage(keyWords, topic string, limit, page int) (result model.Articles, err error) {
	var articles model.Articles
	if err = DB.Preload("User").Preload("ArticleTags").Preload("ArticlePic").
		Where("topic = ? and content like ?", topic, fmt.Sprintf("%%%s%%", keyWords)).
		Order("created_at desc").
		Limit(limit).
		Offset((page - 1) * limit).Find(&articles).Error; err != nil {
		fmt.Println("SelectArticleAndUserListByPageFirstPage() dao.mysql.sql_nzx")
		return nil, err
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

// UpdateArticleCommentNum 设置文章评论数
func UpdateArticleCommentNum(aid, num int) error {
	if err := DB.Where("id = ?", aid).Update("comment_amount", num).Error; err != nil {
		fmt.Println("UpdateArticleCommentNum() dao.mysql.sql_nzx")
		return err
	}
	return nil
}

// QueryArticleCommentNum 获取文章评论数
func QueryArticleCommentNum(aid int) (int, error) {
	article := model.Article{}
	if err := DB.Where("id = ?", aid).First(&article).Error; err != nil {
		fmt.Println("QueryArticleCommentNum() dao.mysql.sql_nzx")
		return -1, err
	}
	return article.CommentAmount, nil
}

// UpdateArticleLikeNum 设置文章点赞数
func UpdateArticleLikeNum(aid, num int) error {
	if err := DB.Where("id = ?", aid).Update("like_amount", num).Error; err != nil {
		fmt.Println("UpdateArticleLikeNum() dao.mysql.sql_nzx")
		return err
	}
	return nil
}

// QueryArticleLikeNum 获取文章点赞数
func QueryArticleLikeNum(aid int) (int, error) {
	article := model.Article{}
	if err := DB.Where("id = ?", aid).First(&article).Error; err != nil {
		fmt.Println("QueryArticleLikeNum() dao.mysql.sql_nzx")
		return -1, err
	}
	return article.LikeAmount, nil
}

// UpdateArticleCollectNum 设置文章收藏数
func UpdateArticleCollectNum(aid, num int) error {
	if err := DB.Where("id = ?", aid).Update("collect_amount", num).Error; err != nil {
		fmt.Println("UpdateArticleLikeNum() dao.mysql.sql_nzx")
		return err
	}
	return nil
}

// QueryArticleCollectNum 获取文章收藏数
func QueryArticleCollectNum(aid int) (int, error) {
	article := model.Article{}
	if err := DB.Where("id = ?", aid).First(&article).Error; err != nil {
		fmt.Println("QueryArticleLikeNum() dao.mysql.sql_nzx")
		return -1, err
	}
	return article.CollectAmount, nil
}

// InsertArticleContent 插入文章内容
func InsertArticleContent(content, topic string, uid, wordCount int, tags []string, picPath []string, videoPath string) (int, error) {
	article := model.Article{
		UserID:    uint(uid),
		Content:   content,
		Topic:     topic,
		Video:     videoPath,
		WordCount: wordCount,
	}
	if err := DB.Create(&article).Error; err != nil {
		zap.L().Error("InsertArticleContent() dao.mysql.sql_article", zap.Error(err))
		return -1, err
	}
	// 同步标签中间表
	for _, tagName := range tags {
		tag := model.Tag{}
		if err := DB.Where("topic = ? and tag_name = ?", topic, tagName).First(&tag).Error; err != nil {
			zap.L().Error("InsertArticleContent() dao.mysql.sql_article", zap.Error(err))
			return -1, err
		}
		if err := DB.Create(&model.ArticleTag{
			ArticleID: article.ID,
			TagID:     tag.ID,
		}).Error; err != nil {
			zap.L().Error("InsertArticleContent() dao.mysql.sql_article", zap.Error(err))
			return -1, err
		}
	}

	if len(picPath) > 0 {
		for _, pic := range picPath {
			if err := DB.Create(model.ArticlePic{
				ArticleID: article.ID,
				Pic:       pic,
			}).Error; err != nil {
				zap.L().Error("InsertArticleContent() dao.mysql.sql_article", zap.Error(err))
				return -1, err
			}
		}
	}
	return int(article.ID), nil
}

// QueryClassByClassId 根据classid查找class
func QueryClassByClassId(classId int) (string, error) {
	class := model.UserClass{}
	if err := DB.Where("id = ?", classId).First(&class).Error; err != nil {
		zap.L().Error("InsertArticleContent() dao.mysql.sql_article", zap.Error(err))
		return "", err
	}
	return class.Class, nil
}

// QueryArticleByClass 根据班级分页查询文章
func QueryArticleByClass(limit, page int, class, keyWord string) (model.Articles, error) {
	var articles model.Articles
	if err := DB.Preload("User", "class = ?", class).
		Where("content like ?", keyWord).
		Order("created_at desc").
		Limit(limit).Offset((page - 1) * limit).Find(&articles).Error; err != nil {
		zap.L().Error("InsertArticleContent() dao.mysql.sql_article", zap.Error(err))
		return nil, err
	}
	return articles, nil
}
