# day06 课上笔记

《北京折叠》

# 并发

### 并发和并行的区别

并发：同一时间段同时在做多个事情

并行：同一时刻同时在做多个事情



### 进程、线程、协程

进程：一个程序启动之后就创建了一个进程

线程：操作系统调度的最小单位

协程：用户态的线程（Python\PHP）

## goroutine

go运行时  runtime(Go语言作者帮我做的一些事情)

一个goroutine对应一个任务（函数）。

启动goroutine使用`go`关键字

`sync.WaitGroup`

`Add(i)`：计数器+i

`Done()`:计数器-1，最好用defer注册

`Wait()`:等待

### goroutine什么时候结束

一个goroutine对应一个任务（函数）。

函数结束了（return）goroutine就结束了

### goroutine和线程的区别

m:n

将m个goroutine调度到n个线程上执行。

### GOMAXPROCS

`runtime.GOMAXPROCS(n)`



## channel

通道是引用类型，必须初始化才能使用。

```go
ch1 := make(chan int, 10)
```



### 通道的基本操作

发送

接收

关闭

#### 关闭的通道有哪些特点：

1. 
2. 



### 无缓冲通道和缓冲通道

无缓冲的通道：4*100交接棒，又称为同步通道

有缓冲的通道：可以让程序实现异步操作

### 取值时判断通道关闭

方法一： ok判断

```go
// 利用for循环去通道ch1中接收值
for {
	ret, ok := <-ch1 // 使用 value, ok := <-ch1 取值方式，当通道关闭的时候 ok = false
	if !ok {
		break
	}
	fmt.Println(ret)
}
```



方法二： for range 循环

```go
// 利用for range 循环去通道ch1中接收值
for ret := range ch1 {
	fmt.Println(ret)
}
```

### 随机数

`math/rand`

```go

	// 给rand加随机数种子实现每一次执行都能产生真正的随机数
	rand.Seed(time.Now().UnixNano())
	ret1 := rand.Int63() // int64正整数
	fmt.Println(ret1)
	ret2 := rand.Intn(101) // [1, 101)
	fmt.Println(ret2)
```



### select

```go
select {
    case ch1 <- 1:
    	...
    case <- ch1:
    	...
    case <- ch2:
    	...
    
}
```

### 单向通道



```go
// chan *item :既能接收也能发送的一个队列
// chan<- *item :只能发送的一个队列（单向通道）
// <-chan *item :只能接收的一个队列（单向通道）
```





## 并发控制与锁



### 互斥锁



### 读写锁

只有在读的操作远远多于写的操作时使用读写锁才能提高性能。



### sync.Once 和 sync.Map



# 内容回顾

## 反射

## 上周作业



# 本周作业

利用channel和goroutine实现异步写日志

1. logger.Error(“xxxx”)的时候 将日志信息发送到channel,
   1. 日志信息是什么格式
   2. channel缓冲区设置为50000，来不及写的日志直接丢弃掉？？？
2. 后台起一个goroutine默默的从channel取日志然后往文件里面写

ini 配置解析库：[ini](<https://ini.unknwon.io/docs/intro/getting_started>)









