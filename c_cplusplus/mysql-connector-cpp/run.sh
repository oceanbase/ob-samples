#!/bin/bash

# environment info
echo "=== Environment Info ==="
echo "C++ compiler:"
g++ --version
echo "MySQL Connector/C++:"
find /usr -name "jdbc.h" 2>/dev/null

# set include path
INCLUDE_PATH="/usr/include/mysql-cppconn-8"

# check if the path exists
if [ ! -d "$INCLUDE_PATH/mysql" ]; then
    echo "Warning: Default include path not found, searching alternatives..."
    # search alternatives
    ALTERNATIVE_PATH=$(find /usr -name "jdbc.h" 2>/dev/null | grep "mysql/jdbc.h" | xargs dirname | xargs dirname)
    if [ -n "$ALTERNATIVE_PATH" ]; then
        INCLUDE_PATH="$ALTERNATIVE_PATH"
        echo "Using alternative path: $INCLUDE_PATH"
    else
        echo "Error: Could not find MySQL Connector/C++ include path"
        exit 1
    fi
fi

# compile
echo "=== Compiling ==="
echo "Using include path: $INCLUDE_PATH"
g++ -std=c++11 -I$INCLUDE_PATH src/mysql_connector_test.cpp -o mysql_connector_test -lmysqlcppconn

# check compile result
if [ $? -ne 0 ]; then
    echo "Compilation failed"
    exit 1
fi

# run
echo "=== Running ==="
./mysql_connector_test "$@"