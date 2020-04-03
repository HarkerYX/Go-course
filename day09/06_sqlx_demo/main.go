package main

// sqlx 示例

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

// 查询单条
func queryRowDemo() {
	sqlStr := "select id,name from user where id=?"
	var user User
	err := DB.Get(&user, sqlStr, 1)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	fmt.Printf("user:%#v\n", user)
}

// 查询多行
func queryMultiDemo() {
	sqlStr := "select id,name,age from user where id >?"
	var users []User
	err := DB.Select(&users, sqlStr, 0)
	if err != nil {
		fmt.Printf("select failed, err:%v\n", err)
		return
	}
	for _, user := range users {
		fmt.Printf("user:%#v\n", user)
	}
}

// 事务操作
func transDemo() {
	tx, err := DB.Beginx()
	if err != nil {
		if tx != nil {
			tx.Rollback()
		}
		fmt.Printf("begin trnas failed, err:%v\n", err)
		return
	}
	sql1 := "update user set age=age-? where id=?"
	tx.MustExec(sql1, 2, 1) // 名字带Must的一般表示出错就panic:
	sql2 := "updateheihei user set age=age+? where id=?"
	tx.MustExec(sql2, 2, 4) // 名字带Must的一般表示出错就panic:
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		fmt.Printf("commit failed, err:%v\n", err)
	}
	fmt.Println("两条数据更新成功！")
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed, err:%v\n", err)
		return
	}
	// 查询单条
	// queryRowDemo()
	// queryMultiDemo()
	transDemo()
}
