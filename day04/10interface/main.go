package main

import "fmt"

type Cat struct{}

// 类型断言
// type nullInterface interface {}// 太繁琐

func ShowType(x interface{}) {
	// 因为我这个函数可以接收任意类型的变量
	// 类型断言
	v1, ok := x.(int)
	if !ok {
		// 说明猜错了
		fmt.Println("不是int")
	} else {
		fmt.Println("x就是一个int类型", v1)
	}
	v2, ok := x.(string)
	if !ok {
		fmt.Println("不是string")
	} else {
		fmt.Println("x是一个string类型", v2)
	}
}

// type nullInterface interface{}

func justifyType(x interface{}) {
	switch v := x.(type) {
	case string:
		fmt.Printf("x is a string，value is %v\n", v)
	case int:
		fmt.Printf("x is a int,value  is %v\n", v)
	case bool:
		fmt.Printf("x is a bool, value is %v\n", v)
	case Cat:
		fmt.Printf("x is a Cat struct,value is %v\n", v)
	case *string:
		fmt.Printf("x is a string poninter,value is %v\n", v)
	default:
		fmt.Println("unsupport type！")
	}
}

func main() {
	// var x interface{}
	// x = 100
	// ShowType(100)
	// ShowType("哈哈")
	// switch版类型断言
	justifyType(100)

	justifyType(Cat{})
	s := "哈哈"
	justifyType(&s)
}
