package AllData

import (
	"errors"
	jsonvalue "github.com/Andrew-M-C/go.jsonvalue"
	"mime/multipart"
)

type AllData struct {
	J        *jsonvalue.V
	FormData *multipart.Form
}

// NewAllData 创建对象
func NewAllData() AllData {
	return AllData{
		nil,
		nil,
	}
}

// GetJsonValue 获取*jsonvalue.V对象
func (allData *AllData) GetJsonValue() (*jsonvalue.V, error) {
	if allData.J == nil {
		return nil, errors.New("数据类型错误")
	} else {
		return allData.J, nil
	}
}

// GetFormAllData 获取*multipart.Form对象
func (allData *AllData) GetFormAllData() (*multipart.Form, error) {
	if allData.J == nil {
		return nil, errors.New("数据类型错误")
	} else {
		return allData.FormData, nil
	}
}

// SetJsonValue 设置*jsonvalue.V对象
func (allData *AllData) SetJsonValue(j *jsonvalue.V) {
	allData.J = j
}

// SetFormAllData 设置*multipart.Form对象
func (allData *AllData) SetFormAllData(formData *multipart.Form) {
	allData.FormData = formData
}
