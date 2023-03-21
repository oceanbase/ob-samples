DROP TABLE IF EXISTS `t_test`;
CREATE TABLE `t_test` (
    `id`   int(10) NOT NULL AUTO_INCREMENT,
    `name` varchar(20) DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE = utf8_bin;

INSERT INTO `t_test` VALUES (default, 'Hello OceanBase');