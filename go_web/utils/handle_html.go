package utils

import "net/http"

func Handle_HTML(w http.ResponseWriter, sss interface{}) {
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")             //返回数据格式是json
	a := sss.([]byte)                                              //空接口 注意空接口的类型与切片啥的不一样
	w.Write(a)
}
