# Connect OceanBase with C++ (mysql-connector-cpp)

English | [简体中文](./README-CN.md)

## Introduction

This example demonstrates how to connect to OceanBase database using MySQL Connector/C++. The example shows basic database operations including:
- Creating a connection
- Creating a table
- Inserting data
- Querying data
- Error handling

## Prerequisites

- **OceanBase Database**:
  - You can use OceanBase Docker image for testing:
    ```bash
    docker pull oceanbase/oceanbase-ce
    docker run -p 2881:2881 --name oceanbase-ce -d oceanbase/oceanbase-ce
    ```

- **Development Environment**:
  - C++ compiler with C++11 support
  - MySQL Connector/C++ (libmysqlcppconn)

## Installation

First, install the MySQL Connector/C++ package onto your local machine. Refer to the [official documentation](https://dev.mysql.com/doc/dev/connector-cpp/latest/) for installation instructions for your specific operating system.

## Quick Start

1. **Clone the repository**:
   ```bash
   git clone https://github.com/oceanbase/ob-samples.git
   cd ob-samples/c_cplusplus/mysql-connector-cpp
   ```

2. **Compile and run the example**:
   ```bash
   ./run.sh
   ```

   Or with custom connection parameters:
   ```bash
   ./run.sh [url] [username] [password] [database]
   ```

## Connection Parameters

The example accepts the following optional parameters (all parameters are optional and will use default values if not provided):
- `url`: Database server URL (default: tcp://127.0.0.1:2881)
- `username`: Database username (default: root@test)
- `password`: Database password (default: "")
- `database`: Database name (default: test)

Example:
```bash
# Use default parameters
./run.sh

# Use custom parameters
./run.sh tcp://127.0.0.1:2881 root@test "" test
```

**Note**: Parameters must be provided in order, from left to right. You cannot skip parameters in the middle.

## Example Code

The example [mysql_connector_test](./src/mysql_connector_test.cpp) demonstrates basic database operations:

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

## Example Output

When you run the example with default parameters, you should see output as follows:

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

## Build Options

The example can be built using the provided `run.sh` script or manually:

```bash
# Using run.sh
./run.sh

# Manual build
g++ -std=c++11 -I/usr/include/mysql-cppconn-8 src/mysql_connector_test.cpp -o mysql_connector_test -lmysqlcppconn
```

**Note**: You may need to replace `/usr/include/mysql-cppconn-8` (in [run.sh](./run.sh)) with your actual MySQL Connector/C++ installation path.

## Troubleshooting

1. **Compilation Errors**:
   - Ensure MySQL Connector/C++ development package is properly installed
   - Check if the include path in `run.sh` is correct
   - Verify that your compiler supports C++11 standard

2. **Connection Errors**:
   - Verify that the OceanBase database service is running
   - Check if the database connection parameters are correct
   - Ensure network connectivity and port accessibility

3. **Runtime Errors**:
   - Check the specific error message for the root cause
   - Verify that the database user has sufficient permissions
   - Ensure the database schema and tables are properly created






