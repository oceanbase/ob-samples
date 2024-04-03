#!/usr/bin/python3

import pymysql

if __name__ == "__main__":
    conn = pymysql.connect(host="127.0.0.1", port=2881,
                           user="root@test", passwd="", db="test")
    print("Success to connect to OceanBase with pymysql")

    sqls = ['DROP TABLE IF EXISTS `t_test`',
            'CREATE TABLE `t_test` ('
            '  `id` int(10) NOT NULL AUTO_INCREMENT, '
            '  `name` varchar(20) DEFAULT NULL, '
            '  PRIMARY KEY (`id`)'
            ') ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE = utf8_bin',
            'INSERT INTO `t_test` VALUES (default, "Hello OceanBase")']

    # execute sql script
    with conn.cursor() as cur:
        for sql in sqls:
            print("Exec sql: " + sql)
            cur.execute(sql)

    # query on test table
    with conn.cursor() as cur:
        sql = 'SELECT * FROM `t_test`'
        print("Query sql: ", sql)
        cur.execute(sql)
        ans = cur.fetchall()
        print("Get rows: ")
        print(ans)
    conn.close()
