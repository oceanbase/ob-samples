# Connect to OceanBase with Sequelize

English | [简体中文](README-CN.md)

This document describes how to connect to [OceanBase](https://www.oceanbase.com) with [Sequelize](https://sequelize.org).

## Preparation

Make sure `Node.js` and `npm` are installed.

## Usage

Clone the project and navigate to the appropriate directory:

```bash
git clone git@github.com:oceanbase/ob-samples.git
cd javascript/sequelize
```

Install dependencies:

```bash
npm install
```

Modify the connection string in the `index.js` file:

```javascript
const sequelize = new Sequelize("mysql://root:@127.0.0.1:2881/test", {
  dialect: "mysql",
  logging: false,
});
```

Execute `index.js`:

```bash
node index.js
```

The output should be as follows, indicating successful execution:

```bash
[
  {
    "id": 1,
    "email": "alice@oceanbase.com",
    "name": "Alice"
  }
]
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
