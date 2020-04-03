package main

import "fmt"

// 接口实现一个洗衣机
// 只要一个类型它实现了 wash() 和 dry() 方法，我们就称这个类型实现了xiyiji这个接口

type xiyiji interface {
	wash()
	dry()
}

type Haier struct {
	name  string
	price float64
	mode  string
}

// 田螺姑娘
type tianluo struct {
	name string
}

func (t tianluo) wash() {
	fmt.Println("田螺姑娘可以洗衣服~")
}
func (t tianluo) dry() {
	fmt.Println("田螺姑娘可以把衣服拧干~")
}

func (h Haier) wash() {
	fmt.Println("海尔洗衣机能洗衣服~")
}
func (h Haier) dry() {
	fmt.Println("海尔洗衣机自带甩干~")
}

func main() {
	var a xiyiji // 声明一个xiyijie类型的变量a
	h1 := Haier{ // 实例化了一个Haier结构体对象
		name:  "小神童",
		price: 998.98,
		mode:  "滚筒",
	}
	fmt.Printf("%T\n", h1)
	a = h1 // 接口是一种类型，一种抽象的类型。
	fmt.Println(a)
	tl := tianluo{
		name: "螺蛳粉",
	}
	a = tl
	fmt.Println(a)
}
