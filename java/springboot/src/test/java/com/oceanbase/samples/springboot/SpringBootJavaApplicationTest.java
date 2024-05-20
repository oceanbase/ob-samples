package com.oceanbase.samples.springboot;

import com.oceanbase.samples.springboot.repository.StaffEntityRepository;
import org.junit.jupiter.api.Test;
import org.springframework.boot.test.context.SpringBootTest;

import javax.annotation.Resource;

@SpringBootTest(classes = SpringBootJavaApplication.class)
public class SpringBootJavaApplicationTest {

    @Resource
    private StaffEntityRepository staffEntityRepository;


    @Test
    public void findByTestName() {
      assert staffEntityRepository.findByTestName("test") != null;
    }

    @Test
    public void findByTestNameContaining() {
        assert staffEntityRepository.findByTestNameContaining("test") != null;
    }

    @Test
    public void findById() {
        boolean b = staffEntityRepository.findById(1) == null;
    }
}
