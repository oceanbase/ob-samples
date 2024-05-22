# OceanBase Examples

English | [简体中文](README-CN.md)

## Introduction

This repository provides the sample projects for OceanBase.

For basic sample projects, we group them by the programming language, and the sample directories are named after the component name.

For sample projects that works as standalone applications, we put them under `applications` directory, and the module name typically contains the component name and purpose.

You can copy the sample code to local environment and use it directly after changing the connection info, please see the README file in the sample directory for more details.

## Contribution

We welcome contributions from anyone, thanks to all [contributors](https://github.com/oceanbase/ob-samples/graphs/contributors)!

We have created some issues on GitHub for some samples which are simple and good for first time contributors: https://github.com/oceanbase/ob-samples/issues/16.

The following are samples that have been added or planned to be added in the near future:

- application
  - `hertzbeat` `sveltekit`
  - TODO: `seatunnel`
- c
  - TODO: `obconnector-c`
- c_plusplus
  - TODO: `mysql-connector-cpp`
- c_sharp
  - TODO: `mysql-connector-net` `MySqlConnector`
- golang
  - `go-sql-driver`
  - TODO: `gorm` `xorm` `obkv-table-client-go`
- java
  - `mybatis` `mybatis-plus` `mysql-connector-java` `oceanbase-client` `spring-jdbc` `springboot` `testcontainers-java`
  - TODO: `hibernate` `spring-data-jpa` `hikaricp` `flink-cdc` `oblogclient` `obkv-table-client-java`
- php
  - TODO: `mysqli` `pdo`
- python
  - `mysql-connector-python` `pymysql`
  - TODO: `mysqlclient` `sqlalchemy`
- rust
  - TODO: `sqlx` `rust-mysql-simple` `obkv-table-client-rust`

### Add a sample

In this repository, every sample project will be as an independent module. The directory name of the module should be same with the tool used in the sample. For example, the sample of `mysql-connector-java` is named `mysql-connector-java`.

There are many category directories in the root directory of this repository. For samples that can be run directly through simple commands, we recommend add them to the directory which corresponding to the programming language environment required by the sample. For example, `mysql-connector-java` needs to be added to the `java` directory. For samples that require more complex configuration to run, we recommend placing the sample project in the `applications` directory.

#### Simple samples

For a simple sample under the programming language category, the directory should contain the following content

- code files
- `run.sh` script to run code
- `README.md` documentation for component usage.

If it's hard to use English for you, you can use your native language in the documentation, and we can improve it later.

To ensure that the sample works, please add your module to the GitHub CI workflow. This project provides a standardized workflow for simple samples. For details, please refer to [basic-workflow.yml](./.github/workflows/basic-workflow.yml). You only need to add the following options to add your module to the `basic` job in [ci.yml](./.github/workflows/ci.yml):

- `module.name`: the name of new module, should be same with the module directory name.
- `module.language`: the programming language, should be same with the directory name under project root.
- `module.with_oceanbase_container`: whether to use a pre-deployed OceanBase container, optional, set 'true' by default. If it's 'true', you can connect to it using username 'root@sys' or 'root@test' with empty password at localhost.

#### Complex samples

For complex samples that need to be placed in the applications directory, the directory also needs to contain the `README.md` document. In addition, its project files, ci workflow, etc. will be added by yourself.

## References

Refer to the [community website](https://open.oceanbase.com) for more details about OceanBase.
