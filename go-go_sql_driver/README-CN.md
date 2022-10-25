# Golang 连接 OceanBase 指南（使用 Go-SQL-Driver/MySQL）

本文介绍如何通过 Golang 驱动连接 OceanBase 数据库。

## 如何安装

### 通过 go get 安装（适用于Go V1.13 - V1.16）

安装命令如下：

```bash
go get -u github.com/go-sql-driver/mysql
```

关于 Go-SQL-Driver/MySQL 的详细信息，您可参考 [Github](https://github.com/go-sql-driver/mysql)。

### 通过 go install 安装

如果由于版本或网络的原因，无法通过 go get 命令安装时，可通过如下方法进行 go-sql-driver/mysql 安装。

1. 在 src 目录克隆 github 中的 go-sql-driver/mysql 仓库。

    ```bash
    cd /usr/local/go/src   
    git clone https://github.com/go-sql-driver/mysql.git 
    ```

2. 通过 go install 进行安装。

    ```bash
    go install mysql
    ```

>**注意**
>
>部分版本 go install 的默认执行目录可能不是 `/src`，可以通过 go install 执行后的报错判断实际目录。例如，报错 `cannot find package "mysql" in: /usr/local/go/src/vendor/mysql`，则应该将 mysql 文件夹放在 `/src/vendor` 目录下再执行安装命令。

## 如何使用

您可使用 `username:password@protocol(address)/dbname?param=value` 设置链接字符串。

以 test.go 代码为例：

```go
package main

import (
    "database/sql"
    "fmt"
    "log"
    
    _ "github.com/go-sql-driver/mysql" //填写 go-sql-driver/mysql 安装的准确路径。如果安装在 src 目录下，可以直接填 "mysql"。
)

type City struct {
    Id         int
    Name       string
    Population int
}

func main() {
    select_all()
    }

func select_all() {
    conn := "root:@tcp(127.0.0.1:2881)/testdb"
    db, err := sql.Open("mysql", conn)
    if err != nil {
        log.Fatal(err)
    }
    
    defer db.Close()
    
    if err != nil {
        log.Fatal(err)
    }
    
    res, err := db.Query("SELECT * FROM cities")
    if err != nil {
        log.Fatal(err)
    }
    
    defer res.Close()
    
    if err != nil {
        log.Fatal(err)
    }
    
    for res.Next() {
        
        var city City
        err := res.Scan(&city.Id, &city.Name, &city.Population)
        
        if err != nil {
            log.Fatal(err)
        }
        
        fmt.Printf("%v\n", city)
    }
}
```

参数说明：

```go
conn := "{username}:{password}@tcp({hostname}:{port})/{dbname}"
```

* **username**：租户的连接账户，租户的管理员用户名默认是 `root`。

* **password**：账户密码。

* **hostname**：OceanBase 数据库连接 IP，通常是一个 OBProxy 地址。

* **port**：OceanBase 数据库连接端口，也是 OBProxy 的监听端口，默认是 2883，可以自定义。

* **dbname**：需要访问的数据库名称。

示例：`conn := "root:@tcp(127.0.0.1:2881)/testdb"`

代码编辑完成后，可以通过如下命令运行：

```bash
//环境配置，根据 go 实际安装路径填写
export PATH=$PATH:/usr/local/go/bin

//运行 go 语言代码
go run test.go

//生成二进制文件后运行
go build test.go
./test
```

### 关于驱动 PreparedStatement 报错

使用 root 用户登录 OceanBase，运行如下命令：

```bash
alter system set _ob_enable_prepared_statement = true;
```
