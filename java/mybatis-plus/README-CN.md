# MybatisPlus Java

[English](README.md) | 简体中文

本文介绍了如何通过 `mybatisplus-java` 启动和测试 OceanBase Docker
容器，更多详细信息可以参见 https://github.com/baomidou/mybatis-plus
以及 https://java.testcontainers.org/modules/databases/oceanbase 。

## 快速开始

将 OceanBase 驱动、TestContainers OceanBase、MybatisPlusStarter、SpringBootStarter Test 模块添加到 POM。

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

以 [MybatisPlusJavaApplicationTests.java](src/test/java/com/oceanbase/samples/mybatisplus/MybatisPlusJavaApplicationTests.java)
代码为例。

以下代码不仅实现了`OceanBaseCEContainer`的生命周期管理。 它将在执行任何测试用例之前启动容器实例，并在执行所有测试用例后停止容器，而且还在期间使用
ScriptUtils.executeSqlScript 执行数据库初始化 SQL。

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

您可以在测试用例中使用 MybatisPlus 操作 OceanBase 实例，如下所示：

```java

@Test
void testSelectList() {
  List<Person> persons = personMapper.selectList(null);
  assertFalse(persons.isEmpty());
}
```
