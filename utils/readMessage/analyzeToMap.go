package readMessage

import (
	"fmt"
	"github.com/goccy/go-json"
)

func AnalyzeToMap(b []byte) (m map[string]string) {
	err := json.Unmarshal(b, &m)

	if err != nil {
		fmt.Println("analyzeToMap() utils.readMessage.Unmarshal() err = ", err)
		return
	}
	return m
}
