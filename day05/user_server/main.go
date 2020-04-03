package main

import "code.oldboy.com/studygolang/day05/mylogger"

var logger mylogger.Logger

// 一个使用自定义日志库的用户程序
func main() {
	// logger = mylogger.NewFileLogger("Info", "./", "xxx.log")
	// defer logger.Close()
	logger = mylogger.NewConsoleLogger("info")
	for {
		sb := "管大妈"
		logger.Debug("%s是个好捧哏", sb)
		logger.Info("Info 这是一条测试的日志")
		logger.Error("Error 这是一条测试的日志")
	}
}

// 按照时间切分（一个小时一切）
// 文件日志的结构体里面记录一下上一次切分的小时
// 每一次写日志的时候 拿当前的时间的小时数和上一次切分的小时数作比较
// 如果小时数不一致就切分 否则就不切分
