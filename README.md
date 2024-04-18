# OceanBase Examples

English | [简体中文](README-CN.md)

## Introduction

This repository provides the example projects for OceanBase， which are grouped by programming language. The example directory is named after the component name.

This repository uses the [oceanbase-ce](https://hub.docker.com/r/oceanbase/oceanbase-ce) image to start the OceanBase database instance in the GitHub workflow for continuous integration. Users can also copy the example code to local environment and use it after changing the connection info.

## Contribution

We welcome contributions from anyone, thanks to all [contributors](https://github.com/oceanbase/ob-samples/graphs/contributors)!

If you want to add a new module, please make sure to place it in the directory of corresponding programming language and name it after the component name.

In the module, make sure to include the following files:

- code files
- `run.sh` script to run code
- `README.md` documentation for component usage
- Add your module's name to the github CI workflow ( please refer to https://github.com/oceanbase/ob-samples/.github/workflows/ci.yml )

## References

Refer to the [community website](https://open.oceanbase.com) for more details about OceanBase.
