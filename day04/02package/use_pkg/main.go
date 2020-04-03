package main

// 这个文件是干什么的
import "fmt"

// 导入包时起别名
import m "code.oldboy.com/studygolang/day04/02package/math_pkg"

// import (
// 	"fmt"

// 	"code.oldboy.com/studygolang/day04/02package/math_pkg"
// )

const Mode = 1

func main() {
	m.Add(100, 200)
	stu := m.Student{Name: "haojie", Age: 18}
	fmt.Println(stu.Name)
	fmt.Println(stu.Age)
}
