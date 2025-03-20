package ttutils

import (
	"strconv"
	"strings"
	"time"
)

// GetFutureTimestamp 获取未来时间戳, 单位为小时  1726766060。
// Args: hours int 当前时间的未来/过去小时的数量。
// Return: int 时间戳。
// e.g.: Args 2, Returns 1726766060。
func GetFutureTimestamp(hours int) int {
	currentTime := time.Now()
	futureTime := currentTime.Add(time.Duration(hours) * time.Hour)
	return int(futureTime.Unix())
}

// GetNowTime 获取当前时间 2024-09-29 16:42:19
// Return: string 当前时间。
// e.g.: Returns 2024-09-29 16:42:19。
func GetNowTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// GetNowTimestamp 获取当前时间戳 1726736060
// Return: string 当前时间戳。
// e.g.: Returns 1726736060。
func GetNowTimestamp() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}

// GetTimestampByTime 传入时间字符串，如，获取时间戳
// Args: timeStr string 时间字符串。
// Return: int 时间戳, error 错误。
// e.g.: Args 2006-01-02 15:04:05, Returns 1726736060,nil。
func GetTimestampByTime(timeStr string) (int, error) {
	timeStr = strings.ReplaceAll(timeStr, "T", " ")
	if strings.Contains(timeStr, ".") {
		timeStr = strings.Split(timeStr, ".")[0]
	} else {
		timeStr = strings.ReplaceAll(timeStr, "000", " ")
		timeStr = strings.ReplaceAll(timeStr, "Z", " ")
	}
	// 设定时间格式
	layout := "2006-01-02 15:04:05"
	// 解析时间字符串
	parsedTime, err := time.Parse(layout, timeStr)
	if err != nil {
		return 0, err
	}
	// 转换为 Unix 时间戳
	return int(parsedTime.Unix()), nil
}
