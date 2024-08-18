package mysql

import "fmt"

// SelMenuId 通过请求方式查询菜单id
//func SelMenuId(requestUrl, requestMethod string) (string, error) {
//	var menuId string
//	err := DB.Table("menus").Select("id").Where("request_url = ?", requestUrl).Where("request_method = ?", requestMethod).Scan(&menuId).Error
//	fmt.Println(menuId)
//	return menuId, err
//}

func SelMenuId(requestUrl, requestMethod string) (string, error) {
	var menuId string
	// 使用模糊查询，例如：查找所有以requestUrl开头的记录
	err := DB.Table("menus").Select("id").Where("request_url LIKE ?", requestUrl+"%").Where("request_method LIKE ?", requestMethod+"%").Scan(&menuId).Error
	fmt.Println(menuId)
	return menuId, err
}
