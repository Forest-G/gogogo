package main

import (
	"go_web/service"
	"go_web/utils"
	"net/http"
)

func main() {
	service.Handle_control()
	err := http.ListenAndServe(":"+utils.InitConfig("server.port"), nil) //调动一个监听器
	if err != nil {
		utils.Log.Fatal("main.go文件,ListenAndServe函数监听失败,无法监听。") //会调用os.Exit(1) 退出程序
	}
}
