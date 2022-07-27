package main

import (
	"backend/backend/Middleware"
	"log"
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

//路由组，将路由分组，方便统一配置路由前缀
func RouterGroup() {
	//生成路由器
	r := gin.Default()

	//注册路由组userGroup，将路由分组，方便统一配置路由前缀
	userGroup := r.Group("/v1")
	{
		userGroup.GET("/index", func(c *gin.Context) {})
		userGroup.GET("/login", func(c *gin.Context) {})
		userGroup.POST("/login", func(c *gin.Context) {})

	}

	//注册路由组v1
	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, "v1")
		})
	}

	//启动路由
	//gin 源码中的engine.Run()
	r.Run()
}

//注册一个针对user的路由
func userGroup() {
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
}

//为没有配置处理函数的路由添加处理程序，默认情况下它返回404代码
//此处为没有匹配到路由的请求都返回views/404.html页面。
func NotFound() {
	router := gin.Default()
	router.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "views/404.html", nil)
	})
}

func main() {
	// 新建一个没有任何默认中间件的路由
	r := gin.New()
	// 注册一个全局中间件StatCost，用于统计请求耗时
	r.Use(Middleware.StatCost())

	r.GET("/test", func(c *gin.Context) {
		name := c.MustGet("name").(string) // 从上下文取值
		log.Println(name)
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world!",
		})
	})

	//启动路由，默认监听8080端口
	r.Run(":8000")

}
