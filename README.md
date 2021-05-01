## 分块上传与断点上传逻辑流程示意图

<img src='/doc/mpupload_process.png' width="600px"></img>

## 快速启动应用

- 配置 MySQL 连接 (db/mysql/conn.go)

```golang
// 根据实际情况修改连接参数
db, err = sql.Open("mysql", "test:test@tcp(127.0.0.1:3306)/fileserver?charset=utf8")
```

- 启动应用

```bash
# 去到工程根目录下
# go1.13或以上开启go modules (初始化一次)
go mod init filestore-server
# go run 启动
go run main.go
```

- 前端主页效果

<img src="/doc/home.png"></img>

## 以下为代码更新提示:-)

- (2020-03-31)已将原有加载的远程资源文件(css, js 等)下载到项目本地，可以加快页面的加载速度

资源引用的 html 也同步跟新，比如 home.html

```html
<head>
  <script src="/static/js/jquery.min.js"></script>
  <link rel="stylesheet" href="/static/css/bootstrap.min.css" />
  <link rel="stylesheet" href="/static/css/bootstrap-theme.min.css" />
  <script src="/static/js/bootstrap.min.js"></script>
  <script src="/static/js/auth.js"></script>
  <script src="/static/js/layer.js"></script>
</head>
```

- (2020-04-02)完善文件合并逻辑

详细参考：https://git.imooc.com/coding-323/filestore-server/src/charter6/handler/mpupload.go

```go
// ...
	// 4. 合并分块 (备注: 更新于2020/04/01; 此合并逻辑非必须实现，因后期转移到ceph/oss时也可以通过分块方式上传)
	if mergeSuc := util.MergeChuncksByShell(ChunkDir+uploadID, MergeDir+filehash, filehash); !mergeSuc {
		w.Write(util.NewRespMsg(-3, "complete upload failed", nil).JSONBytes())
		return
	}
// ...
```

- (2020-04-04)完善取消分块上传逻辑

详细参考：https://git.imooc.com/coding-323/filestore-server/src/charter6/handler/mpupload.go

```go
// CancelUploadHandler : 通知取消上传
// 更新于2020-04
func CancelUploadHandler(w http.ResponseWriter, r *http.Request) {
    // ...
}
```

- (2020-04-15)完善断点续传与测试脚本

详细参考: https://git.imooc.com/coding-323/filestore-server/src/charter6/test
