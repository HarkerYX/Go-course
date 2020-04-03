package main

import (
	"net/http"
	"text/template"
)

// 作业

func login(w http.ResponseWriter, r *http.Request) {
	// 判断用户是获取网页呢还是填完了要提交数据
	t, err := template.ParseFiles("./demo.html")
	if err != nil {
		panic(err)
	}
	if r.Method == "POST" {
		// 获取到用户提交的数据
		r.ParseForm() // 解析表单
		username := r.FormValue("username")
		pwd := r.FormValue("password")
		// 校验用户名和密码是否正确
		// 正常网站的登录逻辑应该是连接数据库然后校验用户名密码
		if username == "alex" && pwd == "alexdsb" {
			// 登陆成功
			http.Redirect(w, r, "http://www.oldboyedu.com", 302)
		} else {
			// 登陆失败
			// 在login页面显示错误提示信息
			t.Execute(w, "用户名或密码错误")
		}
	} else {
		t.Execute(w, nil)
	}
}

func main() {
	http.HandleFunc("/login", login)
	http.ListenAndServe(":9090", nil)
}
