package mysql

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	model "studentGrow/models/gorm_model"
	myErr "studentGrow/pkg/error"
)

// InsertIntoCommentsForArticle 向数据库插入评论数据(回复文章)
func InsertIntoCommentsForArticle(content string, aid int, uid int, db *gorm.DB) (int, error) {
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
	if err := db.Create(&comment).Error; err != nil {
		zap.L().Error("InsertIntoCommentsForArticle() dao.mysql.nzx_sql.Create err=", zap.Error(err))
		return -1, err
	}

	return int(comment.ID), nil
}

// InsertIntoCommentsForComment 向数据库插入评论数据(回复评论)
func InsertIntoCommentsForComment(content string, uid int, pid int, db *gorm.DB) (int, error) {
	// 找到父级评论的文章
	pComment := model.Comment{}
	if err := DB.Preload("Article").Where("id = ?", pid).First(&pComment).Error; err != nil {
		zap.L().Error("InsertIntoCommentsForComment() dao.mysql.nzx_sql.First err=", zap.Error(err))
		return -1, err
	}

	//content;id;username
	comment := model.Comment{
		Model:      gorm.Model{},
		Content:    content,
		LikeAmount: 0,
		IsRead:     false,
		UserID:     uint(uid),
		Pid:        uint(pid),
		ArticleID:  pComment.Article.ID,
	}

	if err := db.Create(&comment).Error; err != nil {
		zap.L().Error("InsertIntoCommentsForComment() dao.mysql.nzx_sql.Create err=", zap.Error(err))
		return -1, err
	}
	return int(comment.ID), nil
}

// QueryLevelOneComments 查询一级评论
func QueryLevelOneComments(aid, limit, page int) ([]model.Comment, error) {
	var comments []model.Comment
	if err := DB.Preload("User").Where("article_id = ? AND pid = ?", aid, 0).
		Order("created_at desc").
		Limit(limit).Offset((page - 1) * limit).
		Find(&comments).
		Error; err != nil {
		zap.L().Error("QueryLevelOneComments() dao.mysql.nzx_sql.Find err=", zap.Error(err))
		return nil, err
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
		zap.L().Error("QueryLevelSonComments() dao.mysql.nzx_sql.Find err=", zap.Error(err))
		return nil, err
	}

	return comments, nil
}

// QuerySonCommentNum 查询子评论数量
func QuerySonCommentNum(cid int) (int, error) {
	var count int64
	if err := DB.Model(&model.Comment{}).Where("pid = ?", cid).Count(&count).Error; err != nil {
		zap.L().Error("QuerySonCommentNum() dao.mysql.nzx_sql.Count err=", zap.Error(err))
		return -1, err
	}
	return int(count), nil

}

// DeleteComment 删除评论
func DeleteComment(cid int, username string) error {
	comment := model.Comment{}
	if err := DB.Preload("User").Where("id = ?", cid).First(&comment).Error; err != nil {
		zap.L().Error("DeleteComment() dao.mysql.nzx_sql.First err=", zap.Error(err))
		return err
	}

	// 只有自己能删除自己的评论
	if username != comment.User.Username {
		return myErr.OverstepCompetence()
	}

	err := DB.Transaction(func(db *gorm.DB) error {
		if comment.Pid == 0 {
			// 若为一级评论
			// 删除子评论
			result := db.Where("pid = ?", comment.ID).Delete(&model.Comment{})
			if result.Error != nil {
				zap.L().Error("DeleteComment() dao.mysql.nzx_sql.Delete err=", zap.Error(result.Error))
				return result.Error
			}
			// 删除父级评论
			if err := db.Delete(&comment).Error; err != nil {
				zap.L().Error("DeleteComment() dao.mysql.nzx_sql.Delete err=", zap.Error(err))
				return err
			}

			//	减少文章评论数
			num, err := QueryArticleCommentNum(int(comment.ArticleID))
			if err != nil {
				zap.L().Error("DeleteComment() dao.mysql.nzx_sql.QueryArticleCommentNum err=", zap.Error(err))
				return err
			}
			err = UpdateArticleCommentNum(int(comment.ArticleID), num-int(result.RowsAffected)-1, db)
			if err != nil {
				zap.L().Error("DeleteComment() dao.mysql.nzx_sql.UpdateArticleCommentNum err=", zap.Error(err))
				return err
			}

		} else {
			// 若为二级评论
			if err := DB.Where("id = ?", comment.ID).Delete(&model.Comment{}).Error; err != nil {
				zap.L().Error("DeleteComment() dao.mysql.nzx_sql.Delete err=", zap.Error(err))
				return err
			}

			//	减少文章评论数
			num, err := QueryArticleCommentNum(int(comment.ArticleID))
			if err != nil {
				zap.L().Error("DeleteComment() dao.mysql.nzx_sql.QueryArticleCommentNum err=", zap.Error(err))
				return err
			}
			err = UpdateArticleCommentNum(int(comment.ArticleID), num-1, db)
			if err != nil {
				zap.L().Error("DeleteComment() dao.mysql.nzx_sql.UpdateArticleCommentNum err=", zap.Error(err))
				return err
			}
		}
		return nil
	})
	if err != nil {
		zap.L().Error("DeleteComment() dao.mysql.nzx_sql.Transaction err=", zap.Error(err))
		return err
	}
	return nil
}

// QueryArticleIdByCommentId 查询评论的文章id
func QueryArticleIdByCommentId(cid int) (int, error) {
	var aid int
	if err := DB.Model(&model.Comment{}).Select("article_id").Where("id = ?", cid).First(&aid).Error; err != nil {
		zap.L().Error("QueryArticleIdByCommentId() dao.mysql.nzx_sql.Transaction err=", zap.Error(err))
		return -1, err
	}
	return aid, nil
}

// QueryCommentNumForLel1 获取一级评论的评论数
func QueryCommentNumForLel1(cid int) (int, error) {
	var count int64
	if err := DB.Model(&model.Comment{}).Where("pid = ?", cid).Count(&count).Error; err != nil {
		zap.L().Error("QueryCommentNumForLel1() dao.mysql.nzx_sql.Count err=", zap.Error(err))
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
		zap.L().Error("QueryUserAllComments() dao.mysql.sql_comment.Find err=", zap.Error(myErr.ErrNotFoundError))
		return nil, myErr.ErrNotFoundError
	}

	return comments, nil
}

// UpdateCommentLikeNum 设置评论点赞数
func UpdateCommentLikeNum(cid, num int, db *gorm.DB) error {
	if err := db.Model(&model.Comment{}).Where("id = ?", cid).Update("like_amount", num).Error; err != nil {
		zap.L().Error("UpdateCommentLikeNum() dao.mysql.sql_comment.Update err=", zap.Error(myErr.ErrNotFoundError))
		return err
	}
	return nil
}

// QueryCommentLikeNum 获取评论点赞数
func QueryCommentLikeNum(cid int) (int, error) {
	comment := model.Comment{}
	if err := DB.Where("id = ?", cid).First(&comment).Error; err != nil {
		zap.L().Error("QueryCommentLikeNum() dao.mysql.sql_comment.First err=", zap.Error(myErr.ErrNotFoundError))
		return -1, err
	}
	return comment.LikeAmount, nil
}
