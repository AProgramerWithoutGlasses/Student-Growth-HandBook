package readMessage

import (
	"fmt"
	jsonvalue "github.com/Andrew-M-C/go.jsonvalue"
	"github.com/gin-gonic/gin"
	data "studentGrow/pkg/data"
)

func AnalyzeDataToMyData(c *gin.Context) (d data.AllData, err error) {
	b, _ := c.GetRawData()
	contentType := c.GetHeader("Content-Type")
	d = data.NewAllData()
	switch contentType {
	case "application/json":
		j, err := jsonvalue.Unmarshal(b)
		d.SetJsonValue(j)
		if err != nil {
			fmt.Println("analyzeToMap() utils.readMessage.Unmarshal() err = ", err)
			return d, err
		}
	case "multipart/form-data":
		form, err := c.MultipartForm()
		d.SetFormAllData(form)
		if err != nil {
			fmt.Println("analyzeToMap() utils.readMessage.MultipartForm() err = ", err)
			return d, err
		}
	}
	return d, nil
}
