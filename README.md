## 划重点

- 秒传原理

场景： 同一文件有人上传过了后，其他人再要上传时就可以不用重复传了，只需要在数据库中增加一条映射记录即可，这样就算是完成了一次上传。

原理： 文件 hash 具有全局唯一性，不同文件的 sha1 值碰撞几率极低；前端上传前计算文件的 sha1 值，然后提交到服务端进行比对，不存在时才真正开始上传，否则秒传。

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
go mod init <指定一个modeul名，如工程名>
```

在项目根目录下成功初始化后，会生成一个 go.mod 和 go.sum 文件。在之后执行 go build 或 go run 时，会触发依赖解析，并且下载对应的依赖包。

更具体的用法可以参考网上其他教程哦(如https://github.com/golang/go/wiki/Modules)。
