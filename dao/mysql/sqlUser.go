package mysql

// 根据用户名查询密码
func SelPassword(username string) (string, error) {
	var sPassword string
	err := DB.Table("users").Select("password").Where("username = ?", username).Scan(&sPassword).Error
	return sPassword, err
}

// 根据用户名查询用户ID
func SelId(username string) (int, error) {
	var id int
	err := DB.Table("users").Select("id").Where("username = ?", username).Scan(&id).Error
	return id, err
}

// 根据用户id查询对应角色id
func SelCasId(id int) (string, error) {
	var code string
	err := DB.Table("users_casbin_rules").Select("casbin_id").Where("user_id = ?", id).Scan(&code).Error
	return code, err
}

// 根据角色id查询角色
func SelRole(id string) (string, error) {
	var role string
	err := DB.Table("casbin_rules").Select("v1").Where("v0 = ?", id).Scan(&role).Error
	return role, err
}
