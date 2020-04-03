package main

import "fmt"

func f(ch chan int) {
	ret := <-ch // 接收值
	fmt.Println(ret)
}

// 内容回顾
func main() {
	var ch chan int
	ch = make(chan int, 100)

	// ch := make(chan int)
	// go f(ch) // 开启一个goroutine去执行f函数

	ch <- 100 // 无缓冲区的通道，没有任何接收 100就发送不进去
	fmt.Println("Hello 沙河！")
}
