package utils //开一个单独的线程 现在是一分钟连接一次  可以改成一天连接一次 死循环

import (
	"database/sql"
	"time"
)

func Many_connect_mysql() {
	// 1.获取ticker对象
	ticker := time.NewTicker(1 * time.Minute)
	// 子协程
	go func() {
		for {
			<-ticker.C
			Db, err = sql.Open("mysql", "root:"+InitConfig("data.password")+"@tcp(localhost:3306)/"+InitConfig("data.who")) //连接数据库
			if err != nil {
				Log.Error("数据库没有连接成功")
			}
			err = Db.Ping()
			if err != nil {
				Log.Error("数据库没有ping通")
			}
			Log.Info("数据库message连接成功")
			time.Sleep(time.Minute * 10)
		}
	}()
}
