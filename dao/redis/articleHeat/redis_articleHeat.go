package articleHeat

import (
	"fmt"
	"github.com/go-redis/redis"
	"strconv"
	redis2 "studentGrow/dao/redis"
)

// SaveArticleInZSet 将文章加入ZSet
func SaveArticleInZSet(aid uint, listName string) error {
	if err := redis2.RDB.ZAdd(listName, redis.Z{
		Score:  0,
		Member: aid,
	}).Err(); err != nil {
		fmt.Println("添加文章至zset失败 err=", err)
		return err
	}
	return nil
}

// UpdateArticleHeat 更新文章热度值
func UpdateArticleHeat(aid uint, increment float64, listName string) error {
	if err := redis2.RDB.ZIncrBy(listName, increment, strconv.Itoa(int(aid))).Err(); err != nil {
		return err
	}
	return nil
}

// QueryArticleHeat 查询某个文章的热度值
func QueryArticleHeat(aid uint, listName string) error {
	id := strconv.Itoa(int(aid))
	if err := redis2.RDB.ZScore(listName, id).Err(); err != nil {
		return err
	}
	return nil
}

// RangeTopN 获取某日热度排行前n的对象
func RangeTopN(listName string, n int) ([]string, error) {
	list, err := redis2.RDB.ZRevRange(listName, 0, int64(n-1)).Result()
	if err != nil {
		return nil, err
	}
	return list, nil
}
