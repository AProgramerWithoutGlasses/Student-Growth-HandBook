package redis

import (
	"go.uber.org/zap"
	"strconv"
)

/*
主要解决广播类消息，用户是否确认的问题
*/

const mgrAck = "manager_ack" // set存储用户已读消息列表
const sysAck = "system_ack"

// AckManagerNotification 用户确认管理员消息
func AckManagerNotification(userId, msgId int) error {
	_, err := RDB.SAdd(mgrAck+strconv.Itoa(userId), msgId).Result()
	if err != nil {
		zap.L().Error("AckManagerNotification() dao.redis.redis_ack.SAdd err=", zap.Error(err))
		return err
	}
	return nil
}

// AckSystemNotification 用户确认广播类系统消息
func AckSystemNotification(userId, msgId int) error {
	res, err := RDB.SIsMember(sysAck+strconv.Itoa(userId), msgId).Result()
	if err != nil {
		zap.L().Error("AckSystemNotification() dao.redis.redis_ack.SIsMember err=", zap.Error(err))
		return err
	}
	// 若确认该消息
	if !res {
		_, err := RDB.SAdd(sysAck+strconv.Itoa(userId), msgId).Result()
		if err != nil {
			zap.L().Error("AckSystemNotification() dao.redis.redis_ack.SAdd err=", zap.Error(err))
			return err
		}
	}
	return nil
}

// IsUserAckedManagerNotification 检查用户是否已经确认管理员消息
func IsUserAckedManagerNotification(userId, msgId int) (bool, error) {
	result, err := RDB.SIsMember(mgrAck+strconv.Itoa(userId), msgId).Result()
	if err != nil {
		zap.L().Error("IsUserAckedManagerNotification() dao.redis.redis_ack.SIsMember err=", zap.Error(err))
		return false, err
	}
	return result, nil
}

// IsUserAckedSystemNotification 检查用户是否已经确认系统消息
func IsUserAckedSystemNotification(userId, msgId int) (bool, error) {
	result, err := RDB.SIsMember(sysAck+strconv.Itoa(userId), msgId).Result()
	if err != nil {
		zap.L().Error("IsUserAckedSystemNotification() dao.redis.redis_ack.SIsMember err=", zap.Error(err))
		return false, err
	}
	return result, nil
}

// GetUserAckedSystemNum 获取用户已读系统消息数量
func GetUserAckedSystemNum(userId int) (int, error) {
	result, err := RDB.SCard(sysAck + strconv.Itoa(userId)).Result()
	if err != nil {
		zap.L().Error("GetUserAckedSystemNum() dao.redis.redis_ack.SCard err=", zap.Error(err))
		return -1, err
	}
	return int(result), nil
}

// GetUserAckedManagerNum 获取用户已读管理员消息数量
func GetUserAckedManagerNum(userId int) (int, error) {
	result, err := RDB.SCard(mgrAck + strconv.Itoa(userId)).Result()
	if err != nil {
		zap.L().Error("GetUserAckedManagerNum() dao.redis.redis_ack.SCard err=", zap.Error(err))
		return -1, err
	}
	return int(result), nil
}
