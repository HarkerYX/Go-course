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

func prepareInsertDemo() {
	sqlStr := "insert into user (name,age) values(?,?)"
	stmt, err := DB.Prepare(sqlStr) // 把要执行的命令发送给MySQL服务端做预处理
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	defer stmt.Close()
	// 执行重复的插入命令
	for i := 0; i < 10; i++ {
		name := fmt.Sprintf("stu%02d", i)
		stmt.Exec(name, i)
	}
}

// 预处理查询
func prepareQueryDemo() {
	sqlStr := "select id,name,age from user where id=?"
	stmt, err := DB.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	defer stmt.Close()
	for i := 0; i < 10; i++ {
		rows, err := stmt.Query(i)
		if err != nil {
			fmt.Printf("query failed, err:%v\n", err)
			continue
		}
		defer rows.Close()
		var user User
		for rows.Next() {
			err := rows.Scan(&user.id, &user.name, &user.age)
			if err != nil {
				fmt.Printf("scan failed, err:%v\n", err)
				return
			}
			fmt.Printf("user:%#v\n", user)
		}
	}
}

func main() {

	dsn := "root:123456@tcp(127.0.0.1:3306)/go_test"
	err := initDB(dsn)
	if err != nil {
		fmt.Printf("init DB failed, err:%v\n", err)
		return
	}

	// 插入10条数据
	// prepareInsertDemo()

	// 查询10次
	prepareQueryDemo()

}
