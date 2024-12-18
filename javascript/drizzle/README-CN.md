# 使用 Drizzle 连接 OceanBase

[English](README.md) | 简体中文

本文介绍如何通过 [Drizzle](https://orm.drizzle.team) 连接 [OceanBase](https://www.oceanbase.com) 数据库。

## 准备工作

确保 Node.js 和 npm 已经安装。

## 项目使用

拉取项目并进入相应目录:

```bash
git clone git@github.com:oceanbase/ob-samples.git
cd javascript/drizzle
```

安装依赖:

```bash
npm install
```

修改 `.env` 中的数据库连接串:

```bash
DATABASE_URL="mysql://root:@127.0.0.1:2881/test"
```

将 `db/schema.ts` 中定义的 `users` 模型同步到数据库中:

```bash
npx drizzle-kit push
```

执行 `index.ts` 中的示例代码:

```bash
npx ts-node index.ts
```

输出以下内容，说明执行成功:

```bash
[ { id: 1, email: 'alice@oceanbase.com', name: 'Alice' } ]
```

查看对应的 `users` 表，数据已正常插入:

```bash
mysql> select * from users;
+----+---------------------+-------+
| id | email               | name  |
+----+---------------------+-------+
|  1 | alice@oceanbase.com | Alice |
+----+---------------------+-------+
1 row in set (0.01 sec)
```
