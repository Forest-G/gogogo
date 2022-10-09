package logic

import (
	"fmt"
	"go_web/model"
	"go_web/mysql"
	"go_web/utils"
)

func Handle_Goregister(in_message model.Login_in) model.Login_back {
	var back_message model.Login_back
	b := mysql.Get_user_name(in_message) //验证用户名
	if b {                               //已经有名字,已经存在
		back_message.Statuscode = 400
		back_message.Id = "user"
		back_message.Msg = "用户名重复"
	} else { //能正常注册成功
		back_message.Statuscode = 200
		back_message.Id = "user"
		back_message.Msg = "注册成功"
		utils.Log.Info(in_message.Name + "用户注册成功")
	}
	return back_message
}

func Handle_GetUserInformation(s model.S) []model.UserInformation {
	message := mysql.Get_user_message(s.Uuid)
	return message
}

func Handle_AddUserInformation(ss model.Add_message) (b bool) {
	//更新到数据库然后返回成功刷新界面
	var aa model.UserInformation
	aa.Time = utils.CreatTime()
	aa.UserName = ss.Uuid
	aa.Information = ss.Information
	b = mysql.Add_user_message(aa)
	return b
}

func Handle_DeleteUserinformation(demessage model.UserInformation) { //删除信息
	mysql.Delete_user_message(demessage)
}

func Handle_ReviseUserinformation(remessage model.UserInformation) bool { //修改信息
	time := utils.CreatTime()                        //获取时间
	bb := mysql.Revise_user_message(remessage, time) //发给数据库
	return bb
}

func Handle_FindUserinformation(finmessage model.Find_message) model.UserInformation { //查找信息
	a := mysql.FInd_user_message(finmessage)
	if a.Information == "" {
		fmt.Println("没有找到信息")
		a.Information = "没有找到信息"
	}
	return a
}
