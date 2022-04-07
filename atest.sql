drop table if exists atest;

CREATE TABLE `atest` (
                         `id` int NOT NULL AUTO_INCREMENT,
                         `code` int NOT NULL DEFAULT '0' COMMENT 'unique code',
                         `type` int NOT NULL DEFAULT '0' COMMENT 'type',
                         PRIMARY KEY (`id`),
                         UNIQUE KEY `uk_code` (`code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;