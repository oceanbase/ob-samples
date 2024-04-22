# TestContainers Java

[English](README.md) | 简体中文

本文介绍了如何通过 testcontainers-java 启动和测试 OceanBase Docker 容器，更多详细信息可以参见 https://java.testcontainers.org/modules/databases/oceanbase 。

## 快速开始

将 TestContainers OceanBase 模块添加到 POM。

```xml
<dependency>
  <groupId>org.testcontainers</groupId>
  <artifactId>oceanbase</artifactId>
  <version>1.19.7</version>
  <scope>test</scope>
</dependency>
```

以 [OceanBaseCEContainerTest.java](src/test/java/com/oceanbase/samples/OceanBaseCEContainerTest.java) 代码为例。

以下代码实现了`OceanBaseCEContainer`的生命周期管理。 它将在执行任何测试用例之前启动容器实例，并在执行所有测试用例后停止容器。

```java
private static final Logger LOG = LoggerFactory.getLogger(OceanBaseCEContainerTest.class);

public static final OceanBaseCEContainer CONTAINER =
    new OceanBaseCEContainer("oceanbase/oceanbase-ce:latest")
        .withEnv("MODE", "slim")
        .withEnv("FASTBOOT", "true")
        .withLogConsumer(new Slf4jLogConsumer(LOG));

@BeforeClass
public static void startContainers() {
    Startables.deepStart(Stream.of(CONTAINER)).join();
    LOG.info(
        "OceanBase docker container started, image: {}, host: {}, sql port: {}, rpc port:{}.",
        CONTAINER.getDockerImageName(),
        CONTAINER.getHost(),
        CONTAINER.getMappedPort(2881),
        CONTAINER.getMappedPort(2882));
}

@AfterClass
public static void closeContainers() {
    CONTAINER.close();
    LOG.info("OceanBase docker container stopped.");
}
```

您可以在测试用例中使用容器实例，如下所示：

```java
@Test
public void test() {
    try (Connection connection = CONTAINER.createConnection("?useSSL=false")) {
      ...
    }
}
```
