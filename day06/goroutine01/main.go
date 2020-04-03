package main

import (
	"fmt"
	"sync"
)

// 启动goroutine
// 利用sync.WaitGroup 实现优雅的等待
var wg sync.WaitGroup // 是一个结构体，它里面有一个计数器

func hello(i int) {
	defer wg.Done() // 计数器-1
	fmt.Println("Hello 沙河！", i)
	// if i == 8 {
	// 	panic("报错啦")
	// }
}

func main() {

	defer fmt.Println("哈哈哈")
	wg.Add(10) // 计数器+10
	for i := 0; i < 10; i++ {
		go hello(i) // 1. 创建一个goroutine 2. 在新的goroutine中执行hello函数

	}
	fmt.Println("Hello main func.")
	// time.Sleep(time.Second)
	// 等hello执行完（执行hello函数的那个goroutine执行完）
	wg.Wait() // 阻塞，一直等待所有的goroutine结束
	fmt.Println("main函数结束")
}
