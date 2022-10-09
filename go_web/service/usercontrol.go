package service

import (
	"encoding/json"
	"go_web/logic"
	"go_web/model"
	"go_web/utils"
	"net/http"
)

func Goregister(w http.ResponseWriter, r *http.Request) { //用户注册
	var in_message model.Login_in                      //前端发送的信息
	var back_message model.Login_back                  //后端返回信息
	err := json.NewDecoder(r.Body).Decode(&in_message) //从body获取前端发过来的信息，并且json解码
	if err != nil {
		utils.Log.Error("接收前端传递注册身份信息失败")
		back_message.Statuscode = 400
		data, _ := json.Marshal(back_message)
		utils.Handle_HTML(w, data)
		return
	}
	back_message = logic.Handle_Goregister(in_message)
	data, _ := json.Marshal(back_message)
	utils.Handle_HTML(w, data)
}

func GetUserInformation(w http.ResponseWriter, r *http.Request) { //获取某个身份的数据发给前端
	var s model.S
	err := json.NewDecoder(r.Body).Decode(&s) //获取身份
	if err != nil {
		utils.Log.Error("后端接收前端发送身份不成功")
		var backk []model.UserInformation
		data, _ := json.Marshal(backk)
		utils.Handle_HTML(w, data)
		return
	}
	message := logic.Handle_GetUserInformation(s) //处理身份，获取返回值,所有信息
	data, _ := json.Marshal(message)
	utils.Handle_HTML(w, data)
}

func AddUserInformation(w http.ResponseWriter, r *http.Request) { //用户添加信息
	//前端传姓名 信息 时间后端获取添加
	var ss model.Add_message //接收前端发送的信息
	var sss model.Login_back //返回给后端信息
	err := json.NewDecoder(r.Body).Decode(&ss)
	if err != nil {
		utils.Log.Error("后端接收前端传递添加的日志信息失败")
		sss.Statuscode = 400
		sss.Msg = "添加失败"
		data, _ := json.Marshal(sss)
		utils.Handle_HTML(w, data)
		return
	}
	b := logic.Handle_AddUserInformation(ss)
	if b {
		sss.Statuscode = 200
		sss.Msg = "添加成功"
	} else {
		sss.Statuscode = 400
		sss.Msg = "添加失败"
	}
	data, _ := json.Marshal(sss)
	utils.Handle_HTML(w, data)
}

func DeleteUserinformation(w http.ResponseWriter, r *http.Request) { //用户删除存储信息
	var demessage model.UserInformation               //前端发送的信息传了四个，后端收到了三个
	err := json.NewDecoder(r.Body).Decode(&demessage) //从body获取前端发过来的信息，并且json解码
	if err != nil {
		utils.Log.Error("DeleteUserinformation()函数后端接收前端传输的需要删除的信息出现错误")
	}
	logic.Handle_DeleteUserinformation(demessage)
}

func ReviseUserinformation(w http.ResponseWriter, r *http.Request) { //用户修改存储信息
	var backk model.Login_back
	var remessage model.UserInformation //前端发送的信息
	err := json.NewDecoder(r.Body).Decode(&remessage)
	if err != nil {
		utils.Log.Error("后端获取前端传输的修改信息出现错误")
		backk.Statuscode = 400
		backk.Msg = "修改没有成功"
		data, _ := json.Marshal(backk)
		utils.Handle_HTML(w, data)
		return
	}
	bb := logic.Handle_ReviseUserinformation(remessage)
	if bb {
		backk.Statuscode = 200
		backk.Msg = "修改成功"
	} else {
		backk.Statuscode = 400
		backk.Msg = "修改没有成功"
	}
	data, _ := json.Marshal(backk)
	utils.Handle_HTML(w, data)
}

func FindUserinformation(w http.ResponseWriter, r *http.Request) { //用户查询信息
	var finmessage model.Find_message //前端发送信息
	err := json.NewDecoder(r.Body).Decode(&finmessage)
	if err != nil {
		utils.Log.Error(err)
	}
	a := logic.Handle_FindUserinformation(finmessage)
	data, _ := json.Marshal(a)
	utils.Handle_HTML(w, data)
}
