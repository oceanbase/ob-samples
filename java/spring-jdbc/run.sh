#!/usr/bin/env bash
mvn test -Dtest exec:java -Dexec.cleanupDaemonThreads=false -Dexec.mainClass=com.oceanbase.samples.OceanBaseSpringJdbcApplicationTest
