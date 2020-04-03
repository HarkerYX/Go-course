package world

import "testing"

// world的测试文件

func TestIsPalindrome(t *testing.T) {
	// 定义一个表示测试用例的结构体
	type test struct {
		str  string
		want bool
	}
	// 用map表示一个测试组
	tests := map[string]test{
		"simple":       {"沙河有沙又有河", false},
		"englishFalse": {"abc", false},
		"englishTrue":  {"abcba", true},
		"ChineseTrue":  {"油灯少灯油", true},
		"withXx":       {"Madam,I’mAdam", true},
	}
	// 执行测试组里的每一个测试用例
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := IsPalindrome(tc.str) // 执行测试函数得到结果
			// 拿着得到的结果和期望的结果比较
			if got != tc.want {
				t.Errorf("want:%#v, got:%#v", tc.want, got)
			}
		})
	}
}
