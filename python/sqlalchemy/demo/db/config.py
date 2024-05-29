#!/usr/bin/env python
# -*- coding:utf-8 -*-
#
#   Author  :   XueWeiHan
#   E-mail  :   595666367@qq.com
#   Date    :   2024-05-29 14:28
#   Desc    :   配置文件


class Config(object):

    @staticmethod
    def oceanbase_db_url():
        # OceanBase 数据库连接参数
        username = 'root'
        password = ''
        host = 'localhost'
        port = '3306'
        database = 'test'
        db_url = 'mysql+pymysql://{username}:{password}@{host}:{port}/{database}'.format(
            username=username, password=password, host=host, port=port, database=database)
        return db_url


ob_db_url = Config.oceanbase_db_url()
