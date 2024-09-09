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
		zap.L().Error("SelectUserById() dao.mysql.sql_nzx.First err=", zap.Error(err))
		return err, nil
	}

	return nil, user

}

// SelectUserByUsername 通过username查找uid
func SelectUserByUsername(username string) (uid int, err error) {
	//select id from users where username = username
	var user model.User
	if err := DB.Model(model.User{}).Select("id").Where("username = ?", username).First(&user).Error; err != nil {
		zap.L().Error("SelectUserByUsername() dao.mysql.sql_nzx.First err=", zap.Error(err))
		return int(user.ID), err
	} else {
		return int(user.ID), nil
	}
}

// QueryArticleIsExist 查询文章是否存在
func QueryArticleIsExist(aid int) (bool, error) {
	var count int64
	if DB.Where("id = ? and ban = ? and status = ?", aid, false, true).Count(&count).RowsAffected > 0 {
		return true, nil
	}
	return false, nil
}

// QueryArticleById 通过id查找文章(普通用户)
func QueryArticleById(aid int, uid uint) (err error, article *model.Article) {
	//查询用户 select * from articles where id = aid
	if err := DB.Preload("ArticlePics").Preload("ArticleTags.Tag").Preload("User").
		Where("id = ? and ban = ? and status = ?", aid, false, true).Or("user_id = ? and ban = ? and status = ?", uid, false, false).First(&article).Error; err != nil {
		zap.L().Error("SelectArticleById() dao.mysql.sql_nzx.First err=", zap.Error(err))
		return err, nil
	} else {
		return nil, article
	}
}

// QueryArticleByIdOfManager QueryArticleById 通过id查找文章(管理员)
func QueryArticleByIdOfManager(aid int) (*model.Article, error) {
	var article model.Article
	if err := DB.Preload("ArticlePics").Preload("ArticleTags.Tag").Preload("User").
		Where("id = ?", aid).First(&article).Error; err != nil {
		zap.L().Error("SelectArticleById() dao.mysql.sql_nzx.First err=", zap.Error(err))
		return nil, err
	}
	return &article, nil
}

// QueryArticleNumByDay 查询当日的相应话题的文章发表数量
func QueryArticleNumByDay(topic string, startOfDay time.Time, endOfDay time.Time, uid int) (int, error) {
	var count int64
	if err := DB.Model(&model.Article{}).Where("created_at >= ? AND	created_at < ? AND topic = ? AND user_id = ? AND status = ? AND ban = ?", startOfDay, endOfDay, topic, uid, true, false).
		Count(&count).Error; err != nil {
		zap.L().Error("SelectArticleById() dao.mysql.sql_nzx.First err=", zap.Error(err))
		return -1, err
	}
	return int(count), nil
}

// QueryArticleByAdvancedFilter 后台分页高级筛选文章结果
func QueryArticleByAdvancedFilter(startAtString, endAtString, topic, keyWords string, isBan bool) (query *gorm.DB, err error) {
	// 解析时间
	var startAt time.Time
	if startAtString != "" {
		startAt, err = time.Parse(time.RFC3339, startAtString)
		if err != nil {
			zap.L().Error("QueryArticleByAdvancedFilter() service.article.Parse err=", zap.Error(err))
			return nil, err
		}
	}
	var endAt time.Time
	if endAtString != "" {
		endAt, err = time.Parse(time.RFC3339, endAtString)
		if err != nil {
			zap.L().Error("QueryArticleByAdvancedFilter() service.article.Parse err=", zap.Error(err))
			return nil, err
		}
	}

	// 筛选
	if startAtString != "" && endAtString != "" {
		query = DB.Where("articles.created_at between ? and ? and topic like ? and content like ? and articles.ban = ?",
			startAt, endAt, fmt.Sprintf("%%%s%%", topic), fmt.Sprintf("%%%s%%", keyWords), isBan)
	} else if startAtString == "" && endAtString != "" {
		query = DB.Where("articles.created_at < ? and topic like ? and content like ? and articles.ban = ?",
			endAt, fmt.Sprintf("%%%s%%", topic), fmt.Sprintf("%%%s%%", keyWords), isBan)
	} else if startAtString != "" && endAtString == "" {
		query = DB.Where("articles.created_at > ? and topic like ? and content like ? and articles.ban = ?",
			startAt, fmt.Sprintf("%%%s%%", topic), fmt.Sprintf("%%%s%%", keyWords), isBan)
	} else if startAtString == "" && endAtString == "" {
		query = DB.Where("topic like ? and content like ? and articles.ban = ?",
			fmt.Sprintf("%%%s%%", topic), fmt.Sprintf("%%%s%%", keyWords), isBan)
	}

	return query, nil

}

// QueryArticleAndUserListByPageForClass 后台分页查询文章及用户列表并模糊查询 - 班级
func QueryArticleAndUserListByPageForClass(page, limit int, sort, order, startAtString, endAtString, topic, keyWords, name string, isBan bool, class string) (result []model.Article, err error) {
	query, err := QueryArticleByAdvancedFilter(startAtString, endAtString, topic, keyWords, isBan)

	var articles []model.Article
	if err := query.InnerJoins("User").Where("name like ? and class = ?", fmt.Sprintf("%%%s%%", name), class).Preload("ArticleTags.Tag").
		Order(fmt.Sprintf("%s %s", sort, order)).Limit(limit).Offset((page - 1) * limit).Find(&articles).Error; err != nil {
		zap.L().Error("SelectArticleAndUserListByPage() dao.mysql.sql_nzx.Find err=", zap.Error(err))
		return nil, err
	}

	return articles, nil
}

// QueryArticleAndUserListByPageForGrade 后台分页查询文章及用户列表并模糊查询 - 年级
func QueryArticleAndUserListByPageForGrade(page, limit int, sort, order, startAtString, endAtString, topic, keyWords, name string, isBan bool, grade int) (result []model.Article, err error) {
	query, err := QueryArticleByAdvancedFilter(startAtString, endAtString, topic, keyWords, isBan)

	var articles []model.Article
	if err := query.InnerJoins("User").Where("name like ? and grade = ?", fmt.Sprintf("%%%s%%", name), grade).Preload("ArticleTags.Tag").
		Order(fmt.Sprintf("%s %s", sort, order)).Limit(limit).Offset((page - 1) * limit).Find(&articles).Error; err != nil {
		zap.L().Error("SelectArticleAndUserListByPage() dao.mysql.sql_nzx.Find err=", zap.Error(err))
		return nil, err
	}

	return articles, nil
}

// QueryArticleAndUserListByPageForSuperman 后台分页查询文章及用户列表并模糊查询 - 院级(超级)
func QueryArticleAndUserListByPageForSuperman(page, limit int, sort, order, startAtString, endAtString, topic, keyWords, name string, isBan bool) (result []model.Article, err error) {
	query, err := QueryArticleByAdvancedFilter(startAtString, endAtString, topic, keyWords, isBan)

	var articles []model.Article
	if err := query.InnerJoins("User").Where("name like ?", fmt.Sprintf("%%%s%%", name)).Preload("ArticleTags.Tag").
		Order(fmt.Sprintf("%s %s", sort, order)).Limit(limit).Offset((page - 1) * limit).Find(&articles).Error; err != nil {
		zap.L().Error("SelectArticleAndUserListByPage() dao.mysql.sql_nzx.Find err=", zap.Error(err))
		return nil, err
	}

	return articles, nil
}

// QueryArticleAndUserListByPageFirstPageByTopic 前台模糊查询文章列表(根据话题查询)
func QueryArticleAndUserListByPageFirstPageByTopic(keyWords, topic string, limit, page int) (result model.Articles, err error) {
	var articles model.Articles
	if err = DB.Preload("User").Preload("ArticleTags.Tag").Preload("ArticlePics").
		Where("topic = ? and content like ? and ban = ? and status = ?", topic, fmt.Sprintf("%%%s%%", keyWords), false, true).
		Order("created_at desc").
		Limit(limit).
		Offset((page - 1) * limit).Find(&articles).Error; err != nil {
		zap.L().Error("SelectArticleAndUserListByPageFirstPage() dao.mysql.sql_nzx err=", zap.Error(err))
		return nil, err
	}

	return articles, nil
}

// QueryArticleAndUserListByPageFirstPage 前台模糊查询文章列表(全部)
func QueryArticleAndUserListByPageFirstPage(keyWords string, limit, page int) (result model.Articles, err error) {
	var articles model.Articles
	if err = DB.Preload("User").Preload("ArticleTags.Tag").Preload("ArticlePics").
		Where("content like ? and ban = ? and status = ?", fmt.Sprintf("%%%s%%", keyWords), false, true).
		Order("created_at desc").
		Limit(limit).
		Offset((page - 1) * limit).Find(&articles).Error; err != nil {
		zap.L().Error("SelectArticleAndUserListByPageFirstPage() dao.mysql.sql_nzx err=", zap.Error(err))
		return nil, err
	}

	return articles, nil
}

// BannedArticleByIdForClass 通过文章id对文章进行封禁或解封 - 班级
func BannedArticleByIdForClass(articleId int, isBan bool, username string, db *gorm.DB) error {
	// 查询班级管理员信息
	user := model.User{}
	if err := DB.Where("username = ?", username).First(&user).Error; err != nil {
		zap.L().Error("BannedArticleByIdForClass() dao.mysql.sql_nzx.First err=", zap.Error(err))
		return err
	}

	// 查询待封禁的文章;若查询不到，则返回
	article := model.Article{}
	if err := DB.Preload("User", "class = ?", user.Class).Where("id = ?", articleId).First(&article).Error; err != nil {
		zap.L().Error("BannedArticleByIdForClass() dao.mysql.sql_nzx.First err=", zap.Error(err))
		return myErr.OverstepCompetence()
	}

	// 修改文章状态
	if err := db.Model(&model.Article{}).Where("id = ?", articleId).Updates(model.Article{Ban: true}).Error; err != nil {
		zap.L().Error("BannedArticleByIdForClass() dao.mysql.sql_nzx.Updates err=", zap.Error(err))
		return err
	}

	return nil
}

// BannedArticleByIdForGrade 通过文章id对文章进行封禁或解封 - 年级
func BannedArticleByIdForGrade(articleId int, grade int, db *gorm.DB) error {
	// GetUnreadReportsForGrade
	year, err := timeConverter.GetEnrollmentYear(grade)
	if err != nil {
		zap.L().Error("BannedArticleByIdForClass() dao.mysql.sql_nzx.GetEnrollmentYear err=", zap.Error(err))
		return err
	}

	// 获取需要被封禁的文章；若找不到则返回
	article := model.Article{}
	if err = DB.Preload("User", "plus_time between ? and ?",
		fmt.Sprintf("%s-01-01", year.Year()), fmt.Sprintf("%s-12-31", year.Year())).
		Where("id = ?", articleId).First(&article).Error; err != nil {
		zap.L().Error("BannedArticleByIdForClass() dao.mysql.sql_nzx.First err=", zap.Error(err))
		return myErr.OverstepCompetence()
	}

	// 修改文章状态
	if err := db.Model(&model.Article{}).Where("id = ?", articleId).Updates(model.Article{Ban: true}).Error; err != nil {
		zap.L().Error("BannedArticleByIdForClass() dao.mysql.sql_nzx.Updates err=", zap.Error(err))
		return err
	}
	return nil
}

// BannedArticleByIdForSuperman 通过文章id对文章进行封禁或解封 - 院级(超级)
func BannedArticleByIdForSuperman(articleId int, db *gorm.DB) error {
	// 修改文章状态
	if err := db.Model(&model.Article{}).Where("id = ?", articleId).Updates(model.Article{Ban: true}).Error; err != nil {
		zap.L().Error("BannedArticleByIdForSuperman() dao.mysql.sql_nzx.Updates err=", zap.Error(err))
		return err
	}
	return nil
}

// QueryIsBanByArticleId 通过文章id查询该文章的封禁状态
func QueryIsBanByArticleId(aid int) (bool, error) {
	var isBan bool
	if err := DB.Model(&model.Article{}).Select("ban").Where("id = ?", aid).First(&isBan).Error; err != nil {
		zap.L().Error("QueryIsBanByArticleId() dao.mysql.sql_nzx.Updates err=", zap.Error(err))
		return false, err
	}
	return isBan, nil
}

// DeleteArticleReportMsg 已读举报信息
func DeleteArticleReportMsg(aid int, db *gorm.DB) error {
	if err := db.Model(model.UserReportArticleRecord{}).Where("article_id = ?", aid).Update("is_read", true).Error; err != nil {
		zap.L().Error("GetUnreadReportsController() dao.mysql.sql_article err=", zap.Error(err))
		return err
	}
	return nil
}

// QueryIsExistArticleIdByReportMsg 查询举报信箱中是否存在被举报的文章id
func QueryIsExistArticleIdByReportMsg(aid int) (bool, error) {
	var count int64
	if err := DB.Model(&model.UserReportArticleRecord{}).Where("article_id = ?", aid).Count(&count).Error; err != nil {
		zap.L().Error("QueryIsExistArticleIdByReportMsg() dao.mysql.sql_nzx.First err=", zap.Error(err))
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

// DeleteArticleByIdForClass 通过文章id删除文章 - 班级
func DeleteArticleByIdForClass(articleId int, username string, db *gorm.DB) error {
	// 查询班级管理员信息
	user := model.User{}
	if err := DB.Where("username = ?", username).First(&user).Error; err != nil {
		zap.L().Error("DeleteArticleByIdForClass() dao.mysql.sql_nzx.First err=", zap.Error(err))
		return err
	}

	// 查询待删除的文章
	article := model.Article{}
	if err := DB.Preload("User", "class = ?", user.Class).Where("id = ?", articleId).First(&article).Error; err != nil {
		zap.L().Error("DeleteArticleByIdForClass() dao.mysql.sql_nzx.First err=", zap.Error(err))
		return err
	}

	if err := db.Delete(&model.Article{}, article.ID).Error; err != nil {
		zap.L().Error("DeleteArticleByIdForClass() dao.mysql.sql_nzx.Delete err=", zap.Error(err))
		return err
	}

	return nil
}

// DeleteArticleByIdForGrade 通过文章id删除文章 - 年级
func DeleteArticleByIdForGrade(articleId int, grade int, db *gorm.DB) error {
	// 将年级转化为入学年份
	year, err := timeConverter.GetEnrollmentYear(grade)
	if err != nil {
		zap.L().Error("DeleteArticleByIdForGrade() dao.mysql.sql_nzx.GetEnrollmentYear err=", zap.Error(err))
		return err
	}

	// 获取需要被删除的文章
	article := model.Article{}
	if err = DB.Preload("User", "plus_time between ? and ?",
		fmt.Sprintf("%d-01-01", year.Year()), fmt.Sprintf("%d-12-31", year.Year())).
		Where("id = ?", articleId).First(&article).Error; err != nil {
		zap.L().Error("DeleteArticleByIdForGrade() dao.mysql.sql_nzx.First err=", zap.Error(err))
		return err
	}

	// 删除文章
	if err = db.Delete(&model.Article{}, article.ID).Error; err != nil {
		zap.L().Error("DeleteArticleByIdForGrade() dao.mysql.sql_nzx.Delete err=", zap.Error(err))
		return err
	}
	return nil
}

// DeleteArticleByIdForSuperman 通过id删除文章 - 院级(超级)
func DeleteArticleByIdForSuperman(articleId int, db *gorm.DB) error {
	article := model.Article{}
	if err := DB.Where("id = ?", articleId).First(&article).Error; err != nil {
		zap.L().Error("DeleteArticleByIdForSuperman() dao.mysql.sql_nzx.First err=", zap.Error(err))
		return err
	}
	if err := db.Delete(&model.Article{}, article.ID).Error; err != nil {
		zap.L().Error("DeleteArticleByIdForSuperman() dao.mysql.sql_nzx.Delete err=", zap.Error(err))
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
			bg.Rollback()
		}
	}()

	// 获取被举报文章举报量，并对举报量+1操作
	article := model.Article{}
	if err := DB.Where("id = ?", uint(aid)).First(&article).Error; err != nil {
		zap.L().Error("ReportArticleById() dao.mysql.sql_nzx.First err=", zap.Error(err))
		return err
	}
	article.ReportAmount += 1
	result := DB.Model(model.Article{}).Select("report_amount").Where("id = ?", aid).Save(&article)

	if result.Error != nil {
		bg.Rollback()
		zap.L().Error("ReportArticleById() dao.mysql.sql_nzx.Save err=", zap.Error(result.Error))
		return result.Error
	}
	// 查询更新结果
	if result.RowsAffected <= 0 {
		zap.L().Error("ReportArticleById() dao.mysql.sql_nzx.Save err=", zap.Error(myErr.ErrNotFoundError))
		return myErr.ErrNotFoundError
	}

	// 检查举报记录：不允许重复举报
	var report []model.UserReportArticleRecord
	if err := DB.Where("user_id = ? and article_id = ?", uid, aid).Find(&report).Error; err != nil {
		zap.L().Error("ReportArticleById() dao.mysql.sql_nzx.Find err=", zap.Error(myErr.ErrNotFoundError))
		bg.Rollback()
		return err
	}

	//如果数据库有重复记录，则拒绝重复提交
	if len(report) > 0 {
		zap.L().Error("ReportArticleById() dao.mysql.sql_nzx.Find err=", zap.Error(myErr.RejectRepeatSubmission()))
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
		zap.L().Error("ReportArticleById() dao.mysql.sql_nzx.Create err=", zap.Error(err))
		bg.Rollback()
		return err
	}

	// 提交
	if err := bg.Commit().Error; err != nil {
		zap.L().Error("ReportArticleById() dao.mysql.sql_nzx.Commit err=", zap.Error(err))
		bg.Rollback()
		return err
	}
	return nil
}

// SearchHotArticlesOfDay 查找今日热门文章
func SearchHotArticlesOfDay(startOfDay time.Time, endOfDay time.Time) (model.Articles, error) {
	var articles model.Articles
	if err := DB.Where("created_at >= ? and created_at < ? and ban = ? and status = ?", startOfDay, endOfDay, false, true).Preload("ArticleTags.Tag").
		Find(&articles).Error; err != nil {
		zap.L().Error("SearchHotArticlesOfDay() dao.mysql.sql_nzx.Find err=", zap.Error(err))
		return nil, err
	}
	return articles, nil
}

// UpdateArticleCommentNum 设置文章评论数
func UpdateArticleCommentNum(aid, num int, db *gorm.DB) error {
	if err := db.Model(&model.Article{}).Where("id = ?", aid).Update("comment_amount", num).Error; err != nil {
		zap.L().Error("UpdateArticleCommentNum() dao.mysql.sql_nzx.Update err=", zap.Error(err))
		return err
	}
	return nil
}

// QueryArticleCommentNum 获取文章评论数
func QueryArticleCommentNum(aid int) (int, error) {
	article := model.Article{}
	if err := DB.Where("id = ?", aid).First(&article).Error; err != nil {
		zap.L().Error("QueryArticleCommentNum() dao.mysql.sql_nzx.First err=", zap.Error(err))
		return -1, err
	}
	return article.CommentAmount, nil
}

// UpdateArticleLikeNum 设置文章点赞数
func UpdateArticleLikeNum(aid, num int, db *gorm.DB) error {
	if err := db.Model(&model.Article{}).Where("id = ?", aid).Update("like_amount", num).Error; err != nil {
		zap.L().Error("UpdateArticleLikeNum() dao.mysql.sql_nzx.Update err=", zap.Error(err))
		return err
	}
	return nil
}

// QueryArticleLikeNum 获取文章点赞数
func QueryArticleLikeNum(aid int) (int, error) {
	article := model.Article{}
	if err := DB.Where("id = ?", aid).First(&article).Error; err != nil {
		zap.L().Error("QueryArticleLikeNum() dao.mysql.sql_nzx.First err=", zap.Error(err))
		return -1, err
	}
	return article.LikeAmount, nil
}

// UpdateArticleCollectNum 设置文章收藏数
func UpdateArticleCollectNum(aid, num int) error {
	if err := DB.Model(&model.Article{}).Where("id = ?", aid).Update("collect_amount", num).Error; err != nil {
		zap.L().Error("UpdateArticleCollectNum() dao.mysql.sql_nzx.Update err=", zap.Error(err))
		return err
	}
	return nil
}

// QueryArticleCollectNum 获取文章收藏数
func QueryArticleCollectNum(aid int) (int, error) {
	article := model.Article{}
	if err := DB.Where("id = ?", aid).First(&article).Error; err != nil {
		zap.L().Error("QueryArticleCollectNum() dao.mysql.sql_nzx.First err=", zap.Error(err))
		return -1, err
	}
	return article.CollectAmount, nil
}

// InsertArticleContent 插入文章内容
func InsertArticleContent(content, topic string, uid, wordCount int, tags []string, picPath []string, videoPath string, status bool, db *gorm.DB) (int, error) {
	article := model.Article{
		UserID:    uint(uid),
		Content:   content,
		Topic:     topic,
		Video:     videoPath,
		WordCount: wordCount,
		Status:    status,
	}
	print(article.Status)
	if err := db.Create(&article).Error; err != nil {
		zap.L().Error("InsertArticleContent() dao.mysql.sql_article", zap.Error(err))
		return -1, err
	}
	// 同步标签表
	for _, tagName := range tags {
		tag := model.Tag{}
		if err := db.Where("topic = ? and tag_name = ?", topic, tagName).First(&tag).Error; err != nil {
			zap.L().Error("InsertArticleContent() dao.mysql.sql_article", zap.Error(err))
			return -1, err
		}
		if err := db.Create(&model.ArticleTag{
			ArticleID: article.ID,
			TagID:     tag.ID,
		}).Error; err != nil {
			zap.L().Error("InsertArticleContent() dao.mysql.sql_article", zap.Error(err))
			return -1, err
		}
	}

	// 同步图片
	if len(picPath) > 0 {
		for _, pic := range picPath {
			if err := db.Create(&model.ArticlePic{
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
	if err := DB.Where("class = ?", class).First(&class).Error; err != nil {
		zap.L().Error("InsertArticleContent() dao.mysql.sql_article", zap.Error(err))
		return "", err
	}
	return class.Class, nil
}

// QueryArticleByClass 根据班级分页查询文章
func QueryArticleByClass(limit, page int, class, keyWord string) (model.Articles, error) {
	var uids []int
	if err := DB.Model(&model.User{}).Where("class = ?", class).Pluck("id", &uids).Error; err != nil {
		zap.L().Error("QueryArticleByClass() dao.mysql.sql_article", zap.Error(err))
		return nil, err
	}

	var articles model.Articles
	if err := DB.Preload("User").Preload("ArticleTags.Tag").
		Where("content like ? and ban = ? and status = ? AND user_id IN ?", fmt.Sprintf("%%%s%%", keyWord), false, true, uids).
		Order("created_at desc").
		Limit(limit).Offset((page - 1) * limit).Find(&articles).Error; err != nil {
		zap.L().Error("QueryArticleByClass() dao.mysql.sql_article", zap.Error(err))
		return nil, err
	}

	return articles, nil
}

// QueryArticleStatusById 通过id查询文章的私密状态
func QueryArticleStatusById(aid int) (bool, error) {
	var article model.Article
	if err := DB.Where("id = ?", aid).First(&article).Error; err != nil {
		zap.L().Error("QueryArticleStatusById() dao.mysql.sql_article", zap.Error(err))
		return false, err
	}
	return article.Status, nil
}

// UpdateArticleStatusById 通过id修改文章的私密状态
func UpdateArticleStatusById(aid int, status bool, db *gorm.DB) error {
	if err := db.Model(&model.Article{}).Where("id = ?", aid).Update("status", status).Error; err != nil {
		zap.L().Error("UpdateArticleStatusById() dao.mysql.sql_article", zap.Error(err))
		return err
	}
	return nil
}

// QueryArticlePoint 查询文章分数
func QueryArticlePoint(aid int) (int, error) {
	var point int
	if err := DB.Model(&model.Article{}).Select("point").Where("id = ?", aid).First(&point).Error; err != nil {
		zap.L().Error("QueryArticlePoint() dao.mysql.sql_article", zap.Error(err))
		return -1, err
	}
	return point, nil
}

// UpdateArticlePoint 修改文章分数
func UpdateArticlePoint(aid int, point int) error {
	if err := DB.Model(&model.Article{}).Where("id = ?", aid).Update("point", point).Error; err != nil {
		zap.L().Error("UpdateArticlePoint() dao.mysql.sql_article", zap.Error(err))
		return err
	}
	return nil
}

// QueryContentByArticleId 通过文章id获取文章内容
func QueryContentByArticleId(aid int) (string, error) {
	var content string
	if err := DB.Model(&model.Article{}).Select("content").Where("id = ?", aid).First(&content).Error; err != nil {
		zap.L().Error("QueryContentByArticleId() dao.mysql.sql_article", zap.Error(err))
		return "", err
	}
	return content, nil
}

// QueryArticleIdsByUserId 通过用户id获取该用户的所有文章id
func QueryArticleIdsByUserId(uid int) ([]int, error) {
	// 查询该用户的所有文章id
	var aids []int
	if err := DB.Model(&model.Article{}).Select("id").Where("user_id = ? AND ban = ? AND status = ?", uid, false, true).Find(&aids).Error; err != nil {
		zap.L().Error("QueryCollectRecordByUserArticles() dao.mysql.sql_msg.Find err=", zap.Error(err))
		return nil, err
	}
	return aids, nil
}

// QueryUserByArticleId 通过文章获取用户User
func QueryUserByArticleId(aid int) (*model.User, error) {
	var article model.Article
	if err := DB.Preload("User").Where("id = ?", aid).First(&article).Error; err != nil {
		zap.L().Error("QueryUserByArticleId() dao.mysql.sql_article", zap.Error(err))
		return nil, err
	}
	return &article.User, nil
}

// QueryUserIsManager 查询用户是否为管理员
func QueryUserIsManager(uid uint) (bool, error) {
	var isManager bool
	if err := DB.Model(&model.User{}).Select("is_manager").Where("id = ?", uid).First(&isManager).Error; err != nil {
		zap.L().Error("QueryUserByArticleId() dao.mysql.sql_article", zap.Error(err))
		return false, err
	}
	return isManager, nil
}

// DeleteArticlePicByArticleId 根据文章id删除文章关联的图片表
func DeleteArticlePicByArticleId(aid int) error {
	if err := DB.Where("article_id = ?", aid).Delete(&model.ArticlePic{}).Error; err != nil {
		zap.L().Error("DeleteArticlePicByArticleId() dao.mysql.sql_article", zap.Error(err))
		return err
	}
	return nil
}

// DeleteArticleTagByArticleId 根据文章id删除文章关联的标签表
func DeleteArticleTagByArticleId(aid int) error {
	if err := DB.Where("article_id = ?", aid).Delete(&model.ArticleTag{}).Error; err != nil {
		zap.L().Error("DeleteArticlePicByArticleId() dao.mysql.sql_article", zap.Error(err))
		return err
	}
	return nil
}

// QueryArticleNum 查看文章总数量
func QueryArticleNum() (int, error) {
	var count int64
	if err := DB.Model(&model.Article{}).Count(&count).Error; err != nil {
		zap.L().Error("QueryArticleNum() dao.mysql.sql_article.Count", zap.Error(err))
		return -1, err
	}
	return int(count), nil
}
