package mysql

// SelRoleMessage 获取角色id,名称,状态码
func SelRoleMessage(id int) (string, string, error) {
	var role string
	var code string
	err := DB.Table("casbin_rule").Where("id = ?", id).Select("v1").Scan(&role).Error
	err = DB.Table("casbin_rule").Where("id = ?", id).Select("v0").Scan(&code).Error
	if err != nil {
		return "", "", err
	}
	return role, code, nil
}

// SelRoleId 查询代表角色的id切片
func SelRoleId() ([]int, error) {
	var id []int
	err := DB.Table("casbin_rule").Where("ptype = ?", "g").Select("id").Scan(&id).Error
	if err != nil {
		return nil, err
	}
	return id, nil
}
