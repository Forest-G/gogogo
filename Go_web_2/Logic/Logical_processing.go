package logic

import (
	"encoding/json"
	utils "go_web_2/Db"
	logmes "go_web_2/Log" //自定义包，日志
	"go_web_2/Public"
)

// **********************************************************接收表示层传来的数据并作逻辑处理
func Receive_Regist(username string, passwd string) []byte { //处理注册逻辑
	if utils.Check_from_Database(username) { //传参，给处理数据库函数，判断数据库内是否含有相同的username
		utils.SaveUser(username, passwd) //保存账号密码
		var s Public.Send
		s.Scode = 200
		s.Smessage = "success"    //发送状态码：200，发送响应信息success
		t, err := json.Marshal(s) //定义t，和err，t是方便传参给写入前端的函数（interf）
		if err != nil {
			logmes.Log.Error(err)
		}
		return t
	} else { //否则就不进去
		var s Public.Send
		s.Scode = 406
		s.Smessage = "failed"
		t, err := json.Marshal(s)
		if err != nil {
			logmes.Log.Error(err)
		}
		return t
	}
}

func Receive_Login(username string, passwd string, typ string, uuid string) []byte { //处理登录逻辑
	var s Public.Send                                                //通俗来讲uuid就是唯一的一个标识符，用来标记不同的账号
	if utils.Check_Login_From_Database(username, passwd, typ) == 1 { //如果处理逻辑返回1，证明是管理员登录
		s.Scode = 200             //设置状态码
		s.Smessage = uuid         //让 uuid赋值给Smessage
		s.Typ = "administrators"  //类型为管理员
		t, err := json.Marshal(s) //打包发给前端
		if err != nil {
			logmes.Log.Error(err)
		}
		return t
	} else if utils.Check_Login_From_Database(username, passwd, typ) == -1 { //返回-1证明是普通用户登录
		s.Scode = 200
		s.Smessage = uuid
		s.Typ = "users"
		t, err := json.Marshal(s)
		if err != nil {
			logmes.Log.Error(err)
		}
		return t
	} else {
		s.Scode = 406
		s.Smessage = "failed"
		s.Smessage = "users"
		t, err := json.Marshal(s)
		if err != nil {
			logmes.Log.Error(err)
		}
		return t
	}
}

func Receive_Add(usname string, timeStr string, content string) []byte { //处理增加逻辑
	if utils.Add_Message_To_Database(usname, timeStr, content) {
		var s Public.Send
		s.Scode = 406
		s.Smessage = "success"
		t, err := json.Marshal(s)
		if err != nil {
			logmes.Log.Error(err)
		}
		return t
	} else { //否则就不进去
		var s Public.Send
		s.Scode = 406
		s.Smessage = "failed"
		t, err := json.Marshal(s)
		if err != nil {
			logmes.Log.Error(err)
		}
		return t
	}
}

func Receive_Delete(usname string, tim string) []byte { //处理删除逻辑
	var s Public.Send
	if utils.Dlelete_from_Database(usname, tim) {
		s.Scode = 200
		s.Smessage = "success"
		t, err := json.Marshal(s)
		if err != nil {
			logmes.Log.Error(err)
		}
		return t
	} else {
		s.Scode = 406
		s.Smessage = "failed"
		t, err := json.Marshal(s)
		if err != nil {
			logmes.Log.Error(err)
		}
		return t
	}
}

func Receive_Modify(usname string, tim string, content string) []byte { //处理更改逻辑
	var s Public.Send
	if utils.Modify_from_Database(usname, tim, content) {
		s.Scode = 200
		s.Smessage = "success"
		t, err := json.Marshal(s)
		if err != nil {
			logmes.Log.Error(err)
		}
		return t
	} else {
		s.Scode = 406
		s.Smessage = "failed"
		t, err := json.Marshal(s)
		if err != nil {
			logmes.Log.Error(err)
		}
		return t
	}
}
func Receive_Check(usname string, id string, content string) []byte { //处理检查逻辑
	var chec string = utils.Check_Message_From_Database(usname, id, content)
	var a Public.Send
	if chec == "check_err" {
		a.Scode = 406
		a.Smessage = chec
		t, err := json.Marshal(a)
		if err != nil {
			logmes.Log.Error(err)
		}
		return t
	} else {
		a.Scode = 200
		a.Smessage = chec
		t, err := json.Marshal(a)
		if err != nil {
			logmes.Log.Error(err)
		}
		return t
	}
}

func Receive_Check_All(usnam string) []byte { //处理所有信息返回逻辑
	uuid := usnam
	rece := utils.Check_All_From_Database(uuid)
	var chec []Public.User_Operation = rece
	var a Public.Send
	if chec != nil {
		a.Scode = 200
		a.Smessage = chec
		t, err := json.Marshal(a)
		if err != nil {
			logmes.Log.Error(err)
		}
		return t
	} else {
		a.Scode = 406
		a.Smessage = chec
		t, err := json.Marshal(a)
		if err != nil {
			logmes.Log.Error(err)
		}
		return t
	}
}
