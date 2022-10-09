package utils //连数据库工具 只连接一次
import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" //引用驱动，github库里边的
)

var (
	Db  *sql.DB
	err error
)

func init() { //因为没有使用journal.go文件 所以init()函数没有调用，从新写一个kkk()函数用来保存
	KKK()
	Db, err = sql.Open("mysql", "root:"+InitConfig("data.password")+"@tcp(localhost:3306)/"+InitConfig("data.who")) //连接数据库
	if err != nil {
		Log.Error("连接message数据库没有连接成功")
	}
	err = Db.Ping()
	if err != nil {
		Log.Error("message数据库没有被ping通")
	}
	Log.Info("连接message数据库连接成功")
	//最大连接数
	Db.SetMaxOpenConns(10)
	//设置连接池中的最大闲置连接数
	Db.SetMaxIdleConns(10)
}
