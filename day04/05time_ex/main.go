package main

import (
	"fmt"
	"time"
)

// time包练习题
// 1. 编写一个函数，接收Time类型的参数，函数内部将传进来的时间格式化输出为`2017/06/19 20:30:05`格式。
// 2. 编写程序统计一段代码的执行耗时时间，单位精确到微秒。

func printTime(t time.Time) {
	// `2017/06/19 20:30:05`
	//       "yyyy/m/d H:i:s"
	timeStr := t.Format("2006/01/02 15:04:05")
	fmt.Println(timeStr)
}

func calcTime() {
	start := time.Now()
	startTimestamp := start.UnixNano() / 1000
	fmt.Println("《钗头凤》红酥手 黄藤酒 满园春色宫墙柳。")
	time.Sleep(time.Millisecond * 30)
	end := time.Now()
	endTimestamp := end.UnixNano() / 1000
	// fmt.Println(time.Now().Unix())
	// fmt.Println(end)
	fmt.Printf("耗费了%d微秒\n", endTimestamp-startTimestamp)
}

func calcTime2() {
	start := time.Now()
	fmt.Println("《钗头凤》红酥手 黄藤酒 满园春色宫墙柳。")
	time.Sleep(time.Millisecond * 30)
	fmt.Println("耗费了", time.Since(start))
}

func main() {
	// now := time.Now()
	// printTime(now)
	// 2.
	calcTime()
	calcTime2()
}
