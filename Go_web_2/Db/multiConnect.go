package utils

import (
	"database/sql"
	Logmes "go_web_2/Log" //自定义包，日志
	"go_web_2/Public"
	"time"
)

func MultiConnect() {
	ticker := time.NewTicker(2 * time.Minute) //每隔2分钟执行一次,周期性的触发一个事件，通过Ticker本身提供的管道将事件传递出去。
	go func() {                               //参数列表，匿名函数创建goroutine
		for {
			<-ticker.C                                                                                                                                   //管道，通过ticker本身提供的管道开启                                                                                                                                //通过管道进行                                                                                                              //通过管道执行                                                                                                                                                                                                                    //时间到了一分钟后
			Db, err := sql.Open("mysql", "root:"+Public.ReadConfiguration("data.password")+"@tcp(localhost:3306)/"+Public.ReadConfiguration("data.who")) //读取yml文件，并连接数据库
			if err != nil {
				Logmes.Log.Error(err)
			} else {
				Logmes.Log.Info("数据库company连接成功!")
			}
			err = Db.Ping() //Ping 验证与数据库的连接是否仍处于活动状态，并在必要时建立连接。
			if err != nil {
				Logmes.Log.Error(err)
			}
			time.Sleep(time.Minute * 10) //睡眠十分钟
		}
	}() //调用参数列表
}
