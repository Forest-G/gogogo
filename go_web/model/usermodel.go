package model

type UserInformation struct { //用户存储信息
	Information string `json:"information"`
	UserName    string `json:"username"`
	Time        string `json:"time"`
}

//获取前端过来的uuid
type S struct {
	Uuid string `json:"uuid"`
}

type Add_message struct { //用户添加信息
	Uuid        string `json:"uuid"`
	Information string `json:"information"`
}

type Find_message struct { //用户查询信息
	Uuid string `json:"uuid"`
	Time string `json:"time"`
}
