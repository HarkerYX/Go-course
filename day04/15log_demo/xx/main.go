package main

import (
	"code.oldboy.com/studygolang/day04/15log_demo/mylog"
)

// 写了一个项目想要在代码中记录日志
// 要使用mylog这个包
func main() {
	fl := mylog.NewFileLogger(mylog.DEBUG, "./", "test.log")
	userID := 10
	fl.Debug("id是%d的用户一直在尝试登陆", userID)
}
