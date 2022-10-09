package logic

import (
	"go_web/model"
	"go_web/mysql"
	"go_web/utils"
)

func Handle_Goadorder(in_message model.Login_in) model.Login_back {
	var back_message model.Login_back //返回给前端数据 没有被json
	var b bool                        //接收数据库返回值
	//发送给数据库
	if in_message.Id == "administrators" {
		b = mysql.Get_administrator(in_message)
		back_message.Id = "administrators"

	} else if in_message.Id == "user" {
		b = mysql.Get_user(in_message)
		back_message.Id = "user"
	} else {
		utils.Log.Error("登录身份不确定")
		b = false
	}
	if b {
		//成功返回
		back_message.Statuscode = 200
		back_message.Msg = in_message.Name //携带姓名，返回前端
	} else {
		//失败返回
		back_message.Statuscode = 400
		back_message.Msg = "密码账号错误"
	}
	return back_message
}
