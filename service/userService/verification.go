package userService

import (
	"fmt"
	"studentGrow/dao/mysql"
)

// 后台登录验证
func BVerify(username, password string) bool {
	//查询数据库中相应账号对应的密码
	sPassword, err := mysql.SelPassword(username)
	if err != nil {
		fmt.Println("BVerify mysql.SelPassword() err")
		return false
	}
	if sPassword == password {
		return true
	}
	return false
}
