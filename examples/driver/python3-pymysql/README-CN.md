# Python 连接 OceanBase 指南（PyMySQL）

[English](README.md) | 简体中文

本文介绍如何通过 PyMySQL 驱动连接 OceanBase 数据库。

## 快速开始

为了防止环境问题，推荐使用 anaconda 配置 python 3.x 环境。

在开始之前，需要先确保 PyMySQL 已安装。PyMySQL 是在 Python 3.x 版本中用于连接 MySQL 服务器的依赖库。有关 PyMySQL 的安装和使用等详细信息，您可参考 [官方文档](https://pypi.org/project/PyMySQL/) 和 [相关 API 参考文档](https://pymysql.readthedocs.io/en/latest/modules/index.html)。

以 [Test.py](Test.py) 为例。

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

在 Gitpod 环境中，可以直接使用 run.sh 运行示例代码。

```bash
sh run.sh
```
