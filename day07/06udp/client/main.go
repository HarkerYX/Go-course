package main

import (
	"fmt"
	"net"
)

// UDP client

func main() {
	conn, err := net.Dial("udp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("连接server失败，err：", err)
		return
	}
	defer conn.Close()
	n, err := conn.Write([]byte("约吗？"))
	if err != nil {
		fmt.Println("发送消息失败，err:", err)
		return
	}
	// 收消息
	var buf [1024]byte
	n, err = conn.Read(buf[:])
	if err != nil {
		fmt.Println("读取消息失败,err:", err)
		return
	}
	fmt.Println("收到回复：", string(buf[:n]))

}
