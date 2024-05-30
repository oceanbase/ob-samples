# Ruby Connection OceanBase Guide (Mysql2)

English | [简体中文](README-CN.md)

This article introduces how to connect to OceanBase database through mysql2 driver.

## Quick Start

Before starting, it is necessary to ensure that mysql2 is installed.

Installation command

```
gem install mysql2
```

Taking [example.rb](example.rb)  as an example.

```python
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

Modify the connection information in the code, and then you can directly run the example code using the command line.

```bash
sh run.sh
```
