#!/usr/bin/env python
# -*- coding:utf-8 -*-
#
#   Date    :   2024-05-29 14:28
#   Desc    :   操作表的数据
from sqlalchemy.orm import Session

from db.models import User


def insert_user(db: Session, name: str, age: int):
    # 新增用户数据
    user = User(name=name, age=age)
    db.add(user)
    # 提交事务, 保存数据。get_db 中做了所以这里不需要再提交
    db.commit()


def query_all_user(db: Session):
    users = db.query(User).all()
    for user in users:
        print(user)
