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

func helloHandler(c *gin.Context){
	// 声明一个结构体类型
	type userinfo struct{
		Name string `json:"username"`
		Password string `json:"pwd"`
	}
	// 实例化一个userinfo对象
	u1 := userinfo{
		Name:"haojie",
		Password:"123456",
	}

	c.JSON(http.StatusOK, u1)
}

func xmlHandler(c *gin.Context){
	c.XML(http.StatusOK, gin.H{
		"msg": "xml",
	})
}

func main(){
	r := gin.Default()
	r.GET("/index", indexHandler)
	r.GET("/hello", helloHandler)
	r.GET("/xml", xmlHandler)
	r.GET("/yaml", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{"message": "ok", "status": http.StatusOK})
	})
	r.Run()
}