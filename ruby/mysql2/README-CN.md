# Ruby连接 OceanBase 指南（mysql2）

[English](README.md) | 简体中文

本文介绍如何通过 mysql2驱动连接 OceanBase 数据库。

## 快速开始

在开始之前，需要先确保 mysql2已安装。

安装命令

```
gem install mysql2
```

以 [example.rb](example.rb) 为例。

```
require 'mysql2'

def connect_to_oceanbase(host, port, username, password, database)
  begin
    # 创建数据库连接
    client = Mysql2::Client.new(
      host: host,
      port: port,
      username: username,
      password: password,
      database: database
    )

    puts "连接到OceanBase数据库成功"

    # 执行一个简单的查询
    results = client.query("SELECT DATABASE();")

    results.each do |row|
      puts "连接到的数据库: #{row['DATABASE()']}"
    end

    # 你可以在这里执行其他的SQL查询或操作

    client.close
    puts "MySQL连接已关闭"

  rescue Mysql2::Error => e
    puts "连接到OceanBase数据库失败: #{e.error}"
  end
end

if __FILE__ == $0
  connect_to_oceanbase(
    '127.0.0.1',
    2881,
    'root',
    '',
    'test'
  )
end

```

修改代码中的连接信息，之后你就可以直接使用命令行运行示例代码。

```bash
sh run.sh
```
