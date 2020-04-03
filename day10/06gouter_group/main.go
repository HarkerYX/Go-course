package main
import (

	"github.com/gin-gonic/gin"
)


// 路由分组
func main(){
	r := gin.Default()
	
	shoppingGroup := r.Group("/shopping")
	{
		shoppingGroup.GET("/index", shopIndexHandler)
		shoppingGroup.GET("/home", shopHomeHandler)
	}
	
	liveGroup := r.Group("/live")
	{
		liveGroup.GET("/index", liveIndexHandler)
		liveGroup.GET("/home", liveHomeHandler)
	}

	v1 := r.Group("/v1")
	{
		v1Shopping := v1.Group("/shopping")
		{
		v1Shopping.GET("/home", shopHomeHandler)
		}
	}
	r.Run()
}