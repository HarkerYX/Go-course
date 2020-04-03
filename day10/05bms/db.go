package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)


// 跟数据库相关的操作
var db *sqlx.DB

func initDB()(err error){
	dsn := "root:123456@tcp(127.0.0.1:3306)/go_test"
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return err
	}
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(16)
	return
}

// 查数据
func queryAllBook()(bookList []*Book, err error){
	sqlStr := "select id,title,price from book"
	err = db.Select(&bookList, sqlStr)
	if err != nil {
		fmt.Println("查询所有书籍信息失败！")
		return
	}
	return 
}

// 插入数据
func insertBook(title string, price float64)(err error){
	sqlStr := "insert into book(title,price) values(?,?)"
	_, err = db.Exec(sqlStr, title, price)
	if err != nil {
		fmt.Println("插入书籍信息失败！")
		return
	}
	return 
}


// 删除数据
func deleteBook(id int64)(err error){
	sqlStr := "delete from book where id = ?"
	_, err = db.Exec(sqlStr, id)
	if err != nil {
		fmt.Println("删除书籍信息失败！")
		return
	}
	return 
}
