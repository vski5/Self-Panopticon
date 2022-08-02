package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

//用viper管理配置文件,初始化配置文件
func Init() (err error) {

	viper.SetConfigName("config") // 配置文件名称(无扩展名)
	viper.SetConfigType("yaml")   // 如果配置文件的名称中没有扩展名，则需要配置此项
	viper.AddConfigPath(".")      // 查找配置文件所在的路径

	err = viper.ReadInConfig() // 查找并读取配置文件
	if err != nil {            // 处理读取配置文件的错误
		fmt.Printf("出错的函数:viper.ReadInConfig() warning err:%v\n", err)
		return
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	viper.WatchConfig()                           // 监控配置文件变化
	viper.OnConfigChange(func(e fsnotify.Event) { // 配置文件变化时重新加载配置
		fmt.Println("Config file changed:", e.Name) // 打印配置文件变化的文件名,e.Name是配置文件的路径
	})
	return
}
