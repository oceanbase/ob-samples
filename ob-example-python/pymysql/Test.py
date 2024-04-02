#!/usr/bin/python3
import os

import pymysql

workspace = "/workspace/ob-example"
sql_file = "tests/sql/test.sql"
table_name = "t_test"


def load_sql_from_script(filename):
    fd = open(filename, 'r')
    f = fd.read()
    fd.close()
    return f.split(';')


if __name__ == "__main__":
    conn = pymysql.connect(host="127.0.0.1", port=2881,
                           user="root@test", passwd="", db="test")
    print("Success to connect to OceanBase with pymysql")

    # execute sql script
    for sql in load_sql_from_script(os.path.join(workspace, sql_file)):
        sql = sql.strip()
        if sql == "":
            continue
        with conn.cursor() as cur:
            print("Exec sql: " + sql)
            cur.execute(sql)

    # query on test table
    with conn.cursor() as cur:
        sql = 'SELECT * FROM ' + table_name
        print("Query sql: ", sql)
        cur.execute(sql)
        ans = cur.fetchall()
        print("Get rows: ")
        print(ans)
    conn.close()
