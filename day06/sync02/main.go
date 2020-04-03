package main

import (
	"fmt"
	"sync"
	"time"
)

// 读比写多的时候要使用读写锁 能够提高性能

var x int64
var wg sync.WaitGroup
var lock sync.Mutex     // 互斥锁
var rwLock sync.RWMutex // 读写互斥锁 : 多个goroutine同时读加的是读锁 写的时候加的是写锁

func read() {
	defer wg.Done()
	// lock.Lock() // 互斥锁
	rwLock.RLock()                   // 加读锁
	time.Sleep(time.Millisecond * 1) // 模拟读操作耗费1毫秒
	// lock.Unlock()                    // 释放互斥锁
	rwLock.RUnlock() // 释放读锁

}

func write() {
	defer wg.Done()
	rwLock.Lock() // 加写锁
	// lock.Lock() // 加互斥锁
	x = x + 1
	time.Sleep(time.Millisecond * 5) // 模拟写操作耗费了5毫秒
	// lock.Unlock()                    // 释放互斥锁
	rwLock.Unlock() // 释放写锁

}

func main() {
	start := time.Now()
	// 写10次
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}
	// 读1000次
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}
	wg.Wait()
	end := time.Now()
	fmt.Printf("耗费了%v.", end.Sub(start))
}
