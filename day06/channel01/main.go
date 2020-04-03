package main

import "fmt"

// channel

func main() {
	// 定义一个ch1变量
	// 是一个channel类型
	// 这个channel内部传递的数据是int类型
	var ch1 chan int
	var ch2 chan string
	// channel是引用类型
	fmt.Println("ch1:", ch1)
	fmt.Println("ch2:", ch2)
	// make函数初始化(分配内存)：slice map channel
	ch3 := make(chan int, 10)
	// 通道的操作：发送、 接收、关闭
	// 发送和接收都用一个符号： <-
	ch3 <- 10 // 把10发送到ch3中
	// ch3 <- 20
	// <-ch3        // 从ch3中接收值，直接丢弃
	ret := <-ch3 // 从ch3中接收值，保存到变量ret中
	fmt.Println(ret)
	ch3 <- 9
	ch3 <- 8
	ch3 <- 7
	// 关闭
	close(ch3)
	// 1. 关闭的通道再接收，能取到对应类型的零值
	ret2 := <-ch3
	fmt.Println(ret2)
	// 2. 往关闭的通道中发送值 会引发panic
	// ch3 <- 20
	// 3. 关闭一个已经关闭的通道会引发panic
	// close(ch3)
}
