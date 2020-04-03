package math_pkg

import "fmt"

type Student struct {
	Name string
	Age  int
}

func init() {
	fmt.Println("我是一个无聊的计算器！init")
}

// Add 加法
func Add(x, y int) int {
	return x + y
}
