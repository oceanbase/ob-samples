# Connect OceanBase with Java (mysql-connector-java)

English | [简体中文](README-CN.md)

This article describes how to connect to the OceanBase database through `mysql-connector-java`.

## Quick Start

Add MySQL JDBC driver to POM.

```xml
<dependency>
  <groupId>mysql</groupId>
  <artifactId>mysql-connector-java</artifactId>
  <version>5.1.47</version>
</dependency>
```

When using the MySQL driver, you need to provide the JDBC Url. Please refer to [MySQL Documentation](https://dev.mysql.com/doc/connector-j/8.0/en/connector-j-reference-jdbc-url-format.html) for details.

Take [MySqlConnectorTest.java](src/main/java/com/oceanbase/example/MySqlConnectorTest.java) code as an example.

```java
public class MySqlConnectorTest {
    public static void main(String[] args) {
        String workspace = "/workspace/ob-example";
        String sqlFile = "tests/sql/test.sql";
        String tableName = "t_test";

        Properties properties = new Properties();
        properties.put("user", "root@test");
        properties.put("password", "");
        String jdbcUrl = "jdbc:mysql://127.0.0.1:2881/test";

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