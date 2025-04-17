# 使用 C++ 连接 OceanBase (mysql-connector-cpp)

[English](./README.md) | 简体中文

## 简介

本示例演示了如何使用 MySQL Connector/C++ 连接 OceanBase 数据库。示例展示了基本的数据库操作，包括：
- 创建连接
- 创建表
- 插入数据
- 查询数据
- 错误处理

## 环境要求

- **OceanBase 数据库**：
  - 可以使用 OceanBase Docker 镜像进行测试：
    ```bash
    docker pull oceanbase/oceanbase-ce
    docker run -p 2881:2881 --name oceanbase-ce -d oceanbase/oceanbase-ce
    ```

- **开发环境**：
  - 支持 C++11 的 C++ 编译器
  - MySQL Connector/C++ (libmysqlcppconn)

## 安装

首先，在本地机器上安装 MySQL Connector/C++ 包。请参考[官方文档](https://dev.mysql.com/doc/dev/connector-cpp/latest/)获取针对您特定操作系统的安装说明。

## 快速开始

1. **克隆仓库**：
   ```bash
   git clone https://github.com/oceanbase/ob-samples.git
   cd ob-samples/c_cplusplus/mysql-connector-cpp
   ```

2. **编译并运行示例**：
   ```bash
   ./run.sh
   ```

   或者使用自定义连接参数：
   ```bash
   ./run.sh [url] [username] [password] [database]
   ```

## 连接参数

示例接受以下可选参数（所有参数都是可选的，如果不提供将使用默认值）：
- `url`: 数据库服务器 URL（默认：tcp://127.0.0.1:2881）
- `username`: 数据库用户名（默认：root@test）
- `password`: 数据库密码（默认：""）
- `database`: 数据库名称（默认：test）

示例：
```bash
# 使用默认参数
./run.sh

# 使用自定义参数
./run.sh tcp://127.0.0.1:2881 root@test "" test
```

**注意**：参数必须按顺序从左到右提供，不能跳过中间参数。

## 示例代码

示例 [mysql_connector_test](./src/mysql_connector_test.cpp) 演示了基本的数据库操作：

```cpp
/* Standard C++ includes */
#include <iostream>
#include <memory>

#include <mysql/jdbc.h>

#define DEFAULT_URI "tcp://127.0.0.1:2881"
#define EXAMPLE_USER "root@test"
#define EXAMPLE_PASSWORD ""
#define EXAMPLE_DB "test"

using namespace std;

int main(int argc, char *argv[]) {
    try {
        const char   *url = (argc > 1 ? argv[1] : DEFAULT_URI);
        const string user(argc >= 3 ? argv[2] : EXAMPLE_USER);
        const string password(argc >= 4 ? argv[3] : EXAMPLE_PASSWORD);
        const string database(argc >= 5 ? argv[4] : EXAMPLE_DB);
        cout << "url: " << url << endl;
        cout << "user: " << user << endl;
        cout << "password: " << password << endl;
        cout << "database: " << database << endl;
        sql::mysql::MySQL_Driver *driver = sql::mysql::get_mysql_driver_instance();

        cout << "Connected to OceanBase server..." << endl;
        unique_ptr< sql::Connection > con{driver->connect(url, user, password)};
        
        
        con->setSchema(database);
        
        unique_ptr< sql::Statement > stmt{con->createStatement()};
        
        // Drop table if exists
        cout << "Dropping table if exists..." << endl;
        stmt->execute("DROP TABLE IF EXISTS `t_test`");
        
        // Create table
        cout << "Creating table..." << endl;
        stmt->execute("CREATE TABLE `t_test` ("
                     "    `id`   int(10) NOT NULL AUTO_INCREMENT,"
                     "    `name` varchar(20) DEFAULT NULL,"
                     "    PRIMARY KEY (`id`)"
                     ") ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE = utf8_bin");
        
        // Insert data
        cout << "Inserting data..." << endl;
        stmt->execute("INSERT INTO `t_test` VALUES (default, 'Hello OceanBase')");
        
        // Query data
        cout << "Querying data..." << endl;
        std::unique_ptr< sql::ResultSet > res{stmt->executeQuery("SELECT * FROM `t_test`")};
        
        cout << "Query results:" << endl;
        while (res->next()) {
            cout << "id: " << res->getInt("id") 
                 << ", name: " << res->getString("name") << endl;
        }
        
        cout << "Done!" << endl;
    }
    catch (sql::SQLException &e)
  {
    /*
      The JDBC API throws three different exceptions:
 
    - sql::MethodNotImplementedException (derived from sql::SQLException)
    - sql::InvalidArgumentException (derived from sql::SQLException)
    - sql::SQLException (derived from std::runtime_error)
    */
 
    cout << "# ERR: SQLException in " << __FILE__;
    cout << "(" << "EXAMPLE_FUNCTION" << ") on line " << __LINE__ << endl;
 
    /* Use what() (derived from std::runtime_error) to fetch the error message */
 
    cout << "# ERR: " << e.what();
    cout << " (MySQL error code: " << e.getErrorCode();
    cout << ", SQLState: " << e.getSQLState() << " )" << endl;
 
    return EXIT_FAILURE;
  }
 

  return EXIT_SUCCESS;
}
```

## 示例输出

使用默认参数运行示例时，您将看到如下输出：

```bash
$ ./run.sh
Connected to OceanBase server...
Dropping table if exists...
Creating table...
Inserting data...
Querying data...
Query results:
id: 1, name: Hello OceanBase
Done!
```

## 编译选项

示例可以使用提供的 `run.sh` 脚本或手动编译：

```bash
# 使用 run.sh
./run.sh

# 手动编译
g++ -std=c++11 -I/usr/include/mysql-cppconn-8 src/mysql_connector_test.cpp -o mysql_connector_test -lmysqlcppconn
```

**注意**：您可能需要将 `/usr/include/mysql-cppconn-8`（在 [run.sh](./run.sh) 中）替换为您实际的 MySQL Connector/C++ 安装路径。

## 故障排除

1. **编译错误**：
   - 确保已正确安装 MySQL Connector/C++ 开发包
   - 检查 `run.sh` 中的头文件包含路径是否正确
   - 确保编译器支持 C++11 标准

2. **连接错误**：
   - 确认 OceanBase 数据库服务是否正常运行
   - 检查数据库连接参数是否正确
   - 确认网络连接是否正常，端口是否开放

3. **运行时错误**：
   - 查看具体的错误信息，了解问题原因
   - 确认数据库用户是否有足够的权限执行操作
   - 确保数据库 schema 和表已正确创建
