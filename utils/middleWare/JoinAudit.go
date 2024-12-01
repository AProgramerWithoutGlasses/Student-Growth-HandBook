package middleWare

import (
	"github.com/gin-gonic/gin"
	"studentGrow/service/JoinAudit"
	"studentGrow/service/userService"
	token2 "studentGrow/utils/token"
)

// 入团申请权限判断中间件
func JoinAuditMiddle() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := token2.NewToken(c)
		user, _ := token.GetUser()
		isAdmin := userService.BVerifyExit(user.Username)
		isJoinAudit := JoinAudit.GetUserJoinAuditRoel(c)
		if !isJoinAudit && !isAdmin {
			c.JSON(500, gin.H{"code": 500, "msg": "驳回"})
			c.Abort()
			return
		}
		c.Next()
	}
}
