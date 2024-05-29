# -*- coding:utf-8 -*-
#
#   Author  :   XueWeiHan
#   E-mail  :   595666367@qq.com
#   Date    :   2024-05-29 14:28
#   Desc    :   运行入口
from db.base import get_db
from db.curd import insert_user, query_all_user


if __name__ == '__main__':
    with get_db() as db:
        # 新增用户
        insert_user(db, 'HelloGitHub', 8)
        # 查询所有用户
        query_all_user(db)
