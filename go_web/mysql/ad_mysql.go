package mysql

import (
	"go_web/model"
	"go_web/utils"
)

func Get_administrator(a model.Login_in) bool { //管理员验证密码
	var b bool
	tx, err := utils.Db.Begin() //事务开启
	if err != nil {
		utils.Log.Error("Get_administrator()函数事务开启出现错误")
		b = false
		return b
	}
	//sql语句
	sqlStr := "select name,password from administrator where name = ? and password = ?"
	row := tx.QueryRow(sqlStr, a.Name, a.Password) //执行
	ad := model.Login{}
	row.Scan(&ad.Name, &ad.Password) //赋值
	if ad.Name == "" {               //为空找不到，返回false
		b = false
	} else {
		b = true
	}
	tx.Commit() //提交事务
	return b
}
