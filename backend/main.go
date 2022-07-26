package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router() {
	//生成路由器
	r := gin.Default()
	//注册路由
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})
	//启动路由
	//gin 源码中的engine.Run()
	r.Run()
}

func main() {
	fmt.Println("Hello, world!")
}
