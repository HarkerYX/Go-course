# day07 课上笔记



坚持这个两个字说起来很空，但是真正做到的人往往都能成功！



# 内容回顾

### goroutine

1. 并发和并行
2. goroutine是什么？
   1. 用户态线程
3. 启动goroutine 
   1. 使用go关键字 
4. goroutine的原理
   1. M:N ：将m个goroutine调度到n个操作系统线程上。n默认是操作系统的逻辑核心数
5. goroutine和OS线程的区别
   1. goroutine是用户态的线程，初始开销很小（2KB），可以随着需求扩充，最大能达1GB。
   2. 一个OS线程是很重量级的，通常内存开销达2MB。
   3. 通过`runtime.GOMAXPROCS`()  用来设置Go并发使用的CPU核数，1.5版本之后默认跑满CPU。
6. Goroutine的特点
   1. 一个goroutine对应一个函数，这个函数就是要做的事情
   2. main函数就是一个goroutine
   3. 当goroutine对应的函数返回的时候goroutine就结束了
   4. main函数所在的goroutine结束了，由它启动的那些goroutine也就结束了
7. sync.WaitGroup
   1. 它是一个结构体类型 var wg sync.WaitGroup
   2. 三个方法
      1. wg.Add(n)
      2. wg.Done()
      3. wg.Wait()

### channel

1. channel是一种类型，一种引用的类型

2. channel的声明和初始化

   1. `var ch chan int`
   2. channel声明之后要使用make函数初始化之后才能使用 , `ch = make(chan int, [cap])`

3. channel的3个操作

   1. 发送： `ch  <- 100`
   2. 接收：`<- ch`,可以使用变量接收值：`a:= <- ch` 也可以直接丢弃：`<-ch`
   3. 关闭：`close(ch)` 
      1. 关闭后的通道还是可以取值，取完之后返回的是类型零值。
      2. 关闭的通道不能再发送值
      3. 关闭的通道不能再关闭

4. 无缓冲区channel和有缓冲区channel

   1. 无缓冲区channel又称为同步channel，必须有接收才能发送，否则会一致阻塞。
   2. 有缓冲区的channel, 超过容量就阻塞。

5. 优雅地从通道取值(能判断通道是否被关闭)

   1. `v, ok := <-ch` :如果通道被关闭ok返回的是false
   2. `for v := range ch{}`

6. 单向通道

   1. 只能接收(`<-ch`)或者只能发送(`ch<-`)的通道
   2. 多用在函数传参的时候，限制某个通道在函数中只能做什么类型的操作

7. select多路复用

   1. 同一时刻可以对多个通道做发送和接收操作

      ```go
      
      ch := make(chan int, 1)
      for i := 0; i < 10; i++ {
      	select {
      	case ch <- i:
      	case ret := <-ch:
      		fmt.Println(ret)
      	}
      }
      ```

8. 通道是线程安全的。

### 并发控制与锁

1. 很多并发的场景下需要goroutine之间做协同处理

2. 如果多个goroutine操作同一个全局变量的时候，就会存在数据竞争

3. 互斥锁

   1. sync.Mutex,，它是一个结构体类型
   2. 声明锁
      1. `var lock sync.Mutex`
   3. 操作
      1. 加锁：`lock.Lock()`
      2. 解锁：`lock.Unlock()`

4. 读写锁

   1. `sync.RWMutex`适用于**读多写少**的场景。类比网站数据库的读写分离。
   2. 声明读写锁
      1. `var rwLock sync.RWMutex`
   3. 操作
      1. 加读锁：`rwLock.RLock()`
      2. 解读锁：`rwLock.RUnlock()`
      3. 加写锁：`rwLock.Lock()`
      4. 解写锁：`rwLock.Unlock()`

5. sync.Map

   1. 内置的map不是并发安全的
   2. 并发场景下推荐使用`sync.Map` 
   3. sync.Map的使用详见[博客](<https://www.liwenzhou.com/posts/Go/14_concurrence/#autoid-4-3-0>)

6. sync.Once

   1. 详见[博客](<https://www.liwenzhou.com/posts/Go/14_concurrence/#autoid-4-2-0>)

   2. 闭包的应用

      ​	

### 原子操作

​	内置整数支持的操作

### 上周作业

详见课上代码 homework文件夹



# 今日内容

## Go测试

[博客地址](<https://www.liwenzhou.com/posts/Go/16_test/>)

单元测试：开发人员自测。

一个大的程序十分很多功能单元的，开发完一个单元/模块之后我们自己进行测试。

单元测试很重要！ TDD:测试驱动开发

Go内置testing包做单元测试



测试相关的文件名为：`xxx_test.go`

### testing包

 	1. 单元测试 函数名以`TestXxx`
 	2. 基准测试
 	3. 示例函数



### 单元测试(TestXxx)

#### 单个测试

```go

func TestSplit(t *testing.T) {
   
	t.Log("测试字符串中包含分隔符的情形")
	got := Split("a:b:c", ":") // 调用函数得到的结果
	want := []string{"a", "b", "c"}  //期望得到的结果

	if ok := reflect.DeepEqual(got, want); !ok {
		t.Fatalf("期望得到:%v，实际得到：%v/n", want, got)
	}
        
}
```



#### 测试组

```go
func TestSplit(t *testing.T) {
	// 定义一个存放测试数据的结构体
	type test struct {
		str  string
		sep  string
		want []string
	}

	// 创建一个存放所有测试用例的map
	var tests = map[string]test{
		"normal": test{"a:b:c", ":", []string{"a", "b", "c"}},
		"none":   test{"a:b:c", "*", []string{"a:b:c"}},
		"multi":  test{"abcfbcabcd", "bc", []string{"a", "f", "a", "d"}},
		"num":    test{"1231", "1", []string{"", "23", ""}},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			ret := Split(tc.str, tc.sep)
			if !reflect.DeepEqual(ret, tc.want) {
				t.Errorf("期望得到:%#v，实际得到：%#v/n", tc.want, ret)
			}
		})
	}
}
```



#### 覆盖率

查看测试代码的覆盖率：`go test -cover`

将覆盖率的日志文件输出到文件：`go test -cover -coverprofile=c.out`

以HTML方式打开上一步生成的文件：`go tool cover -html=c.out`



#### 推荐阅读

[搞定Go单元测试(一)——基础原理 ](https://juejin.im/post/5ce93447e51d45775746b8b0)

### 基准测试（BenchmarkXxx)

#### 基本格式

```go
func BenchmarkName(b *testing.B){
    // ...
}
```

## Go网络编程

互联网协议：

​	OSI七层模型

​	IP

​	TCP

​		三次握手 四次挥手

​	UDP

​	HTTP



Go语言实现TCP客户端和服务端

粘包

Go语言实现UDP客户端和服务端

Go语言实现HTTP客户端和服务端



# 后面课程安排

## 前端 ：<https://www.cnblogs.com/liwenzhou/p/9959979.html>

![1558864965001](./assets/1558864965001.png)

## 数据库

MySQL 

## gin框架



# 本周作业

写一个回文判断的函数，编写单元测试和基准测试

abcba

油灯少灯油

Madam,I’mAdam

