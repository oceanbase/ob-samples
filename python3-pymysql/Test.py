#!/usr/bin/python3
  
import pymysql

conn = pymysql.connect(host="localhost", port=2881,
                       user="root", passwd="", db="test")
print("success to connect OceanBase with pymysql")
try:
    with conn.cursor() as cur:
        cur.execute('drop table if exists test')
        cur.execute('create table test(str varchar(256))')
        cur.execute('insert into test values ("hello Oceanbase")')
        cur.execute('select * from test')
        ans = cur.fetchall()
        print(ans[0][0])

finally:
    conn.close()
