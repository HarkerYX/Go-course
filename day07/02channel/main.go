package main

import (
	"fmt"
	"sync"
)

// 内容回顾 select多路复用

func main() {
	// ch1 := make(chan int, 10)
	// ch2 := make(chan int, 1)

	// select {
	// case ch1 <- 100:
	// 	fmt.Println("111")
	// case <-ch2:
	// 	fmt.Println("<-ch2")
	// default:
	// }

	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case ch <- i:
		case ret := <-ch:
			fmt.Println(ret)
		}
	}

	var lock sync.Mutex
	lock.Lock()
	lock.Unlock()
}
