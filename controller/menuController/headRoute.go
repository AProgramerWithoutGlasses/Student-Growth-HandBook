package menuController

import (
	"github.com/gin-gonic/gin"
	"studentGrow/dao/mysql"
	"studentGrow/pkg/response"
	token2 "studentGrow/utils/token"
)

func HeadRoute(c *gin.Context) {
	token := c.GetHeader("token")
	role, err := token2.GetRole(token)
	username, err := token2.GetUsername(token)
	if err != nil {
		response.ResponseError(c, 400)
		return
	}
	//1.查找用户姓名
	name, err := mysql.SelName(username)
	//2.查找用户头像
	avatar, err := mysql.SelHead(username)
	//3.查找权限下按钮的所有权限标识
	perms, err := mysql.SelPerms(role)
	if err != nil {
		response.ResponseErrorWithMsg(c, 400, "侧边栏获取信息失败")
		return
	}
	data := map[string]any{
		"name":   name,
		"avatar": avatar,
		"perms":  perms,
	}
	response.ResponseSuccess(c, data)
}
