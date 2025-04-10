package eventBus

import (
	"fmt"
	"github.com/asaskevich/EventBus"
	"studentGrow/pkg/global"
	"studentGrow/pkg/subscribes/articleHeat"
)

var Bus EventBus.Bus

func InitEventBus() error {
	Bus = EventBus.New()

	// 文章热度实时计算事件订阅
	err := Bus.Subscribe(global.ArticleHeatTopic, articleHeat.OptArticleHeatCal)
	if err != nil {
		panic(fmt.Errorf("initEventBus Error--> %v", err))
		return err
	}

	return nil
}
