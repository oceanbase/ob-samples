<p align="center">
  <a href="https://hertzbeat.apache.org">
     <img alt="hertzbeat" src="/home/static/img/hertzbeat-brand.svg" width="260">
  </a>
</p>

<h4 align="center">
<a href="README_CN.md">中文文档</a> | <a href="README.md">English Document</a>
</h4>

> 实时监控系统，无需 Agent，性能集群，兼容 Prometheus，自定义监控和状态页构建能力。

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

## JWT 签名密钥

使用本源码构建时，将 `sureness.jwt.secret` 留空，服务会在首次启动时生成随机签名密钥，写入元数据库，并在重启后复用。请保留元数据库；也可以通过 `sureness.jwt.secret` 或环境变量 `SURENESS_JWT_SECRET` 提供经过 Base64 编码的随机密钥。

显式配置的自定义密钥优先生效，不会改动已存储密钥。配置为空或为旧默认值时，服务会复用非默认的已存储密钥；若存储为空或包含两个公开旧默认值之一，则生成替代密钥。密钥替换后，使用旧密钥签发的令牌会失效，受影响的用户需要重新登录。Helm 使用外部镜像，必须显式提供密钥，具体步骤见 [Chart 说明](script/helm/hertzbeat/README.md#configure-the-jwt-signing-key)。

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
