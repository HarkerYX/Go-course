package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func shopIndexHandler(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg": "shopping/index",
	})
}
func shopHomeHandler(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg": "shopping/home",
	})
}
