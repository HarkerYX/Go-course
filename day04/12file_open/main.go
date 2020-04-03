package main

import (
	"fmt"
	"io"
	"os"
)

// 打开和关闭文件
func main() {
	file, err := os.Open("./xx.txt")
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	// 文件能打开
	defer file.Close() // 使用defer延迟关闭文件
	// 读文件
	var tmp [128]byte //定义一个字节数组
	// var s = make([]byte, 0, 128)
	for {
		n, err := file.Read(tmp[:]) // 基于数组得到一个切片
		if err == io.EOF {          // End Of File
			fmt.Println("文件已经读完了")
			return
		}
		if err != nil {
			fmt.Println("read from file failed, err:", err)
			return
		}
		fmt.Printf("读取了%d个字节\n", n)
		fmt.Println(string(tmp[:]))
	}

}
