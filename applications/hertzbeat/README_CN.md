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
