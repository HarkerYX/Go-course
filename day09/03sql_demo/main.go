package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// 使用连接池方式连接MySQL

// DB 数据库连接句柄,全局变量
var DB *sql.DB

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

// User 用户结构体
type User struct {
	id   int
	name string
	age  string
}

// 查询单条
func queryRowDemo() {
	// 查询单行数据
	// sql:select id,name,age from user where id=1;
	var user User
	sqlStr := "select id,name,age from user where id=1"
	err := DB.QueryRow(sqlStr).Scan(&user.id, &user.name, &user.age)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	fmt.Printf("查询结果：%#v\n", user)
}

// 演示查询单行数据不调用row的Scan方法，会一直占用连接
func queryRowFaultDemo() {
	// 查询单行数据
	for i := 0; i < 100; i++ {
		fmt.Printf("第%d次查询。\n", i)
		// sql:select id,name,age from user where id=1;
		sqlStr := "select id,name,age from user where id=1"
		// 查询但是没有取结果 row会一直占用连接
		row := DB.QueryRow(sqlStr)
		fmt.Println(row)
		continue
	}
	fmt.Println("查询结束！")
}

// 查询多条
func queryMultiDemo() {
	var user User
	sqlStr := "select id, name, age from user where id > ?"
	rows, err := DB.Query(sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	defer func() {
		rows.Close() // 会释放数据库连接
	}()

	// 循环读取数据
	for rows.Next() {
		err := rows.Scan(&user.id, &user.name, &user.age)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("user:%#v\n", user)
	}
}

// 插入数据示例
// sql:insert into user (name, age) values("赵凯", 18);
func insertDemo() {
	sqlStr := "insert into user(name, age) values(?,?)"
	name := "赵凯"
	age := 18
	// Exec:执行
	ret, err := DB.Exec(sqlStr, name, age)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	// 拿到刚插入的数据id值（不同的数据库有不同的实现）
	theID, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get lastinsertid failed, err:%v\n", err)
		return
	}
	fmt.Println(theID)
}

// 更新数据
// 把官大妈那条数据的age字段 改成48
// sql:update user set age=? where id=?;
func updateDemo() {
	sqlStr := "update user set age=? where id=?"
	ret, err := DB.Exec(sqlStr, 48, 2)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	// 拿到受影响的行数
	num, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get affected row failed, err:%v\n", err)
		return
	}
	fmt.Println("受影响行数：", num)
}

// 删除数据
// sql: delete from user where id=2;
func deleteDemo() {
	sqlStr := "delete from user where id=?"
	ret, err := DB.Exec(sqlStr, 2)
	if err != nil {
		fmt.Printf("deleter failed, err:%v\n", err)
		return
	}
	num, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get affected row failed, err:%v\n", err)
		return
	}
	fmt.Println("受影响行数：", num)
}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/go_test"
	err := initDB(dsn)
	if err != nil {
		fmt.Printf("init DB failed, err:%v\n", err)
		return
	}
	// CRUD
	// 查询单条
	// queryRowDemo()
	//
	// queryRowFaultDemo()
	// 查询多条
	// queryMultiDemo()
	// 插入数据
	// insertDemo()
	// 更新数据
	// updateDemo()
	// 删除数据
	deleteDemo()
}
