package articleHeat

import (
	"errors"
	"github.com/go-redis/redis"
	"studentGrow/dao/redis/articleHeat"
)

type ArticleHeatTopic struct {
	ListName string
	Aid      uint
	Type     string
}

// OptArticleHeatCal 实时计算文章热度
func OptArticleHeatCal(a ArticleHeatTopic) error {
	// 判断ZSet中是否存在该文章,若没有则需要先添加
	err := articleHeat.QueryArticleHeat(a.Aid, a.ListName)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			err := articleHeat.SaveArticleInZSet(a.Aid, a.ListName)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}

	// 开始更新
	mapping := map[string]float64{
		"like":      1.75,
		"unLike":    -1.75,
		"collect":   3.2,
		"unCollect": -3.2,
		"comment":   1,
		"unComment": -1,
	}
	// 计算热度值
	increment := mapping[a.Type]

	err = articleHeat.UpdateArticleHeat(a.Aid, increment, a.ListName)
	if err != nil {
		return err
	}
	return nil
}
