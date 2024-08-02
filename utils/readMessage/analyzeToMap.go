package readMessage

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
)

func AnalyzeDataToMap(c *gin.Context) (m map[string]any, err error) {
	b, _ := c.GetRawData()
	contentType := c.GetHeader("Content-Type")
	switch contentType {
	case "application/json":
		err = json.Unmarshal(b, &m)
		if err != nil {
			fmt.Println("analyzeToMap() utils.readMessage.Unmarshal() err = ", err)
			return nil, err
		}
	case "multipart/form-data":
		//form, _ := c.MultipartForm()
	}
	return m, nil
}
