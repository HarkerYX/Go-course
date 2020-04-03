package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func liveIndexHandler(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg": "live/index",
	})
}

func liveHomeHandler(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg": "live/Home",
	})
}