package split

import (
	"fmt"
	"strings"
)

// 定义一个切割字符串的包

// Split 用sep分割s
// a:b:c : --> [a b c]
func Split(s, sep string) []string {
	count := strings.Count(s, sep)       // 数一下字符串s中包含多少个sep
	result := make([]string, 0, count+1) // 根据sep的数量初始化切片
	index := strings.Index(s, sep)
	for index >= 0 {
		result = append(result, s[:index])
		s = s[index+len(sep):]
		index = strings.Index(s, sep)
	}
	result = append(result, s)
	return result
}

// Add 实现字符串相加的函数
func Add(a, b string) string {
	return fmt.Sprintf("%s%s", a, b)
}
