drop database if exists vg_env_config;
drop database if exists vg_game_config;
drop database if exists vg_master_data;
drop database if exists vg_login_data;
drop database if exists vg_game_data_1;
drop database if exists vg_game_oa_1;
drop database if exists vg_recharge_data;

create database vg_env_config default character set utf8mb4 collate utf8mb4_general_ci;
create database vg_game_config character set utf8mb4 collate utf8mb4_general_ci;
create database vg_master_data default character set utf8mb4 collate utf8mb4_general_ci;
create database vg_login_data default character set utf8mb4 collate utf8mb4_general_ci;
create database vg_game_data_1 default character set utf8mb4 collate utf8mb4_general_ci;
create database vg_game_oa_1 default character set utf8mb4 collate utf8mb4_general_ci;
create database vg_recharge_data default character set utf8mb4 collate utf8mb4_general_ci;

use vg_env_config;

DROP TABLE IF EXISTS `c_login_server`;
CREATE TABLE  `c_login_server`(
    `server_id` int(11) NOT NULL,
    `listen_addr` char(32) NOT NULL,
    `cluster_addr` char(32) NOT NULL,
    `master_addr` char(32) NOT NULL,
    `data_dsn` varchar(255) NOT NULL,
    `data_db_conn_num` int(11) NOT NULL DEFAULT 10,
    `client_key` char(32) NOT NULL,
    `auth_key` char(32) NOT NULL,
    `check_time` int(11) NOT NULL,
    `debug` int(11) NOT NULL,
    PRIMARY KEY (`server_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


DROP TABLE IF EXISTS `c_game_server`;
CREATE TABLE  `c_game_server`(
    `server_id` int(11) NOT NULL,
    `listen_addr` char(32) NOT NULL,
    `cluster_addr` char(32) NOT NULL,
    `master_addr` char(32) NOT NULL,
    `data_dsn` varchar(255) NOT NULL,
    `data_db_conn_num` int(11) NOT NULL DEFAULT 10,
    `oa_dsn` varchar(255) NOT NULL,
    `oa_db_conn_num` int(11) NOT NULL DEFAULT 10,
    `conf_dsn` varchar(255) NOT NULL,
    `auth_key` char(32) NOT NULL,
    `debug` int(11) NOT NULL,
    PRIMARY KEY (`server_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


DROP TABLE IF EXISTS `c_master_server`;
CREATE TABLE  `c_master_server`(
    `server_id` int(11) NOT NULL,
    `listen_addr` char(32) NOT NULL,
    `cluster_addr` char(32) NOT NULL,
    `data_dsn` varchar(255) NOT NULL,
    `data_db_conn_num` int(11) NOT NULL DEFAULT 10,
    `auth_key` char(32) NOT NULL,
    `debug` int(11) NOT NULL,
    PRIMARY KEY (`server_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


DROP TABLE IF EXISTS `c_recharge_server`;
CREATE TABLE  `c_recharge_server`(
    `server_id` int(11) NOT NULL,
    `listen_addr` char(32) NOT NULL,
    `cluster_addr` char(32) NOT NULL,
    `master_addr` char(32) NOT NULL,
    `data_dsn` varchar(255) NOT NULL,
    `data_db_conn_num` int(11) NOT NULL DEFAULT 10,
    `client_key` char(32) NOT NULL,
    `auth_key` char(32) NOT NULL,
    `check_time` int(11) NOT NULL,
    `debug` int(11) NOT NULL,
    PRIMARY KEY (`server_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


insert into c_login_server values(1,':8000','127.0.0.1:8001','127.0.0.1:8101','root:root@tcp(127.0.0.1:3306)/vg_login_data?charset=utf8mb4&parseTime=True&loc=Local&multiStatements=true',10,'vgcli','MWIzODNlYTIxYmU5MDAwNjMwNjIxYjRk',0,1);

insert into c_master_server values(1,':8100','127.0.0.1:8101','root:root@tcp(127.0.0.1:3306)/vg_master_data?charset=utf8mb4&parseTime=True&loc=Local&multiStatements=true',10,'MWIzODNlYTIxYmU5MDAwNjMwNjIxYjRk',1);

insert into c_game_server values(1,':9010','127.0.0.1:9011','127.0.0.1:8101','root:root@tcp(127.0.0.1:3306)/vg_game_data_1?charset=utf8mb4&parseTime=True&loc=Local&multiStatements=true',10,'root:root@tcp(127.0.0.1:3306)/vg_game_oa_1?charset=utf8mb4&parseTime=True&loc=Local&multiStatements=true',10,'root:root@tcp(127.0.0.1:3306)/vg_game_config?charset=utf8mb4&parseTime=True&loc=Local&multiStatements=true','MWIzODNlYTIxYmU5MDAwNjMwNjIxYjRk',1);

insert into c_recharge_server values(1,':8200','127.0.0.1:8201','127.0.0.1:8101','root:root@tcp(127.0.0.1:3306)/vg_recharge_data?charset=utf8mb4&parseTime=True&loc=Local&multiStatements=true',10,'vgcli','MWIzODNlYTIxYmU5MDAwNjMwNjIxYjRk',0,1);


use vg_recharge_data;

DROP TABLE IF EXISTS `recharge_order`;
CREATE TABLE  `recharge_order`(
    `local_order_id` bigint(20) unsigned NOT NULL,
    `pf_id` int(11) NOT NULL,
    `pf_order_id` char(64) NOT NULL,
    `receive_date` datetime NOT NULL,
    `source` varchar(32) NOT NULL,
    `currency` varchar(32) NOT NULL,
    `amount` bigint(20) NOT NULL,
    `pf_product_id` varchar(32) NOT NULL,
    `local_product_id` int(11) NOT NULL,
    `account_id` bigint(20) NOT NULL,
    `server_id` int(11) NOT NULL,
    `player_id` bigint(20) NOT NULL,
    `status` int(11) NOT NULL,
    `sandbox` int(11) NOT NULL,
    PRIMARY KEY (`local_order_id`),
    INDEX `pf_order`(`pf_id`,`pf_order_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


DROP TABLE IF EXISTS `sdk_param`;
CREATE TABLE  `sdk_param`(
    `id` int(11) NOT NULL,
    `name` varchar(32) NOT NULL,
    `app_id` varchar(64) NOT NULL,
    `key_1` varchar(255) NOT NULL,
    `key_2` varchar(255) NOT NULL,
    `key_3` varchar(255) NOT NULL,
    `key_4` varchar(255) NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


use vg_master_data;

DROP TABLE IF EXISTS `global_player_name`;
CREATE TABLE  `global_player_name`(
    `name` char(64) NOT NULL,
    PRIMARY KEY (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


DROP TABLE IF EXISTS `global_guild_name`;
CREATE TABLE  `global_guild_name`(
    `name` char(64) NOT NULL,
    PRIMARY KEY (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


DROP TABLE IF EXISTS `node_list`;
CREATE TABLE  `node_list`(
    `server_type` int(11) NOT NULL,
    `server_id` varchar(16000),
    `ip` char(64) NOT NULL,
    PRIMARY KEY (`ip`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


use vg_login_data;

DROP TABLE IF EXISTS `account_info`;
CREATE TABLE  `account_info`(
    `account_id` bigint(20) NOT NULL,
    `password` varchar(64) NOT NULL,
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

use vg_game_data_1;

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

use vg_game_oa_1;

DROP TABLE IF EXISTS `log_create_character`;
CREATE TABLE  `log_create_character`(
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `player_id` bigint(20) NOT NULL,
    `account_id` bigint(20) NOT NULL,
    `server_id` int(11) NOT NULL,
    `name` char(64) NOT NULL,
    `head` int(11) NOT NULL,
    `sex` int(11) NOT NULL,
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

