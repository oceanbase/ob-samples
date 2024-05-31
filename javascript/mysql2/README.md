# JavaScript Connection OceanBase Guide

English | [简体中文](README-CN.md)

This article introduces how to connect to the OceanBase database through JavaScript.

## prepare

We need to create a project to confirm nodejs, NPM, mysql2 has been installed and installed.

command

```
mkdir example
cd example
npm init -y

npm install mysql2

```

create [index.js](index.js) file

```
const mysql = require('mysql2');

const connection = mysql.createConnection({
  host: '127.0.0.1',  // OceanBase address
  port: 2881,            // OceanBase port
  user: 'root',          // username
  password: '',  // passwd
  database: 'test'    // database
});

// Connection OceanBase server
connection.connect(error => {
  if (error) {
    return console.error('Connection OceanBase faild: ' + error.message);
  }

  console.log('Connection OceanBase Successd');

  // Other Database Operations

  // Close Connection
  connection.end(err => {
    if (err) {
      return console.error('Close Connection Faild: ' + err.message);
    }
    console.log('Close Connection Successd');
  });
});
```

Modify the connection information in the code, and then you can directly run the example code using the command line.

```bash
sh run.sh
```
