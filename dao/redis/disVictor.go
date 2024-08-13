package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

// UpdateVictor 记录用户登录
func UpdateVictor(username string) {
	today := time.Now().Format("20060102")
	key := "user_victor_status" + today
	//将用户登陆状态设为true
	err := RDB.HSet(key, username, "true").Err()
	if err != nil {
		fmt.Println("UpdateVictor HSet err", err)
	}
	//设置过期时间为72小时
	err = RDB.Expire(key, 72*time.Hour).Err()
	if err != nil {
		fmt.Println("UpdateVictor Expire err", err)
	}
}

// IfVictor 查询用户是否登录
func IfVictor(username string, data string) bool {
	key := "user_victor_status" + data

	//使用HGet获取用户登录状态
	loginStatus, err := RDB.HGet(key, username).Result()

	if err != nil && err != redis.Nil {
		fmt.Println("IfVictor HGet err", err)
		return false
	}
	//为false则返回false
	if loginStatus != "true" {
		return false
	}
	return true
}
