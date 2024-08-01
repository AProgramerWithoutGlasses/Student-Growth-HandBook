package mysql

import (
	model "studentGrow/models/gorm_model"
)

// GetArticleAndUserByArticleId 通过文章id获取文章信息和用户信息
func GetArticleAndUserByArticleId(aid int) (model.User, model.Article) {
	//获取用户信息 select * from users where id = (select user_id from articles where id = aid)
	//获取文章信息 select * from articles where id = aid
	var user model.User
	var article model.Article
	DB.Model(model.User{}).Where("id = (?)", DB.Model(model.Article{}).Select("user_id")).Find(&user)
	DB.Model(model.Article{}).Where("id = ?", aid).Find(&article)
	return user, article
}
