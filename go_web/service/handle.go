package service

import (
	"go_web/utils"
	"net/http"
)

func Handle_control() {
	go utils.Many_connect_mysql()
	http.HandleFunc("/register", Goregister)                         //普通用户注册函数
	http.HandleFunc("/order", Goadorder)                             //所有成员登录函数
	http.HandleFunc("/userinformation", GetUserInformation)          //获取某个用户存储的信息
	http.HandleFunc("/adduserinformation", AddUserInformation)       //用户添加信息
	http.HandleFunc("/deleteuserinformation", DeleteUserinformation) //用户删除信息
	http.HandleFunc("/reviseuserinformation", ReviseUserinformation) //用户修改信息
	http.HandleFunc("/findinformation", FindUserinformation)         //用户查询信息
}
