package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// HTTP server端

func sayHello(w http.ResponseWriter, t *http.Request) {
	// fmt.Fprint(w, "Hello 沙河！")
	// w.Write([]byte("Hello 沙河"))

	// 从hello.txt文件中读取数据写入到w中
	data, err := ioutil.ReadFile("./hello.html")
	if err != nil {
		fmt.Println("read from file failed, err:", err)
		return
	}
	w.Write(data)
}
func main() {
	http.HandleFunc("/", sayHello) // 注册一个处理 / 的函数
	// 启动服务
	err := http.ListenAndServe("127.0.0.1:9090", nil)
	if err != nil {
		panic(err)
	}
}
