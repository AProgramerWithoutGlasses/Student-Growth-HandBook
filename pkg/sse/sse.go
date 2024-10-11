package sse

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"studentGrow/models/gorm_model"
	"sync"
)

var isInitChannelsMap = false
var ChannelsMap sync.Map

func AddChannel(userId int) {
	if !isInitChannelsMap {
		ChannelsMap = sync.Map{}
		isInitChannelsMap = true
	}
	newChannel := make(chan string, 10)
	ChannelsMap.Store(userId, newChannel)
	fmt.Println("Build SSE connection for user = ", userId)
}

func BuildNotificationChannel(userId int, c *gin.Context) error {
	AddChannel(userId)
	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")

	w := c.Writer
	flusher, _ := w.(http.Flusher)

	curChan, _ := ChannelsMap.Load(userId)

	// 监听客户端通道是否被关闭
	closeNotify := c.Request.Context().Done()

	go func() {
		<-closeNotify
		close(curChan.(chan string))
		ChannelsMap.Delete(userId)
		fmt.Println("SSE close for user = ", userId)
		return
	}()

	for msg := range curChan.(chan string) {
		_, err := fmt.Fprintf(w, "data:%s\n\n", msg)
		if err != nil {
			return err
		}
		flusher.Flush()
	}
	return nil
}

// SendInterNotification 互动消息推送
func SendInterNotification(n gorm_model.InterNotification) {
	fmt.Println("Send interNotification to user = ", n.TarUserId)
	// 若对方用户不在线，则不推送消息
	if _, ok := ChannelsMap.Load(n.TarUserId); !ok {
		return
	}

	msg, err := json.Marshal(n)
	if err != nil {
		return
	}
	ChannelsMap.Range(func(key, value any) bool {
		k := key.(uint)
		if k == n.TarUserId {
			channel := value.(chan string)
			channel <- string(msg)
		}
		return true
	})
}

// SendSysNotification 广播消息推送
func SendSysNotification(n gorm_model.SysNotification) {
	fmt.Println("Send sysNotification is user = ", n.OwnUserId)
	msg, err := json.Marshal(n)
	if err != nil {
		return
	}
	ChannelsMap.Range(func(key, value any) bool {
		channel := value.(chan string)
		channel <- string(msg)
		return true
	})
}
