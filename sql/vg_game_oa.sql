DROP TABLE IF EXISTS `log_create_character`;
CREATE TABLE  `log_create_character`(
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `player_id` bigint(20) NOT NULL,
    `account_id` bigint(20) NOT NULL,
    `server_id` int(11) NOT NULL,
    `name` char(64) NOT NULL,
    `head` int(11) NOT NULL,
    `time` datetime NOT NULL,
    PRIMARY KEY (`id`),
    INDEX `player_id`(`player_id`),
    INDEX `account_id`(`account_id`),
    INDEX `time`(`time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


DROP TABLE IF EXISTS `log_player_login`;
CREATE TABLE  `log_player_login`(
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `player_id` bigint(20) NOT NULL,
    `account_id` bigint(20) NOT NULL,
    `ip` char(32) NOT NULL,
    `time` datetime NOT NULL,
    PRIMARY KEY (`id`),
    INDEX `player_id`(`player_id`),
    INDEX `account_id`(`account_id`),
    INDEX `time`(`time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


DROP TABLE IF EXISTS `log_player_logout`;
CREATE TABLE  `log_player_logout`(
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `player_id` int(11) NOT NULL,
    `account_id` bigint(20) NOT NULL,
    `online_ts` bigint(20) NOT NULL,
    `time` datetime NOT NULL,
    PRIMARY KEY (`id`),
    INDEX `player_id`(`player_id`),
    INDEX `account_id`(`account_id`),
    INDEX `time`(`time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


DROP TABLE IF EXISTS `log_item`;
CREATE TABLE  `log_item`(
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `player_id` bigint(20) NOT NULL,
    `source` int(11) NOT NULL,
    `source_ext` varchar(128) NOT NULL,
    `item_id` int(11) NOT NULL,
    `delta_num` bigint(20) NOT NULL,
    `total_num` bigint(20) NOT NULL,
    `time` datetime NOT NULL,
    PRIMARY KEY (`id`),
    INDEX `player_id`(`player_id`),
    INDEX `time`(`time`),
    INDEX `source`(`source`),
    INDEX `item_id`(`item_id`),
    INDEX `delta_num`(`delta_num`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


DROP TABLE IF EXISTS `log_lev_up`;
CREATE TABLE  `log_lev_up`(
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `player_id` bigint(20) NOT NULL,
    `old_lv` int(11) NOT NULL,
    `new_lv` int(11) NOT NULL,
    `time` datetime NOT NULL,
    PRIMARY KEY (`id`),
    INDEX `player_id`(`player_id`),
    INDEX `time`(`time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
