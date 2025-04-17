#!/bin/bash

# show environment info
echo "=== Environment Info ==="
echo "C++ compiler:"
g++ --version

echo "=== Package Info ==="
echo "Checking installed packages..."
dpkg -l | grep mysql
dpkg -L libmysqlcppconn-dev

echo "=== File Search ==="
echo "Searching for MySQL files..."
find /usr -name "*mysql*" 2>/dev/null
find /usr -name "*.h" | grep -i mysql 2>/dev/null

# set possible paths
POSSIBLE_PATHS=(
    "/usr/include/mysql-cppconn-8"
    "/usr/include/mysql"
    "/usr/include/mysql-cppconn-8/mysql"
    "/usr/include/mysql-cppconn-8/jdbc"
)

# check each possible path
for path in "${POSSIBLE_PATHS[@]}"; do
    if [ -d "$path" ]; then
        echo "Found directory: $path"
        ls -la $path
        if [ -f "$path/mysql/jdbc.h" ] || [ -f "$path/jdbc.h" ]; then
            INCLUDE_PATH=$path
            echo "Using include path: $INCLUDE_PATH"
            break
        fi
    fi
done

if [ -z "$INCLUDE_PATH" ]; then
    echo "Error: Could not find MySQL Connector/C++ include path"
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