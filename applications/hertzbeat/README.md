<p align="center">
  <a href="https://hertzbeat.apache.org">
     <img alt="hertzbeat" src="/home/static/img/hertzbeat-brand.svg" width="260">
  </a>
</p>

<h4 align="center">
<a href="README.md">English Document</a> | <a href="README_CN.md">中文文档</a>
</h4>

> A real-time monitoring system with agentless, performance cluster, prometheus-compatible, custom monitoring and status page building capabilities.


## HertzBeat Start With OceanBase

1. Start the OceanBase Database

> here we use docker to start an ob standalone.

```shell
docker run -p 2881:2881 --name obstandalone -e MINI_MODE=1 -d oceanbase/oceanbase-ce
```

2. Create the database name `hertzbeat`

```shell
create database if not exists hertzbeat default charset utf8mb4;
```

3. Build the HertzBeat

```shell
mvn clean install
```

4. Start the HertzBeat Backend

Start SpringBoot Instance `org.apache.hertzbeat.manager.Manager`

5. Start the HertzBeat Webapp

```shell
cd  web-app

yarn install

yarn start
```

6. Access `http://localhost:4200` to start, account `admin/hertzbeat`

## Custom Database Configuration

### Custom Local or Prod Env Configuration

Modify the `application.yml` file in the `manager` module `manager/src/main/resources`

```yaml
spring:
  config:
    activate:
      on-profile: prod

  datasource:
    driver-class-name: com.mysql.cj.jdbc.Driver
    username: root
    password:
    url: jdbc:mysql://localhost:2881/hertzbeat?createDatabaseIfNotExist=true&useUnicode=true&characterEncoding=utf-8&useSSL=false
    hikari:
      max-lifetime: 120000
```

### Custom Test Env Configuration

This if for the test environment, `mvn clean test`

Modify the `application-test.yml` file in the `manager` module `manager/src/main/resources`

```yaml
spring:
  datasource:
    driver-class-name: com.mysql.cj.jdbc.Driver
    username: root
    password:
    url: jdbc:mysql://localhost:2881/test?createDatabaseIfNotExist=true&useUnicode=true&characterEncoding=utf-8&useSSL=false
    hikari:
      max-lifetime: 120000
```
