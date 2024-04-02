# Java 连接 OceanBase 指南（使用 mysql-connector-java）

[English](README.md) | 简体中文

本文介绍如何通过 MySQL 官方 Java 驱动连接 OceanBase 数据库。

## 快速开始

在 Maven 中加入 MySQL 驱动

```xml
<dependency>
  <groupId>mysql</groupId>
  <artifactId>mysql-connector-java</artifactId>
  <version>5.1.47</version>
</dependency>
```

使用 MySQL 驱动时，需要提供相应的 JDBC Url，详细信息请参考 [MySQL文档](https://dev.mysql.com/doc/connector-j/8.0/en/connector-j-reference-jdbc-url-format.html)。

以 [MySqlConnectorTest.java](src/main/java/com/oceanbase/example/MySqlConnectorTest.java) 为例

```java
import java.sql.Connection;
import java.sql.DriverManager;
import java.sql.ResultSet;
import java.sql.ResultSetMetaData;
import java.sql.SQLException;
import java.sql.Statement;
import java.util.ArrayList;
import java.util.List;

public class MySqlConnectorTest {

  private static final String JDBC_URL = "jdbc:mysql://127.0.0.1:2881/test";
  private static final String USERNAME = "root@test";
  private static final String PASSWORD = "";

  public static void main(String[] args) {
    try (Connection connection = DriverManager.getConnection(JDBC_URL, USERNAME, PASSWORD);
         Statement statement = connection.createStatement()) {
      statement.execute("DROP TABLE IF EXISTS `t_test`");
      statement.execute("CREATE TABLE `t_test` (" +
        "    `id`   int(10) NOT NULL AUTO_INCREMENT," +
        "    `name` varchar(20) DEFAULT NULL," +
        "    PRIMARY KEY (`id`)" +
        ") ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE = utf8_bin");
      statement.execute("INSERT INTO `t_test` VALUES (default, 'Hello OceanBase')");

      ResultSet rs = statement.executeQuery("SELECT * FROM `t_test`");
      ResultSetMetaData metaData = rs.getMetaData();

      List<String> result = new ArrayList<>();
      while (rs.next()) {
        StringBuilder sb = new StringBuilder();
        for (int i = 0; i < metaData.getColumnCount(); i++) {
          if (i != 0) {
            sb.append(",");
          }
          Object value = rs.getObject(i + 1);
          sb.append(value == null ? "null" : value.toString());
        }
        result.add(sb.toString());
      }
      System.out.println(result);
    } catch (SQLException e) {
      throw new RuntimeException(e);
    }
  }
}
```

修改代码中的连接信息，之后你就可以直接使用 run.sh 运行示例代码。

```bash
sh run.sh
```
