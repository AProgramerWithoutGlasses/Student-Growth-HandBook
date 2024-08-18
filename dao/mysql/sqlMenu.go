package mysql

// SelFMenu 查询权限下父ID是0的目录及菜单ID
func SelFMenu(role string) ([]int, error) {
	var DId []int
	err := DB.Table("menus").Where("type <> ?", 2).Where("roles LIKE ?", "%"+role+"%").Select("id").Scan(&DId).Error
	if err != nil {
		return nil, err
	}
	return DId, nil
}

// SelValueInt 查询menus菜单所有int类型的数据
func SelValueInt(id int, column string) (int, error) {
	var fId int
	err := DB.Table("menus").Where("id = ?", id).Select(column).Scan(&fId).Error
	if err != nil {
		return 0, err
	}
	return fId, nil
}

// SelValueString 查询menus菜单所有string类型的数据
func SelValueString(id int, column string) (string, error) {
	var value string
	err := DB.Table("menus").Where("id = ?", id).Select(column).Scan(&value).Error
	if err != nil {
		return "", err
	}
	return value, nil
}

// SelParamId 查询params表中菜单所有的参数的id
func SelParamId(id int) ([]int, error) {
	var pid []int
	err := DB.Table("params").Where("menu_id =?", id).Select("id").Scan(&pid).Error
	if err != nil {
		return nil, err
	}
	return pid, nil
}

func SelParamKeyVal(id int) (string, string, error) {
	var key string
	var value string
	err := DB.Table("params").Where("id = ?", id).Select("params_key").Scan(&key).Error
	err = DB.Table("params").Where("id = ?", id).Select("params_value").Scan(&value).Error
	if err != nil {
		return "", "", err
	}
	return key, value, nil
}

func SelIcon(id int) (string, error) {
	var icon string
	err := DB.Table("menus").Where("id = ?", id).Select("icon").Scan(&icon).Error
	if err != nil {
		return "", err
	}
	return icon, nil
}
