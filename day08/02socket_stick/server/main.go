package main

import (
	"bufio"
	"fmt"
	"net"

	"code.oldboy.com/studygolang/day08/02socket_stick/proto"
)

// 粘包现象 服务端
// socket_stick/server/main.go

func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	// 循环读
	for {
		msg, err := proto.Decode(reader) // 调用自定义的协议 proto.Decode去解包
		if err != nil {
			fmt.Println("decode msg failed, err:", err)
			return
		}
		fmt.Println("收到client发来的数据：", msg)
	}
}

func main() {

	listen, err := net.Listen("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		go process(conn)
	}
}
