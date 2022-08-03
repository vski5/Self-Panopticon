package routes

import (
	"backend/backend/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Setup 路由初始化，返回路由，用于启动服务
func Setup() *gin.Engine {

	r := gin.New() //创建一个gin实例
	//设置日志格式
	//GinLogger 接收gin框架默认的日志
	//GinRecovery recover掉项目可能出现的panic，并使用zap记录相关日志
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	return r
}
