package mysql

import (
	"fmt"
	"gorm.io/gorm"
	model "studentGrow/models/gorm_model"
)

// SelectUserById 查询数据库是否存在该用户
func SelectUserById(uid int) (err error, user *model.User) {
	//select * from users where id = uid
	// 查询用户
	if err := DB.Where("id = ?", uid).First(&user).Error; err != nil {
		fmt.Println("SelectUserById() dao.mysql err=", err)
		return err, nil
	} else {
		return nil, user
	}
}

// SelectUserByUsername 通过username查找uid
func SelectUserByUsername(username string) (uid int, err error) {
	//select id from users where username = username
	var user model.User
	if err := DB.Model(model.User{}).Select("id").Where("username = ?", username).Find(&user).Error; err != nil {
		fmt.Println("Error:", err)
		return int(user.ID), err
	} else {
		fmt.Println("user.ID", int(user.ID))
		return int(user.ID), nil
	}
}

// SelectArticleById 通过id查找文章
func SelectArticleById(aid int) (err error, article *model.Article) {
	//查询用户 select * from articles where id = aid
	fmt.Println("id:", aid)
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
	comment := model.Comment{
		Model:        gorm.Model{},
		Content:      content,
		UpvoteAmount: 0,
		IsRead:       false,
		Del:          false,
		Uid:          uid,
		Pid:          0,
		Aid:          aid,
		Upvote:       nil,
	}
	DB.Create(&comment)

	return nil
}

// InsertIntoCommentsForComment 向数据库插入评论数据(回复评论)
func InsertIntoCommentsForComment(content string, uid int, pid int) (err error) {
	//content;id;username
	comment := model.Comment{
		Model:        gorm.Model{},
		Content:      content,
		UpvoteAmount: 0,
		IsRead:       false,
		Del:          false,
		Uid:          uid,
		Pid:          pid,
		Aid:          0,
		Upvote:       nil,
	}

	DB.Create(&comment)
	return nil
}
