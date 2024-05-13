# MybatisPlus Java

English | [简体中文](README-CN.md)

This document introduces how to start and test the OceanBase Docker container using `mybatisplus-java`. For more
details, please see https://github.com/baomidou/mybatis-plus
and https://java.testcontainers.org/modules/databases/oceanbase.

## Quick Start

Add the OceanBase driver, TestContainers OceanBase, MybatisPlusStarter, and SpringBootStarter Test modules to your POM.

```xml

<dependencies>
  <dependency>
    <groupId>com.oceanbase</groupId>
    <artifactId>oceanbase-client</artifactId>
    <version>2.4.9</version>
  </dependency>
  <dependency>
    <groupId>org.testcontainers</groupId>
    <artifactId>oceanbase</artifactId>
    <version>1.19.7</version>
    <scope>test</scope>
  </dependency>
  <dependency>
    <groupId>com.baomidou</groupId>
    <artifactId>mybatis-plus-boot-starter</artifactId>
    <version>3.5.6</version>
    <scope>test</scope>
  </dependency>
  <dependency>
    <groupId>org.springframework.boot</groupId>
    <artifactId>spring-boot-starter</artifactId>
  </dependency>

  <dependency>
    <groupId>org.projectlombok</groupId>
    <artifactId>lombok</artifactId>
    <optional>true</optional>
  </dependency>
  <dependency>
    <groupId>org.springframework.boot</groupId>
    <artifactId>spring-boot-starter-test</artifactId>
    <scope>test</scope>
  </dependency>
  <dependency>
    <groupId>org.testcontainers</groupId>
    <artifactId>testcontainers</artifactId>
    <version>1.19.7</version>
    <scope>test</scope>
  </dependency>
  <dependency>
    <groupId>org.testcontainers</groupId>
    <artifactId>junit-jupiter</artifactId>
    <version>1.19.7</version>
    <scope>test</scope>
  </dependency>
</dependencies>
```

Take the code
in [MybatisPlusJavaApplicationTests.java](src/test/java/com/oceanbase/samples/mybatisplus/MybatisPlusJavaApplicationTests.java)
as an example.

The following code not only implements lifecycle management for the `OceanBaseCEContainer`. It will start the container
instance before executing any test cases and stop it after all test cases are completed, and also uses
ScriptUtils.executeSqlScript to perform database initialization SQL during the period.

```java

@SpringBootTest
@Testcontainers
class MybatisPlusJavaApplicationTests {

  @Container
  public static OceanBaseCEContainer oceanBaseContainer = new OceanBaseCEContainer(DockerImageName.parse("oceanbase/oceanbase-ce:latest"))
    .withEnv("MODE", "slim")
    .withEnv("FASTBOOT", "true");
  @Autowired
  private PersonMapper personMapper;

  @DynamicPropertySource
  static void oceanBaseProperties(DynamicPropertyRegistry registry) {
    registry.add("spring.datasource.url", oceanBaseContainer::getJdbcUrl);
    registry.add("spring.datasource.username", oceanBaseContainer::getUsername);
    registry.add("spring.datasource.password", oceanBaseContainer::getPassword);
    registry.add("spring.datasource.driver-class-name", oceanBaseContainer::getDriverClassName);
  }

  @BeforeAll
  static void setup(@Autowired DataSource dataSource) throws Exception {
    JdbcTemplate jdbcTemplate = new JdbcTemplate(dataSource);
    assertNotNull(jdbcTemplate.getDataSource());
    ScriptUtils.executeSqlScript(jdbcTemplate.getDataSource().getConnection(), new ClassPathResource("init.sql"));
  }
}
```

You can use MybatisPlus to operate the OceanBase instance in your test cases as follows:

```java

@Test
void testSelectList() {
  List<Person> persons = personMapper.selectList(null);
  assertFalse(persons.isEmpty());
}
```
