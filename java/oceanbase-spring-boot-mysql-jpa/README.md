# Guide to Connecting Spring Boot to OceanBase (Using Spring Data JPA)

English | [简体中文](README-CN.md)

This document introduces how to connect to the OceanBase database through Spring's official Spring Data JPA.
Since OceanBase supports MySQL mode and Oracle mode, you can use the MySQL driver to connect to OceanBase.

## Quick Start

### First, add the Spring Boot, Spring Data JPA, and MySQL driver dependencies to the pom.xml file, referring to the [OceanBase SpringBoot connection example](https://www.oceanbase.com/docs/community-observer-cn-10000000000900914).

```xml
<project xmlns="http://maven.apache.org/POM/4.0.0" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd">
  <modelVersion>4.0.0</modelVersion>

  <groupId>com.oceanbase.samples</groupId>
  <artifactId>oceanbase-spring-boot-mysql-jpa</artifactId>
  <version>1.0-SNAPSHOT</version>
  <parent>
    <groupId>org.springframework.boot</groupId>
    <artifactId>spring-boot-starter-parent</artifactId>
    <version>2.0.1.RELEASE</version>
  </parent>
  <name>oceanbase-spring-boot</name>

  <properties>
    <project.build.sourceEncoding>UTF-8</project.build.sourceEncoding>
    <maven.compiler.source>1.8</maven.compiler.source>
    <maven.compiler.target>1.8</maven.compiler.target>
  </properties>

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
      <artifactId>spring-boot-starter-data-jpa</artifactId>
    </dependency>
    <dependency>
      <groupId>junit</groupId>
      <artifactId>junit</artifactId>
      <version>3.8.1</version>
      <scope>test</scope>
    </dependency>
    <dependency>
      <groupId>com.fasterxml.jackson.core</groupId>
      <artifactId>jackson-databind</artifactId>
      <version>2.8.5</version>
    </dependency>
  </dependencies>
</project>
```

### Add the database connection information and other configurations in the application.yml file.

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

### Test Cases:

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
  // @GeneratedValue(strategy=GenerationType.AUTO) //Oracle does not have an auto-increment strategy. Adding this annotation can automatically generate a sequence to provide an auto-increment primary key. If the database already has a relevant sequence, you can omit this annotation.
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

  List<TestEntity> findByTestNameContaining(String testName);

  TestEntity findById(int id);
}
```

#### 3.Create test cases in the controller for CRUD operations:
```java
package com.oceanbase.samples.controller;

import com.oceanbase.samples.entity.TestEntity;
import com.oceanbase.samples.repository.TestEntityRepository;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

import java.util.List;
import java.util.Objects;
import java.util.stream.Collectors;
import java.util.stream.StreamSupport;

@RestController
@RequestMapping("/")
public class TestController {
  private static final Logger log = LoggerFactory.getLogger(TestController.class);

  @Autowired
  private TestEntityRepository testEntityRepository;

  @PostMapping("/save")
  public Boolean save(@RequestBody TestEntity testSaveEntity) {
    TestEntity resultEntity = testEntityRepository.save(testSaveEntity);
    log.info("save result: {}", resultEntity);
    return resultEntity != null;
  }

  @GetMapping("/{id}")
  public TestEntity findById(@PathVariable("id") int id) {
    log.info("find by id: {}", id);
    TestEntity resultEntity = testEntityRepository.findById(id);
    log.info("find result: {}", resultEntity);
    return resultEntity;
  }

  @GetMapping("/findAll")
  public List<TestEntity> findAll() {
    log.info("find all");
    return StreamSupport.stream(testEntityRepository.findAll().spliterator(), false)
      .filter(Objects::nonNull)
      .peek(entity -> log.info(entity.toString()))
      .collect(Collectors.toList());
  }

  @GetMapping("/findByName")
  public List<TestEntity> findByName(@RequestParam("name") String name) {
    log.info("find by name: {}", name);
    return testEntityRepository.findByTestNameContaining(name);
  }

  @DeleteMapping("/{id}")
  public void deleteById(@PathVariable("id") int id) {
    log.info("delete by id: {}", id);
    testEntityRepository.deleteById(id);
  }

  @PutMapping("/update")
  public void update(@RequestBody TestEntity testUpdateEntity) {
    log.info("update: {}", testUpdateEntity);
    testEntityRepository.save(testUpdateEntity);
  }

}
```

#### 4.Create the application class and run it:
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

Modify the connection info in code, and use `run.sh` to run the example code.

```bash
sh run.sh
```
