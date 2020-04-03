package main

import "fmt"

// 实现接口时使用指针接受者和使用值接受者的区别

type animal interface {
	speak()
	move()
}
type cat struct {
	name string
}

// cat类型要实现animal的接口
// 方法一：使用值接收者
// func (c cat) speak() {
// 	fmt.Println("喵喵喵")
// }
// func (c cat) move() {
// 	fmt.Println("猫会动")
// }
// 方法二：使用指针接收者
func (c *cat) speak() {
	fmt.Println("喵喵喵")
}
func (c *cat) move() {
	fmt.Println("猫会动")
}

func main() {
	var x animal

	// hh := cat{"花花"} // cat
	// x = hh          // 实现animal接口的是*cat类型，hh此时是一个cat类型

	tom := &cat{"汤姆猫"} // *cat
	x = tom
	tom.speak() // (*tom).speak()
	tom.move()  // (*tom).move()
	fmt.Println(x)
}
