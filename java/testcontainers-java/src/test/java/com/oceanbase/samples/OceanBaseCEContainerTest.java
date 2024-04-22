package com.oceanbase.samples;

import org.junit.AfterClass;
import org.junit.BeforeClass;
import org.junit.Test;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.testcontainers.containers.output.Slf4jLogConsumer;
import org.testcontainers.lifecycle.Startables;
import org.testcontainers.oceanbase.OceanBaseCEContainer;

import java.sql.Connection;
import java.sql.PreparedStatement;
import java.sql.ResultSet;
import java.sql.SQLException;
import java.sql.Statement;
import java.util.stream.Stream;

public class OceanBaseCEContainerTest {

    private static final Logger LOG = LoggerFactory.getLogger(OceanBaseCEContainerTest.class);

    public static final OceanBaseCEContainer CONTAINER =
        new OceanBaseCEContainer("oceanbase/oceanbase-ce:latest")
            .withEnv("MODE", "slim")
            .withEnv("FASTBOOT", "true")
            .withLogConsumer(new Slf4jLogConsumer(LOG));

    @BeforeClass
    public static void startContainers() {
        Startables.deepStart(Stream.of(CONTAINER)).join();
        LOG.info(
            "OceanBase docker container started, image: {}, host: {}, sql port: {}, rpc port:{}.",
            CONTAINER.getDockerImageName(),
            CONTAINER.getHost(),
            CONTAINER.getMappedPort(2881),
            CONTAINER.getMappedPort(2882));
    }

    @AfterClass
    public static void closeContainers() {
        CONTAINER.close();
        LOG.info("OceanBase docker container stopped.");
    }

    @Test
    public void test() {
        String database = "testcontainers";
        String table = "person";
        String tableName = String.format("`%s`.`%s`", database, table);

        LOG.info(
            "Try to connect to OceanBase docker container with url: {}.",
            CONTAINER.getJdbcUrl());
        try (Connection connection = CONTAINER.createConnection("?useSSL=false")) {
            LOG.info("Connect to OceanBase docker container successfully.");

            LOG.info("Prepare database and table.");
            try (Statement statement = connection.createStatement()) {
                statement.execute("CREATE DATABASE IF NOT EXISTS " + database);
                statement.execute("USE " + database);
                statement.execute(
                    "CREATE TABLE IF NOT EXISTS " + table + " (name VARCHAR(50), age INT)");
            } catch (SQLException e) {
                throw new RuntimeException(e);
            }

            LOG.info("Insert data to table {}.", tableName);
            try (PreparedStatement ps =
                     connection.prepareStatement("INSERT INTO " + tableName + " values(?, ?)")) {
                ps.setString(1, "Adam");
                ps.setInt(2, 28);
                ps.executeUpdate();
                ps.setString(1, "Eve");
                ps.setInt(2, 26);
                ps.executeUpdate();
            }

            LOG.info("Query rows from {}.", tableName);
            try (PreparedStatement ps =
                     connection.prepareStatement(
                         "SELECT * from " + tableName,
                         ResultSet.TYPE_FORWARD_ONLY,
                         ResultSet.CONCUR_READ_ONLY)) {
                ResultSet rs = ps.executeQuery();
                int count = 0;
                while (rs.next()) {
                    LOG.info("Row {}: name {}, age {}.", count++, rs.getString(1), rs.getInt(2));
                }
                assert count == 2;
            }
        } catch (SQLException e) {
            throw new RuntimeException(e);
        }
    }
}
