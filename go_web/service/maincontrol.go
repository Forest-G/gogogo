package service

import (
	"encoding/json"
	"go_web/logic"
	"go_web/model"
	"go_web/utils"
	"net/http"
)

//登录函数
func Goadorder(w http.ResponseWriter, r *http.Request) {
	var in_message model.Login_in                      //前端发送的信息
	var back_message model.Login_back                  //后端要发送的信息
	err := json.NewDecoder(r.Body).Decode(&in_message) //从body获取前端发过来的信息，并且json解码
	if err != nil {
		utils.Log.Error("获取前端输入的登录信息出现错误")
		//既然错误了就要返回前端给400
		back_message.Statuscode = 400
		data, _ := json.Marshal(back_message)
		utils.Handle_HTML(w, data) //访问域
		return
	}
	back_message = logic.Handle_Goadorder(in_message) //逻辑处理
	if back_message.Statuscode == 200 {
		utils.Log.Info(back_message.Msg + "登录成功")
	}
	data, _ := json.Marshal(back_message)
	utils.Handle_HTML(w, data) //访问域
}
