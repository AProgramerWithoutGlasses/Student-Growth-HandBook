package readMessage

import (
	"fmt"
	jsonvalue "github.com/Andrew-M-C/go.jsonvalue"
	"github.com/gin-gonic/gin"
	data "studentGrow/pkg/data"
)

func AnalyzeDataToMyData(c *gin.Context) (d data.AllData, err error) {
	contentType := c.GetHeader("Content-Type")
	d = data.NewAllData()
	switch contentType {
	case "application/json":
		b, _ := c.GetRawData()
		j, err := jsonvalue.Unmarshal(b)
		d.SetJsonValue(j)
		if err != nil {
			fmt.Println("analyzeToMap() utils.readMessage.Unmarshal() err = ", err)
			return d, err
		}
	default:
		fmt.Println(contentType)
		form, err := c.MultipartForm()
		fmt.Println("formfile:", form.File)
		d.SetFormAllData(form)
		if err != nil {
			fmt.Println("analyzeToMap() utils.readMessage.MultipartForm() err = ", err)
			return d, err
		}
	}
	return d, nil
}
