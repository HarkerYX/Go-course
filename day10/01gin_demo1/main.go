package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func indexHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "这是index页面！",
	})
}
func main() {
	// 启动一个默认的路由
	router := gin.Default()
	// 给/hello配置一个处理函数
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "Hello 沙河！",
		})
	})

	router.GET("/index", indexHandler)
	// 启动webserver
	router.Run(":9090")
}
