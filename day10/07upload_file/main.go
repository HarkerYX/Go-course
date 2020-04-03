package main

import (
	"net/http"
	"fmt"

	"github.com/gin-gonic/gin"
)

func uploadHandler(c *gin.Context){
	// 提取用户上传的文件
	fileObj,err := c.FormFile("filename")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err,
		})
		return
	}
	// fileobj：上传的文件对象
	// fileobj.filename // 拿到上传文件的文件名
	filePath := fmt.Sprintf("./%s", fileObj.Filename)
	// 保存文件到本地的路径
	c.SaveUploadedFile(fileObj, filePath)
	c.JSON(http.StatusOK, gin.H{
		"msg": "OK",
	})
}

// 100 0   00 00   0 000   0 00 0   00 000 000 
//   1byte = 8bit
//   1KB = 1024Bit
//   1MB = 1024KB

// 上传文件示例
func main(){
	r := gin.Default()
	r.LoadHTMLFiles("./upload.html")
	r.GET("/upload", func(c *gin.Context){
		c.HTML(http.StatusOK, "upload.html", nil)
	})
	r.POST("/upload", uploadHandler)
	r.Run()
}