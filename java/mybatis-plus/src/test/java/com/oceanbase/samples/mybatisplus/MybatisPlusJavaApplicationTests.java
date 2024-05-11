package com.oceanbase.samples.mybatisplus;

import org.junit.jupiter.api.BeforeAll;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.core.io.ClassPathResource;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.jdbc.datasource.init.ScriptUtils;
import org.springframework.test.context.DynamicPropertyRegistry;
import org.springframework.test.context.DynamicPropertySource;
import org.testcontainers.junit.jupiter.Container;
import org.testcontainers.junit.jupiter.Testcontainers;
import org.testcontainers.oceanbase.OceanBaseCEContainer;
import org.testcontainers.utility.DockerImageName;

import javax.sql.DataSource;
import java.util.List;

import static org.junit.jupiter.api.Assertions.assertFalse;
import static org.junit.jupiter.api.Assertions.assertNotNull;

@SpringBootTest
@Testcontainers
class MybatisPlusJavaApplicationTests {

    @Container
    public static OceanBaseCEContainer oceanBaseContainer = new OceanBaseCEContainer(DockerImageName.parse("oceanbase/oceanbase-ce:latest"))
        .withEnv("MODE", "slim")
        .withEnv("FASTBOOT", "true");
    @Autowired
    private PersonMapper personMapper;

    @DynamicPropertySource
    static void oceanBaseProperties(DynamicPropertyRegistry registry) {
        registry.add("spring.datasource.url", oceanBaseContainer::getJdbcUrl);
        registry.add("spring.datasource.username", oceanBaseContainer::getUsername);
        registry.add("spring.datasource.password", oceanBaseContainer::getPassword);
        registry.add("spring.datasource.driver-class-name", oceanBaseContainer::getDriverClassName);
    }

    @BeforeAll
    static void setup(@Autowired DataSource dataSource) throws Exception {
        JdbcTemplate jdbcTemplate = new JdbcTemplate(dataSource);
        assertNotNull(jdbcTemplate.getDataSource());
        ScriptUtils.executeSqlScript(jdbcTemplate.getDataSource().getConnection(), new ClassPathResource("init.sql"));
    }

    @Test
    void testSelectList() {
        List<Person> persons = personMapper.selectList(null);
        assertFalse(persons.isEmpty());
    }
}
