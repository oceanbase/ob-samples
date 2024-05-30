require 'sequel'

# 连接到OceanBase数据库
DB = Sequel.connect(
  adapter:  'mysql2',
  host:     '127.0.0.1',
  port:     2881,
  user:     'root',
  password: '',
  database: 'test'
)

if DB.test_connection
  puts "成功连接到OceanBase数据库"
else
  puts "连接到OceanBase数据库失败"
end
