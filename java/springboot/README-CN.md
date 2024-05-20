# Spring Boot 连接 OceanBase 指南（使用 Spring Data JPA）

[English](README.md) | 简体中文

本文介绍如何通过 OceanBase 官方 SpringBoot 连接示例连接 OceanBase 数据库。
由于 OceanBase 支持 MySQL 模式与 Oracle 模式，因此可以使用 MySQL 驱动连接 OceanBase。
## 快速开始

### 在 pom.xml 中首先加入 Spring Boot 与 Spring Data JPA 相关的驱动, 以及 MySQL 驱动，pom.xml 参考[OceanBase SpringBoot 连接示例](https://www.oceanbase.com/docs/community-observer-cn-10000000000900914) 示例。

```xml
  <dependencies>
  <dependency>
    <groupId>org.springframework.boot</groupId>
    <artifactId>spring-boot-starter-web</artifactId>
    <exclusions>
      <exclusion>
        <artifactId>spring-boot-starter-json</artifactId>
        <groupId>org.springframework.boot</groupId>
      </exclusion>
    </exclusions>
  </dependency>
  <dependency>
    <groupId>mysql</groupId>
    <artifactId>mysql-connector-java</artifactId>
    <version>8.0.25</version>
  </dependency>
  <dependency>
    <groupId>org.springframework.boot</groupId>
    <artifactId>spring-boot-starter-test</artifactId>
    <scope>test</scope>
  </dependency>
  <dependency>
    <groupId>org.springframework.boot</groupId>
    <artifactId>spring-boot-starter-data-jpa</artifactId>
  </dependency>
  <dependency>
    <groupId>org.projectlombok</groupId>
    <artifactId>lombok</artifactId>
    <optional>true</optional>
  </dependency>
</dependencies>
```

### 在 application.yml 文件加入数据库连接信息等。

```yaml
spring:
  jpa:
    database: mysql # mysql, oracle
    show-sql: true
  datasource:
    driver-class-name: com.mysql.cj.jdbc.Driver
    url: jdbc:mysql://localhost:2881/test?characterEncoding=UTF-8
    username: root@test
    password:
  sql:
    init:
      mode: always
      data-locations: classpath:/init.sql
```
### 测试用例：

#### 1.定义简单实体：
```java
/*
 * CREATE TABLE TEST(id integer, name varchar2(50))
 *
 */
@Data
@Entity
@Table( name = "staff" )
public class StaffEntity implements Serializable {

  private static final long serialVersionUID = -6578740021873269176L;

  @Id
  @GeneratedValue(strategy=GenerationType.IDENTITY)
  // @GeneratedValue(strategy=GenerationType.AUTO) //oracle 没有自增策略，添加该注解可以自动生成一个序列，提供自增主键，若数据库已有相关序列，可以忽 //略该注解。
  @Column(name = "id")
  private Integer testId;

  @Column( name = "name" )
  private String testName;
}
```
#### 2.创建简单查询：
```java
import com.oceanbase.samples.springboot.entity.StaffEntity;
import org.springframework.data.repository.CrudRepository;
import org.springframework.stereotype.Repository;

import java.util.List;

@Repository
public interface StaffEntityRepository extends CrudRepository<StaffEntity, Integer> {
  List<StaffEntity> findByTestName(String lastName);

  List<StaffEntity> findByTestNameContaining(String testName);

  StaffEntity findById(int id);
}
```

#### 3.为Demo创建单元测试，并运行：
```java
import com.oceanbase.samples.springboot.repository.StaffEntityRepository;
import org.junit.jupiter.api.Test;
import org.springframework.boot.test.context.SpringBootTest;

import javax.annotation.Resource;

@SpringBootTest(classes = SpringBootJavaApplication.class)
public class SpringBootJavaApplicationTest {

  @Resource
  private StaffEntityRepository staffEntityRepository;


  @Test
  public void findByTestName() {
    assert staffEntityRepository.findByTestName("test") != null;
  }

  @Test
  public void findByTestNameContaining() {
    assert staffEntityRepository.findByTestNameContaining("test") != null;
  }

  @Test
  public void findById() {
    assert staffEntityRepository.findById(1) != null;
  }
}
```

修改代码中的连接信息，之后你就可以直接使用 run.sh 运行示例代码。

```bash
sh run.sh
```
