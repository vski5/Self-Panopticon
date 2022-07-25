package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func router() {
	//生成路由器
	r := gin.Default()
	//注册路由
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, world!",
		})
	})
}

func main() {
	fmt.Println("Hello, world!")
}
