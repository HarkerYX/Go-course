package main

import (
	"fmt"
	"io"
	"net"
)

// HTTP Client
func main() {
	conn, err := net.Dial("tcp", "www.liwenzhou.com:80")
	if err != nil {
		fmt.Println("访问jd.com")
		return
	}
	defer conn.Close()
	// 发送数据
	conn.Write([]byte("GET / HTTP/1.0\r\n\r\n"))
	// 接收数据
	var buf [1024]byte
	for {
		n, err := conn.Read(buf[:])
		if err == io.EOF {
			fmt.Print(string(buf[:n]))
			return
		}
		if err != nil {
			fmt.Println("接收数据失败,err:", err)
			return
		}
		fmt.Print(string(buf[:n]))
	}
}
