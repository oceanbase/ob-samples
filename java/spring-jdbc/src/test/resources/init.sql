CREATE DATABASE IF NOT EXISTS test;
USE test;

CREATE TABLE IF NOT EXISTS staff
(
  id   INT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(50)
  );

INSERT INTO staff (name) VALUES ('Bruce');
INSERT INTO staff (name) VALUES ('Jack');
INSERT INTO staff (name) VALUES ('Tom');
