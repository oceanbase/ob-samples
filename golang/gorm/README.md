# Connect OceanBase with Golang (gorm.io/driver/mysql)

English | [简体中文](README-CN.md)

This article describes how to connect to the OceanBase database through `gorm.io/driver/mysql`.

For details about `gorm.io/driver/mysql`, you can refer to [gorm.io/driver/mysql](https://gorm.io/driver/mysql).

## Quick Start

You can use `conf` to create a database connection, please refer to [gorm.io/docs/](https://gorm.io/docs) for details.

Take [example.go](example.go) code as an example.

Modify input parameters in `run.sh`, e.g. `host, port, username, password, database`, and use `run.sh` to run the example code.

```bash
sh run.sh
```

### Use PreparedStatement

Log in to OceanBase with a root user and run the following command:

```bash
alter system set _ob_enable_prepared_statement = true;
```
