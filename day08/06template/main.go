package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

// template 示例

type User struct {
	UserName string
	Password string
	Age      int
}

func info(w http.ResponseWriter, r *http.Request) {
	// 打开一个模板文件
	htmlByte, err := ioutil.ReadFile("./info.html")
	if err != nil {
		fmt.Println("read html failed, err:", err)
		return
	}
	// 添加自定义的方法要在parse模板文件之前添加

	// 1. 自定义一个函数
	// 自定义一个夸人的模板函数
	kuaFunc := func(arg string) (string, error) {
		return arg + "真帅", nil
	}
	// 2. 把自定义的函数告诉模板系统
	// template.New("info") // 创建一个Template对象
	// template.New("info").Funcs(template.FuncMap{"kua": kuaFunc}) // 给模板系统追加自定义函数
	// 解析模板

	// 链式操作？
	// 原理：每一次执行完方法之后返回操作的对象本身
	t, err := template.New("info").Funcs(template.FuncMap{"kua": kuaFunc}).Parse(string(htmlByte))
	if err != nil {
		fmt.Println("parse html file failed, err:", err)
		return
	}
	// template.ParseFiles("./info.html")
	// 用数据去渲染模板
	// data := "《我的世界》"
	// t.Execute(w, data)

	// user := User{
	// 	"豪杰",
	// 	"1234",
	// 	18,
	// }

	// t.Execute(w, user)

	userMap := map[int]User{
		1: {"豪杰", "1234", 18},
		2: {"韩鑫", "1234", 28},
		3: {"管大妈", "1234", 38},
	}
	t.Execute(w, userMap)
	// userSlice := []User{
	// 	{"豪杰", "1234", 18},
	// 	{"韩鑫", "1234", 28},
	// 	{"管大妈", "1234", 38},
	// }
	// t.Execute(w, userSlice)
}

func main() {
	http.HandleFunc("/info", info)
	http.ListenAndServe("127.0.0.1:8080", nil)
}
