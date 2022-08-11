package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

//用struct定义程序配置项 全局变量，保存程序的所有配置信息
//用settings.Conf拿到配置信息
var Conf = new(AppConfig)

type AppConfig struct {
	Name         string `yaml:"name"`
	Mode         string `yaml:"mode"`
	Port         int    `yaml:"port"`
	Version      string `yaml:"version"`
	*MysqlConfig `yaml:"mysql"`
	*RedisConfig `yaml:"redis"`
	*LogConfig   `yaml:"log"`
}

type MysqlConfig struct {
	User         string `yaml:"user"`
	Password     string `yaml:"password"`
	Dbname       string `yaml:"dbname"`
	MaxOpenConns int    `yaml:"max_open_conns"`
	MaxIdleConns int    `yaml:"max_idle_conns"`
	Host         string `yaml:"host"`
	Port         int    `yaml:"port"`
}

type RedisConfig struct {
	Port     int    `yaml:"port"`
	Db       int    `yaml:"db"`
	PoolSize int    `yaml:"pool_size"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
}

type LogConfig struct {
	Level      string `yaml:"level"`
	Filename   string `yaml:"filename"`
	MaxSize    int    `yaml:"max_size"`
	MaxAge     int    `yaml:"max_age"`
	MaxBackups int    `yaml:"max_backups"`
}

//用viper管理配置文件,初始化配置文件
func Init() (err error) {

	viper.SetConfigName("config") // 配置文件名称(无扩展名)
	viper.SetConfigType("yaml")   // 如果配置文件的名称中没有扩展名，则需要配置此项
	viper.AddConfigPath(".")      // 查找配置文件所在的路径

	err = viper.ReadInConfig() // 查找并读取配置文件
	if err != nil {            // 处理读取配置文件的错误
		fmt.Printf("出错的函数:viper.ReadInConfig() warning err:%v\n", err)

		return

	}
	//读取的配置信息反序列化到Conf
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper.Unmarshal错误,%v\n", err)
	}
	viper.WatchConfig() // 监控配置文件变化

	viper.OnConfigChange(func(e fsnotify.Event) { // 配置文件变化时重新加载配置
		fmt.Println("Config file changed:", e.Name) // 打印配置文件变化的文件名,e.Name是配置文件的路径
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Printf("viper.Unmarshal错误,%v\n", err)
		}
	})
	return
}
