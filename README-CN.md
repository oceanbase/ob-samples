# ob-example

[English](README.md) | 简体中文

本仓库基于 Gitpod 建立了快速在线体验平台, 点击下面按钮一键体验（建议使用 Chrome 浏览器）

[![Open in Gitpod](https://gitpod.io/button/open-in-gitpod.svg)](https://gitpod.io/#https://github.com/oceanbase/ob-example)

进入 Gitpod 后，在工作区内会自动部署一个 OceanBase 本地实例，请等待左侧终端界面显示 "ob boot success!"，之后您可以在右侧终端进行操作。

![示意图](./tools/scripts/gitpod1.png)

本仓库提供了不同语言和工具连接 OceanBase 的示例，您可以按下面步骤进行操作，通过 `run.sh` 在 Gitpod 环境直接运行示例代码。

```bash
// 进入目录
cd xxxx
// 执行示例代码
sh run.sh
```

这里以 python3-pymysql 为例

```bash
cd python3-pymysql
sh run.sh
```

关于更多 OceanBase 的细节请参考 [社区官网](https://open.oceanbase.com).
