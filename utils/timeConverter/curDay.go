package timeConverter

import "time"

func GetCurDay() string {
	now := time.Now()
	curDay := now.Format("2006-01-02")
	return curDay
}
