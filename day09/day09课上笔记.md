# day09课上笔记



生活大于一切。

# 内容回顾

## HTML基础

HTML:超文本标记语言

HTTP：超文本传输协议



前端页面：

HTML + CSS + JavaScript



### HTML标签

head标签：给浏览器看的信息

body标签：给用户看的内容

 * 给用户展示信息用的标签： div、span、h1~h6、p、 a 、 img
 * 获取用户输入的标签
    * form表单
    * input、select、texarea

## Go template语法

归根结底就是字符串替换。

具体语法看博客。

## 作业

`net/http`写一个登录页面

详见课上代码`day09/01HTML` 

# 今日内容

## MySQL

关系型数据库：MySQL、postgreSQL、SQL Server、Oracle、sqlite

表的概念：

​	一行记录

​	表之间关联关系：外键关联、多对多关联、一对一关联

​	《漫画数据库》   讲数据库设计的三大范式

SQL(架构化查询语言):

​	操作数据库

​		修改用户名密码、权限

​	操作数据表

​		创建表、删除表、修改表、查看表

​	操作数据行

​		增删改查



#### 增删改查

 	1. 写原生的SQL
 	2. 使用ORM
      	1. 把增删改查映射成编程语言中对象的增删改查

![1560654113748](D:/Go/src/code.oldboy.com/studygolang/day09/assets/1560654113748.png)



​		2.Go语言ORM框架推荐使用GORM 

非关系型数据库：mongoDB、Redis、Memcached、ETCD、TiDB

#### 常用引擎

MySQL支持插件式的引擎。

MyISAM: 不支持事务，只有表级锁 | 查询快

InnoDB：支持事务、支持行级锁 ，现在主流用这个



#### MySQL主从

为什么要分主从？

​	读写分离

主从数据同步？

​	binlog

主从数据一致性问题

​	用缓存来做

#### 优化

​	索引

​	

## Go操作MySQL

本地登录MySQL

```bash
mysql -uroot -p123456
```

退出

```bash
mysql>quit;
```



查看MySQL中有哪些数据库：

show databases;

创建数据库

```sql
create database go_test;
```

切换数据库：

```sql
use go_test;
```

查看数据库中的表：

```sql
show tables;
```

创建表：

```sql
CREATE TABLE `user` (
    `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(20) DEFAULT '',
    `age` INT(11) DEFAULT '0',
    PRIMARY KEY(`id`)
)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;
```

查看表结构：

```sql
desc user;
```

查看详细建表语句：

```sql
show create table user/G
```

### 下载驱动

```bash
go get -u github.com/go-sql-driver/mysql
```

如果下载因为超时而失败，可以设置`goproxy`，详见[blog](<https://www.liwenzhou.com/posts/Go/fix_go_get/>)



### 注册驱动

```go
// 只用到了它这个包里面的init()
import _ "github.com/go-sql-driver/mysql"
```



### 连接数据库

```go
func main(){
	// dsn:"user:password@tcp(ip:port)/databasename"
	dsn := "root:123456@tcp(127.0.0.1:3306)/go_test"
	// 调用标准库中的方法
	// 前提是要注册对应数据库的驱动
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("open mysql failed, err:%v/n", err)
		return
	}
	// 尝试链接一下数据库，校验用户名密码是否正确...
	err = db.Ping()
	if err != nil {
		fmt.Printf("connect MySQL failed, err:%v/n", err)
		return
	}
	fmt.Println("连接数据库成功！")
}

```

### CRUD

查询单行

```go
DB.QueryRow()
```

查询多行

```go
DB.Query()
```

插入、更新、删除

```go
DB.Exec()
```

### 预处理

```go
DB.Prepare()
```

事务

```go
tx, err := DB.Begin()
```



### sqlx





## Redis



## NSQ



# 今日作业

1. 结合`net/http`和`sqlx`包实现一个注册及登陆的web程序。



