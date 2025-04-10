package article

import (
	"fmt"
	"studentGrow/dao/mysql"
	model "studentGrow/models/gorm_model"
	"studentGrow/pkg/eventBus"
	"studentGrow/pkg/global"
	"studentGrow/pkg/subscribes/articleHeat"
)

func getListName(day string) string {
	return fmt.Sprintf("articleHeatList:%s", day)
}

// 计算ZSet key
func getListNameByCreatedAt(aid uint) (string, error) {
	// 查询文章的发布日期
	article := new(model.Article)
	if err := mysql.DB.Model(article).Select("created_at").Where("id = ?", aid).First(article).Error; err != nil {
		return "", err
	}

	day := article.CreatedAt.Format("2006-01-02")
	return getListName(day), nil
}

// SendOptHeatEvent 发送热度实时计算event
func SendOptHeatEvent(aid uint, t string) error {
	// 计算key
	listName, err := getListNameByCreatedAt(aid)
	if err != nil {
		return err
	}

	eventBus.Bus.Publish(global.ArticleHeatTopic, articleHeat.ArticleHeatTopic{
		ListName: listName,
		Aid:      aid,
		Type:     t,
	})
	return nil
}
