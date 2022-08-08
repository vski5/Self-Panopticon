package routes

import (
	"backend/backend/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {

	r := gin.New() //创建一个gin实例
	//不用自己写的中间件，下面用三方库
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	//注册中间件使得Gin日志记录到zap中
	//r.Use(ginzap.Ginzap(zap.L(), time.RFC3339, true))
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	return r
}
