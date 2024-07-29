package isEmptyData

import "reflect"

// IsEmptyStruct 判断结构体是否为空
func IsEmptyStruct(s interface{}) bool {
	// 使用反射获取结构体的值
	v := reflect.ValueOf(s)

	// 判断是否为结构体
	if v.Kind() != reflect.Struct {
		return false
	}

	// 检查每个字段是否为零值
	for i := 0; i < v.NumField(); i++ {
		if !v.Field(i).IsZero() {
			return false
		}
	}
	return true
}
