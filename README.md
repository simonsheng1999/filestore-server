## 划重点

- 理解和使用中间件

一个好的中间件是能够根据需求实现可插拔的。这就意味着可以在接口级别嵌入一个中间件，他就能直接正常的工作。它不会污染原有的代码，不会改变原始的逻辑，也不会影响原有的编码方式。需要它的时候就把它安插在指定的位置，不需要的时候就把它移除掉。

在 Go 语言中，在标准库中，中间件的应用是非常普遍的。一个比较典型的场景，在 net/http 里的某些业务 handler，可以配置 TimeoutHandler 作为中间件，用来处理请求超时的问题；这样业务 handler 专注于处理业务逻辑即可，无需亲自管理超时。

同样，在 API 鉴权方面，中间件也是发挥了不可或缺的作用：API 请求首先经过配置好的中间件，只有校验通过了才会将请求转发到下一级中间件或具体的业务 handler 进行逻辑的处理。以下是 API 拦截器使用示例:

```go
func HTTPInterceptor(h http.HandlerFunc) http.HandlerFunc {
    return http.HandlerFunc(
        func(w http.ResponseWriter, r *http.Request) {
            r.ParseForm()
            // TODO: 进行身份验证，比如校验cookie或token
            h(w, r)
        })
}

// 然后在创建路由时类似这么调用:
// http.HandleFunc("/test", HTTPInterceptor(YourCustomHandler))
```

## 关于 go modules

在 go1.11 之后，官方推出了 Go Modules 这原生的包管理工具；相对于 vendor, go mod 可以更有效的进行包版本管理。

- 在 1.12 及之前， go modules 需要配置环境变量来开启此功能:

```bash
# 分别有 on off auto
go env -w GO111MODULE=on
```

- 配置代理。因为众所周知的原因，有些包我们国内无法访问，一般需要通过代理(如 goproxy.cn):

```bash
go env -w GOSUMDB=sum.golang.google.cn
go env -w GOPROXY=https://goproxy.cn,direct
# 查看是否成功(go env的输出中包含代理信息)
go env
```

- go mod 初始化

```
go mod init <指定一个module名，如工程名>
```

在项目根目录下成功初始化后，会生成一个 go.mod 和 go.sum 文件。在之后执行 go build 或 go run 时，会触发依赖解析，并且下载对应的依赖包。

更具体的用法可以参考网上其他教程哦(如https://github.com/golang/go/wiki/Modules)。
