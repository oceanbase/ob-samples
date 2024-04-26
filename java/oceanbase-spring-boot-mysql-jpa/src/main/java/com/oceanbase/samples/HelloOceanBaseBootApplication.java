package com.oceanbase.samples;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@SpringBootApplication
public class HelloOceanBaseBootApplication {

    private static final Logger log = LoggerFactory.getLogger(HelloOceanBaseBootApplication.class);

    public static void main( String[] args )
    {
        SpringApplication.run(HelloOceanBaseBootApplication.class, args);
    }


}
