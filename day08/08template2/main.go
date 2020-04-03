package main

import (
	"fmt"
	"net/http"
	"text/template"
)

// 模板嵌套

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./index.html", "./ul.html")
	if err != nil {
		fmt.Println("parse failed, err:", err)
		return
	}
	t.Execute(w, nil)
}

func main() {
	http.HandleFunc("/index", index)
	http.ListenAndServe("127.0.0.1:9090", nil)
}
