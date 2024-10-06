package NotificationPush

import (
	"studentGrow/dao/mysql"
	"studentGrow/models/gorm_model"
)

func BuildLikeNotification(username, tarUsername string, objId, likeType int) (*gorm_model.Notification, error) {
	ownUserId, err := mysql.QueryUserIdByUsername(username)
	if err != nil {
		return nil, err
	}
	tarUserId, err := mysql.QueryUserIdByUsername(tarUsername)
	if err != nil {
		return nil, err
	}

	notification := gorm_model.Notification{
		TarUserId:  uint(tarUserId),
		OwnUserId:  uint(ownUserId),
		NoticeType: 0,
		SuperType:  likeType,
		SuperId:    objId,
		IsRead:     false,
		Status:     0,
	}

	return &notification, nil
}
