package com.oceanbase.example;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.sql.*;
import java.util.Properties;
import java.util.stream.Collectors;

public class OceanBaseClientTest {
    public static void main(String[] args) {
        String workspace = "/workspace/ob-example";
        String sqlFile = "tests/sql/test.sql";
        String tableName = "t_test";

        Properties properties = new Properties();
        properties.put("user", "root@test");
        properties.put("password", "");
        String jdbcUrl = "jdbc:oceanbase://127.0.0.1:2881/test";

        Connection connection;
        Statement statement;
        try {
            connection = DriverManager.getConnection(jdbcUrl, properties);
            statement = connection.createStatement();
            System.out.println("Success to connect to OceanBase");
        } catch (SQLException e) {
            System.out.println("Failed to connect to OceanBase, exception: " + e.getMessage());
            return;
        }

        String[] sqlArray;
        Path sqlPath = Paths.get(workspace, sqlFile);
        try {
            sqlArray = Files.readAllLines(sqlPath).stream().map(String::trim).collect(Collectors.joining("\n"))
                .split(";");
        } catch (IOException e) {
            System.out.println("Failed to load sql file from " + sqlPath + ", exception: " + e.getMessage());
            return;
        }

        try {
            for (int i = 0; i < sqlArray.length; i++) {
                String sql = sqlArray[i];
                System.out.println("Execute sql: " + sql);
                statement.execute(sql);
            }
        } catch (SQLException e) {
            System.out.println("Execute sql exception: " + e.getMessage());
            return;
        }

        String selectSql = "SELECT * FROM " + tableName;
        System.out.println("Query sql: " + selectSql);
        try {
            ResultSet rs = statement.executeQuery(selectSql);
            ResultSetMetaData metaData = rs.getMetaData();
            System.out.println("Get rows:");
            int count = 0;
            while (rs.next()) {
                System.out.printf("## row %d: { ", count++);
                for (int i = 0; i < metaData.getColumnCount(); i++) {
                    System.out.print(metaData.getColumnName(i + 1) + ": " + rs.getObject(i + 1) + "; ");
                }
                System.out.println("}");
            }
        } catch (SQLException e) {
            System.out.println("Failed to query table " + tableName + ", exception: " + e.getMessage());
            return;
        }

        try {
            if (statement != null) {
                statement.close();
            }
            if (connection != null) {
                connection.close();
            }
        } catch (SQLException e) {
            System.out.println("Failed to close statement and connection, exception: " + e.getMessage());
        }
    }
}
