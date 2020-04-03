package main

import "fmt"

// 关于defer的面试题
func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x // 1. 返回值=5  2. x++ 3. RET指令  ==> 5
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5 // 1. (汇编)返回值=x(5) 2. x++  3.(汇编)RET ==> 6
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x // 1. (汇编)返回值=y(5)  2. x++  3.(汇编)RET  ==> 5
}
func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5 //1. (汇编)返回值=x(5)  2. x++(函数内部的x)  3.(汇编)RET  ==> 5
}
func main() {
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
}
