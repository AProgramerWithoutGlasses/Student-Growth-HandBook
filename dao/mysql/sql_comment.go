package mysql

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"strconv"
	"studentGrow/dao/redis"
	model "studentGrow/models/gorm_model"
	myErr "studentGrow/pkg/error"
)

// InsertIntoCommentsForArticle 向数据库插入评论数据(回复文章)
func InsertIntoCommentsForArticle(content string, aid int, uid int) (err error) {
	//content;id;username
	comment := model.Comment{
		Model:      gorm.Model{},
		Content:    content,
		LikeAmount: 0,
		IsRead:     false,
		UserID:     uint(uid),
		Pid:        0,
		ArticleID:  uint(aid),
	}
	if err = DB.Create(&comment).Error; err != nil {
		fmt.Println("InsertIntoCommentsForArticle() dao.mysql.nzx_sql err=", err)
		return err
	}

	return nil
}

// InsertIntoCommentsForComment 向数据库插入评论数据(回复评论)
func InsertIntoCommentsForComment(content string, uid int, pid int) error {
	// 找到父级评论的文章
	article := model.Article{}
	if err := DB.Preload("Article").Where("id = ?", pid).First(&article).Error; err != nil {
		fmt.Println("InsertIntoCommentsForComment() dao.mysql.nzx_sql err=", err)
		return err
	}

	//content;id;username
	comment := model.Comment{
		Model:      gorm.Model{},
		Content:    content,
		LikeAmount: 0,
		IsRead:     false,
		UserID:     uint(uid),
		Pid:        uint(pid),
		ArticleID:  article.ID,
	}

	if err := DB.Create(&comment).Error; err != nil {
		fmt.Println("InsertIntoCommentsForComment() dao.mysql.nzx_sql err=", err)
		return err
	}
	return nil
}

// QueryLevelOneComments 查询一级评论
func QueryLevelOneComments(aid, limit, page int) ([]model.Comment, error) {
	var comments []model.Comment
	if err := DB.Preload("User").Where("article_id = ?", aid).
		Order("created_at desc").
		Limit(limit).Offset((page - 1) * limit).
		Find(&comments).
		Error; err != nil {
		fmt.Println("QueryLevelOneComments() dao.mysql.nzx_sql err=", err)
		return nil, err
	}

	if len(comments) == 0 {
		return nil, myErr.NotFoundError()
	}
	return comments, nil
}

// QueryLevelSonComments 查询子评论
func QueryLevelSonComments(pid, limit, page int) ([]model.Comment, error) {
	var comments []model.Comment
	if err := DB.Preload("User").Where("pid = ?", pid).
		Order("created_at desc").
		Limit(limit).Offset((page - 1) * limit).
		Find(&comments).
		Error; err != nil {
		fmt.Println("QueryLevelTwoComments() dao.mysql.nzx_sql err=", err)
		return nil, err
	}

	if len(comments) == 0 {
		return nil, myErr.NotFoundError()
	}
	return comments, nil
}

// QuerySonCommentNum 查询子评论数量
func QuerySonCommentNum(cid int) (int, error) {
	var count int64
	if err := DB.Model(&model.Comment{}).Where("pid = ?", cid).Count(&count).Error; err != nil {
		fmt.Println("QuerySonCommentNum() dao.mysql.nzx_sql err=", err)
		return -1, err
	}
	return int(count), nil

}

// DeleteComment 删除评论
func DeleteComment(cid int, username string) error {
	comment := model.Comment{}
	if err := DB.Preload("User").Where("id = ?", cid).First(&comment).Error; err != nil {
		fmt.Println("DeleteComment() dao.mysql.nzx_sql err=", err)
		return err
	}

	if username != comment.User.Username {
		return myErr.OverstepCompetence()
	}

	if comment.Pid == 0 {
		// 若为一级评论
		// 删除子评论
		result := DB.Where("pid = ?", comment.ID).Delete(&model.Comment{})
		if result.Error != nil {
			fmt.Println("DeleteComment() dao.mysql.nzx_sql err=", result.Error)
			return result.Error
		}
		// 删除父级评论
		if err := DB.Delete(&comment).Error; err != nil {
			fmt.Println("DeleteComment() dao.mysql.nzx_sql err=", err)
			return err
		}

		//	减少文章评论数
		num, err := QueryArticleCommentNum(int(comment.ArticleID))
		if err != nil {
			fmt.Println("DeleteComment() dao.mysql.nzx_sql.QueryArticleCommentNum err=", err)
			return err
		}
		err = UpdateArticleCommentNum(int(comment.ArticleID), num-int(result.RowsAffected)-1)
		if err != nil {
			fmt.Println("DeleteComment() dao.mysql.nzx_sql.UpdateArticleCommentNum err=", err)
			return err
		}

	} else {
		// 若为二级评论
		if err := DB.Where("id = ?", comment.ID).Delete(&model.Comment{}).Error; err != nil {
			fmt.Println("DeleteComment() dao.mysql.nzx_sql err=", err)
			return err
		}

		//	减少文章评论数
		num, err := QueryArticleCommentNum(int(comment.ArticleID))
		if err != nil {
			fmt.Println("DeleteComment() dao.mysql.nzx_sql.QueryArticleCommentNum err=", err)
			return err
		}
		err = UpdateArticleCommentNum(int(comment.ArticleID), num-1)
		if err != nil {
			fmt.Println("DeleteComment() dao.mysql.nzx_sql.UpdateArticleCommentNum err=", err)
			return err
		}
	}
	return nil
}

// QueryCommentLikeNum 查询评论点赞数
func QueryCommentLikeNum(cid int) (int, error) {
	num, err := redis.GetObjLikes(strconv.Itoa(cid), 1)
	if err != nil {
		fmt.Println("QueryCommentLikeNum() dao.mysql.nzx_sql.GetObjLikes err=", err)
		return 0, err
	}
	return num, nil
}

// UpdateCommentLikeNum 设置评论点赞数
func UpdateCommentLikeNum(cid, num int) error {
	err := redis.SetObjLikes(strconv.Itoa(cid), num, 1)
	if err != nil {
		fmt.Println("UpdateCommentLikeNum() dao.mysql.nzx_sql.SetObjLikes err=", err)
		return err
	}
	return nil
}

// QueryCommentNumForLel1 获取一级评论的评论数
func QueryCommentNumForLel1(cid int) (int, error) {
	var count int64
	if err := DB.Model(&model.Comment{}).Where("pid = ?", cid).Count(&count).Error; err != nil {
		fmt.Println("QueryCommentNumForLel1() dao.mysql.nzx_sql err=", err)
		return -1, err
	}
	return int(count), nil
}

// QueryUserAllComments 查找用户的所有一级评论
func QueryUserAllComments(uid int) (model.Comments, error) {
	comments := model.Comments{}
	if err := DB.Where("user_id = ? and pid = ?", uid, 0).Order("created_at desc").
		Find(&comments).Error; err != nil {
		zap.L().Error("QueryUserAllComments() dao.mysql.mysql_like.Find err=", zap.Error(err))
		return nil, err
	}

	if len(comments) == 0 {
		zap.L().Error("QueryUserAllComments() dao.mysql.sql_comment.Find err=", zap.Error(myErr.NotFoundError()))
		return nil, myErr.NotFoundError()
	}

	return comments, nil
}
