#!/bin/bash

# find mysql connector c++ headers
echo "Searching for MySQL Connector/C++ headers..."
POSSIBLE_PATHS=(
    "/usr/include/mysql-cppconn-8/jdbc"
    "/usr/include/mysql-cppconn-8"
    "/usr/include/"
)

# find actual header file location
for path in "${POSSIBLE_PATHS[@]}"; do
    if [ -f "$path/cppconn/driver.h" ] || [ -f "$path/driver.h" ]; then
        INCLUDE_PATH=$path
        echo "Found headers in: $INCLUDE_PATH"
        break
    fi
done

if [ -z "$INCLUDE_PATH" ]; then
    echo "Error: Could not find MySQL Connector/C++ headers"
    echo "Please check if libmysqlcppconn-dev is properly installed"
    echo "You can install it with: sudo apt-get install libmysqlcppconn-dev"
    exit 1
fi

# compile
echo "=== Compiling ==="
g++ -std=c++11 -I$INCLUDE_PATH src/mysql_connector_test.cpp -o mysql_connector_test -lmysqlcppconn

# check compile result
if [ $? -ne 0 ]; then
    echo "Compilation failed"
    exit 1
fi

# run
echo "=== Running ==="
./mysql_connector_test "$@"