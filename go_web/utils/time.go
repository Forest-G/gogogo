package utils //获取时间的工具 格式2006/01/02 15:04:05
import (
	"time"
)

func CreatTime() (hour string) {
	now := time.Now()
	hour = now.Format("2006/01/02 15:04:05")
	return
}
