package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// MySQL预处理

// DB 数据库连接句柄,全局变量
var DB *sql.DB

// User 用户结构体
type User struct {
	id   int
	name string
	age  string
}

func initDB(dsn string) (err error) {
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	err = DB.Ping()
	if err != nil {
		return err
	}
	// 连接上数据库了
	// 设置最大连接数
	DB.SetMaxOpenConns(50)
	// 设置最大的空闲连接数
	// DB.SetMaxIdleConns(20)
	return nil
}

// 事务
func transDemo() {
	// 把id=1的用户年龄+2岁；把id=4的用户的年龄-2岁
	tx, err := DB.Begin() // 开启事务
	if err != nil {
		fmt.Printf("begin trnas failed, err:%v\n", err)
		if tx != nil {
			tx.Rollback()
		}
		return
	}

	// 开始执行事务操作
	sql1 := "update user set age=age+? where id=?"
	_, err = tx.Exec(sql1, 2, 1) // 注意是tx.Exec()
	if err != nil {
		// 出错了，先回滚
		tx.Rollback()
		fmt.Printf("exec sql1 failed, err:%v\n", err)
		return
	}

	sql2 := "update user set age=age-? where id=?"
	_, err = tx.Exec(sql2, 2, 4)
	if err != nil {
		// 出错了，先回滚
		tx.Rollback()
		fmt.Printf("exec sql2 failed, err:%v\n", err)
		return
	}
	// 如果走到这里说明上面两条sql都执行成功，可以提交事务了
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		fmt.Printf("commit failed, err:%v\n", err)
		return
	}
	fmt.Println("两行记录更新成功！")
}

func main() {

	dsn := "root:123456@tcp(127.0.0.1:3306)/go_test"
	err := initDB(dsn)
	if err != nil {
		fmt.Printf("init DB failed, err:%v\n", err)
		return
	}

	transDemo()
}
