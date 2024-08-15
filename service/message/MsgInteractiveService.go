package message

import (
	"go.uber.org/zap"
	"studentGrow/dao/mysql"
	"studentGrow/models/gorm_model"
)

// GetSystemMsgService 获取系统消息通知
func GetSystemMsgService(limit, page int, username string) ([]gorm_model.MsgRecord, int, error) {
	msg, err := mysql.QuerySystemMsg(page, limit, username)
	if err != nil {
		zap.L().Error("GetSystemMsgService() service.message.QuerySystemMsg", zap.Error(err))
		return nil, -1, err
	}

	// 查询未读消息条数
	count, err := mysql.QueryUnreadSystemMsg(username)
	if err != nil {
		zap.L().Error("GetSystemMsgService() service.message.QueryUnreadSystemMsg", zap.Error(err))
		return nil, -1, err
	}

	return msg, count, nil
}

// GetManagerMsgService 获取管理员消息
func GetManagerMsgService(limit, page int, username string) ([]gorm_model.MsgRecord, int, error) {
	msg, err := mysql.QueryManagerMsg(page, limit, username)
	if err != nil {
		zap.L().Error("GetManagerMsgService() service.message.QueryManagerMsg", zap.Error(err))
		return nil, -1, err
	}

	// 查询未读消息条数
	count, err := mysql.QueryUnreadManagerMsg(username)
	if err != nil {
		zap.L().Error("GetManagerMsgService() service.message.QueryUnreadManagerMsg", zap.Error(err))
		return nil, -1, err
	}

	return msg, count, nil
}
