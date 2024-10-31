# Connect to OceanBase with Drizzle

English | [简体中文](README-CN.md)

This document describes how to connect to [OceanBase](https://www.oceanbase.com) with [Drizzle](https://orm.drizzle.team).

## Preparation

Make sure `Node.js` and `npm` are installed.

## Usage

Clone the project and navigate to the appropriate directory:

```bash
git clone git@github.com:oceanbase/ob-samples.git
cd javascript/drizzle
```

Install dependencies:

```bash
npm install
```

Modify the connection string in the `.env` file:

```bash
DATABASE_URL="mysql://root:@127.0.0.1:2881/test"
```

Synchronize the `users` model defined in `db/schema.ts` to the database:

```bash
npx drizzle-kit push
```

Execute `index.ts`:

```bash
npx ts-node index.ts
```

The output should be as follows, indicating successful execution:

```bash
[ { id: 1, email: 'alice@oceanbase.com', name: 'Alice' } ]
```

Check the corresponding `users` table and the data has been inserted:

```bash
mysql> select * from users;
+----+---------------------+-------+
| id | email               | name  |
+----+---------------------+-------+
|  1 | alice@oceanbase.com | Alice |
+----+---------------------+-------+
1 row in set (0.01 sec)
```
