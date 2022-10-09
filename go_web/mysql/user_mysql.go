package mysql

import (
	"go_web/model"
	"go_web/utils"
)

func Get_user(a model.Login_in) bool { //用户验证密码
	var b bool
	tx, err := utils.Db.Begin() //事务开启
	if err != nil {
		utils.Log.Error("事务开启出现问题")
		b = false
		return b
	}
	//sql语句
	sqlStr := "select username,password from users where username = ? and password = ?"
	row := tx.QueryRow(sqlStr, a.Name, a.Password) //执行
	ad := model.Login{}
	row.Scan(&ad.Name, &ad.Password) //赋值
	if ad.Name == "" {               //判断失败 没有该用户
		b = false
	} else {
		b = true
	}
	tx.Commit() //提交事务
	return b
}

func Get_user_name(a model.Login_in) bool { //根据用户名查找是否存在
	tx, err := utils.Db.Begin() //事务开启
	if err != nil {
		utils.Log.Error("查找用户名事务开启出现问题")
	}
	sqlStr := "select username,password from users where username = ?"
	row := tx.QueryRow(sqlStr, a.Name)
	ad := model.Login{}
	row.Scan(&ad.Name, &ad.Password) //赋值
	var b bool
	if ad.Name == "" {
		c := Add_user_name(a)
		if c {
			b = false
		} else {
			b = true
		}
	} else {
		b = true
	}
	tx.Commit() //提交事务
	return b
}

func Add_user_name(a model.Login_in) (c bool) { //用户注册后存储信息
	tx, err := utils.Db.Begin() //事务开启
	if err != nil {
		utils.Log.Error("注册存储信息事务开启出现问题")
		c = false
		return c
	}
	sql := "insert into users (username,password) values(?,?)"
	_, err = tx.Exec(sql, a.Name, a.Password)
	if err != nil {
		tx.Rollback() //错误回滚事务
		utils.Log.Error("注册存储信息事务回滚出现问题")
		c = false
		return c
	}
	tx.Commit() //提交事务
	c = true
	return c
}

func Get_user_message(name string) []model.UserInformation { //根据用户名获取存储的信息
	tx, err := utils.Db.Begin() //事务开启
	var d []model.UserInformation
	if err != nil {
		utils.Log.Error("获取存储信息事务开启出现问题")
		return d
	}
	sqlStr := "select username,information,time from news "
	rows, err := tx.Query(sqlStr)
	if err != nil {
		utils.Log.Error("获取存储信息事务开启出现问题")
		return d
	}
	for rows.Next() {
		b := model.UserInformation{}
		rows.Scan(&b.UserName, &b.Information, &b.Time)
		if b.UserName == name {
			d = append(d, b)
		}
	}
	tx.Commit() //提交事务
	return d
}

func Add_user_message(a model.UserInformation) (b bool) { //用户添加信息
	tx, err := utils.Db.Begin() //事务开启
	if err != nil {
		utils.Log.Error("Add_user_message()函数事务开启出现问题")
		b = false
		return b
	}
	sql := "insert into news (username,information,time) values(?,?,?)"
	_, err = tx.Exec(sql, a.UserName, a.Information, a.Time)
	if err != nil {
		tx.Rollback() //错误回滚事务
		b = false
		return b
	}
	tx.Commit() //提交事务
	b = true
	return b
}

func Delete_user_message(a model.UserInformation) { //根据用户存储的信息删除东西，一定能找到
	tx, err := utils.Db.Begin() //事务开启
	if err != nil {
		utils.Log.Error("删除信息事务开启出现问题")
		return
	}
	//写sql语句
	sqlStr := "delete from news where username = ? and information = ? and time = ? "
	//执行
	_, err = tx.Exec(sqlStr, a.UserName, a.Information, a.Time)
	if err != nil {
		tx.Rollback() //错误回滚事务
		utils.Log.Error("删除信息事务开启出现问题")
	}
	tx.Commit() //提交事务
}

func Revise_user_message(a model.UserInformation, b string) (bb bool) { //用户修改保存的内容
	bb = true
	tx, err := utils.Db.Begin() //事务开启
	if err != nil {
		utils.Log.Error("Revise_user_message()函数事务开启出现问题")
		bb = false
		return bb
	}
	//写sql语句
	sql := "update news set information = ?, time = ? where  time = ? and username = ?"
	//执行
	_, err = tx.Exec(sql, a.Information, b, a.Time, a.UserName)
	if err != nil {
		tx.Rollback() //错误回滚事务
		utils.Log.Error("Revise_user_message()函数事务开启出现问题")
		bb = false
	}
	tx.Commit() //提交事务
	return bb
}

func FInd_user_message(a model.Find_message) model.UserInformation { //查找返回信息
	d := model.UserInformation{}
	tx, err := utils.Db.Begin() //事务开启
	if err != nil {
		utils.Log.Error("事务开启出现问题")
		return d
	}
	sqlStr := "select username,information,time from news "
	rows, err := tx.Query(sqlStr)
	if err != nil {
		utils.Log.Error("信息查找出现错误")
		return d
	}
	for rows.Next() {
		b := model.UserInformation{}
		rows.Scan(&b.UserName, &b.Information, &b.Time)
		if b.Time == a.Time {
			if b.UserName == a.Uuid {
				d = b
			}
		}
	}
	tx.Commit() //提交事务
	return d
}
