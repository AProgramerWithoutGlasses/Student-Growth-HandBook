package userService

import (
	"studentGrow/dao/mysql"
)

// 后台登录验证
func BVerify(username, password string) bool {
	number, err := mysql.SelPassword(username, password)
	if err != nil || number != 1 {
		return false
	}
	return true
}

// 验证用户是否为管理员
func BVerifyExit(username string) bool {
	number, err := mysql.SelIfexit(username)
	if err != nil || number != 1 {
		return false
	}
	return true
}
