#!/usr/bin/env python
# -*- coding:utf-8 -*-
#
#   Author  :   XueWeiHan
#   E-mail  :   595666367@qq.com
#   Date    :   2024-05-29 14:28
#   Desc    :   SQLAlchemy 引擎和 OceanBase 基表定义
from contextlib import contextmanager
from typing import Generator

from sqlalchemy import create_engine, Column, Integer
from sqlalchemy.orm import sessionmaker
from sqlalchemy.orm import DeclarativeBase

from .config import ob_db_url as db_url


# 创建SQLAlchemy引擎
engine = create_engine(db_url)
# 创建一个 session 类型
Session = sessionmaker(bind=engine)


class Base(DeclarativeBase):
    id = Column(Integer, primary_key=True, autoincrement=True)


@contextmanager
def get_db() -> Generator:
    session = Session()
    try:
        yield session
    except Exception:
        session.rollback()
        raise
    finally:
        session.close()
