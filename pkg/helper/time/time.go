package utiltime

import (
	"fmt"
	"time"
)

const LAYOUT = "2006-01-02 15:04:05"

func CurrentDateTime() string {
	// 设置北京时间（GMT+8）时区
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println("错误: 无法加载时区信息:", err)
		return ""
	}
	return time.Now().In(location).Format(LAYOUT)
}

func GetNowTimestamp() int64 {
	// TODO 注意时区问题
	return time.Now().Unix()
}

func GetNowDateTime() string {
	beijingLocation, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return ""
	}
	// 将UTC时间转换为北京时间
	currentTimeBeijing := time.Now().UTC().In(beijingLocation)
	// 格式化时间为"2006-01-02 15:04:05"格式
	return currentTimeBeijing.Format("2006-01-02 15:04:05")
}

func ParseTo3339Time(rfc3339Time string) string {
	t, err := time.Parse(time.RFC3339, rfc3339Time)
	if err != nil {
		return ""
	}
	return t.Format(LAYOUT)
}
