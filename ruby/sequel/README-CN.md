# Ruby连接 OceanBase 指南（sequel）

[English](README.md) | 简体中文

本文介绍如何通过sequel连接 OceanBase 数据库。

## 快速开始

在开始之前，需要先确保 mysql2,sequel已安装。

安装命令

```
gem install sequel
gem install mysql2
```

以 [example.rb](example.rb) 为例。

require 'sequel'

# 连接到OceanBase数据库
DB = Sequel.connect(
  adapter:  'mysql2',
  host:     '127.0.0.1',
  port:     2881,
  user:     'root',
  password: 'password',
  database: 'test'
)

if DB.test_connection
  puts "成功连接到OceanBase数据库"
else
  puts "连接到OceanBase数据库失败"
end

```

修改代码中的连接信息，之后你就可以直接使用命令行运行示例代码。

```bash
ruby  example.rb
```
