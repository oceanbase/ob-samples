# Connect OceanBase with Python (PyMySQL)

English | [简体中文](README-CN.md)

This article describes how to connect to the OceanBase database through `pymysql`.

## Quick Start

To prevent environmental problems, it is recommended to use anaconda to set up the python 3.x environment.

Before starting, you need to make sure PyMySQL is installed. PyMySQL is a package used to connect to MySQL server with Python 3.x. For details about the installation and usage of PyMySQL, you can refer to the [Official Documentation](https://pypi.org/project/PyMySQL/) and [API Documentation](https://pymysql.readthedocs.io/en/latest/modules/index.html).

Take [Test.py](Test.py) for example.

```python
import pymysql

conn = pymysql.connect(host="127.0.0.1", port=2881,
                       user="root@test", passwd="", db="test")
with conn.cursor() as cur:
    sql = 'SELECT * FROM t_test'
    print("Query sql: ", sql)
    cur.execute(sql)
    ans = cur.fetchall()
    print(ans)
conn.close()
```

In the Gitpod environment, you can directly use `run.sh` to run the demo code.

```bash
sh run.sh
```
