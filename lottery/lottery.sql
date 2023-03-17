CREATE DATABASE IF NOT EXISTS lottery;
USE lottery;

# 奖品
# NOTE：art_name 可以存放ascii-art格式数据，可以以图片的形式展示奖品
CREATE TABLE prize (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `type` varchar(60) NOT NULL COMMENT '奖品级别',
  `name` varchar(200) NOT NULL COMMENT '奖品名称',
  `count` bigint NOT NULL COMMENT '奖品个数',
  `art_name` LONGTEXT NOT NULL COMMENT '奖品ascii-art展示',
  PRIMARY KEY (`id`)
);
CREATE UNIQUE INDEX prize_index_name ON prize(`name`);

# 抽奖候选人
CREATE TABLE `candidate` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `name` varchar(20) NOT NULL COMMENT '姓名',
    `prize_id` bigint unsigned COMMENT '获奖',
    PRIMARY KEY (`id`)
);

# 导入数据前调整一些参数
# global 参数调整后需要重新连接客户端才会生效
SET GLOBAL secure_file_priv = "";
set global ob_query_timeout=36000000000;

# 导入奖品数据
# 由于奖品的"图片"比较大，直接放文件中导入
source prizes/keyboard.sql;
source prizes/airpods.sql;
source prizes/charger.sql;

# 查看抽奖人数
SELECT COUNT(1) from candidate;

# 查看候选人列表
SELECT name from candidate;

# 展示三等奖奖品
SELECT name from prize where type='三等奖' \G;
select art_name from prize where type='三等奖' \G;

# 筛选三等奖
# 抽1个
UPDATE candidate set prize_id=(select id from prize where type='三等奖') where prize_id is null order by rand() limit 1;

# 抽3个
UPDATE candidate set prize_id=(select id from prize where type='三等奖') where prize_id is null order by rand() limit 3;

# 查看三等奖获得者
SELECT name as '三等奖获得者' from candidate where prize_id in (select id from prize where type='三等奖');

# 展示二等奖奖品
SELECT name from prize where type='二等奖'\G;
SELECT art_name from prize where type='二等奖' \G;

#筛选二等奖
# 抽1个
UPDATE candidate set prize_id=(select id from prize where type='二等奖') where prize_id is null order by rand() limit 1;

# 抽2个
UPDATE candidate set prize_id=(select id from prize where type='二等奖') where prize_id is null order by rand() limit 2;

# 查看2等奖获得者
SELECT name as '二等奖获得者' from candidate where prize_id in (select id from prize where type='二等奖');

#展示一等奖奖品
SELECT name from prize where type='一等奖'\G;
SELECT art_name from prize where type='一等奖' \G;


#筛选一等奖
UPDATE candidate set prize_id=(select id from prize where type='一等奖') where prize is null order by rand() limit 1;

# 查看1等奖获得者
SELECT name as '一等奖获得者' from candidate where prize_id in (select id from prize where type='一等奖');

# 红色展示
# echo -n -e '\e[31m' ; echo select name from lottery.prize | obclient -N -h127.0.0.1 -P55800 -uroot -Doceanbase -A ; echo -n -e '\e[0m'
