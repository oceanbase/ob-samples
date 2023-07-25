# Java 连接 OceanBase 指南（使用 oceanbase-client）

[English](README.md) | 简体中文

本文介绍如何通过 OceanBase Java 驱动 `oceanbase-client` 连接 OceanBase 数据库。

关于 `oceanbase-client` 的详细信息，请参考 [https://github.com/oceanbase/obconnector-j](https://github.com/oceanbase/obconnector-j)。

## 快速开始

在 Maven 中加入 OceanBase JDBC 驱动

```xml
<dependency>
  <groupId>com.oceanbase</groupId>
  <artifactId>oceanbase-client</artifactId>
  <version>2.4.2</version>
</dependency>
```

使用 OceanBase 驱动时，需要提供相应的 JDBC Url，详细信息请参考 [OceanBase Connector/J 文档](https://www.oceanbase.com/docs/oceanbase-connector-j-cn)。

以 [InsertAndSelectExample.java](src/main/java/com/oceanbase/example/InsertAndSelectExample.java) 为例

```java
public class InsertAndSelectExample {

  public static void main(String[] args) throws ClassNotFoundException, SQLException {
    // connect to your database
    String url = "jdbc:oceanbase://127.0.0.1:2881/test?characterEncoding=utf-8&useServerPrepStmts=true";
    String user = "root@test";
    String password = "";
    Class.forName("com.oceanbase.jdbc.Driver");
    Connection conn = DriverManager.getConnection(url, user, password);

    // create a table
    Statement stmt = conn.createStatement();
    try {
      stmt.execute("drop table person");
    } catch (Exception ignore) {
    }
    stmt.execute("create table person (name varchar(50), age int)");

    // insert records
    PreparedStatement ps = conn.prepareStatement("insert into person values(?, ?)");
    ps.setString(1, "Adam");
    ps.setInt(2, 28);
    ps.executeUpdate();
    ps.setString(1, "Eve");
    ps.setInt(2, 26);
    ps.executeUpdate();

    // fetch all records
    ps = conn.prepareStatement("select * from person", ResultSet.TYPE_FORWARD_ONLY, ResultSet.CONCUR_READ_ONLY);
    ResultSet rs = ps.executeQuery();
    while (rs.next()) {
      System.out.println(rs.getString(1) + " is " + rs.getInt(2) + " years old.");
    }

    // release all resources
    ps.close();
    stmt.close();
    conn.close();
  }

}
```

在 Gitpod 环境中，可以直接使用 run.sh 运行示例代码。

```bash
sh run.sh
```
