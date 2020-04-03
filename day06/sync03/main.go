package main

import (
	"fmt"
	"strconv"
	"sync"
)

// sync.Map

var m = make(map[string]int)

func get(key string) int {
	return m[key]
}

func set(key string, value int) {
	m[key] = value
}

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			// 将int类型转换成字符串类型
			set(key, n)                                // 给map设置键值对
			fmt.Printf("k=:%v,v:=%v\n", key, get(key)) // 打印键值对
			wg.Done()
		}(i)
	}
	wg.Wait()
}
