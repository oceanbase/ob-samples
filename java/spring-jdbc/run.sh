#!/usr/bin/env bash
mvn test -Dtest=com.oceanbase.samples.OceanBaseSpringJdbcApplicationTest exec:java -Dexec.cleanupDaemonThreads=false -Dexec.mainClass=com.oceanbase.samples.OceanBaseSpringJdbcApplicationTest
