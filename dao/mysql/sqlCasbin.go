package mysql

import "fmt"

// 通过请求方式查询菜单id
func SelMenuId(requestUrl, requestMethod string) (string, error) {
	var menuId string
	err := DB.Table("menus").Select("id").Where("request_url = ?", requestUrl).Where("request_method = ?", requestMethod).Scan(&menuId).Error
	fmt.Println(menuId)
	return menuId, err
}
