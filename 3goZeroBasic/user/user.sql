create database if not exists `goZero`; #创建数据库
USE goZero; #选择数据酷

create table if not exists `user` # 创建表
(
    `id`         varchar(24) COLLATE utf8_general_ci  not null,
    `avatar`     varchar(291) COLLATE utf8_general_ci   NOT NULL DEFAULT '',
    `name`       varchar(24) COLLATE utf8_general_ci  not null,
    `phone`      varchar(20) COLLATE utf8_general_ci  not null,
    `password`   varchar(191) COLLATE utf8_general_ci not null,
    `status`     varchar(10) COLLATE utf8_general_ci  not null,
    `created_at` timestamp                              NULL     DEFAULT NULL,
    `update_at`  timestamp                              NULL     DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8
  COLLATE = utf8_general_ci;