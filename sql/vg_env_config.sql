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


insert into c_login_server values(1,':8000','127.0.0.1:8001','127.0.0.1:8101','root:root@tcp(127.0.0.1:3306)/vg_login_data?charset=utf8mb4&parseTime=True&loc=Local&multiStatements=true',10,'vgcli','MWIzODNlYTIxYmU5MDAwNjMwNjIxYjRk',0,1);

insert into c_master_server values(1,':8100','127.0.0.1:8101','root:root@tcp(127.0.0.1:3306)/vg_master_data?charset=utf8mb4&parseTime=True&loc=Local&multiStatements=true',10,'MWIzODNlYTIxYmU5MDAwNjMwNjIxYjRk',1);

insert into c_game_server values(1,':9010','127.0.0.1:9011','127.0.0.1:8101','root:root@tcp(127.0.0.1:3306)/vg_game_data_1?charset=utf8mb4&parseTime=True&loc=Local&multiStatements=true',10,'root:root@tcp(127.0.0.1:3306)/vg_game_oa_1?charset=utf8mb4&parseTime=True&loc=Local&multiStatements=true',10,'root:root@tcp(127.0.0.1:3306)/vg_game_config?charset=utf8mb4&parseTime=True&loc=Local&multiStatements=true','MWIzODNlYTIxYmU5MDAwNjMwNjIxYjRk',1);
