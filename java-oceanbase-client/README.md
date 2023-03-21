# Connect OceanBase with Java (oceanbase-client)

English | [简体中文](README-CN.md)

This article describes how to connect to the OceanBase database through `oceanbase-client`.

For more details about `oceanbase-client`, see [https://github.com/oceanbase/obconnector-j](https://github.com/oceanbase/obconnector-j).

## Quick Start

Add OceanBase JDBC driver to POM.

```xml
<dependency>
  <groupId>com.oceanbase</groupId>
  <artifactId>oceanbase-client</artifactId>
  <version>2.4.2</version>
</dependency>
```

When using the OceanBase driver, you need to provide the JDBC Url. Please refer to [OceanBase Connector/J Documentation](https://www.oceanbase.com/docs/oceanbase-connector-j-cn) for details.

Take [OceanBaseClientTest.java](src/main/java/com/oceanbase/example/OceanBaseClientTest.java) code as an example.

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

In the Gitpod environment, you can directly use `run.sh` to run the demo code.

```bash
sh run.sh
```
