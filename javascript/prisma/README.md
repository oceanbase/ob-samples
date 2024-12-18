# Connect to OceanBase with Prisma

English | [简体中文](README-CN.md)

This document describes how to connect to [OceanBase](https://www.oceanbase.com) with [Prisma](https://www.prisma.io).

## Preparation

Make sure `Node.js` and `npm` are installed.

## Configuration

Clone the project and navigate to the appropriate directory:

```bash
git clone git@github.com:oceanbase/ob-samples.git
cd javascript/prisma
```

Set the global environment variables (for reasons see https://open.oceanbase.com/blog/15137753618):

```bash
export PRISMA_ENGINES_MIRROR=https://oceanbase-prisma-builds.s3.ap-southeast-1.amazonaws.com
export BINARY_DOWNLOAD_VERSION=96fa66f2f130d66795d9f79dd431c678a9c7104e
```

Install dependencies (note that the versions of `prisma` and `@prisma/client` should be `^5.20.0` or higher):

```bash
npm install
```

Modify the connection string in the `.env` file, formatted as follows. Note that `prefer_socket=false` must be set to avoid errors when connecting to OceanBase.

```bash
DATABASE_URL="mysql://root:@127.0.0.1:2881/test?prefer_socket=false&a=.psdb.cloud"
```

Execute the following command to synchronize the `User`, `Post` and `Profile` data models defined in `prisma/schema.prisma` to the database:

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

Execute `index.ts`:

```bash
npx ts-node index.ts
```

The output should be as follows, indicating successful execution:

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

Check the corresponding `users`, `posts` and `profiles` tables and the data has been inserted:

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
