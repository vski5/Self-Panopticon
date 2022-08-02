package routes

import (
	"backend/backend/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {

	r := gin.New() //创建一个gin实例
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	return r
}
