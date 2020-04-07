DROP TABLE IF EXISTS `account_info`;
CREATE TABLE  `account_info`(
    `account_id` bigint(20) NOT NULL,
    `password` varchar(64) NOT NULL,
    `salt` varchar(16) NOT NULL,
    `create_time` datetime NOT NULL,
    `is_ban` int(11) NOT NULL DEFAULT 0,
    `ban_ts` bigint(20) NOT NULL DEFAULT 0,
    `ban_type` int(11) NOT NULL DEFAULT 0,
    `ban_duration` bigint(20) NOT NULL DEFAULT 0,
    `online_server` int(11) NOT NULL,
    PRIMARY KEY (`account_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


DROP TABLE IF EXISTS `account_name`;
CREATE TABLE  `account_name`(
    `login_type` int(11) NOT NULL,
    `account_name` char(64) NOT NULL,
    `account_id` bigint(20) NOT NULL,
    `create_time` datetime NOT NULL,
    PRIMARY KEY (`login_type`,`account_name`),
    INDEX `account_id`(`account_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


DROP TABLE IF EXISTS `character_info`;
CREATE TABLE  `character_info`(
    `player_id` bigint(20) NOT NULL,
    `account_id` bigint(20) NOT NULL,
    `server_id` int(11) NOT NULL,
    `name` char(64) NOT NULL,
    `combat` bigint(20) NOT NULL,
    `update_ts` bigint(20) NOT NULL,
    PRIMARY KEY (`player_id`),
    INDEX `account_id`(`account_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


DROP TABLE IF EXISTS `server_list`;
CREATE TABLE  `server_list`(
    `server_id` int(11) NOT NULL,
    `name` varchar(128) NOT NULL,
    `status` int(11) NOT NULL,
    `addr` varchar(128) NOT NULL, 
    PRIMARY KEY (`server_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


insert into server_list values(1,'test',0,'192.168.1.31:9010');
