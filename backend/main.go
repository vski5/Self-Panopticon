package main

import (
	"backend/backend/logger"
	"backend/backend/settings"
	"fmt"

	"go.uber.org/zap"
)

func main() {
	//加载配置文件，判断setting是否出错
	if err := settings.Init(); err != nil {
		fmt.Printf("setting加载配置文件错了,err:%v\n", err)
		return
	}
	//初始化日志
	if err := logger.Init(); err != nil {
		fmt.Printf("logger初始化日志错了,err:%v\n", err)
		return
	}
	//拿zap库中全局定义的logger
	zap.L().Debug("logger init ok")
	//初始化mysql
	if err := mysql.Init(); err != nil {
		fmt.Printf("mysql初始化日志错了,err:%v\n", err)
		return
	}
	//初始化redis连接
	if err := redis.Init(); err != nil {
		fmt.Printf("redis初始化日志错了,err:%v\n", err)
		return
	}
	//注册路由
	//启动服务 and 优雅关机
}
