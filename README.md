# OceanBase Examples

English | [简体中文](README-CN.md)

## Introduction

This repository provides the example projects for OceanBase. It contains the following directories:

- [`examples`](examples): the directory contains all the example projects.
- [`tests`](tests): the directory contains resources for testing.
- [`tools`](tools): the directory contains scripts and other tools.

### Examples

For now, the repository contains examples for the following components:

#### [Driver](examples/driver)

- (golang) [go-sql-driver](examples/driver/golang-go-sql-driver)
- (java) [mysql-connector-java](examples/driver/java-mysql-connector-java)
- (java) [oceanbase-client](examples/driver/java-oceanbase-client)
- (python3) [pymysql](examples/driver/python3-pymysql)

## Quick Start

This repo builds an online platform for fast use based on Gitpod, click the following button to have a try.

[![Open in Gitpod](https://gitpod.io/button/open-in-gitpod.svg)](https://gitpod.io/#https://github.com/oceanbase/ob-example)

Note that it is necessary to choose a large-class workspace for OceanBase, otherwise the database may fail to deploy due to insufficient disk space.

Open the newly created workspace, Gitpod will Automatically deploy a standalone OceanBase server. Please wait until the following information about successful deployment appears on the terminal, after that you can try it with our examples referencing [examples/README.md](examples/README.md).

![Boot Success](tools/gitpod/boot.png)

## Contribution

We welcome contributions from anyone, thanks to all [contributors](https://github.com/oceanbase/ob-example/graphs/contributors)!

In this repository, the example projects in the same type are placed in the same directory, and the project directories are named in format `{programming language}-{component name}`.

Before you submit a Pull Request, we recommend that you first create a [workspace](https://gitpod.io/workspaces) on Gitpod to test and verify your fork branch.

Finally, in the newly added directory, there should be at least the following files:

- code files
- `run.sh` script to run code
- `README.md` documentation for component usage

## References

Refer to the [community website](https://open.oceanbase.com) for more details about OceanBase.
