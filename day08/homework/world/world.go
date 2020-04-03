package world

import "unicode"

// 跟文本处理相关的函数

// IsPalindrome 是一个回文判断函数，返回布尔值
// func IsPalindrome(s string) bool {
// 	s2 := []rune(s)
// 	lenStr := len(s2)
// 	for i := 0; i < lenStr/2; i++ {
// 		// abcde  lenStr:5
// 		if s2[i] != s2[lenStr-i-1] {
// 			return false
// 		}
// 	}
// 	return true
// }

func IsPalindrome(s string) bool {
	var letters []rune
	for _, l := range s {
		// 判断l是不是一个letter
		if unicode.IsLetter(l) {
			// unicode.ToLower(l)转换成小写
			letters = append(letters, unicode.ToLower(l))
		}
	}

	for i := 0; i < len(letters); i++ {
		if letters[i] != letters[len(letters)-i-1] {
			return false
		}
	}
	return true
}
