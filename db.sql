CREATE DATABASE IF NOT EXISTS `users` DEFAULT CHARACTER SET utf8mb4

USE `users`;

DROP TABLE IF EXISTS `user_info`;

CREATE TABLE `user_info` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `user_name` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `created_time` timestamp NULL DEFAULT NULL,
  `updated_time` timestamp NULL DEFAULT NULL,
  `deleted_time` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_name` (`user_name`),
  KEY `idx_user_info_deleted_time` (`deleted_time`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4;

LOCK TABLES `user_info` WRITE;
INSERT INTO `user_info` VALUES (0, 'admin', '$2a$10$veGcArz47VGj7l9xN7g2iuT9TF21jLI1YGXarGzvARNdnt4inC9PG', '2018-05-27 16:25:33', '2018-05-27 16:25:33', NULL);
UNLOCK TABLES;