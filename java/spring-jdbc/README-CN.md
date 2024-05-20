# SpringJDBC 连接 OceanBase 指南
[English](README.md) | 简体中文

本文介绍如何通过 SpringJDBC 连接 OceanBase 数据库。
由于 OceanBase 支持 MySQL 模式与 Oracle 模式，因此可以使用 MySQL 驱动连接 OceanBase。
### 快速开始
1. 在 pom.xml 中首先加入 MySQL 驱动，pom.xml 参考[OceanBase Spring JDBC 连接示例](https://www.oceanbase.com/docs/community-observer-cn-10000000000900916) 示例。

```xml
    <dependency>
      <groupId>org.springframework</groupId>
      <artifactId>spring-jdbc</artifactId>
      <version>5.0.9.RELEASE</version>
    </dependency>
    <dependency>
      <groupId>mysql</groupId>
      <artifactId>mysql-connector-java</artifactId>
      <version>8.0.25</version>
    </dependency>
    <dependency>
      <groupId>com.alibaba</groupId>
      <artifactId>druid</artifactId>
      <version>1.2.18</version>
    </dependency>
    <dependency>
      <groupId>junit</groupId>
      <artifactId>junit</artifactId>
      <version>4.13.2</version>
      <scope>test</scope>
    </dependency>
```
2.新建测试类，使用 Druid 连接池实例化 JdbcTemplate。

```java
public class OceanBaseSpringJdbcApplicationTest {
  private static JdbcTemplate jdbcTemplate;
  private String sql;

  static {
    Map<String, String> map = new HashMap<String, String>();
    map.put("url", "jdbc:mysql://localhost:2881/test?characterEncoding=utf-8&useSSL=false&serverTimezone=UTC");
    map.put("driverClassName", "com.mysql.cj.jdbc.Driver");
    map.put("username", "root@test");
    map.put("password", "");
    try {
      Class.forName(map.get("driverClassName"));
      jdbcTemplate = new JdbcTemplate(DruidDataSourceFactory.createDataSource(map));
      //防止异常语句,没有这两句，会出错(Prevent abnormal statements, without which errors will occur)
      jdbcTemplate.execute("set transaction_isolation = 'READ-COMMITTED';"); // MySQL 8.0 之后，系统变量 tx_isolation 被更改为 transaction_isolation (After MySQL 8.0, the system variable tx_isolation was changed to transaction_isolation)
      // jdbcTemplate.execute("set tx_isolation = 'READ-COMMITTED';"); // MySQL 8.0 之前的版本使用 tx_isolation (tx_isolation is used in versions before MySQL 8.0)
    } catch (Exception e) {
      e.printStackTrace();
    }
  }
}
```
3.编写测试方法，执行 SQL 语句。

```java
package com.oceanbase.samples;

import com.alibaba.druid.pool.DruidDataSourceFactory;
import org.junit.Before;
import org.junit.Test;
import org.springframework.core.io.ClassPathResource;
import org.springframework.core.io.Resource;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.jdbc.core.RowMapper;
import org.springframework.jdbc.datasource.init.ScriptUtils;

import java.sql.ResultSet;
import java.sql.SQLException;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * OceanBaseSpringJdbcApplication 简单单元测试
 * Unit test for simple OceanBaseSpringJdbcApplication.
 */
public class OceanBaseSpringJdbcApplicationTest
{
  private static JdbcTemplate jdbcTemplate;
  private String sql;
  static {
    Map<String, String> map = new HashMap<String, String>();
    map.put("url", "jdbc:mysql://localhost:2881/test?characterEncoding=utf-8&useSSL=false&serverTimezone=UTC");
    map.put("driverClassName", "com.mysql.cj.jdbc.Driver");
    map.put("username", "root@test");
    map.put("password", "");
    try {
      Class.forName(map.get("driverClassName"));
      jdbcTemplate = new JdbcTemplate(DruidDataSourceFactory.createDataSource(map));
      //防止异常语句,没有这两句，会出错(Prevent abnormal statements, without which errors will occur)
      jdbcTemplate.execute("set transaction_isolation = 'READ-COMMITTED';"); // MySQL 8.0 之后，系统变量 tx_isolation 被更改为 transaction_isolation (After MySQL 8.0, the system variable tx_isolation was changed to transaction_isolation)
      // jdbcTemplate.execute("set tx_isolation = 'READ-COMMITTED';"); // MySQL 8.0 之前的版本使用 tx_isolation (tx_isolation is used in versions before MySQL 8.0)
    } catch (Exception e) {
      e.printStackTrace();
    }
  }

  @Before
  public void initializeDatabase() {
    Resource initSchema = new ClassPathResource("init.sql");
    try {
      // Execute init.sql script to create and populate the database
      ScriptUtils.executeSqlScript(jdbcTemplate.getDataSource().getConnection(), initSchema);
    } catch (Exception e) {
      throw new RuntimeException("Failed to execute init.sql script", e);
    }
  }


  // insert
  @Test
  public void testInsert() {
    String sql = "INSERT INTO staff (name) VALUES (?)";
    int rowsAffected = jdbcTemplate.update(sql, "New Staff");
    System.out.println("rowsAffected: " + rowsAffected);
  }

  // select
  @Test
  public void testSelect() {
    String sql = "SELECT * FROM staff WHERE name = ?";
    RowMapper<String> rowMapper = new RowMapper<String>() {
      @Override
      public String mapRow(ResultSet rs, int rowNum) throws SQLException {
        return rs.getString("name");
      }
    };
    List<String> names = jdbcTemplate.query(sql, rowMapper, "New Staff");
    System.out.println("names: " + names);
  }
}
```

修改代码中的连接信息，之后你就可以直接使用 run.sh 运行示例代码。

  ```bash
  sh run.sh
  ```
