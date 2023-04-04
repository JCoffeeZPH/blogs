package utils

import "time"

func NowTime() uint32 {
	return uint32(time.Now().Unix())
}

func TimeFormat(timestamp uint32) string {
	return time.Unix(int64(timestamp), 0).Format("2006-01-02 15:04:05")
}
