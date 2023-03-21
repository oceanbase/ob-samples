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

以 [OceanBaseClientTest.java](src/main/java/com/oceanbase/example/OceanBaseClientTest.java) 为例

```java
public class OceanBaseClientTest {
    public static void main(String[] args) {
        String workspace = "/workspace/ob-example";
        String sqlFile = "tests/sql/test.sql";
        String tableName = "t_test";

        Properties properties = new Properties();
        properties.put("user", "root@test");
        properties.put("password", "");
        String jdbcUrl = "jdbc:oceanbase://127.0.0.1:2881/test";

        Connection connection;
        Statement statement;
        try {
            connection = DriverManager.getConnection(jdbcUrl, properties);
            statement = connection.createStatement();
            System.out.println("Success to connect to OceanBase");
        } catch (SQLException e) {
            System.out.println("Failed to connect to OceanBase, exception: " + e.getMessage());
            return;
        }

        String selectSql = "SELECT * FROM " + tableName;
        System.out.println("Query sql: " + selectSql);
        try {
            ResultSet rs = statement.executeQuery(selectSql);
            ResultSetMetaData metaData = rs.getMetaData();
            System.out.println("Get rows:");
            int count = 0;
            while (rs.next()) {
                System.out.printf("## row %d: { ", count++);
                for (int i = 0; i < metaData.getColumnCount(); i++) {
                    System.out.print(metaData.getColumnName(i + 1) + ": " + rs.getObject(i + 1) + "; ");
                }
                System.out.println("}");
            }
        } catch (SQLException e) {
            System.out.println("Failed to query table " + tableName + ", exception: " + e.getMessage());
            return;
        }

        try {
            if (statement != null) {
                statement.close();
            }
            if (connection != null) {
                connection.close();
            }
        } catch (SQLException e) {
            System.out.println("Failed to close statement and connection, exception: " + e.getMessage());
        }
    }
}
```

在 Gitpod 环境中，可以直接使用 run.sh 运行示例代码。

```bash
sh run.sh
```
