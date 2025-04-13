package userService

import (
	"errors"
	"fmt"
	redises "github.com/go-redis/redis"
	"log"
	"studentGrow/dao/mysql"
	"studentGrow/dao/redis"
	"time"
)

// BVerify 后台登录验证
func BVerify(username, password string) bool {
	number, err := mysql.SelPassword(username, password)
	if err != nil || number != 1 {
		return false
	}
	return true
}

// BVerifyExit 验证用户是否为管理员
func BVerifyExit(username string) bool {
	number, err := mysql.SelIfexit(username)
	if err != nil || number != 1 {
		return false
	}
	return true
}

// BVerifyBan 验证用户是否被封
func BVerifyBan(username string) (bool, error) {
	ban, err := mysql.SelBan(username)
	if err != nil {
		return false, err
	}
	return ban, nil
}

// CheckUserLock 1.检查账户是否被锁定
func CheckUserLock(blacklistKey string) bool {
	if redis.RDB.Exists(blacklistKey).Val() == 1 { //锁定
		return true
	}
	return false
}

// 查找登录失败次数
func selectFailure(failedAttemptsKey string) (int, error) {
	failedAttempts, err := redis.RDB.Get(failedAttemptsKey).Int()
	if errors.Is(err, redises.Nil) { //说明不存在
		return 0, nil
	} else if err != nil {
		return 0, err
	}
	return failedAttempts, nil
}

// 进入黑名单
func lockAccount(blacklistKey string) {
	redis.RDB.Set(blacklistKey, "lock", 1*time.Hour)
}

// 进黑名单后删除登录失败
func delKeyFailure(failedAttemptsKey string) {
	redis.RDB.Del(failedAttemptsKey)
}

// AddFailure Add 登录失败添加1
func AddFailure(failedAttemptsKey, blacklistKey string) (string, error) {
	amount, err := selectFailure(failedAttemptsKey)
	if err != nil {
		log.Println("userService AddFailure err =", err)
		return "", err
	}
	if amount == 5 { //可以锁定了
		lockAccount(blacklistKey)
		delKeyFailure(failedAttemptsKey)
		return "账户已锁定请稍后再试", nil
	} else if amount < 5 {
		redis.RDB.Incr(failedAttemptsKey)
		redis.RDB.Expire(failedAttemptsKey, 1*time.Hour)
		return fmt.Sprintf("密码错误,还有%v次尝试机会", 4-amount), nil
	}
	return "", nil
}
