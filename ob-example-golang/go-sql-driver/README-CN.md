# Golang 连接 OceanBase 指南（使用 go-sql-driver/mysql）

[English](README.md) | 简体中文

本文介绍如何通过 `go-sql-driver/mysql` 连接 OceanBase 数据库。

关于 `go-sql-driver/mysql` 的详细信息，您可参考 [https://github.com/go-sql-driver/mysql](https://github.com/go-sql-driver/mysql)。

## 快速开始

您需要使用 dataSourceName 来创建数据库连接，详细信息请参考 [go-sql-driver/mysql 文档](https://github.com/go-sql-driver/mysql#dsn-data-source-name)。

以 [example.go](example.go) 代码为例

```go
var (
   host     = "127.0.0.1"
   port     = 2881
   dbName   = "test"
   username = "root@test"
   password = ""
)
dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", username, password, host, port, dbName)

db, err := sql.Open("mysql", dataSourceName)
if err != nil {
    log.Fatal(err)
}
defer db.Close()
```

修改代码中的连接信息，之后你就可以直接使用 run.sh 运行示例代码。

```bash
sh run.sh
```

### 使用 PreparedStatement

使用 root 用户登录 OceanBase，运行如下命令：

```bash
alter system set _ob_enable_prepared_statement = true;
```
