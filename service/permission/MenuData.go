package service

import (
	"studentGrow/dao/mysql"
	"studentGrow/models"
)

// MenuIdClass 班级管理员所有的菜单
func MenuIdClass() ([]models.Sidebar, error) {
	//返回前端的切片
	var Menu []models.Sidebar
	//1.查询权限下所有的菜单和目录
	DId, err := mysql.SelFMenu("class")
	if err != nil {
		return nil, err
	}
	for _, id := range DId {
		var params []models.Params
		//1.查询父id
		Fid, err := mysql.SelValueInt(id, "parent_id")
		//2.查询路由路径
		path, err := mysql.SelValueString(id, "path")
		//3.查询component
		component, err := mysql.SelValueString(id, "component")
		//4.查询跳转路径
		redirect, err := mysql.SelValueString(id, "redirect")
		//5.查询名字
		name, err := mysql.SelValueString(id, "name")
		//6.查询是否可见
		visible, err := mysql.SelValueInt(id, "visible")
		//7.查询菜单下所属参数的所有id
		pids, err := mysql.SelParamId(id)
		for _, pid := range pids {
			key, value, err := mysql.SelParamKeyVal(pid)
			if err != nil {
				return nil, err
			}
			param := models.Params{
				ParamsKey:   key,
				ParamsValue: value,
			}
			params = append(params, param)
		}
		if err != nil {
			return nil, err
		}
		mesa := models.Message{
			Name:    name,
			Visible: visible,
		}
		sidebar := models.Sidebar{
			ParentId:  Fid,
			Path:      path,
			Component: component,
			Redirect:  redirect,
			Meta:      mesa,
			Params:    params,
		}
		Menu = append(Menu, sidebar)
	}
	return Menu, nil
}
