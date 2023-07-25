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

Take [InsertAndSelectExample.java](src/main/java/com/oceanbase/example/InsertAndSelectExample.java) code as an example.

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

In the Gitpod environment, you can directly use `run.sh` to run the demo code.

```bash
sh run.sh
```
