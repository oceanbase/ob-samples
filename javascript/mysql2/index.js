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
