package userService

import "time"

// IntervalInDays 获取时间间隔的天数
func IntervalInDays(t time.Time) int {
	now := time.Now()
	delta := now.Sub(t)

	// 如果时间差大于或等于24小时，计算天数
	if delta >= 24*time.Hour {
		return int(delta.Hours() / 24)
	}
	// 如果时间差小于一天，返回0
	return 0
}
