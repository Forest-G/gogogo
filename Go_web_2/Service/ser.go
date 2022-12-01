package Service

import (
	//json数据的解码编码包

	//自己定义包的导入
	"encoding/json"
	logmes "go_web_2/Log" //自定义包，日志
	logic "go_web_2/Logic"
	"go_web_2/Public"
	"io/ioutil" //读取请求体所需要的包，至于为何会警告，查到的是说有另外一个包跟它有相同的功能
	"net/http"
	"time"
)

func RegisterRoutes() {
	http.HandleFunc("/handleregist", handleRegist)                 // 处理注册请求的处理器，handleRegist是处理器函数
	http.HandleFunc("/order", handleLogin)                         // 处理登录请求的处理器
	http.HandleFunc("/adduserinformation", handleAddMessage)       //处理增加用户请求
	http.HandleFunc("/deleteuserinformation", handleDeleteMessage) // 处理删除信息请求
	http.HandleFunc("/reviseuserinformation", handleModifyMessage) //处理更改数据请求
	http.HandleFunc("/findinformation", handleCheckMessage)        //查找信息
	http.HandleFunc("/userinformation", handleSendMessage)         //传输数据库数据到前端页面
}
func interf(w http.ResponseWriter, receive interface{}) { //封装定义写入前端数据的函数，其中receive空接口的目的是：为了方便传入不同类型的结构体
	w.Header().Set("Access-Control-Allow-Origin", "*") //跨域请求，个人理解：可以在其他平台给web页面传信息
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("content-type", "application/json") //设置传输格式是json格式
	a := receive.([]byte)                              //强转，因为传过来的都是编好的json数据类型，所以直接强制类型转换为切片类型直接传给前端
	w.Write(a)                                         //w.write函数写入数据给前端
}

func handleRegist(w http.ResponseWriter, r *http.Request) { //处理注册响应
	body, err := ioutil.ReadAll(r.Body) //用ioutil.read函数读取请求体的内容
	if err != nil {
		logmes.Log.Error(err)
	}
	println("json:", string(body)) //打印获取的json数据
	var a Public.User              //用结构体接收获取的json数据
	json.Unmarshal(body, &a)
	interf(w, logic.Receive_Regist(a.Username, a.Passwd))
}
func handleLogin(w http.ResponseWriter, r *http.Request) { //处理登录响应
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logmes.Log.Error(err)
	}
	println("json:", string(body))
	var a Public.User
	json.Unmarshal(body, &a)
	uuid := a.Username
	interf(w, logic.Receive_Login(a.Username, a.Passwd, a.Typ, uuid))
}

// **********************************************************************以下是普通用户
func handleAddMessage(w http.ResponseWriter, r *http.Request) { //处理增加信息响应
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logmes.Log.Error(err)
	}
	println("json:", string(body))
	var a Public.User_Operation
	json.Unmarshal(body, &a)
	timeStr := time.Now().Format("2006-01-02 15:04:05")      //固定写法
	a.Tim = timeStr                                          //把获取的时间赋给结构体
	interf(w, logic.Receive_Add(a.Usname, a.Tim, a.Content)) //把需要的数据传给逻辑处理层
}

func handleDeleteMessage(w http.ResponseWriter, r *http.Request) { //处理删除信息响应
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logmes.Log.Error(err)
	}
	println("json:", string(body))
	var a Public.User_Operation
	json.Unmarshal(body, &a)
	interf(w, logic.Receive_Delete(a.Usname, a.Tim))
}

func handleModifyMessage(w http.ResponseWriter, r *http.Request) { //处理更改信息响应
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logmes.Log.Error(err)
	}
	println("json:", string(body))
	var a Public.User_Operation
	json.Unmarshal(body, &a)
	interf(w, logic.Receive_Modify(a.Usname, a.Tim, a.Content))
}

func handleCheckMessage(w http.ResponseWriter, r *http.Request) { //处理查看信息响应
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logmes.Log.Error(err)
	}
	println("jsonfind:", string(body))
	var a Public.User_Operation
	json.Unmarshal(body, &a)
	interf(w, logic.Receive_Check(a.Usname, a.ID, a.Content))
}

func handleSendMessage(w http.ResponseWriter, r *http.Request) { //发送所有信息给前端
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logmes.Log.Error(err)
	}
	println("jsonuuid", string(body))
	var a Public.User_Operation
	json.Unmarshal(body, &a)
	interf(w, logic.Receive_Check_All(a.Usname))
}
