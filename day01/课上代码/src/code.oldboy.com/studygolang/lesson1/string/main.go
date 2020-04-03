package main

import (
	"fmt"
	"strings"
)
// 字符串操作
func main(){
	s1:= "alexdsb"
	fmt.Println(len(s1))
	// 字符串拼接
	s2 := "Python"
	fmt.Println(s1+s2)
	s3 := fmt.Sprintf("%s----%s", s1, s2)
	fmt.Println(s3)
	// 分割
	ret := strings.Split(s1, "x")
	fmt.Println(ret)
	// 判断是否包含
	ret2 := strings.Contains(s1, "sb")
	fmt.Println(ret2)
	// 判断前缀和后缀
	ret3 := strings.HasPrefix(s1, "alex")
	ret4 := strings.HasSuffix(s1, "sb")
	fmt.Println(ret3, ret4)
	// 求子串的位置
	s4 := "applepen"
	fmt.Println(strings.Index(s4, "p"))
	fmt.Println(strings.LastIndex(s4, "p"))
	// join
	a1 := []string{"Python", "PHP", "JavaScript", "Ruby", "Golang"}
	fmt.Println(strings.Join(a1, "-"))
}