package main

import (
	"fmt"
	"time"
)

// 内置的time包

func timestamp2Timeobj(timestamp int64) {
	timeObj := time.Unix(timestamp, 0) //将时间戳转为时间格式
	fmt.Println(timeObj)
	year := timeObj.Year()     //年
	month := timeObj.Month()   //月
	day := timeObj.Day()       //日
	hour := timeObj.Hour()     //小时
	minute := timeObj.Minute() //分钟
	second := timeObj.Second() //秒
	fmt.Printf("%4d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}

func tickDemo() {
	ticker := time.Tick(time.Second) //定义一个1秒间隔的定时器

	for i := range ticker {
		fmt.Println(i) //每秒都会执行的任务
		// f1()
	}
}

// yyyy- m-d
// 2006-01-02 15:04:05    200612345

func formatDemo() {
	now := time.Now()
	// 格式化的模板为Go的出生时间2006年1月2号15点04分
	fmt.Println(now.Format("2006-01-02 15:04:05.000"))
	fmt.Println(now.Format("2006/01/02 15:04"))
	fmt.Println(now.Format("15:04 2006/01/02"))
	fmt.Println(now.Format("2006/01/02"))
}

func main() {
	// time.Time struct
	// now := time.Now()
	// fmt.Printf("%#v\n", now)
	// fmt.Println(now.Year())
	// fmt.Println(now.Month())
	// fmt.Println(now.Day())
	// fmt.Println(now.Hour())
	// fmt.Println(now.Minute())
	// fmt.Println(now.Second())
	// fmt.Println(now.Nanosecond())
	// // 时间戳
	// fmt.Println(now.Unix())
	// fmt.Println(now.UnixNano())
	// tm := now.Unix()
	// timestamp2Timeobj(tm)

	// 定时器
	// tickDemo()
	// 日期格式化 时间格式的数据 ==> 字符串格式
	formatDemo()
}
