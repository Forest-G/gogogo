package utils

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Log = logrus.New()

func init() {
	// 设置日志输出为os.Stdout
	Log.Out = os.Stdout
	Log.SetReportCaller(true)
	// 可以设置像文件等任意`io.Writer`类型作为日志输出
	file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		Log.Out = file
	} else {
		Log.Error("日志设置输出位置没有配置成功")
	}
}

func KKK() {
	// 设置日志输出为os.Stdout
	Log.Out = os.Stdout
	Log.SetReportCaller(true)
	// 可以设置像文件等任意`io.Writer`类型作为日志输出
	file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		Log.Out = file
	} else {
		Log.Error("日志设置输出位置没有配置成功")
	}
}
