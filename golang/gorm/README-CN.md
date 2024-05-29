# Golang 连接 OceanBase 指南（使用 gorm.io/driver/mysql）

[English](README.md) | 简体中文

本文介绍如何通过 `gorm.io/driver/mysql` 连接 OceanBase 数据库。

关于 `gorm.io/driver/mysql` 的详细信息，您可参考 [gorm.io/driver/mysql](https://gorm.io/driver/mysql)。

## 快速开始

您需要使用代码中的 `conf` 来创建数据库连接，详细信息请参考 [gorm.io/docs/](https://gorm.io/docs)。

以 [example.go](example.go) 代码为例


修改run.sh代码中的连接信息, 如`host, port, username, password, database`，之后你就可以直接使用 run.sh 运行示例代码。

```bash
sh run.sh
```

### 使用 PreparedStatement

使用 root 用户登录 OceanBase，运行如下命令：

```bash
alter system set _ob_enable_prepared_statement = true;
```
