package com.oceanbase.samples.springboot.repository;

import com.oceanbase.samples.springboot.entity.StaffEntity;
import org.springframework.data.repository.CrudRepository;
import org.springframework.stereotype.Repository;

import java.util.List;

@Repository
public interface StaffEntityRepository extends CrudRepository<StaffEntity, Integer> {
    List<StaffEntity> findByTestName(String lastName);

    List<StaffEntity> findByTestNameContaining(String testName);

    StaffEntity findById(int id);
}
