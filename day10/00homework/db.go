package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// 当你不知道变量名首字母是应该大写还是小写的时候，那你就用小写
// 保证最小暴露原则
var db *sqlx.DB

// 初始化数据库
func initDB() (err error) {
	// dsn := "root:123456@tcp(127.0.0.1:3306)/go_test?parsetime=true"
	dsn := "root:123456@tcp(127.0.0.1:3306)/go_test"
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return err
	}
	// 连接成功
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(16)
	return
}

// 创建用户的函数
func createUser(username, password string) error {
	sqlStr := "insert into userinfo(username,password) values(?,?)"
	_, err := db.Exec(sqlStr, username, password)
	if err != nil {
		fmt.Println("插入userinfo数据失败！")
		return err
	}
	return nil
}

// 查询数据库
func queryUser(username, password string) error {
	sqlStr := "select id from userinfo where username=? and password=? limit 1"
	var id int64
	err := db.Get(&id, sqlStr, username, password)
	if err != nil {
		return err
	}
	return nil
}
