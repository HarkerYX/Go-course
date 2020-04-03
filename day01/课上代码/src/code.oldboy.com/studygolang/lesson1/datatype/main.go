package main

import (
	"fmt"
	"math"
)
func main(){
	var a int = 77
	var b int = 077
	var c int = 0xff
	fmt.Println(a, b, c)
	fmt.Printf("%b\n", a)
	fmt.Printf("%o\n", b)
	fmt.Printf("%x\n", c)
	// 求c变量的内存地址
	fmt.Printf("%p\n", &c)

	// 浮点数常量
	fmt.Println(math.MaxFloat64)
	fmt.Println(math.MaxFloat32)

	// 字符串转义
	fmt.Println("\"c:\\go\"")
	var s1 = "单行文本"
	var s2 = `这
	是 "这里不用转义"
	多行
	文本
	`
	fmt.Println(s1)
	fmt.Println(s2)
}