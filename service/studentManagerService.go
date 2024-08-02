package service

import "studentGrow/utils/readMessage"

// 判断前端发来的结构体中非空字段的内容
func GetNotEmptyFieldMap(queryStuReqMap map[string]any) (map[string]any, map[string]any) {
	var notEmptyPullMap map[string]any  // 下拉框map
	var notEmptyQueryMap map[string]any // 搜索栏map
	i := 1
	for k, v := range queryStuReqMap {
		if !readMessage.IsZero(v) && i <= 4 {
			notEmptyPullMap[k] = v
		}
		if !readMessage.IsZero(v) && i > 4 {
			notEmptyQueryMap[k] = v
		}
		i++
	}

	return notEmptyPullMap, notEmptyQueryMap
}
