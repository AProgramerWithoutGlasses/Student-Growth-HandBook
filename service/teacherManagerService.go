package service

import (
	"errors"
	"strconv"
	"studentGrow/dao/mysql"
	"studentGrow/models/gorm_model"
	"studentGrow/models/jrx_model"
)

// 查询老师
func QueryTeacher(queryTeacherParama jrx_model.QueryTeacherParamStruct) ([]gorm_model.User, int64, error) {
	// 获取老师总数量
	allTeacherCount, err := mysql.GetAllUserCount("老师")
	if err != nil {
		return nil, allTeacherCount, err
	}

	// 获取老师列表
	queryTeacherSql := GetQueryTeacherSql(queryTeacherParama)
	teacherList, err := mysql.GetTeacherList(queryTeacherSql)
	if err != nil {
		return nil, allTeacherCount, err
	}

	// 获取老师属于哪一级管理员

	return teacherList, allTeacherCount, err
}

// 获得查询老师的sql语句
func GetQueryTeacherSql(queryTeacherParama jrx_model.QueryTeacherParamStruct) string {
	querySql := `Select name, username, password, gender, is_manager, ban, from users where identity = '老师'`

	if queryTeacherParama.Gender != "" {
		querySql = querySql + " and gender = '" + queryTeacherParama.Gender + "'"
	}

	if queryTeacherParama.Ban != nil {
		querySql = querySql + " and ban = '" + strconv.FormatBool(*queryTeacherParama.Ban) + "'"
	}

	if len(queryTeacherParama.SearchSelect) > 0 {
		querySql = querySql + " and " + queryTeacherParama.SearchSelect + " like '%" + queryTeacherParama.SearchMessage + "%'"
	}

	// limit 分页查询语句的拼接
	querySql = querySql + " ORDER BY name ASC" + " limit " + strconv.Itoa(queryTeacherParama.Limit) + " offset " + strconv.Itoa((queryTeacherParama.Page-1)*queryTeacherParama.Limit)

	return querySql

}

func GetManagerType(username string) (string, error) {
	CId, err := mysql.GetManagerCId(username)
	if err != nil {
		return "", err
	}

	var managerType string

	switch CId {
	case "2", "3", "4", "5":
		managerType = "年级管理员"
	case "6":
		managerType = "班级管理员"
	case "1":
		managerType = "院级管理员"
	case "0":
		managerType = "超级管理员"
	default:
		managerType = "无"
	}

	return managerType, err
}

// 处理管理员
func SetStuManagerService(username string, ManagerType string, year string) error {
	// 判断这个用户是不是管理员
	isManager, err := mysql.GetIsManagerByUsername(username)
	if err != nil {
		return err
	}

	var casbinCid string
	switch ManagerType {
	case "班级管理员":
		casbinCid = "6"
	case "年级管理员":
		switch year {
		case "2024":
			casbinCid = "2"

		case "2023":
			casbinCid = "3"

		case "2022":
			casbinCid = "4"

		case "2021":
			casbinCid = "5"
		}

	case "取消管理员":
		if isManager {
			err := mysql.CancelStuManager(username, casbinCid)
			return err
		} else {
			return errors.New("该用户不为管理员")
		}

	}

	if !isManager { // 不是管理员
		err := mysql.SetStuManager(username, casbinCid)
		if err != nil {
			return err
		}

	} else { // 是管理员
		existedCasbinCid, err := mysql.GetManagerCId(username)
		if err != nil {
			return err
		}

		if casbinCid == existedCasbinCid {
			return errors.New("该用户已是" + ManagerType)
		}

		err = mysql.ChangeStuManager(username, casbinCid)
		if err != nil {
			return err
		}
	}

	return err
}
