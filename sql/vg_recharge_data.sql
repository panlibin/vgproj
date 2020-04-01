DROP TABLE IF EXISTS `recharge_order`;
CREATE TABLE  `recharge_order`(
    `local_order_id` bigint(20) unsigned NOT NULL,
    `pf_id` int(11) NOT NULL,
    `pf_order_id` varchar(128) NOT NULL,
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
