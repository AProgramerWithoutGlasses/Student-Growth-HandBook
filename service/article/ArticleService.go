package article

import (
	"fmt"
	jsonvalue "github.com/Andrew-M-C/go.jsonvalue"
	"studentGrow/dao/mysql"
	model "studentGrow/models/gorm_model"
)

// GetArticleService 获取文章详情
func GetArticleService(j *jsonvalue.V) (err error, user *model.User, article *model.Article) {
	//获取文章id
	//aid, err := utils.StringToInt(m["article_id"].(string))
	aid, _ := j.GetInt("article_id")
	fmt.Println("aid----:", aid)
	//查找文章信息
	err, article = mysql.SelectArticleById(aid)
	if err != nil {
		fmt.Println("GetArticleLogic() dao.mysql.sqp_nzx.SelectArticleById err=", err)
		return err, nil, nil
	}
	//查找用户信息
	err, user = mysql.SelectUserById(article.UserId)
	fmt.Println(article.UserId)
	if err != nil {
		fmt.Println("GetArticleAndUserByArticleId() dao.mysql.sqp_nzx.SelectUserById err=", err)
		return err, nil, nil
	}

	return nil, user, article
}

// PublishArticleService 发布文章
func PublishArticleService(m map[string]string) {
	//article := model.Article{
	//	Content: m["article_content"],
	//	Topic:   m["article_topic"],
	//	Tag:     m["article_tags"],
	//}

	fmt.Println(m)

}
