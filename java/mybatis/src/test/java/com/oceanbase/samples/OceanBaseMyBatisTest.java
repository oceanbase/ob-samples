package com.oceanbase.samples;


import com.oceanbase.samples.entity.User;
import com.oceanbase.samples.util.SqlSessionUtil;
import org.apache.ibatis.io.Resources;
import org.apache.ibatis.jdbc.ScriptRunner;
import org.apache.ibatis.session.SqlSession;
import org.junit.After;
import org.junit.Before;
import org.junit.Test;

import java.io.InputStreamReader;
import java.sql.Connection;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.Objects;


public class OceanBaseMyBatisTest
{


    @Before
    public void setUp() throws Exception {
        // 执行SQL脚本
            try  {
                Connection connection = SqlSessionUtil.openSession().getConnection();
                ScriptRunner runner = new ScriptRunner(connection);
                runner.runScript(new InputStreamReader(Resources.getResourceAsStream("init.sql")));
                SqlSessionUtil.openSession().commit();
            } catch (Exception e) {
                e.printStackTrace();
            }
    }

    @Test
    public void test() {
        insertTest(SqlSessionUtil.openSession());
        updateTest(SqlSessionUtil.openSession());
        selectTest(SqlSessionUtil.openSession());
        selectWithPagination(SqlSessionUtil.openSession(), 0, 2);
        deleteTest(SqlSessionUtil.openSession());
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

    @After
    public void closeSession() {
        SqlSessionUtil.openSession().close();
    }
}
