package main

import "fmt"

// 接口值 由两部分组成， 类型和值

func main() {
	var x interface{} // <Type,Value>
	var a int64 = 100
	var b int32 = 10
	var c int8 = 1
	x = a     // <int64, 100>
	x = b     // <int32, 10>
	x = c     // <int8, 1>
	x = 12.34 // <float64, 12.34>
	// x = false // <bool, false>
	// fmt.Println(x)
	// 类型断言（猜）
	// 如果猜对了，ok=true,value=对应类型的值
	// 如果猜错了, ok=false,value=对应类型的零值
	value, ok := x.(int)
	fmt.Printf("ok:%t  value:%#v value type:%T\n", ok, value, value)
}
