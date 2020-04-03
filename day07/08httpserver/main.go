package main

import (
	"fmt"
	"net/http" // 专门为HTTP协议写的包
)

// HTTP Server

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1 style='color:green'>Hello 沙河！</h1>")
}

func main() {
	http.HandleFunc("/", sayHello)           // 注册路由：当你访问/ 就执行sayHello函数
	err := http.ListenAndServe(":9090", nil) // 建立监听
	if err != nil {
		fmt.Printf("http server failed, err:%v\n", err)
		return
	}
}
