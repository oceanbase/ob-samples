# Lottery
这是一个基于[OceanBase](https://github.com/oceanbase/oceanbase)数据库的抽奖程序。

lottery.sql 包含了数据库的初始化代码以及查看奖品、抽奖、查看中奖人员的样例代码。

prize 目录下包含了初始化奖品数据的SQL文件，其中包含了使用ascii-art展示图片的数据，因此每条数据都比较大。

**关于ascii-art**
ascii-art是将图片转换为ascii字符，使用字符来展示图片的方法。更详细的信息可以搜索互联网。
这里的数据是使用 Python 库 [ascii_art](https://github.com/jontonsoup4/ascii_art) 转换的。
