# javascript连接OceanBase 指南

[English](README.md) | 简体中文

本文介绍如何通过javascript连接 OceanBase 数据库。

## 准备工作

需要创建一个项目，确认nodejs,npm,mysql2已经安装安装。

命令

```
mkdir example
cd example
npm init -y

npm install mysql2

```

创建 [index.js](index.js) 文件

```
const mysql = require('mysql2');

const connection = mysql.createConnection({
  host: '127.0.0.1',  // OceanBase服务器地址
  port: 2881,            // OceanBase端口
  user: 'root',          // 数据库用户名
  password: '',  // 数据库密码
  database: 'test'    // 数据库名称
});

// 连接到数据库
connection.connect(error => {
  if (error) {
    return console.error('连接到OceanBase数据库失败: ' + error.message);
  }

  console.log('成功连接到OceanBase数据库');

  // 这里可以执行其他数据库操作

  // 关闭连接
  connection.end(err => {
    if (err) {
      return console.error('关闭数据库连接失败: ' + err.message);
    }
    console.log('关闭数据库连接成功');
  });
});
```

修改代码中的连接信息，之后你就可以直接使用命令行运行示例代码。

```bash
sh run.sh
```
