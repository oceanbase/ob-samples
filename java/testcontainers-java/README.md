# TestContainers Java

English | [简体中文](README-CN.md)

This article describes how to start and test OceanBase Docker container through `testcontainers-java`, you can see here for more details https://java.testcontainers.org/modules/databases/oceanbase/.

## Quick Start

Add TestContainers OceanBase module to POM.

```xml
<dependency>
  <groupId>org.testcontainers</groupId>
  <artifactId>oceanbase</artifactId>
  <version>1.19.7</version>
  <scope>test</scope>
</dependency>
```

Take [OceanBaseCEContainerTest.java](src/test/java/com/oceanbase/samples/OceanBaseCEContainerTest.java) code as an example.

The following code implements life cycle management of `OceanBaseCEContainer`. It will start the container instance before executing any test cases and stop the container after all test cases have been executed.

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

You can use the container instance in test cases like below:

```java
@Test
public void test() {
    try (Connection connection = CONTAINER.createConnection("?useSSL=false")) {
      ...
    }
}
```
