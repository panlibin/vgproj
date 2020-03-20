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
