package main

import (
	"fmt"
	"time"
)

// 空接口
func showType(a interface{}) {
	fmt.Printf("type:%T\n", a)
}

func main() {
	// var x interface{}
	// x = 100
	// fmt.Println(x)
	// x = "沙河"
	// fmt.Println(x)
	// x = false
	// fmt.Println(x)
	// x = struct{ name string }{name: "花花"}
	// fmt.Println(x)
	// showType(x)
	// showType(100)
	// showType(99.99)
	// showType(time.Second)

	// 定义一个值为空接口的map
	var stuInfo = make(map[string]interface{}, 100)
	stuInfo["豪杰"] = 100
	stuInfo["韩鑫"] = true
	stuInfo["王展"] = 99.99
	stuInfo["带帽哥"] = "呵呵"
	stuInfo["老学员"] = time.Now()
	fmt.Println(stuInfo)
}
