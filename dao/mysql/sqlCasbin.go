package mysql

import "fmt"

func SelMenuId(role, requestUrl, requestMethod string) (string, error) {
	var menuId string
	requestUrl = requestUrl[0:7]
	// 使用模糊查询，例如：查找所有以requestUrl开头的记录
	err := DB.Table("menus").Select("id").Where("roles LIKE ?", role).Where("request_url LIKE ?", requestUrl+"%").Where("request_method LIKE ?", "%"+requestMethod+"%").Scan(&menuId).Error
	fmt.Println(menuId)
	return menuId, err
}
