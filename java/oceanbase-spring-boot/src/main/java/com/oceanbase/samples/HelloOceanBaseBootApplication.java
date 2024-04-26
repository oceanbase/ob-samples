package com.oceanbase.samples;

import com.oceanbase.samples.entity.TestEntity;
import com.oceanbase.samples.repository.TestEntityRepository;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.boot.CommandLineRunner;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.annotation.Bean;

@SpringBootApplication
public class HelloOceanBaseBootApplication {

    private static final Logger log = LoggerFactory.getLogger(HelloOceanBaseBootApplication.class);

    public static void main( String[] args )
    {
        SpringApplication.run(HelloOceanBaseBootApplication.class, args);
    }

    @Bean
    public CommandLineRunner demo(TestEntityRepository repository) {
        return (args) -> {
            // save a few test entities
            repository.save(new TestEntity("Hello OceanBase"));
            repository.save(new TestEntity("OceanBase is a distributed database"));

            // fetch all test entities
            log.info("Test entities found with findAll():");
            log.info("-------------------------------");
            for (TestEntity entity : repository.findAll()) {
                log.info(entity.toString());
            }
            log.info("");

            // fetch an individual test entity by ID
            TestEntity entity = repository.findById(1);
            log.info("Test entity found with findById(1):");
            log.info("--------------------------------");
            log.info(entity.toString());
            log.info("");

            // fetch test entities by test name
            log.info("Test entity found with findByTestName('OceanBase'):");
            log.info("--------------------------------------------");
            repository.findByTestName("OceanBase").forEach(oceanBase -> {
                log.info(oceanBase.toString());
            });
            log.info("");
        };
    }

}
