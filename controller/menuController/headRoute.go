package menuController

import (
	"github.com/gin-gonic/gin"
	"studentGrow/dao/mysql"
	"studentGrow/models"
	"studentGrow/pkg/response"
)

func HeadRoute(c *gin.Context) {
	claim, _ := c.Get("claim")
	username := claim.(*models.Claims).Username
	role := claim.(*models.Claims).Role
	//1.查找用户姓名
	name, err := mysql.SelName(username)
	//2.查找用户头像
	avatar, err := mysql.SelHead(username)
	//3.查找权限下按钮的所有权限标识
	if role == "grade1" || role == "grade2" || role == "grade3" || role == "grade4" {
		role = "grade"
	}
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
