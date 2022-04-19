package plugins

import (
	"time"
)

func GetDate () string{
	jst, _ := time.LoadLocation("Asia/Tokyo")
	return time.Now().In(jst).Format("2006-01-02-15:04:05")
}