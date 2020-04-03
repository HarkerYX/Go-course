package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

// TCP 客户端
func main() {
	// 1. 根据地址找到那个server
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("连接服务端失败，err：", err)
		return
	}
	defer conn.Close()
	// 2. 向server端发消息
	// fmt.Fprintln(conn, "约吗？")
	// 获取用户输入
	// var input string
	// fmt.Scanf("%s\n", &input)
	reader := bufio.NewReader(os.Stdin)   // 从标准输入获取输入
	input, err := reader.ReadString('\n') // 读取内容直到遇到\n
	if err != nil {
		fmt.Println("获取标准输入失败,err:", err)
		return
	}
	_, err = conn.Write([]byte(input))
	if err != nil {
		fmt.Println("发送消息失败，err:", err)
		return
	}

}
