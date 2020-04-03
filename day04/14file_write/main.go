package main

import (
	"fmt"
	"os"
)

// 打开文件支持文件写入

func main() {
	file, err := os.OpenFile("xx.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("打开文件失败", err)
		return
	}
	defer file.Close()
	str := "Hello 沙河"
	file.Write([]byte("哈哈哈\n"))
	file.WriteString(str)
}
