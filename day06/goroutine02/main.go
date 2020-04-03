package main

import (
	"runtime"
	"sync"
)

var wg sync.WaitGroup

// func a() {
// 	defer wg.Done()
// 	// for i := 1; i < 100; i++ {
// 	// 	time.Sleep(time.Millisecond)
// 	// 	fmt.Println("A:", i)
// 	// }

// }

func b() {
	defer wg.Done()
	// for i := 1; i < 100; i++ {
	// 	fmt.Println("B:", i)
	// }
	var i int64
	for {
		i++
	}

}

func main() {
	runtime.GOMAXPROCS(1) // 设置我的Go程序只用一个逻辑核心 m:n中设置n为1
	wg.Add(10)
	// go a()
	for i := 0; i < 10; i++ {
		go b()
	}
	wg.Wait()
}
