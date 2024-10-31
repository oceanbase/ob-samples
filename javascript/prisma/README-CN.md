# Prisma 连接 OceanBase 指南

[English](README.md) | 简体中文

本文介绍如何通过 [Prisma](https://www.prisma.io) 连接 [OceanBase](https://www.oceanbase.com) 数据库。

## 准备工作

确保 Nodejs 和 npm 已经安装。

## 项目配置

拉取项目并进入相应目录:

```bash
git clone git@github.com:oceanbase/ob-samples.git
cd javascript/prisma
```

全局配置环境变量 (原因详见 https://open.oceanbase.com/blog/15137753618):

```bash
export PRISMA_ENGINES_MIRROR=https://oceanbase-prisma-builds.s3.ap-southeast-1.amazonaws.com
export BINARY_DOWNLOAD_VERSION=96fa66f2f130d66795d9f79dd431c678a9c7104e
```

安装依赖 (注意 `prisma` 和 `@prisma/client` 版本在 `^5.20.0` 及以上):

```bash
npm install
```

修改 `.env` 中的数据库连接串，格式如下。注意需要设置 `prefer_socket=false`，以避免和 OceanBase 建立连接时报错。

```bash
DATABASE_URL="mysql://root:@127.0.0.1:2881/test?prefer_socket=false&a=.psdb.cloud"
```

执行以下命令，将 `prisma/schema.prisma` 中定义的 `User`、`Post` 和 `Profile` 模型同步到数据库中:

```bash
npx prisma migrate dev --name init
```

```sql
mysql> show tables;
+--------------------+
| Tables_in_test     |
+--------------------+
| _prisma_migrations |
| posts              |
| profiles           |
| users              |
+--------------------+
4 rows in set (0.02 sec)
```

执行 `index.ts`:

```bash
npx ts-node index.ts
```

输出以下内容，说明执行成功:

```bash
[
  {
    id: 1,
    email: 'alice@oceanbase.com',
    name: 'Alice',
    posts: [
      {
        id: 1,
        createdAt: 2024-10-31T04:33:45.535Z,
        updatedAt: 2024-10-31T04:33:45.535Z,
        title: 'Hello World',
        content: null,
        published: false,
        authorId: 1
      }
    ],
    profile: { id: 1, bio: 'I like turtles', userId: 1 }
  }
]
```

查看对应的 `users`、`posts` 和 `profiles` 表，数据已正常插入:

```bash
mysql> select * from users;
+----+---------------------+-------+
| id | email               | name  |
+----+---------------------+-------+
|  1 | alice@oceanbase.com | Alice |
+----+---------------------+-------+
1 row in set (0.01 sec)

mysql> select * from posts;
+----+-------------------------+-------------------------+-------------+---------+-----------+----------+
| id | createdAt               | updatedAt               | title       | content | published | authorId |
+----+-------------------------+-------------------------+-------------+---------+-----------+----------+
|  1 | 2024-10-31 04:33:45.535 | 2024-10-31 04:33:45.535 | Hello World | NULL    |         0 |        1 |
+----+-------------------------+-------------------------+-------------+---------+-----------+----------+
1 row in set (0.01 sec)

mysql> select * from profiles;
+----+----------------+--------+
| id | bio            | userId |
+----+----------------+--------+
|  1 | I like turtles |      1 |
+----+----------------+--------+
1 row in set (0.01 sec)
```
