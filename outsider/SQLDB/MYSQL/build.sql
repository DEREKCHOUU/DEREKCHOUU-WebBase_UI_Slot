CREATE DATABASE [IF NOT EXISTS] UI CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci;

CREATE SCHEMA `os`;

CREATE SCHEMA `Device`;

CREATE TABLE `os`.`settings` (
  `ipv4` int(unsigned) NOT NULL,
  `port` int(3) NOT NULL DEFAULT 4041,
  `os` varchar(10) NOT NULL DEFAULT "windows"
)COMMENT='os_setting'; 

CREATE TABLE `Device`.`settings` (
  `id` int(4) UNIQUE PRIMARY KEY,
  `os` varchar(10) NOT NULL DEFAULT "Unix",
  `ip` int(unsigned) NOT NULL,
  `port` int(4) NOT NULL DEFAULT 4045,
  `servicetype` varchar(10) NOT NULL DEFAULT "mysql",
  `sqlac` varchar(32) NOT NULL DEFAULT "./administrator",
  `sqlpw` varchar(32) NOT NULL DEFAULT "P@ssw0rd"
)COMMENT='device_list';

CREATE TABLE `Device`.`Data` (
  `id` int(unsigned) PRIMARY KEY AUTO_INCREMENT,
  `deviceid` int,
  `description` varchar(unsigned),
  `actiontime` datetime
)COMMENT='Data';

ALTER TABLE `Device`.`Data` ADD FOREIGN KEY (`Deviceid`) REFERENCES `Device`.`settings` (`id`);
