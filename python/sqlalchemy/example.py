#!/usr/bin/python3
from sqlalchemy import create_engine, Column, Integer, VARCHAR
from sqlalchemy.orm import sessionmaker
from sqlalchemy.ext.declarative import declarative_base

# OceanBase 数据库连接参数
username = 'root'
password = ''
host = 'localhost'
port = '2881'
database = 'test'

# 创建一个 SQLAlchemy 引擎，连接到 OceanBase 数据库
# 创建数据库连接字符串
connection_string = f'mysql+pymysql://{username}:{password}@{host}:{port}/{database}'
# 创建SQLAlchemy引擎
engine = create_engine(connection_string)


# 创建一个 session 类型
Session = sessionmaker(bind=engine)

# 创建一个基类，用于定义 ORM 模型
Base = declarative_base()


# 定义一个 ORM 模型，映射到数据库中的表
class User(Base):
    __tablename__ = 'users'

    id = Column(Integer, primary_key=True)
    name = Column(VARCHAR(50))
    age = Column(Integer)

    def __repr__(self):
        return f"<User(name={self.name}, age={self.age})>"


# 创建数据库表（如果表不存在）
Base.metadata.create_all(engine)

# 创建 session 实例
session = Session()

# 添加一些用户数据
new_user_1 = User(name='Alice', age=30)
new_user_2 = User(name='Bob', age=25)

session.add(new_user_1)
session.add(new_user_2)

# 提交事务
session.commit()

# 查询数据
users = session.query(User).all()
for user in users:
    print(user)

# 关闭 session
session.close()
