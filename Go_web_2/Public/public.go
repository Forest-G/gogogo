package Public

import (
	Logmes "go_web_2/Log"

	"github.com/spf13/viper"
)

type User struct {
	Username string `json:"name"`     //用户名
	Passwd   string `json:"password"` //用户密码
	Id       int    // 数据库里边的ID
	Typ      string `json:"id"` // TYP是类型，表示身份，tag表示的是身份administrator或者user
}

type User_Operation struct { //普通用户对自己的日记增删改查
	Usname  string `json:"uuid"`        //用户名
	Content string `json:"information"` //内容
	Tim     string `json:"time"`        //时间
	ID      string `json:"id"`          //序号
}

type Send struct {
	Scode    int         `json:"statuscode"`  //状态码
	Smessage interface{} `json:"information"` //返回信息，是成功或者Println()
	Typ      string      `json:"id"`
}
type Info struct { //该结构体主要是针对登录注册的操作
	Name string `db:"username"` //该tag，对应的数据库
	Pswd string `db:"password"`
	Typ  string `db:"identity"`
	Id   int    `db:"id "`
}

type Out struct { //该结构体主要是针对用户的增删改查
	Usnam string `db:"usnam"`   //用户名
	Time  string `db:"time"`    //时间
	Con   string `db:"content"` //内容
	ID    int    `db:"id"`      //Id
}

func ReadConfiguration(name string) string { //读取配置文件
	viper.SetConfigName("Cfig") //获取文件的文件名
	viper.SetConfigType("yml")  //类型是yml的
	viper.AddConfigPath(".")    //在这个目录下查找yml文件
	err := viper.ReadInConfig() //定义错误
	if err != nil {
		Logmes.Log.Error(err)
	} else {
		Logmes.Log.Info("读取yml文件成功")
	}
	a := viper.GetString(name) //返回名字
	return a
}
