package mysql

// SelPassword 根据用户名和密码查询用户是否存在
func SelPassword(username, password string) (int64, error) {
	var number int64
	err := DB.Table("users").Select("password").Where("username = ?", username).Where("password = ?", password).Count(&number).Error
	return number, err
}

// SelCasId 根据用户id查询对应角色id
func SelCasId(username string) (string, error) {
	var code string
	err := DB.Table("user_casbin_rules").Select("casbin_cid").Where("c_username = ?", username).Scan(&code).Error
	return code, err
}

// SelRole 根据角色id查询角色
func SelRole(id string) (string, error) {
	var role string
	err := DB.Table("casbin_rule").Select("v1").Where("v0 = ?", id).Scan(&role).Error
	return role, err
}

// SelClass 根据角色获取班级
func SelClass(username string) (string, error) {
	var class string
	err := DB.Table("users").Select("class").Where("username = ?", username).Scan(&class).Error
	if err != nil {
		return "", err
	}
	return class, nil
}

// SelIfexit 查找用户是否存在
func SelIfexit(username string) (int64, error) {
	var number int64
	err := DB.Table("user_casbin_rules").Where("c_username = ?", username).Count(&number).Error
	if err != nil {
		return 0, err
	}
	return number, nil
}

// SelHead 查找用户头像
func SelHead(username string) (string, error) {
	var headshot string
	err := DB.Table("users").Where("username = ?", username).Select("head_shot").Scan(&headshot).Error
	if err != nil {
		return "", err
	}
	return headshot, nil
}
