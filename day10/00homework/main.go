package main

import (
	"fmt"
	"net/http"
	"text/template"
)

/*
day09作业
注册：
	实现用户注册用户名和密码，并写入数据库
登录：
	实现根据用户名和密码登录，去数据库校验用户名和密码是否正确

*/

func registerHandler(w http.ResponseWriter, r *http.Request) {
	// 根据请求方法的不同，来做不同的处理
	// 如果是POST请求，就提取用户提交的form表单数据，去数据库创建一行数据
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			// w.WriteHeader(http.StatusInternalServerError)
			w.WriteHeader(500)
		}
		username := r.FormValue("username")
		pwd := r.FormValue("password")
		// 往数据库里写
		err = createUser(username, pwd)
		if err != nil {
			// w.WriteHeader(http.StatusInternalServerError)
			w.WriteHeader(500)
			return
		}
		http.Redirect(w, r, "https://www.oldboyedu.com", 301)
	} else {
		// 如果是GET请求，就返回一个HTML页面，供用户输入注册信息
		t, err := template.ParseFiles("./register.html")
		if err != nil {
			// w.WriteHeader(http.StatusInternalServerError)
			w.WriteHeader(500)
		}
		t.Execute(w, nil)
	}

}

// 登录
func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			w.WriteHeader(500)
		}
		username := r.FormValue("username")
		pwd := r.FormValue("password")
		// 去数据库校验
		err = queryUser(username, pwd)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(500)
			return
		}
		http.Redirect(w, r, "https://www.oldboyedu.com", 301)

	} else {
		t, err := template.ParseFiles("./login.html")
		if err != nil {
			w.WriteHeader(500)
		}
		t.Execute(w, nil)
	}
}
func main() {
	err := initDB()
	if err != nil {
		panic(err)
	}
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/login", loginHandler)

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("启动http server失败！")
	}
}
