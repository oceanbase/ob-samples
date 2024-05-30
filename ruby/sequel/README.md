# Ruby Connection OceanBase Guide (sequel)

English | [简体中文](README-CN.md)

This article introduces how to connect to OceanBase database through sequel.

## Quick Start

Before starting, it is necessary to ensure that mysql2,sequel is installed.

Installation command

```
gem install sequel
gem install mysql2
```

Taking [example.rb](example.rb)  as an example.

```
require 'sequel'

# 连接到OceanBase数据库
DB = Sequel.connect(
  adapter:  'mysql2',
  host:     '127.0.0.1',
  port:     2881,
  user:     'root',
  password: '',
  database: 'test'
)

if DB.test_connection
  puts "成功连接到OceanBase数据库"
else
  puts "连接到OceanBase数据库失败"
end

```

Modify the connection information in the code, and then you can directly run the example code using the command line.

```bash
sh run.sh
```
