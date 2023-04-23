# OceanBase Examples

English | [简体中文](README-CN.md)

## Quick Start

This repo builds an online platform for fast use based on Gitpod, click the following button to have a try.

[![Open in Gitpod](https://gitpod.io/button/open-in-gitpod.svg)](https://gitpod.io/#https://github.com/oceanbase/ob-example)

Note that it is necessary to choose a large-class workspace for OceanBase, otherwise the database may fail to deploy due to insufficient disk space.

![workspace](tools/gitpod/new-workspace.png)

Open the newly created workspace, Gitpod will Automatically deploy a standalone OceanBase server. Please wait until the following information about successful deployment appears on the terminal, after that you can try it with our examples.

![demo](tools/gitpod/demo.png)

### Try with Example Code

This repository provides examples for different languages and tools connecting to OceanBase. You can follow the steps below to run the example code directly in the Gitpod environment through `run.sh`.

```bash
// enter the working directory
cd xxxx
// execute the sample code
sh run.sh
```

Here we use python3-pymysql as an example:

```bash
cd python3-pymysql
sh run.sh
```

## Contribution

We welcome contributions from anyone, thanks to all [contributors](https://github.com/oceanbase/ob-example/graphs/contributors)!

In this repository, each subdirectory is an independent project, and the naming format is `{programming language}-{component name}`.

Before you submit a Pull Request, we recommend that you first create a [workspace](https://gitpod.io/workspaces) on Gitpod to test and verify your fork branch.

Finally, in the newly added directory, there should be at least the following files:
- code files
- `run.sh` script to run code
- `README.md` documentation for component usage

## References

Refer to the [community website](https://open.oceanbase.com) for more details about OceanBase.
