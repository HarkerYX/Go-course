package main

import (
	"fmt"
	"net"
)

// UDP server端

func process(listener *net.UDPConn) {
	defer listener.Close()
	// 循环收发数据
	for {
		var buf [1024]byte
		n, addr, err := listener.ReadFromUDP(buf[:])
		if err != nil {
			fmt.Println("接收消息失败：err:", err)
			return
		}
		fmt.Printf("接收到来自%v的消息：%v\n", addr, string(buf[:n]))
		// 回复消息
		n, err = listener.WriteToUDP([]byte("滚"), addr)
		if err != nil {
			fmt.Println("回复消息失败，err：", err)
			return
		}
	}
}

func main() {
	listener, err := net.ListenUDP("udp", &net.UDPAddr{
		// IP: "127.0.0.1"
		IP:   net.ParseIP("127.0.0.1"),
		Port: 30000,
	})
	if err != nil {
		fmt.Println("启动server失败，err:", err)
		return
	}

	process(listener)
}
