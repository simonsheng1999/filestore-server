### 划重点

关于 `net/http` 包：

- `net/http`本身基于 goroutine 实现, 通过新建协程处理新的连接任务;
- 默认是长连接： net/http 客户端发起请求时 header 标记 HTTP/1.1;
- 连接可复用：默认创建连接池；
- 关于连接池使用：池中找不到空闲连接时，会重新 new 一个连接，而不会阻塞等待一个连接；
- 关于连接断开：如果对端关闭连接，由于 Go Runtime 会在底层进行 epoll wait，监听 close 事件并关闭相关 fd 资源，上层应用可以被告知哪些连接已关闭，从而进行相关的逻辑处理；
- 关于`time_wait`与`close_wait`：收到对方的 fin 请求， 内核会将连接置为`close_wait`状态;主动发起关闭连接请求的一方，会将连接置为`time_wait`状态。

net/http 服务监听流程示意图:

<img src='/doc/nethttp.png'></img>

-----

补充：

- 静态资源配置(支持用户访问/static目录下的文件):


```go
func main() {
	// 静态资源处理
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("./static"))))
    // ...
}
```
