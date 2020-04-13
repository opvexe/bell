package time

import "time"

// 时间格式化
func TimeFormat(time time.Time, layout string) string {
	return time.Format(layout)
}
