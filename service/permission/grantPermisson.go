package service

// 对用户权限进行修改
//func UpdateRolePermissions(username, role string, newFunctions []string) (bool, error) {
//		//
//		// 删除旧权限
//		if err := tx.Where("v0 = ?", role).Delete(gorm_model.CasbinRule{}).Error; err != nil {
//			tx.Rollback()
//			return false, err
//		}
//
//		// 创建新权限
//		for _, functionName := range newFunctions {
//			var menuId int
//			if err := tx.Table("menus").Select("id").Where("name = ?", functionName).Scan(&menuId).Error; err != nil {
//				tx.Rollback()
//				return false, err
//			}
//			casbinRule := gorm_model.CasbinRule{V0: role, V1: strconv.Itoa(menuId)}
//			if err := tx.Create(&casbinRule).Error; err != nil {
//				tx.Rollback()
//				return false, err
//			}
//		}
//
//		// 提交事务
//		if err := tx.Commit().Error; err != nil {
//			return false, err
//		}
//
//		return true, nil
//	}
//}

// 添加管理员
func AddAdmin() {

}
