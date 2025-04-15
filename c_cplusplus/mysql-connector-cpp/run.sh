#!/bin/bash

# Usage: ./run.sh [url] [username] [password] [database]
#   url:      database server URL (default: tcp://127.0.0.1:2881)
#   username: database username (default: root@test)
#   password: database password (default: "")
#   database: database name (default: test)

# Note: Parameters must be provided in order, from left to right.
#       You cannot skip parameters in the middle.
#       For example:
#         Valid:   ./run.sh tcp://127.0.0.1:2881 root@test "" test
#         Valid:   ./run.sh tcp://127.0.0.1:2881
#         Invalid: ./run.sh tcp://127.0.0.1:2881 test  (skipping username and password)
# All parameters are optional and will use default values if not provided



# Compile the program
# Note: You may need to modify the include path (-I) based on your MySQL Connector/C++ installation
g++ -std=c++11 -I/usr/include/mysql-cppconn-8 src/mysql_connector_test.cpp -o mysql_connector_test -lmysqlcppconn

# Run the test 
./mysql_connector_test "$@"