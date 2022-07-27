package main

import (
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
	router := gin.Default()

	/*注册路由，设置路由规则，路由规则是一个函数：func(c *gin.Context)，
	这个函数的参数是gin.Context类型的，可以通过c.Param()获取路由参数*/
	router.GET("/user/:name/*action", func(c *gin.Context) {
		//获取路由参数，参数名称是name，参数值是c.Param("name")
		name := c.Param("name")
		action := c.Param("action")
		//获取路由参数，参数名称是name，参数值是c.Query("name")
		c.String(http.StatusOK, "%s is %s", name, action)
	})

	//路由组，将路由分组，方便统一配置路由前缀
	userGroup := router.Group("/user")
	{
		userGroup.GET("/index", func(c *gin.Context) {})
		userGroup.GET("/login", func(c *gin.Context) {})
		userGroup.POST("/login", func(c *gin.Context) {})

	}

	//为没有配置处理函数的路由添加处理程序，默认情况下它返回404代码
	//此处为没有匹配到路由的请求都返回views/404.html页面。
	router.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "views/404.html", nil)
	})

	//启动路由，默认监听8080端口
	router.Run(":8000")
}
