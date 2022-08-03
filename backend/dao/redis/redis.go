package redis

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

//声明一个全局的rdb变量
var rdb *redis.Client

//Init 初始化redis连接，并返回连接对象，如果出错，则返回错误
func Init() (err error) {
	//NewClient是一个初始化连接的方法，可以传入一个连接参数 *redis.Options
	//func redis.NewClient(opt *redis.Options) *redis.Client
	//类型名称前加 * 的方式，即可表示一个对应的指针类型，redis.Options就是一种类型（struct声明的类型）
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s,%d",
			viper.GetString("redis.host"),
			viper.GetInt("redis.port")),

		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),
		PoolSize: viper.GetInt("redis.pool_size"),
	})

	//检查连接是否成功
	_, err = rdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

//对外包装关闭方法
func Close() {
	_ = rdb.Close()
}
