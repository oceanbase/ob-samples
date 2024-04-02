#!/usr/bin/env bash
mvn clean install exec:java -Dexec.cleanupDaemonThreads=false -Dexec.mainClass=com.oceanbase.example.InsertAndSelectExample
mvn clean install exec:java -Dexec.cleanupDaemonThreads=false -Dexec.mainClass=com.oceanbase.example.OceanBaseClientTest
