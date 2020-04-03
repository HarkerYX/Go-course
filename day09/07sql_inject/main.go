package main

// SQL注入
// 过分相信用户
// 拿着用户输入的内容直接拼接字符串去执行

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

// DB 全局数据库连接对象（内置连接池的）
var DB *sqlx.DB

// User user表对应的结构体
type User struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
	Age  int    `db:"age"`
}

func initDB() (err error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/go_test"
	DB, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return
	}
	return nil
}

// sql注入示例
func sqlInjectDemo(name string) {
	// 自己拼接字符串
	sqlStr := fmt.Sprintf("select id, name, age from user where name='%s'", name)
	fmt.Printf("SQL:%s\n", sqlStr)

	var users []User
	err := DB.Select(&users, sqlStr)
	if err != nil {
		fmt.Printf("exec failed, err:%v\n", err)
		return
	}
	for _, u := range users {
		fmt.Printf("user:%#v\n", u)
	}
}
func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed, err:%v\n", err)
		return
	}
	// 正经的输入内容
	// sqlInjectDemo("haojie")
	// 不正经的输入内容
	// "select id, name, age from user where name='%s'"
	// sqlInjectDemo("heyoubin' or 1=1#")
	// sqlInjectDemo("xxx' union select * from user #")
	// sqlInjectDemo("xxx' and (select count(*) from user) <10 #")
}
