package mysql

// 根据角色查询对应的角色状态码
func SelCode(role string) (string, error) {
	var cId string
	err := DB.Table("casbin_rule").Select("v0").Where("v1 = ?", role).Scan(&cId).Error
	return cId, err
}

// 通过请求方式查询菜单id
func SelMenuId(requestUrl, requestMethod string) (string, error) {
	var menuId string
	err := DB.Table("menus").Select("id").Where("request_url = ?", requestUrl).Where("request_method = ?", requestMethod).Scan(&menuId).Error
	return menuId, err
}
