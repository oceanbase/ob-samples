require 'mysql2'

def connect_to_oceanbase(host, port, username, password, database)
  begin
    # 创建数据库连接
    client = Mysql2::Client.new(
      host: host,
      port: port,
      username: username,
      password: password,
      database: database
    )

    puts "连接到OceanBase数据库成功"

    # 执行一个简单的查询
    results = client.query("SELECT DATABASE();")

    results.each do |row|
      puts "连接到的数据库: #{row['DATABASE()']}"
    end

    # 你可以在这里执行其他的SQL查询或操作

    client.close
    puts "MySQL连接已关闭"

  rescue Mysql2::Error => e
    puts "连接到OceanBase数据库失败: #{e.error}"
  end
end

if __FILE__ == $0
  connect_to_oceanbase(
    '172.30.3.244',
    2881,
    'root',
    'password',
    'test'
  )
end
