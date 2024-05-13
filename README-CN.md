# OceanBase Examples

[English](README.md) | 简体中文

## 介绍

本仓库提供了 OceanBase 的示例项目。

对于基础的示例，我们将它们按照编程语言分组，示例目录以工具名称命名。

对于作为独立应用程序工作的示例，我们将它们放在 applications 目录下，示例目录名称一般包含主要工具名和用途。

您可以将示例代码复制到本地环境并在更改连接信息后直接使用，更多详细信息请参阅示例目录中的README文件。

## 贡献

我们欢迎任何人来贡献，感谢所有的[贡献者](https://github.com/oceanbase/ob-samples/graphs/contributors)！

我们在 GitHub 上创建了一些任务，这些任务难度不高，对于首次贡献者来说比较友好，欢迎感兴趣的开发者认领：https://github.com/oceanbase/ob-samples/issues/16 。

### 增加一个示例

在本仓库中，一个示例将作为一个独立的模块存在，模块的目录名称应当与示例所用的工具保持一致，如 `mysql-connector-java` 的示例目录就命名为 `mysql-connector-java`。

在本仓库的根目录内有许多分类目录。对于能够通过简单的命令可以直接运行的示例，我们建议按照工具所需要的编程语言环境来分类，如 `mysql-connector-java` 需要添加到 `java` 目录下。 对于需要比较复杂的配置才能运行的示例，我们建议将示例项目放到 `applications` 目录。

#### 简单示例

对于编程语言分类下的简单示例，其目录内应当包含以下内容

- 代码文件
- `run.sh` 运行代码的脚本
- `README.md` 组件用法的介绍文档

如果您不能提供英文的文档，您可以在文档中使用您的母语，我们会在之后对其进行改进。

为了确保示例能够在 GitHub Action 中运行，您的模块添加到 GitHub CI 工作流程。本项目对简单示例提供了一套标准化的运行流程，详情请参阅 [basic-ci.yml](./.github/workflows/basic-ci.yml)。您只需要添加以下内容，将您贡献的模块增加到 [ci.yml](./.github/workflows/ci.yml) 中的 basic job 中即可：

- `module.name`：新模块的名称，应与模块目录名称相同。
- `module.language`：编程语言，应与项目根目录下的目录名称相同。
- `module.with_oceanbase_container`：是否使用预先部署的 OceanBase 容器，可选，默认设置为 true。如果它是 true，您可以在 localhost 上使用用户名 `root@sys` 或 `root@test` 以及空密码连接到它。

#### 复杂示例

对于需要放到 applications 目录下的复杂示例，其目录内同样需要包含 `README.md` 文档，除此之外，其项目文件、ci 流程等由贡献者自行判断添加。

## 参考信息

关于更多 OceanBase 的细节请参考 [社区官网](https://open.oceanbase.com)。
