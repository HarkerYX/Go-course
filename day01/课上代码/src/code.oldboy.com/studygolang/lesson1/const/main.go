package main
import "fmt"
const pi = 3.14

// 批量声明常量
// const (
// 	a = 100
// 	b = 1000
// 	c = 1000
// 	d = 1000
// )

// iota 枚举
// const (
// 	aa = iota  // 0
// 	bb = iota //1
// 	cc = iota //2
// 	dd = iota //3
// )

// const (
// 	n1 = iota//0
// 	n2 = iota  //1    
// 	_ = iota //2
// 	n4 = iota //3
// )

// const (
// 	n1 = iota  // 0
// 	n2 = 100  // 100
// 	n3 = iota  //2
// 	n4 = iota //3
// )
// const n5 = iota // 0


// const (
// 	_  = iota//0
// 	KB = 1 << (10 * iota)  // 1<<10 ==> 2的10次方 = 1024
// 	MB = 1 << (10 * iota)  // 1<<20  = 1024 * 1024
// 	GB = 1 << (10 * iota)  // 1<<30
// 	TB = 1 << (10 * iota)  // 1<<40
// 	PB = 1 << (10 * iota)  // 1<<50
// )

const (
	a, b = iota + 1, iota + 2 // iota=0; a=1 b=2
	c, d = iota + 1, iota + 2 // iota=1; c=2 d=3                 
	e, f = iota + 1, iota + 2 // iota=2; e=3 f=4                   
)

// 常量
func main(){
	// pi = 3.1415  //常量是不允许修改的
	fmt.Println(pi)
	fmt.Println(a,b,c,d,e,f)
	// fmt.Println(aa,bb,cc,dd)
	// fmt.Println(n1,n2,n3,n4,n5)

}

// iota 
// 0. const声明如果不写，默认就和上一行一样
// 1. 遇到const iota就初始化为零
// 2. const中每新增一行变量声明iota就递增1