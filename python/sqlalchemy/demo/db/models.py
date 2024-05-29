#!/usr/bin/env python
# -*- coding:utf-8 -*-
#
#   Author  :   XueWeiHan
#   E-mail  :   595666367@qq.com
#   Date    :   2024-05-29 14:28
#   Desc    :   配置文件
from sqlalchemy import Column, Integer, VARCHAR

from db.base import Base, engine


# 定义一个 ORM 模型，映射到数据库中的表
class User(Base):
    __tablename__ = 'users'

    id = Column(Integer, primary_key=True)
    name = Column(VARCHAR(50))
    age = Column(Integer)

    def __repr__(self):
        return f"<User(name={self.name}, age={self.age})>"


if __name__ == '__main__':
    # 创建数据库表（如果表不存在）
    Base.metadata.create_all(engine)
