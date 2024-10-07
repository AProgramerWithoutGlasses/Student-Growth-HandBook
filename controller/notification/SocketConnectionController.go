package notification

import (
	"github.com/gin-gonic/gin"
	"studentGrow/pkg/sse"
)

// SocketConnectionController 向客户端建立sse连接
func SocketConnectionController(c *gin.Context) {
	in := struct {
		UserId int `json:"user_id"`
	}{}

	err := c.ShouldBindJSON(&in)
	if err != nil {
		return
	}

	err = sse.BuildNotificationChannel(in.UserId, c)
	if err != nil {
		return
	}
}
