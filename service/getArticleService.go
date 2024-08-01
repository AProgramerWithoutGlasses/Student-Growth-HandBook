package service

import (
	"strconv"
	"studentGrow/dao/mysql"
	model "studentGrow/models/gorm_model"
	"studentGrow/utils/isEmptyData"
)

func GetArticleLogic(m map[string]string) (error, model.User, model.Article) {
	var user model.User
	var article model.Article

	strId := m["article_id"]
	aid, _ := strconv.Atoi(strId)
	user, article = mysql.GetArticleAndUserByArticleId(aid)

	if isEmptyData.IsEmptyStruct(user) || isEmptyData.IsEmptyStruct(article) {
		// return errors.NoFindError, user, article
	}
	return nil, user, article
}
