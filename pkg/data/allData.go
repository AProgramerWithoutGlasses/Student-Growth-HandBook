package AllData

import (
	"errors"
	"fmt"
	jsonvalue "github.com/Andrew-M-C/go.jsonvalue"
	"mime/multipart"
)

type AllData struct {
	j        *jsonvalue.V
	formData *multipart.Form
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
	if allData.j == nil {
		return nil, errors.New("数据类型错误")
	} else {
		return allData.j, nil
	}
}

// GetFormAllData 获取*multipart.Form对象
func (allData *AllData) GetFormAllData() (*multipart.Form, error) {
	fmt.Println(allData.formData)
	if allData.formData == nil {
		return nil, errors.New("数据类型错误")
	} else {
		return allData.formData, nil
	}
}

// SetJsonValue 设置*jsonvalue.V对象
func (allData *AllData) SetJsonValue(j *jsonvalue.V) {
	allData.j = j
}

// SetFormAllData 设置*multipart.Form对象
func (allData *AllData) SetFormAllData(formData *multipart.Form) {
	allData.formData = formData
}
