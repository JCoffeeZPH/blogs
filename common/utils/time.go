package utils

import (
	"time"
)

func NowTime() uint32 {
	return uint32(time.Now().Unix())
}

func TimeFormat(timestamp uint32) string {
	return time.Unix(int64(timestamp), 0).Format("2006-01-02 15:04:05")
}

// GetLast7DayTime 获取过去七天凌晨一点时间戳
func GetLast7DayTime() []uint32 {
	t := time.Now()
	datetime := make([]uint32, 0)
	for i := 0; i < 7; i++ {
		ts := t.AddDate(0, 0, -i)
		addTime := time.Date(ts.Year(), ts.Month(), ts.Day(), 1, 0, 0, 0, t.Location())
		timestamp := uint32(addTime.Unix())
		datetime = append(datetime, timestamp)
	}
	return datetime
}
