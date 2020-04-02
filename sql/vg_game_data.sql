DROP TABLE IF EXISTS `global_system`;
CREATE TABLE  `global_system`(
    `open_server_time` datetime NOT NULL,
    `last_daily_refresh_ts` bigint(20) NOT NULL DEFAULT 0
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


DROP TABLE IF EXISTS `player_data`;
CREATE TABLE  `player_data`(
    `player_id` bigint(20) NOT NULL,
    `account_id` bigint(20) NOT NULL,
    `server_id` int(11) NOT NULL,
    `name` char(64) NOT NULL,
    `head` int(11) NOT NULL,
    `sex` int(11) NOT NULL,
    `lev` int(11) NOT NULL,
    `exp` bigint(20) NOT NULL,
    `create_ts` bigint(20) NOT NULL,
    `last_login_ip` char(32) NOT NULL DEFAULT '',
    `last_login_ts` bigint(20) NOT NULL DEFAULT 0,
    `last_logout_ts` bigint(20) NOT NULL DEFAULT 0,
    `last_daily_refresh_ts` bigint(20) NOT NULL DEFAULT 0,
    PRIMARY KEY (`player_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


DROP TABLE IF EXISTS `player_hero`;
CREATE TABLE  `player_hero`(
    `player_id` bigint(20) NOT NULL,
    `hero_id` int(11) NOT NULL,
    `star` int(11) NOT NULL,
    `lev` bigint(20) NOT NULL,
    PRIMARY KEY (`player_id`, `hero_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


DROP TABLE IF EXISTS `player_property`;
CREATE TABLE  `player_property`(
    `player_id` bigint(20) NOT NULL,
    `prop_id` int(11) NOT NULL,
    `prop_value` bigint(20) NOT NULL,
    `update_ts` bigint(20) NOT NULL,
    PRIMARY KEY (`player_id`, `prop_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


DROP TABLE IF EXISTS `player_item`;
CREATE TABLE  `player_item`(
    `player_id` bigint(20) NOT NULL,
    `item_id` int(11) NOT NULL,
    `item_num` bigint(20) NOT NULL,
    PRIMARY KEY (`player_id`, `item_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


DROP TABLE IF EXISTS `player_mail_ctrl`;
CREATE TABLE  `player_mail_ctrl`(
    `player_id` bigint(20) NOT NULL,
    `global_mail_id` int(11) NOT NULL,
    PRIMARY KEY (`player_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


DROP TABLE IF EXISTS `player_mail_list`;
CREATE TABLE  `player_mail_list`(
    `player_id` bigint(20) NOT NULL,
    `mail_id` bigint(20) NOT NULL,
    `source` int(11) NOT NULL,
    `source_ext` varchar(128) NOT NULL,
    `ts` bigint(20) NOT NULL,
    `first_type` int(11) NOT NULL,
    `second_type` int(11) NOT NULL,
    `title` varchar(255) NOT NULL,
    `title_params` varchar(255) NOT NULL,
    `content` varchar(2048) NOT NULL,
    `content_params` varchar(2048) NOT NULL,
    `attachments` varchar(4096) NOT NULL,
    `is_new` int(11) NOT NULL,
    `is_read` int(11) NOT NULL,
    `is_got` int(11) NOT NULL,
    PRIMARY KEY (`player_id`,`mail_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


DROP TABLE IF EXISTS `global_mail_list`;
CREATE TABLE  `global_mail_list`(
    `global_mail_id` int(11) NOT NULL,
    `source` int(11) NOT NULL,
    `source_ext` varchar(128) NOT NULL,
    `ts` bigint(20) NOT NULL,
    `first_type` int(11) NOT NULL,
    `second_type` int(11) NOT NULL,
    `title` varchar(255) NOT NULL,
    `title_params` varchar(255) NOT NULL,
    `content` varchar(2048) NOT NULL,
    `content_params` varchar(2048) NOT NULL,
    `attachments` varchar(4096) NOT NULL,
    `vip_lev_limit` int(11) NOT NULL,
    PRIMARY KEY (`global_mail_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


DROP TABLE IF EXISTS `custom_lan`;
CREATE TABLE  `custom_lan`(
    `custom_key` char(64) NOT NULL,
    `lan` int(11) NOT NULL,
    `custom_val` varchar(4096) NOT NULL,
    PRIMARY KEY (`custom_key`,`lan`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


DROP TABLE IF EXISTS `player_vip`;
CREATE TABLE  `player_vip`(
    `player_id` bigint(20) NOT NULL,
    `vip_lev` int(11) NOT NULL,
    `vip_exp` int(11) NOT NULL,
    PRIMARY KEY (`player_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


DROP TABLE IF EXISTS `player_vip_gift`;
CREATE TABLE  `player_vip_gift`(
    `player_id` bigint(20) NOT NULL,
    `vip_lev` int(11) NOT NULL,
    PRIMARY KEY (`player_id`,`vip_lev`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


DROP TABLE IF EXISTS `player_settings`;
CREATE TABLE  `player_settings`(
    `player_id` bigint(20) NOT NULL,
    `language` int(11) NOT NULL,
    PRIMARY KEY (`player_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

