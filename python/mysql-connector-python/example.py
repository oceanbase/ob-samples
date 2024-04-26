import json
import mysql.connector

# 替换为你自己的 OceanBase 服务器和数据库信息
db_config = {
    'user': 'root@test',
    'password': '',
    'host': '127.0.0.1',
    'port': '2881',   
    'database': 'test',
}

# 替换为你的 JSON 文件路径
json_file_path = 'example-data.json'

# 要分配给每个记录的 ProductID
product_id = 1

# 读取 JSON 数据
with open(json_file_path, 'r', encoding='utf-8') as file:
    data = json.load(file)

# 连接数据库
cnx = mysql.connector.connect(**db_config)
cursor = cnx.cursor()

# 创建 Products 表和 ProductFiles 表
create_products_table = """
CREATE TABLE IF NOT EXISTS Products (
    ProductID INT AUTO_INCREMENT PRIMARY KEY,
    ProductName VARCHAR(255) NOT NULL,
    ProductVersion VARCHAR(50) NOT NULL,
    UNIQUE (ProductName, ProductVersion)
);
"""
create_productfiles_table = """
CREATE TABLE IF NOT EXISTS ProductFiles (
    FileID INT AUTO_INCREMENT PRIMARY KEY,
    ProductID INT NOT NULL,
    FileName VARCHAR(255) NOT NULL,
    FileSummary TEXT,
    FileKeywords TEXT,
    FileContent MEDIUMTEXT,
    FilePurpose VARCHAR(255),
    FOREIGN KEY (ProductID) REFERENCES Products(ProductID)
      ON DELETE CASCADE ON UPDATE CASCADE
);
"""
# 执行创建表的 SQL 语句
cursor.execute(create_products_table)
cursor.execute(create_productfiles_table)

# 读取 JSON 数据
with open(json_file_path, 'r', encoding='utf-8') as file:
    data = json.load(file)


# SQL 插入语句
insert_stmt = (
    "INSERT INTO ProductFiles (ProductID, FileName, FileContent) "
    "VALUES (%s, %s, %s)"
)

# 初始化数据插入计数器
inserted_count = 0

# 遍历 JSON 数据，准备并执行 SQL 插入语句
for item in data:
    # 将 JSON 的 'path' 和 'content' 字段分别映射到数据库的 'FileName' 和 'FileContent'
    values = (product_id, item["path"], item["content"])
    
    try:
        cursor.execute(insert_stmt, values)
        inserted_count += 1  # 成功插入后，增加计数
    except mysql.connector.Error as err:
        print(f"Failed to insert data: {err}")

# 提交事务，确保数据持久化到数据库
cnx.commit()

# 关闭游标和连接
cursor.close()
cnx.close()

# 输出成功插入的数据条数
print(f"Data has been successfully inserted into the database. Total inserted records: {inserted_count}.")
