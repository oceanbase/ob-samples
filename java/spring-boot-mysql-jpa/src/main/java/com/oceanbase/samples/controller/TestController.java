package com.oceanbase.samples.controller;

import com.oceanbase.samples.entity.TestEntity;
import com.oceanbase.samples.repository.TestEntityRepository;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

import java.util.List;
import java.util.Objects;
import java.util.stream.Collectors;
import java.util.stream.StreamSupport;

@RestController
@RequestMapping("/")
public class TestController {
    private static final Logger log = LoggerFactory.getLogger(TestController.class);

    @Autowired
    private TestEntityRepository testEntityRepository;

    @PostMapping("/save")
    public Boolean save(@RequestBody TestEntity testSaveEntity) {
        TestEntity resultEntity = testEntityRepository.save(testSaveEntity);
        log.info("save result: {}", resultEntity);
        return resultEntity != null;
    }



    @GetMapping("/{id}")
    public TestEntity findById(@PathVariable("id") int id) {
        log.info("find by id: {}", id);
        TestEntity resultEntity = testEntityRepository.findById(id);
        log.info("find result: {}", resultEntity);
        return resultEntity;
    }

    @GetMapping("/findAll")
    public List<TestEntity> findAll() {
        log.info("find all");
        return StreamSupport.stream(testEntityRepository.findAll().spliterator(), false)
            .filter(Objects::nonNull)
            .peek(entity -> log.info(entity.toString()))
            .collect(Collectors.toList());
    }

    @GetMapping("/findByName")
    public List<TestEntity> findByName(@RequestParam("name") String name) {
        log.info("find by name: {}", name);
        return testEntityRepository.findByTestNameContaining(name);
    }

    @DeleteMapping("/{id}")
    public void deleteById(@PathVariable("id") int id) {
        log.info("delete by id: {}", id);
        testEntityRepository.deleteById(id);
    }

    @PutMapping("/update")
    public void update(@RequestBody TestEntity testUpdateEntity) {
        log.info("update: {}", testUpdateEntity);
        testEntityRepository.save(testUpdateEntity);
    }

}
