## 划重点

​ 关于数据库访问，Golang 中提供了标准库 database/sql。不过它不是针对某种具体数据库的逻辑实现，而是一套统一抽象的接口。
真正与数据库打交道的，是各个数据库对应的驱动 Driver；在使用时需要先注册对应的驱动库，然后就能通过标准库 sql 中定义的接口来统一操作数据库。

- 创建 sql.DB 连接池

​ 我们来看一下如何创建 sdl.DB 连接池，以 MySQL 为例:

```go
import (
  "log"
  "os"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
)
​
func main() {
  db, err := sql.Open("mysql",
    "user:pwd@tcp(127.0.0.1:3306)/testdb")
  if err != nil {
    log.Fatal(err)
    os.Exit(1)
  }
  defer db.Close()

  err = db.Ping()
  if err != nil {
     // TODO do something
  }
  // ...
}
```

创建数据库连接是一种比较耗资源的操作，先要完成 TCP 三次握手，连接到数据库后需要进行鉴权和分配连接资源，因此建议使用长连接来避免频繁进行此类操作。在我们的 Go 应用中，sql.DB 自身就是会管理连接池的，一般实现为全局的连接池，不用重复进行 open/close 动作。

- CRUD 接口

```go
// 1. 返回多行数据，手动关闭结果集 (defer rows.Close())
db.Query()
// 2. 返回单行数据，不须手动关闭结果集
db.QueryRow()
// 3. 预先将一条连接(conn)与一条sql语句绑定起来，供重复使用
stmt, err := db.Prepare(sql)
stmt.Query(args)
// 4. 适用于执行增/删/更新等操作(不需要返回结果集)
db.Exec()
```

- 结果集合

Query()方法会返回结果集合，需要通过 rows.Next(), rows.Scan()方法来遍历结果集合。

- 事务

```go
// 开始事务
tx := db.Begin()
// 执行事务
db.Exec()
// ...
// 提交事务
tx.Commit()
// 如果失败的话，回滚
tx.Rollback()
```

- 小结

  用户层面所执行的 sql 语句，在底层其实是将 sql 语句串编码后传输到数据库服务器端再执行。因此本质上可以看作是 C/S 架构程序。

  对于如何访问不同的数据库，Go 的做法是抽离出具体代码逻辑，将数据库操作分为 database/sql 和 driver 两层。database/sql 负责提供统一的用户接口，以及一些不涉及具体数据库的逻辑(如连接池管理)；Driver 层负责实际的数据库通信。

(moxiaomomo)
