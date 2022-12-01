package main

import (
	//自定义包

	utils "go_web_2/Db"
	logmes "go_web_2/Log" //自定义包，日志
	"go_web_2/Service"
	"net/http"
)

func main() {
	Service.RegisterRoutes() //不同的handle处理器
	logmes.Init_Jou()        //日志
	utils.Databaseinit()     //连接数据库
	utils.MultiConnect()     //隔一段时间连接数据库

	err := http.ListenAndServe(":8080", nil) //开启服务
	if err != nil {
		logmes.Log.Panic("监听失败")
		panic(err)
	} else {
		logmes.Log.Info("正在监听:8080端口")
	}
}
