package utils //数据库的操作

import (
	"database/sql"
	Logmes "go_web_2/Log" //要引用的包的别名，防止冲突
	"go_web_2/Public"
	"strconv" //将ID的string转化为int需要的包

	_ "github.com/go-sql-driver/mysql" //go语言数据库驱动,前面加下划线，只调用包里面的init函数。无法调用包中别的函数。
)

var (
	Db *sql.DB
)

func Databaseinit() { //初始化连接数据库
	var err error
	Db, err = sql.Open("mysql", "root:"+Public.ReadConfiguration("data.password")+"@tcp(localhost:3306)/"+Public.ReadConfiguration("data.who")) //连接数据库
	//fmt.Println(Public.ReadConfiguration("data.password") + "@tcp(localhost:3306)/" + Public.ReadConfiguration("data.who"))
	Db.SetMaxOpenConns(100) //设置同时打开的连接数(使用中+空闲)
	//设为10。将此值设置为小于或等于0表示没有限制                                                                                                                     //通过配置文件，获取密码，    //配置文件，获取要连接的数据库
	Db.SetMaxIdleConns(100) // 将最大并发空闲链接数设置为 5.
	// 小于或等于 0 表示不保留任何空闲链接.
	if err != nil {
		Logmes.Log.Error(err)
	} else {
		Logmes.Log.Info("数据库连接成功")
	}
}

func SaveUser(username string, password string) { //保存账号密码到数据库内
	//写sql语句
	tx, _ := Db.Begin()
	sqlStr_user := "insert into users (username, password,identity) values(?,?,?)" //保存数据到数据库
	sqlStr_noteped := "insert into notepad(usnam)values(?)"
	//执行
	_, err := Db.Exec(sqlStr_user, username, password, "user") //增加账号密码到user数据库
	_, err2 := Db.Exec(sqlStr_noteped, username)               //增加账号到用户数据库(notepad)
	if err != nil {
		tx.Rollback()
	}
	if err2 != nil {
		tx.Rollback()
	}
	tx.Commit()
}

func Check_from_Database(name string) bool { //查重（账号），防止增加相同账号
	rows, err := Db.Query("select username from users")
	if err != nil {
		Logmes.Log.Error(err)
	}
	for rows.Next() {
		var s Public.Info
		rows.Scan(&s.Name)  //rows.Scan意思是把数据库内的数据放到s这个结构体内部，便于和外来数据比较
		if name == s.Name { //如果相同，证明数据库内部有这个数据，返回错误
			return false
		}
	}
	rows.Close() //关闭
	return true
}

func Check_Login_From_Database(nam string, pas string, typ string) int { //检查登录是否成功
	rows, err := Db.Query("select username,password,identity from users")
	if err != nil {
		Logmes.Log.Error(err)
	}
	for rows.Next() {
		var us Public.Info
		rows.Scan(&us.Name, &us.Pswd, &us.Typ)
		if us.Name == nam && us.Pswd == pas && typ == us.Typ { //如果传来的用户名和密码都和数据库内的数据匹配上了
			if typ == "administrator" {
				rows.Close()
				return 1
			} else {
				rows.Close()
				return -1
			}
		}
	}
	rows.Close()
	return 0
}

func Dlelete_from_Database(usnam string, ti string) bool { //普通用户删除一条日记
	tx, _ := Db.Begin()
	rows, err := Db.Query("delete from notepad where usnam = ? AND time = ?", usnam, ti)
	if err != nil {
		tx.Rollback()
		rows.Close()
		Logmes.Log.Error(err)
		return false
	} else {
		tx.Commit()
		Logmes.Log.Info("删除成功")
		rows.Close()
		return true
	}
}

func Modify_from_Database(usnam string, ti string, con string) bool { //更改数据库信息
	tx, _ := Db.Begin()
	rows, err := Db.Query("update notepad set content = ? where usnam =? AND time =?", con, usnam, ti)
	if err != nil {
		tx.Rollback()
		rows.Close()
		Logmes.Log.Error(err)
		return false
	} else {
		tx.Commit()
		rows.Close()
		Logmes.Log.Info("更改成功")
		return true
	}
}

func Check_Message_From_Database(Usnam string, id string, con string) string { //查找数据库信息
	a, err := strconv.Atoi(id) //将传来的string 型的id转化为int类型
	if err != nil {
		Logmes.Log.Error(err)
	}
	row := Db.QueryRow("select content from notepad where usnam=? AND id =? ", Usnam, a) //单行查询
	var s Public.Out
	err2 := row.Scan(&s.Con)
	if err2 != nil {
		Logmes.Log.Error(err2)
	}
	return s.Con
}

func Add_Message_To_Database(name string, tim string, con string) bool { //增加信息到数据库内
	tx, _ := Db.Begin()
	sqlStr_noteped := "insert into notepad(usnam,time,content)values(?,?,?)"
	_, err2 := Db.Exec(sqlStr_noteped, name, tim, con)
	if err2 != nil {
		Logmes.Log.Error(err2)
		tx.Rollback()
		return false
	} else {
		tx.Commit()
		Logmes.Log.Info("增加成功")
		return true
	}
}

func Check_All_From_Database(uuid string) []Public.User_Operation { //定义结构体切片类型的返回值
	println(uuid)
	var a []Public.User_Operation                                                             //先定义切片，为了返回                                                                 //开启事务
	rows, err1 := Db.Query("select usnam,time,content,id from notepad where usnam = ?", uuid) //根据uuid获取信息
	if err1 != nil {
		Logmes.Log.Error(err1)
	} //检查错误
	for rows.Next() {
		var s Public.User_Operation
		err2 := rows.Scan(&s.Usname, &s.Tim, &s.Content, &s.ID)
		if err2 != nil {
			Logmes.Log.Error(err2)
		}
		sen := append(a, s) //append追加写入结构体切片中
		a = sen             //把查到的数据放入a里边
	}
	rows.Close()
	return a //返回a
}
