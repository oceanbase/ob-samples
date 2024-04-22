package com.oceanbase.samples;

import java.sql.Connection;
import java.sql.DriverManager;
import java.sql.PreparedStatement;
import java.sql.ResultSet;
import java.sql.SQLException;
import java.sql.Statement;

public class InsertAndSelectExample {

    private static final String JDBC_URL = "jdbc:oceanbase://127.0.0.1:2881/test?characterEncoding=utf-8&useServerPrepStmts=true";
    private static final String USERNAME = "root@test";
    private static final String PASSWORD = "";

    public static void main(String[] args) throws ClassNotFoundException, SQLException {
        // connect to your database
        Class.forName("com.oceanbase.jdbc.Driver");
        Connection conn = DriverManager.getConnection(JDBC_URL, USERNAME, PASSWORD);

        // create a table
        Statement stmt = conn.createStatement();
        try {
            stmt.execute("drop table person");
        } catch (Exception ignore) {
        }
        stmt.execute("create table person (name varchar(50), age int)");

        // insert records
        PreparedStatement ps = conn.prepareStatement("insert into person values(?, ?)");
        ps.setString(1, "Adam");
        ps.setInt(2, 28);
        ps.executeUpdate();
        ps.setString(1, "Eve");
        ps.setInt(2, 26);
        ps.executeUpdate();

        // fetch all records
        ps = conn.prepareStatement("select * from person", ResultSet.TYPE_FORWARD_ONLY, ResultSet.CONCUR_READ_ONLY);
        ResultSet rs = ps.executeQuery();
        while (rs.next()) {
            System.out.println(rs.getString(1) + " is " + rs.getInt(2) + " years old.");
        }

        // release all resources
        ps.close();
        stmt.close();
        conn.close();
    }

}
