package main

import (
	"fmt"
	"net"

	"code.oldboy.com/studygolang/day08/02socket_stick/proto"
)

// 粘包现象 客户端
// socket_stick/client/main.go

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("dial failed, err", err)
		return
	}
	defer conn.Close()
	for i := 0; i < 20; i++ {
		msg := `Hello, Hello. How are you?`
		// 调用我们自己定义的协议，proto.Encode先封装数据包 再发送
		pkg, err := proto.Encode(msg)
		if err != nil {
			fmt.Println("encode msg failed, err:", err)
			return
		}
		conn.Write(pkg)
	}
}
