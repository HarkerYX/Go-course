# day10 课上笔记





# 内容回顾

## MySQL

### database/sql

驱动：

```go
// 只用到了它这个包里面的init()
import	_ "github.com/go-sql-driver/mysql"
```

`sqlx` 



连接数据库

​	设置最大连接数

​	设置最大空闲连接数

查询单条

查询多条

执行SQL

事务

预处理



## 作业

```sql
 create table userinfo(
 id bigint(20) not null auto_increment,
 username varchar(20) not null,
 password varchar(20) not null,
 primary key(id)
 )engine=innoDB default charset=utf8mb4;
```

具体代码详见课上笔记`day10/homework`

## Redis

Redis基本数据类型及命令、与生产环境的结合

go操作Redis

## NSQ

安装和部署非常简单 、分布式、 消息队列



### nsqd

### nsqlookupd

### nsqadmin



# 今日内容

## go module

[博客地址](<https://www.liwenzhou.com/posts/Go/go_dependency/>)

[知乎连接](<https://zhuanlan.zhihu.com/p/59687626>)



## Gin框架

下载失败设置一下这个代理：

要设置`GOPROXY`，必须先开启`GO111MODULE`

```bash
SET GO111MODULE=on
```

设置代理：

```go
SET GOPROXY=https://goproxy.cn
```



### 下载Gin框架

```go
go get -u github.com/gin-gonic/gin
```



### Gin框架基本示例

```go
func indexHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "这是index页面！",
	})
}
func main() {
	// 启动一个默认的路由
	router := gin.Default()
	// 给/hello配置一个处理函数
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "Hello 沙河！",
		})
	})

	router.GET("/index", indexHandler)
	// 启动webserver
	router.Run(":9090")
}
```



### RESTful API

MVC架构

![1561263459176](./assets/1561263459176.png)

前后端分离的架构

MVVM

![1561264028435](./assets/1561264028435.png)





[postman调试接口的工具](<https://www.getpostman.com/downloads/>)



### Gin框架的渲染



#### JSON

`c.JSON(状态码，能够被JSON序列化的数据gin.H{})`

#### HTML

`c.HTML(状态码, 模板文件,数据)`



**加载模板文件的两个方法**

```go

r.LoadHTMLGlob("/模板文件存放的路径/*")
r.LoadHTMLGlob("/模板文件存放的路径/**/*")

r.LoadHTMLFiles("1.html", "2.html", ...)
```



**加载静态文件**

HTML里用到的CSS、JS、image等

```go
r.Static("代码里写的路径", "实际存放静态文件的路径")
```



#### XML

`c.XML(状态码, 数据)`



#### YAML

`c.YAML(状态码，数据)`



#### Protobuf



### 参数解析

#### query string

```go
c.Query(“key”)
c.DefaultQuery(“key”， “默认值”)
```



#### form param

```go
c.PostForm("key")
c.DefaultPostForm("key"， “默认值”)
```



#### path param

```go
url/:key1/:key2

c.Param("key")
```



# 书籍管理的练习

#### 书籍的信息

​	id

​	书名

​	价格

### Web页面版增删改查

增删改查

1. 项目代码要分层

   ![1561283524663](./assets/1561283524663.png)

2. 路由跳转

   1. HTTP跳转

      1. `c.Redirect(状态码,"地址")`

   2. 路由跳转

      1. 在具体的Handler之间跳转

         ```go
         c.Request.URL.Path = "/book/list"
         r.HandleContext(c)
         ```



 	3. Gin取queryString、form表单、path param的方式
 	4. [template语法](<https://www.liwenzhou.com/posts/Go/go_template/>)





## 路由跳转

1. 路由跳转

   1. HTTP跳转

      1. `c.Redirect(状态码,"地址")`

   2. 路由跳转

      1. 在具体的Handler之间跳转

         ```go
         c.Request.URL.Path = "/book/list"
         r.HandleContext(c)
         ```

### Gin路由

#### 路由的方法



### 路由组

```go
liveGroup := r.Group("/live")
{
    liveGroup.GET("/index", Handler)
    ...
}
```



### 路由的原理

![1561285044515](D:/Go4期/day10/assets/1561285044515.png)



### Gin上传文件



#### 上传单个文件

```go
func uploadHandler(c *gin.Context){
	// 提取用户上传的文件
	fileObj,err := c.FormFile("filename")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err,
		})
		return
	}
	// fileobj：上传的文件对象
	// fileobj.filename // 拿到上传文件的文件名
	filePath := fmt.Sprintf("./%s", fileObj.Filename)
	// 保存文件到本地的路径
	c.SaveUploadedFile(fileObj, filePath)
	c.JSON(http.StatusOK, gin.H{
		"msg": "OK",
	})
}
```



#### 上传多个文件

```go
func main() {
	router := gin.Default()
	// 处理multipart forms提交文件时默认的内存限制是32 MiB
	// 可以通过下面的方式修改
	// router.MaxMultipartMemory = 8 << 20  // 8 MiB
	router.POST("/upload", func(c *gin.Context) {
		// Multipart form
		form, _ := c.MultipartForm()
		files := form.File["file"]

		for index, file := range files {
			log.Println(file.Filename)
			dst := fmt.Sprintf("C:/tmp/%s_%d", file.Filename, index)
			// 上传文件到指定的目录
			c.SaveUploadedFile(file, dst)
		}
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("%d files uploaded!", len(files)),
		})
	})
	router.Run()
}
```



# 作业

1. 把BMS的增删查 自己写一遍

2. 把编辑书籍的功能自己实现
3. 难点：
   1. 第一次是GET请求，拿到一个页面，页面上有默认值
   2. 第二次点提交是发POST请求，在数据库更新数据之后要返回/book/list页面
4. Blog项目
   1. 想一下表结构
      1. 用户
      2. 文章
      3. 评论
      4. 分类

# 未来内容

`context` 