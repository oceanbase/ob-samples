# Python 连接 OceanBase 指南

本文介绍如何通过 Python 驱动连接 OceanBase 数据库。

## Python3.x 系列（使用 PyMySQL）
为了防止环境问题，推荐使用anaconda配置python3.x环境

PyMySQL 是在 Python3.x 版本中用于连接 MySQL 服务器的一个库。

PyMySQL 遵循 Python 数据库 API v2.0 规范，并包含了 pure-Python MySQL 客户端库。

有关 PyMySQL 的详细信息，您可参考 [官方文档](https://pypi.org/project/PyMySQL/) 和 [相关 API 参考文档](https://pymysql.readthedocs.io/en/latest/modules/index.html)。

### PyMySQL 安装

在使用 PyMySQL 之前需要确保 PyMySQL 已安装。您可前往 [下载](https://github.com/PyMySQL/PyMySQL) 界面进行下载安装。

PyMySQL 有以下两种安装方式：

1. 使用命令行安装

   ```bash
   python3 -m pip install PyMySQL
   ```

2. 源码编译

    ```bash
    git clone https://github.com/PyMySQL/PyMySQL
    cd PyMySQL/
    python3 setup.py install
    ```

### PyMySQL 使用

您可通过以下命令设置 connect 对象。

```bash
conn = pymysql.connect(
    host="localhost", 
    port=2881,user="root", 
    passwd="", 
    db="testdb"
)
```

示例：

```python
#!/usr/bin/python3

import pymysql

conn = pymysql.connect(host="localhost", port=2881,
                       user="root", passwd="", db="testdb")

try:
    with conn.cursor() as cur:
        cur.execute('SELECT * FROM cities')
        ans = cur.fetchall()
        print(ans)

finally:
    conn.close()
```

## Python2.x 系列（使用 MySQL-python）

MySQL-python 是 Python2.X 版本中用于连接 MySQL 服务器的一个库。

安装 MySQL-python 是为了使用 MySQLdb 连接和操作 OceanBase 数据库。MySQLdb 是 Python 连接 MySQL 数据库的接口，它实现了 Python 数据库 API 规范 V2.0，基于 MySQL C API 建立。

有关 MySQL-python 的详细信息，您可参考 [官方文档](https://pypi.org/project/MySQL-python/) 和 [Github 文档](https://github.com/farcepest/MySQLdb1).

### MySQL-python 安装

首先您需确保计算机上有 Python2.X 的环境，之后使用 pip 安装 MySQL-python。

```bash
pip install MySQL-python
```

### MySQL-python 使用

您可通过以下命令设置 connect 对象。

```bash
conn= MySQLdb.connect(
    host='127.0.0.1',
    port = 2881,
    user='root',
    passwd='',
    db ='testdb'
)    
```

示例：

```python
#!/usr/bin/python2

import MySQLdb

conn= MySQLdb.connect(
    host='127.0.0.1',
    port = 2881,
    user='root',
    passwd='',
    db ='testdb'
)

try:
    cur = conn.cursor()
    cur.execute('SELECT * from cities')
    ans = cur.fetchall()
    print(ans)
    
finally:
    conn.close()
```

​
