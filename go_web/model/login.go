package model

type Login_in struct { //登录时接收前端发过来的信息
	Name     string `json:"name"`     //姓名
	Password string `json:"password"` //密码
	Id       string `json:"id"`       //身份
}

type Login_back struct { //登录后返回给后端
	Statuscode int    `json:"statuscode"` //状态码
	Msg        string `json:"msg"`        //传递的信息
	Id         string `json:"id"`         //传递身份
}

type Login struct { //保存密码和账号的
	Name     string `json:"username"`
	Password string `json:"userpassword"`
}
