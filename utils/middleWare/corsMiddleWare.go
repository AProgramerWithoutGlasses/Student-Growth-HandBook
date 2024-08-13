package middleWare

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*") // 允许所有来源
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Accept")
		c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")

		//如果是预检请求（OPTIONS），直接返回200
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}
