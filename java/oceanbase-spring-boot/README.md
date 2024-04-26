# Guide to Connecting Spring Boot to OceanBase (Using Spring Data JPA)

English | [简体中文](README-CN.md)

This document introduces how to connect to the OceanBase database through Spring's official Spring Data JPA.
Since OceanBase supports MySQL mode and Oracle mode, you can use the MySQL driver to connect to OceanBase.

## Quick Start

### First, add the Spring Boot, Spring Data JPA, and MySQL driver dependencies to your Maven project.

```xml
<parent>
  <groupId>org.springframework.boot</groupId>
  <artifactId>spring-boot-starter-parent</artifactId>
  <version>2.7.16</version>
</parent>
<dependencies>
  <dependency>
    <groupId>org.springframework.boot</groupId>
    <artifactId>spring-boot-starter</artifactId>
  </dependency>
  <dependency>
    <groupId>mysql</groupId>
    <artifactId>mysql-connector-java</artifactId>
    <version>8.0.25</version>
  </dependency>
  <dependency>
    <groupId>org.springframework.boot</groupId>
    <artifactId>spring-boot-starter-data-jpa</artifactId>
  </dependency>
</dependencies>
```

### Add the database connection information to the Spring Boot configuration file.

```yaml
server:
  port: 8081
spring:
  jpa:
    database: mysql
    show-sql: true
  datasource:
    driver-class-name: com.mysql.cj.jdbc.Driver
    url: jdbc:mysql://host:port/test?characterEncoding=UTF-8
    username: *****
    password: *****
#spring.jpa.hibernate.ddl-auto=update
jackson:
  serialization:
    indent_output: true
```

### Next, refer to the [Spring Data JPA quick start example](https://spring.io/guides/gs/accessing-data-jpa) to write a demo:

#### 1.Define a simple entity:
```java
package com.oceanbase.samples.entity;

import javax.persistence.*;
import java.io.Serializable;
import java.util.Objects;

/*
 * CREATE TABLE TEST(id integer, name varchar2(50))
 *
 */
@Entity
@Table( name = "test" )
public class TestEntity implements Serializable {

  private static final long serialVersionUID = -6578740021873269176L;

  @Id
  // @GeneratedValue(strategy=GenerationType.AUTO) //oracle 没有自增策略，添加该注解可以自动生成一个序列，提供自增主键，若数据库已有相关序列，可以忽 //略该注解。
  @Column(name = "id")
  private Integer testId;

  @Column( name = "name" )
  private String TestName;



  public TestEntity(){

  }

  public TestEntity(String bauer) {
    this.TestName = bauer;
  }


  public Integer getTestId() {
    return testId;
  }

  public void setTestId(Integer testId) {
    this.testId = testId;
  }

  public String getTestName() {
    return TestName;
  }

  public void setTestName(String testName) {
    TestName = testName;
  }

  @Override
  public String toString() {
    return "TestEntity{" +
      "testId=" + testId +
      ", TestName='" + TestName + '\'' +
      '}';
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) return true;
    if (o == null || getClass() != o.getClass()) return false;
    TestEntity that = (TestEntity) o;
    return Objects.equals(testId, that.testId) && Objects.equals(TestName, that.TestName);
  }

  @Override
  public int hashCode() {
    return Objects.hash(testId, TestName);
  }
}

```
#### 2.Create a simple query:
```java
package com.oceanbase.samples.repository;

import com.oceanbase.samples.entity.TestEntity;
import org.springframework.data.repository.CrudRepository;
import org.springframework.stereotype.Repository;

import java.util.List;

@Repository
public interface TestEntityRepository extends CrudRepository<TestEntity, Integer> {
    List<TestEntity> findByTestName(String lastName);

    TestEntity findById(int id);
}

```

#### 3.Create an application class:
```java
package com.oceanbase.samples;

import com.oceanbase.samples.entity.TestEntity;
import com.oceanbase.samples.repository.TestEntityRepository;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.boot.CommandLineRunner;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.annotation.Bean;

@SpringBootApplication
public class HelloOceanBaseBootApplication {

    public static void main( String[] args ) {
        SpringApplication.run(HelloOceanBaseBootApplication.class, args);
    }

}

```

#### 4.Create a test in HelloOceanBaseBootApplication:
```java
package com.oceanbase.samples;

import com.oceanbase.samples.entity.TestEntity;
import com.oceanbase.samples.repository.TestEntityRepository;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.boot.CommandLineRunner;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.annotation.Bean;

@SpringBootApplication
public class HelloOceanBaseBootApplication {

    private static final Logger log = LoggerFactory.getLogger(HelloOceanBaseBootApplication.class);

    public static void main( String[] args )
    {
        SpringApplication.run(HelloOceanBaseBootApplication.class, args);
    }

    @Bean
    public CommandLineRunner demo(TestEntityRepository repository) {
        return (args) -> {
            // save a few test entities
            repository.save(new TestEntity("Hello OceanBase"));
            repository.save(new TestEntity("OceanBase is a distributed database"));

            // fetch all test entities
            log.info("Test entities found with findAll():");
            log.info("-------------------------------");
            for (TestEntity entity : repository.findAll()) {
                log.info(entity.toString());
            }
            log.info("");

            // fetch an individual test entity by ID
            TestEntity entity = repository.findById(1);
            log.info("Test entity found with findById(1):");
            log.info("--------------------------------");
            log.info(entity.toString());
            log.info("");

            // fetch test entities by test name
            log.info("Test entity found with findByTestName('OceanBase'):");
            log.info("--------------------------------------------");
            repository.findByTestName("OceanBase").forEach(oceanBase -> {
                log.info(oceanBase.toString());
            });
            log.info("");
        };
    }

}

```

Modify the connection info in code, and use `run.sh` to run the example code.

```bash
sh run.sh
```
