CREATE DATABASE IF NOT EXISTS test;
USE test;

CREATE TABLE IF NOT EXISTS person
(
  id   INT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(50),
  age  INT
);

INSERT INTO person (name, age)
VALUES ('Alice', 30);
INSERT INTO person (name, age)
VALUES ('Bob', 25);
INSERT INTO person (name, age)
VALUES ('Charlie', 35);
