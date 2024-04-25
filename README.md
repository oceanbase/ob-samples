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

In the module, make sure to include the following files:

- code files
- `run.sh` script to run code
- `README.md` documentation for component usage.

If it's hard to use English for you, you can use your native language in the documentation, and we can improve it later.

To ensure the sample works, please add your module to the GitHub CI workflow.

- `module.name`: the name of new module, should be same with the module directory name.
- `module.language`: the programming language, should be same with the directory name under project root.
- `module.with_oceanbase_container`: whether to use a pre-deployed OceanBase container, optional, set 'true' by default. If it's 'true', you can connect to it using username 'root@sys' or 'root@test' with empty password at localhost.

See [ci.yml](./.github/workflows/ci.yml) for more details about CI workflow.

## References

Refer to the [community website](https://open.oceanbase.com) for more details about OceanBase.
