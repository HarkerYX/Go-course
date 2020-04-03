package split

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

/*
func TestSplit(t *testing.T) {
	t.Log("测试字符串中包含分隔符的情形")
	got := Split("a:b:c", ":")
	want := []string{"a", "b", "c"}

	if ok := reflect.DeepEqual(got, want); !ok {
		t.Fatalf("期望得到:%v，实际得到：%v\n", want, got)
	}
}

// 如果字符串中不包含分隔符 测试结果是否正确
func TestNoneSplit(t *testing.T) {
	t.Log("测试字符串中不包含分隔符的情形")
	got := Split("a:b:c", "*")
	want := []string{"a:b:c"}

	if ok := reflect.DeepEqual(got, want); !ok {
		t.Fatalf("期望得到:%v，实际得到：%v\n", want, got)
	}
}

// 分隔符是多个字符的
func TestMultiSepSplit(t *testing.T) {
	got := Split("abcfbcabcd", "bc")
	want := []string{"a", "f", "a", "d"}

	if ok := reflect.DeepEqual(got, want); !ok {
		t.Fatalf("期望得到:%v，实际得到：%v\n", want, got)
	}
}

*/

// 将多个测试用例放到一起组成 测试组
func TestSplit(t *testing.T) {
	// 定义一个存放测试数据的结构体
	type test struct {
		str  string
		sep  string
		want []string
	}

	// 创建一个存放所有测试用例的map
	var tests = map[string]test{
		"normal": test{"a:b:c", ":", []string{"a", "b", "c"}},
		"none":   test{"a:b:c", "*", []string{"a:b:c"}},
		"multi":  test{"abcfbcabcd", "bc", []string{"a", "f", "a", "d"}},
		"num":    test{"1231", "1", []string{"", "23", ""}},
	}

	for name, tc := range tests {
		// ret := Split(tc.str, tc.sep)
		// if !reflect.DeepEqual(ret, tc.want) {
		// 	t.Errorf("测试用例：%v失败,期望得到:%#v，实际得到：%#v\n", name, tc.want, ret)
		// }
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			// 测试之前做点什么
			t.Log("要开始测试啦！")
			defer func() {
				t.Log("沙河出太阳了！")
			}()
			ret := Split(tc.str, tc.sep)
			t.Log("管大妈 SB")
			if !reflect.DeepEqual(ret, tc.want) {
				t.Errorf("期望得到:%#v，实际得到：%#v\n", tc.want, ret)
			}
		})
	}
}

// 基准测试
func BenchmarkSplit(b *testing.B) {
	// b.Log("这是一个基准测试")
	for i := 0; i < b.N; i++ {
		Split("a:b:c", ":")
	}
}

// 并行测试
func BenchmarkSplitParallel(b *testing.B) {
	// b.SetParallelism(1) // 设置使用的CPU数
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Split("沙河有沙又有河", "沙")
		}
	})
}

// 整个测试之前做的事儿和之后做的事儿
func TestMain(m *testing.M) {
	fmt.Println("write setup code here...") // 测试之前的做一些设置
	// 如果 TestMain 使用了 flags，这里应该加上flag.Parse()
	retCode := m.Run()                         // 执行测试
	fmt.Println("write teardown code here...") // 测试之后做一些拆卸工作
	os.Exit(retCode)                           // 退出测试
}

// 示例函数
func ExampleAdd() {
	fmt.Println(Add("管大妈", "DSB"))
	// OutPut: 管大妈DSB
}
