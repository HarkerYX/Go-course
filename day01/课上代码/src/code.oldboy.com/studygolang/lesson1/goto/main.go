package main 

import "fmt"
// goto
func main(){

	// for i:=0;i<5;i++{
	// 	if i ==2 {
	// 		// break  // 跳出for循环
	// 		continue  // 继续下一次循环
	// 	}
	// 	fmt.Println(i)
	// }

	// flag := false
	for i:=0;i<5;i++{
		for j:=0;j<3;j++{
			// 当i=2,j=2时跳出两层for循环
			if i ==2 && j ==2 {
				// flag = true
				goto haojie
				// break
			}
			fmt.Printf("%d--%d\n", i, j)
		}
		// 通过标志位来判断是否跳出外层for循环
		// if flag{
		// 	break
		// }
	}
	haojie:  // 定义一个标签 label
	fmt.Println("两层for循环结束")

	
}