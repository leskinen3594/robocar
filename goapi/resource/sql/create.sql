-- Create database
CREATE DATABASE IF NOT EXISTS `robotic`
    CHARACTER SET UTF8 COLLATE utf8_unicode_ci;


-- USE database
USE `robotic`;


-- Create Robots table
CREATE TABLE `Robots`(
    `rbt_id`        INT             NOT NULL        AUTO_INCREMENT,
    `mac_addr`      CHAR(17)        NOT NULL        UNIQUE          COLLATE utf8_unicode_ci,
    
    PRIMARY KEY(rbt_id)
)ENGINE=INNODB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;


-- Create Users table
CREATE TABLE `Users`(
    `uid`           INT             NOT NULL        AUTO_INCREMENT,
	`urbt_id`       INT             NOT NULL,
    `uname`         VARCHAR(25)     NOT NULL        UNIQUE          COLLATE utf8_unicode_ci,
    `passwd`        VARCHAR(25)     NOT NULL                        COLLATE utf8_unicode_ci,
    `email`         VARCHAR(36)     NOT NULL        UNIQUE          COLLATE utf8_unicode_ci,
    `phone`         CHAR(10)        NULL                            COLLATE utf8_unicode_ci,
    `is_staff`      TINYINT(1)      NOT NULL        DEFAULT '0',

    PRIMARY KEY(uid),
    FOREIGN KEY(urbt_id) REFERENCES `Robots`(rbt_id)
)ENGINE=INNODB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;


-- Create APIKeys table
CREATE TABLE `APIkeys`(
    `api_id`          INT           NOT NULL        AUTO_INCREMENT,
	`api_uid`		  INT           NOT NULL        UNIQUE,
    `api_key`         CHAR(32)      NOT NULL        UNIQUE          COLLATE utf8_unicode_ci,

    PRIMARY KEY(api_id),
    FOREIGN KEY(api_uid) REFERENCES `Users`(uid)
)ENGINE=INNODB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;