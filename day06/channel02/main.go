package main

import "fmt"

// 无缓冲通道和缓冲通道

func recv(ch chan bool) {
	ret := <-ch // 阻塞
	fmt.Println(ret)
}

func main() {
	ch := make(chan bool, 1)
	ch <- false
	// len:获取数据量，cap：获取容量
	fmt.Println(len(ch), cap(ch))
	go recv(ch)
	ch <- true
	fmt.Println("main函数结束")
}
