# Guide to Connecting Spring Boot to OceanBase (Using Spring Data JPA)

English | [简体中文](README-CN.md)

This document introduces how to connect to the OceanBase database through Spring's official Spring Data JPA.
Since OceanBase supports MySQL mode and Oracle mode, you can use the MySQL driver to connect to OceanBase.

## Quick Start

### First, add the Spring Boot, Spring Data JPA, and MySQL driver dependencies to the pom.xml file, referring to the [OceanBase SpringBoot connection example](https://www.oceanbase.com/docs/community-observer-cn-10000000000900914).

```xml
<?xml version="1.0" encoding="UTF-8"?>

<project xmlns="http://maven.apache.org/POM/4.0.0" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd">
  <modelVersion>4.0.0</modelVersion>
  <parent>
    <groupId>org.springframework.boot</groupId>
    <artifactId>spring-boot-starter-parent</artifactId>
    <version>2.5.12</version>
    <relativePath/>
  </parent>
  <groupId>com.oceanbase.samples</groupId>
  <artifactId>springboot</artifactId>
  <version>1.0-SNAPSHOT</version>

  <name>springboot</name>

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

  <build>
    <plugins>
      <plugin>
        <groupId>org.springframework.boot</groupId>
        <artifactId>spring-boot-maven-plugin</artifactId>
        <configuration>
          <excludes>
            <exclude>
              <groupId>org.projectlombok</groupId>
              <artifactId>lombok</artifactId>
            </exclude>
          </excludes>
        </configuration>
      </plugin>
    </plugins>
  </build>
</project>

```

### Add the database connection information and other configurations in the application.yml file.

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

### Test Cases:

#### 1.Define a simple entity:
```java
import lombok.Data;

import javax.persistence.*;
import java.io.Serializable;

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
  // @GeneratedValue(strategy=GenerationType.AUTO) //oracle There is no auto-increment policy, adding the annotation can automatically generate a sequence, provide the auto-increment primary key, and if the database already has a related sequence, you can ignore // omit the annotation.
  @Column(name = "id")
  private Integer testId;

  @Column( name = "name" )
  private String testName;
}
```
#### 2.Create a simple query:
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

#### 3.Create the application class and run it:
```java
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@SpringBootApplication
public class SpringBootJavaApplication {
  public static void main(String[] args) {
    SpringApplication.run(SpringBootJavaApplication.class, args);
  }

}
```

#### 4.Create a test class to test the Demo:
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


```bash
sh run.sh
```
