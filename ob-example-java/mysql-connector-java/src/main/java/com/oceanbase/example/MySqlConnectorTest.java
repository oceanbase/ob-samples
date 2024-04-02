package com.oceanbase.example;

import java.sql.Connection;
import java.sql.DriverManager;
import java.sql.ResultSet;
import java.sql.ResultSetMetaData;
import java.sql.SQLException;
import java.sql.Statement;
import java.util.ArrayList;
import java.util.List;

public class MySqlConnectorTest {

    private static final String JDBC_URL = "jdbc:mysql://127.0.0.1:2881/test";
    private static final String USERNAME = "root@test";
    private static final String PASSWORD = "";

    public static void main(String[] args) {
        try (Connection connection = DriverManager.getConnection(JDBC_URL, USERNAME, PASSWORD);
             Statement statement = connection.createStatement()) {
            statement.execute("DROP TABLE IF EXISTS `t_test`");
            statement.execute("CREATE TABLE `t_test` (" +
                "    `id`   int(10) NOT NULL AUTO_INCREMENT," +
                "    `name` varchar(20) DEFAULT NULL," +
                "    PRIMARY KEY (`id`)" +
                ") ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE = utf8_bin");
            statement.execute("INSERT INTO `t_test` VALUES (default, 'Hello OceanBase')");

            ResultSet rs = statement.executeQuery("SELECT * FROM `t_test`");
            ResultSetMetaData metaData = rs.getMetaData();

            List<String> result = new ArrayList<>();
            while (rs.next()) {
                StringBuilder sb = new StringBuilder();
                for (int i = 0; i < metaData.getColumnCount(); i++) {
                    if (i != 0) {
                        sb.append(",");
                    }
                    Object value = rs.getObject(i + 1);
                    sb.append(value == null ? "null" : value.toString());
                }
                result.add(sb.toString());
            }

            System.out.println(result);

            assert result.size() == 1;
            assert result.get(0).equals("0,Hello OceanBase");
        } catch (SQLException e) {
            throw new RuntimeException(e);
        }
    }
}
