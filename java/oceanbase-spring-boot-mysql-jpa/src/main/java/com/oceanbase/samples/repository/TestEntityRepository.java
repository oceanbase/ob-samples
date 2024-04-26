package com.oceanbase.samples.repository;

import com.oceanbase.samples.entity.TestEntity;
import org.springframework.data.repository.CrudRepository;
import org.springframework.stereotype.Repository;

import java.util.List;

@Repository
public interface TestEntityRepository extends CrudRepository<TestEntity, Integer> {
    List<TestEntity> findByTestName(String lastName);

    List<TestEntity> findByTestNameContaining(String testName);

    TestEntity findById(int id);
}
