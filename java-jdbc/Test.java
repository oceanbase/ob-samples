import java.sql.Connection;
import java.sql.DriverManager;
import java.sql.ResultSet;
import java.sql.SQLException;
import java.sql.Statement;
public class Test {
   public static void main(String[] args) {
       try {

            Class.forName("com.mysql.jdbc.Driver");

            try{
                
                Connection connection = DriverManager.getConnection("jdbc:mysql://127.0.0.1:2881/test?user=root&password=");
                System.out.println("success to connect OceanBase with java jdbc");
                Statement sm = connection.createStatement();
                String q1="drop table if exists test";
                String q2="CREATE TABLE test( name varchar(36) NOT NULL DEFAULT ' ') DEFAULT CHARSET = utf8mb4";
                String q3="insert into test values ('Hello OceanBase')";
                String q4="select * from test limit 1";
                sm.executeUpdate(q1);
                sm.executeUpdate(q2);
                sm.executeUpdate(q3);                  
                ResultSet rs = sm.executeQuery(q4);
                rs.beforeFirst();
                while(rs.next()){
                 String Name = rs.getString("name");
                 System.out.printf("%s\n",Name);
                }
                rs.close();
                } catch (SQLException se) {
                  System.out.println("error!");
                  se.printStackTrace() ;
                 }
            } catch (Exception ex) {
                ex.printStackTrace();
            }
    }
}
