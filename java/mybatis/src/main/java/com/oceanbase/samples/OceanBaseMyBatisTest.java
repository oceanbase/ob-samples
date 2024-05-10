package com.oceanbase.samples;


import com.oceanbase.samples.entity.User;
import org.apache.ibatis.io.Resources;
import org.apache.ibatis.session.SqlSession;
import org.apache.ibatis.session.SqlSessionFactory;
import org.apache.ibatis.session.SqlSessionFactoryBuilder;

import java.io.IOException;
import java.io.InputStream;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.Objects;


public class OceanBaseMyBatisTest
{

    public static void main( String[] args )
    {
        SqlSession sqlSession = null;
        try {
            // Get SqlSessionFactoryBuilder
            SqlSessionFactoryBuilder sqlSessionFactoryBuilder = new SqlSessionFactoryBuilder();
            // Load mybatis-config.xml as InputStream
            InputStream inputStream = Resources.getResourceAsStream("mybatis-config.xml");
            // Get SqlSessionFactory
            SqlSessionFactory sqlSessionFactory = sqlSessionFactoryBuilder.build(inputStream);
            // Get SqlSession
            sqlSession = sqlSessionFactory.openSession();
            // Execute SQL
            // insertTest(sqlSession);
            // updateTest(sqlSession);
            // selectTest(sqlSession);
            // selectWithPagination(sqlSession, 0, 10);
            deleteTest(sqlSession);
            // Commit Transaction
            sqlSession.commit();
        } catch (IOException e) {
            // Rollback Transaction
            if (sqlSession != null) {
                sqlSession.rollback();
            }
            e.printStackTrace();
        } finally {
            // Close SqlSession
            if (sqlSession != null) {
                sqlSession.close();
            }
        }
    }

    public static void insertTest(SqlSession sqlSession) {
        // Insert data
        User user = new User();
        user.setName("Tom");
        int count = sqlSession.insert("com.oceanbase.samples.mapper.UserMapper.insert", user);
        System.out.println("Insert count: " + count);

    }

    public static void updateTest(SqlSession sqlSession) {
        // Update data
        User user = new User();
        user.setId(1);
        user.setName("Jerry");
        int count = sqlSession.update("com.oceanbase.samples.mapper.UserMapper.update", user);
        System.out.println("Update count: " + count);
    }

    public static void selectTest(SqlSession sqlSession) {
        // Select data
        List<User> user = sqlSession.selectList("com.oceanbase.samples.mapper.UserMapper.selectUser", 1);
        user.stream().filter(Objects::nonNull).forEach(System.out::println);
    }


    public static void selectWithPagination(SqlSession sqlSession, int offset, int pageSize) {
        Map<String, Object> params = new HashMap<>();
        params.put("offset", offset);
        params.put("pageSize", pageSize);
        List<User> users = sqlSession.selectList("com.oceanbase.samples.mapper.UserMapper.selectWithPagination", params);
        users.stream().filter(Objects::nonNull).forEach(System.out::println);
    }

    public static void deleteTest(SqlSession sqlSession) {
        // Delete data
        int count = sqlSession.delete("com.oceanbase.samples.mapper.UserMapper.delete", 3);
        System.out.println("Delete count: " + count);
    }
}
