# Spring Boot 连接 OceanBase 指南（使用 Spring Data JPA）

[English](README.md) | 简体中文

本文介绍如何通过 MyBatis 连接 OceanBase 数据库。
由于 OceanBase 支持 MySQL 模式与 Oracle 模式，因此可以使用 MySQL 驱动连接 OceanBase。
## 快速开始

1.在 pom.xml 中首先加入 MySQL 以及 MyBatis 驱动，pom.xml 参考[OceanBase MyBatis 连接示例](https://www.oceanbase.com/docs/community-observer-cn-10000000000900919) 示例。

```xml
  <dependencies>
  <dependency>
    <groupId>mysql</groupId>
    <artifactId>mysql-connector-java</artifactId>
    <version>8.0.25</version>
  </dependency>
  <dependency>
    <groupId>org.mybatis</groupId>
    <artifactId>mybatis</artifactId>
    <version>3.5.4</version>
  </dependency>
</dependencies>
```

 2.在 src/main/resources 文件夹新建 mybatis-config.xml ，并修改数据库连接信息,指定 mapper.xml 路径。

 mybatis-config.xml
```xml
<?xml version="1.0" encoding="UTF8"?>
<!DOCTYPE configuration
  PUBLIC "-//mybatis.org//DTD Config 3.0//EN"
  "http://mybatis.org/dtd/mybatis-3-config.dtd">
<configuration>
  <environments default="development">
    <environment id="development">
      <transactionManager type="JDBC"/>
      <dataSource type="POOLED">
        <property name="driver" value="com.mysql.cj.jdbc.Driver"/>
        <property name="url" value="jdbc:mysql://host:port/dbName?useUnicode=true&amp;characterEncoding=utf-8&amp;useServerPrepStmts=false&amp;useCursorFetch=true"/>
        <property name="username" value="*****"/>
        <property name="password" value="*****"/>
      </dataSource>
    </environment>
  </environments>

  <!--注册mapper（mapper.xml所在地址）-->
  <mappers>
    <mapper resource="UserMapper.xml"></mapper>
  </mappers>
</configuration>
```
3.新建 实体类，以及对应的 Mapper.xml 文件。 本次示例中，我们创建一个 User 实体类。
```java
package com.oceanbase.samples.entity;

import java.util.Objects;

public class User {
    private Integer id;
    private String name;

    public User() {
    }

    public User(Integer id, String name) {
        this.id = id;
        this.name = name;
    }

    public void setId(Integer id) {
        this.id = id;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        User user = (User) o;
        return Objects.equals(id, user.id) && Objects.equals(name, user.name);
    }

    @Override
    public int hashCode() {
        return Objects.hash(id, name);
    }

    @Override
    public String toString() {
        return "User{" +
            "id=" + id +
            ", name='" + name + '\'' +
            '}';
    }
}
```
4.新建 UserMapper 接口，用于与 UserMapper关联相应的增删改查操作
```java
package com.oceanbase.samples.mapper;

public interface UserMapper {
}
```

5.新建 UserMapper.xml ，并绑定 UserMapper 接口， 添加相应的增删改查操作。
```xml
<?xml version="1.0" encoding="UTF-8" ?>
<!DOCTYPE mapper
  PUBLIC "-//mybatis.org//DTD Mapper 3.0//EN"
  "http://mybatis.org/dtd/mybatis-3-mapper.dtd">

<mapper namespace="com.oceanbase.samples.mapper.UserMapper">
  <select id="selectUser" resultType="com.oceanbase.samples.entity.User" fetchSize="40000">
    select * from user;
  </select>

  <select id="selectWithPagination" parameterType="map" resultType="com.oceanbase.samples.entity.User">
    SELECT * FROM User LIMIT #{offset}, #{pageSize}
  </select>

  <delete id="delete" >
    delete from user where id = #{id};
  </delete>

  <insert id="insert" parameterType="com.oceanbase.samples.entity.User">
    insert into user (name) values (#{name});
  </insert>

  <update id="update" parameterType="com.oceanbase.samples.entity.User">
    update user set name = #{name} where id = #{id};
  </update>
</mapper>
```

6.测试用例,用于测试 UserMapper.xml 中的增删改查操作。
```java
package com.oceanbase.samples;


import com.oceanbase.samples.entity.User;
import org.apache.ibatis.io.Resources;
import org.apache.ibatis.session.SqlSession;
import org.apache.ibatis.session.SqlSessionFactory;
import org.apache.ibatis.session.SqlSessionFactoryBuilder;

import java.io.IOException;
import java.io.InputStream;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.Objects;


public class OceanBaseMyBatisTest
{

    public static void main( String[] args )
    {
        SqlSession sqlSession = null;
        try {
            // Get SqlSessionFactoryBuilder
            SqlSessionFactoryBuilder sqlSessionFactoryBuilder = new SqlSessionFactoryBuilder();
            // Load mybatis-config.xml as InputStream
            InputStream inputStream = Resources.getResourceAsStream("mybatis-config.xml");
            // Get SqlSessionFactory
            SqlSessionFactory sqlSessionFactory = sqlSessionFactoryBuilder.build(inputStream);
            // Get SqlSession
            sqlSession = sqlSessionFactory.openSession();
            // Execute SQL
            // insertTest(sqlSession);
            // updateTest(sqlSession);
            // selectTest(sqlSession);
            // selectWithPagination(sqlSession, 0, 10);
            deleteTest(sqlSession);
            // Commit Transaction
            sqlSession.commit();
        } catch (IOException e) {
            // Rollback Transaction
            if (sqlSession != null) {
                sqlSession.rollback();
            }
            e.printStackTrace();
        } finally {
            // Close SqlSession
            if (sqlSession != null) {
                sqlSession.close();
            }
        }
    }

    public static void insertTest(SqlSession sqlSession) {
        // Insert data
        User user = new User();
        user.setName("Tom");
        int count = sqlSession.insert("com.oceanbase.samples.mapper.UserMapper.insert", user);
        System.out.println("Insert count: " + count);

    }

    public static void updateTest(SqlSession sqlSession) {
        // Update data
        User user = new User();
        user.setId(1);
        user.setName("Jerry");
        int count = sqlSession.update("com.oceanbase.samples.mapper.UserMapper.update", user);
        System.out.println("Update count: " + count);
    }

    public static void selectTest(SqlSession sqlSession) {
        // Select data
        List<User> user = sqlSession.selectList("com.oceanbase.samples.mapper.UserMapper.selectUser", 1);
        user.stream().filter(Objects::nonNull).forEach(System.out::println);
    }


    public static void selectWithPagination(SqlSession sqlSession, int offset, int pageSize) {
        Map<String, Object> params = new HashMap<>();
        params.put("offset", offset);
        params.put("pageSize", pageSize);
        List<User> users = sqlSession.selectList("com.oceanbase.samples.mapper.UserMapper.selectWithPagination", params);
        users.stream().filter(Objects::nonNull).forEach(System.out::println);
    }

    public static void deleteTest(SqlSession sqlSession) {
        // Delete data
        int count = sqlSession.delete("com.oceanbase.samples.mapper.UserMapper.delete", 3);
        System.out.println("Delete count: " + count);
    }
}
```


修改代码中的连接信息，之后你就可以直接使用 run.sh 运行示例代码。

```bash
sh run.sh
```
