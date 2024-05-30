require 'active_record'

# 数据库配置
db_config = {
  adapter:  'mysql2',
  host:     '127.0.0.1',
  port:     2881,
  username: 'root',
  password: '',
  database: 'test'
}

begin
  # 建立连接
  ActiveRecord::Base.establish_connection(db_config)

  # 测试连接是否成功
  connection = ActiveRecord::Base.connection
  result = connection.active? # 检查数据库连接是否有效

  if result
    puts "成功连接到OceanBase数据库"
  else
    puts "连接到OceanBase数据库失败"
  end

rescue StandardError => e
  puts "连接到OceanBase数据库时发生错误: #{e.message}"
ensure
  ActiveRecord::Base.connection.close if ActiveRecord::Base.connected?
end
