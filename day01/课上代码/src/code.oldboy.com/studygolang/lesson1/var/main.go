package main

import "fmt"

var alex = "dsb"
// alex := "alex"
// var alex string
// alex = "dsb"

// 定义有两个返回值的函数
func foo()(string, int){
	return "alex", 9000
}


func main(){
	// var name string
	// var age int
	// // 声明变量必须要使用
	// fmt.Println(name)
	// fmt.Println(age)

	// var (
	// 	a string // ""
	// 	b int //0
	// 	c bool//false
	// 	d string//""
	// )
	// a = "沙河"
	// b = 100
	// c = true
	// d = "100"
	// fmt.Println(a,b,c,d)
	// // 声明变量并且初始化
	// var x string = "老男孩"
	// fmt.Println(x)
	// // "欢迎%s登录" name
	// fmt.Printf("%s嘿嘿嘿%d\n", x, b)
	// // 类型推导(编译器根据变量初始值的类型 指定给变量)
	// var y = 200
	// var z = true
	// fmt.Println(y)
	// fmt.Println(z)
	// // 简短变量声明（只能在函数内部使用）
	// nazha := "嘿嘿嘿"  // var nazha string = "嘿嘿嘿"
	// fmt.Println(nazha)

	// // 调用foo函数
	// // _ （匿名变量）用于接收不需要的变量值
	// aa, _ := foo()
	// fmt.Println(aa)

	// 不能重复声明同名变量
	// var name = "alex"
	// // var name = "nazha"
	// fmt.Println(name)

	aa, _ := foo()
	fmt.Println(aa)
	_, bb := foo()
	fmt.Println(bb)

}