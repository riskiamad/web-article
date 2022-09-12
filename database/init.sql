START TRANSACTION;

-- create database
CREATE DATABASE IF NOT EXISTS web_article;
USE `web_article`;



-- create table article
CREATE TABLE `article` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `author` varchar(100) DEFAULT '',
    `title` varchar(250) DEFAULT '',
    `body` longtext,
    `created_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- create database
CREATE DATABASE IF NOT EXISTS web_article_test;
USE `web_article_test`;

-- create table article
CREATE TABLE `article` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `author` varchar(100) DEFAULT '',
    `title` varchar(250) DEFAULT '',
    `body` longtext,
    `created_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

COMMIT;