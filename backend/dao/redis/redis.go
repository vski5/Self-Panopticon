package redis

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

//声明一个全局的rdb变量
var rdb *redis.Client

//Init 初始化redis连接
func Init() (err error) {
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s,%d",
			viper.GetString("redis.host"),
			viper.GetInt("redis.port")),

		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),
		PoolSize: viper.GetInt("redis.pool_size"),
	})
	_, err = rdb.Ping().Result()
	if err != nil {
		return
	}
	return nil
}

//对外包装关闭方法
func Close() {
	_ = rdb.Close()
}
