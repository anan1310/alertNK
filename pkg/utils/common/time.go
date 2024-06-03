package common

import (
	"strings"
	"time"
)

// FormatTime 格式化时间
func FormatTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// TimeStampToTime 时间戳转换为格式话时间
func TimeStampToTime(timeStamp int) string {
	timeLayout := "2006-01-02 15:04:05"
	return time.Unix(int64(timeStamp), 0).Format(timeLayout)
}

func TimeParse(timeStr string) time.Time {
	time2, _ := time.Parse("2006-01-02 15:04:05", timeStr)
	return time2
}

// GetCurrentTime 用于解决tar包的命名
func GetCurrentTime() string {
	StartTimeStr := time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05") //把时间戳转换成时间,并格式化为年月日
	StartTimeStr = strings.Replace(StartTimeStr, " ", "", -1)
	StartTimeStr = strings.Replace(StartTimeStr, ":", "", -1)
	StartTimeStr = strings.Replace(StartTimeStr, "-", "", -1)
	return StartTimeStr
}
