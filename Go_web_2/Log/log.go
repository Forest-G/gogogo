package Logmes

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Log = logrus.New()

func Init_Jou() {
	Log.Out = os.Stdout                                                               //日志是“io”。在互斥锁中复制到此。通常将其设置为文件，或将其保留为默认值，即“os”
	Log.SetReportCaller(true)                                                         //设置在输出日志中添加文件名和方法信息
	file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666) //创建文件名为logrus.log的文件用来书写日志,第二个是可读可写，第三个是追加
	if err == nil {
		Log.Out = file
	} else {
		Log.Error(err)
	}
}
