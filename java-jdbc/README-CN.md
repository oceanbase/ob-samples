# Java 连接 OceanBase 指南

OceanBase 数据库支持通过 MySQL 官方 JDBC 驱动连接。

## 前提条件

* 确保计算机上的 Java 环境为 Java JDK 8 及以上版本。

* 安装 MySQL Connector/J，并配置运行环境。

    详细的下载及安装方法，请参考 [Connector/J Downloads](https://dev.mysql.com/downloads/connector/j/)、[Connector/J Installation](https://dev.mysql.com/doc/connector-j/8.0/en/connector-j-installing.html) 。

## 连接 OceanBase 数据库

下文以 Linux 中通过 Java 连接数据库为例。

在正确安装 MySQL Connector/J 8.0.30 驱动并配置环境之后，可以通过以下 `Test.java` 文件的示例代码进行数据库连接。

>**注意**
>
>如果是 MySQL Connector/J 5.x 版本，`Class.forName("com.mysql.cj.jdbc.Driver")` 中的 `com.mysql.cj.jdbc.Driver` 需要替换成 `com.mysql.jdbc.Driver`。

```java
import java.sql.Connection;
import java.sql.DriverManager;
import java.sql.ResultSet;
import java.sql.SQLException;
import java.sql.Statement;

public class Test {
   public static void main(String[] args) {
       try {

            Class.forName("com.mysql.cj.jdbc.Driver").newInstance();

            try{
                
                Connection connection = DriverManager.getConnection("jdbc:mysql://172.30.xx.xx:2883/test?user=r***&password=");
                System.out.println(connection.getAutoCommit());
                Statement sm = connection.createStatement();
                //执行删除表、新建表、插入数据等操作。
                String q1="drop table if exists t_meta_form";
                sm.executeUpdate(q1);
                String q2="CREATE TABLE t_meta_form ( name varchar(36) NOT NULL DEFAULT ' ', id int NOT NULL ) DEFAULT CHARSET = utf8mb4";
                String q3="insert into t_meta_form (name,id) values ('an','1')";
                sm.executeUpdate(q2);
                sm.executeUpdate(q3);                  

            }catch(SQLException se){
                System.out.println("error!");
                se.printStackTrace() ;
            }
            }catch (Exception ex) {
                ex.printStackTrace();
            }
    }
}
```

参数说明：

```java
connection = DriverManager.getConnection("jdbc:mysql://{hostname}:{port}/{dbname}?user={username}&password={password}")
```

* **hostname**：OceanBase 数据库连接 IP，通常是一个 OBProxy 地址。

* **port**：OceanBase 数据库连接端口，也是 OBProxy 的监听端口，默认是 2883，可以自定义。

* **dbname**：需要访问的数据库名称。

* **username**：租户的连接账户，租户的管理员用户名默认是 `root`。

* **password**：账户密码。

示例：`jdbc:mysql://172.30.xx.xx:2883/test?user=r***&password=***1**`

代码编辑完成后，可以通过如下命令进行编译和执行，返回 true 说明数据库连接成功。

```bash
//手动添加驱动到环境配置，根据 mysql-connector-java.jar 实际安装路径填写
export CLASSPATH=/usr/share/java/mysql-connector-java.jar:$CLASSPATH
//编译
javac Test.java
//运行
java Test
```
