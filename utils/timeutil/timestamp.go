package timeutil

import "time"

// 时间字符串
const (
	PureNumLayout = "20060102150405"
	NormalLayout  = "2006-01-02 15:04:05"
	RFC3339       = time.RFC3339
	RFC822        = time.RFC822
)

var (
	LocalTimeZone       = time.Local
	UTCTimeZone, _      = time.LoadLocation("UTC")
	ShanghaiTimeZone, _ = time.LoadLocation("Asia/Shanghai")
)

func ParseTimeStr(standard, timeStr string) (time.Time, error) {
	return time.Parse(standard, timeStr)
}

func ParseTimeStrInLocation(layout, timeStr string, localTimeZone *time.Location) (time.Time, error) {
	return time.ParseInLocation(layout, timeStr, localTimeZone)
}
