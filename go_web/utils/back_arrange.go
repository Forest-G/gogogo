package utils //获取配置文件里边的信息

import (
	"os"

	"github.com/spf13/viper"
)

func InitConfig(a string) string {
	KKK()
	work, _ := os.Getwd()                // 获取目录路径
	viper.SetConfigName("arrange")       // 设置文件名
	viper.SetConfigType("yml")           // 设置文件类型
	viper.AddConfigPath(work + "/utils") // 执行单文件路径
	//viper.AddConfigPath(work + "/config")         // 执行go run文件路径
	err := viper.ReadInConfig()
	if err != nil {
		Log.Error("配置文件读取时出现错误")
	}
	port := viper.GetString(a)
	return port
}
