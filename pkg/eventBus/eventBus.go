package eventBus

import (
	"github.com/asaskevich/EventBus"
	"studentGrow/service/notificationPush"
)

var Bus EventBus.Bus

func InitEventBus() error {
	Bus = EventBus.New()

	// 点赞消息订阅
	err := Bus.Subscribe("interactive:like", NotificationPush.LikeNotificationPushService)
	if err != nil {
		return err
	}

	return nil
}
