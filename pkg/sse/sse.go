package sse

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"studentGrow/models/gorm_model"
	res "studentGrow/pkg/response"
	"sync"
)

var isInitChannelsMap = false
var ChannelsMap sync.Map

func AddChannel(userId int) {
	if !isInitChannelsMap {
		ChannelsMap = sync.Map{}
		isInitChannelsMap = true
	}
	newChannel := make(chan gorm_model.Notification)
	ChannelsMap.Store(userId, newChannel)
	fmt.Println("Build SSE connection for user = ", userId)
}

func BuildNotificationChannel(userId int, c *gin.Context) {
	AddChannel(userId)
	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")

	curChan, _ := ChannelsMap.Load(userId)

	// 监听客户端通道是否被关闭
	closeNotify := c.Request.Context().Done()

	go func() {
		<-closeNotify
		close(curChan.(chan gorm_model.Notification))
		ChannelsMap.Delete(userId)
		fmt.Println("SSE close for user = ", userId)
		return
	}()

	for msg := range curChan.(chan gorm_model.Notification) {
		res.ResponseSuccess(c, msg)
	}
}

func SendNotification(n gorm_model.Notification) {
	fmt.Println("Send notification to user = ", n.TarUserId)
	ChannelsMap.Range(func(key, value any) bool {
		k := key.(int)
		if k == int(n.TarUserId) {
			channel := value.(chan gorm_model.Notification)
			channel <- n
		}
		return true
	})
}
