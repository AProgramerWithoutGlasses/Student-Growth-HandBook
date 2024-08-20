package mysql

import (
	"fmt"
	"regexp"
	"strings"
)

func SelMenuId(requestUrl, requestMethod string) (string, error) {
	var menuId string
	//获取最后一个/后的信息
	re := regexp.MustCompile("/([^/]+)$")
	matches := re.FindStringSubmatch(requestUrl)
	//判断是不是角色
	ok, err := SelMRole(matches[1])
	if ok {
		requestUrl = strings.Replace(requestUrl, matches[0], "", 1)
	}
	// 使用模糊查询，例如：查找所有以requestUrl开头的记录
	err = DB.Table("menus").Select("id").Where("deleted_at IS NULL").Where("request_url = ?", requestUrl).Where("request_method = ? ", requestMethod).Scan(&menuId).Error
	fmt.Println(menuId)
	return menuId, err
}
