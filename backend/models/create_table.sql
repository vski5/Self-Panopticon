--
CREATE TABLE `user` (
    `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
    `user_id` BIGINT(20) NOT NULL,
    `username` VARCHAR(64) COLLATE UTF8MB4_GENERAL_CI NOT NULL,
    `password` VARCHAR(64) COLLATE UTF8MB4_GENERAL_CI NOT NULL,
    `email` VARCHAR(64) COLLATE UTF8MB4_GENERAL_CI,
    `gender` TINYINT(4) NOT NULL DEFAULT '0',
    `create_time` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_username` (`username`) USING BTREE,
    UNIQUE KEY `idx_user_id` (`user_id`) USING BTREE
)  ENGINE=INNODB DEFAULT CHARSET=UTF8MB4 COLLATE = UTF8MB4_GENERAL_CI;
-- 排序依据：id，user_id，username，password，email，gender，create_time，update_time

DROP TABLE IF EXISTS `community`;
CREATE TABLE `community` (
     `id` int(11) NOT NULL AUTO_INCREMENT,
     `community_id` int(10) unsigned NOT NULL,
     `community_name` varchar(128) COLLATE utf8mb4_general_ci NOT NULL,
     `introduction` varchar(256) COLLATE utf8mb4_general_ci NOT NULL,
     `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
     `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
     PRIMARY KEY (`id`),
     UNIQUE KEY `idx_community_id` (`community_id`),
     UNIQUE KEY `idx_community_name` (`community_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- 设置初始的用户信息
INSERT INTO `community` VALUES ('1', '1', 'userA', 'A1', '2021-11-01 08:10:10', '2021-11-01 08:10:10');
INSERT INTO `community` VALUES ('2', '2', 'userB', 'B1', '2020-01-01 08:00:00', '2020-01-01 08:00:00');
INSERT INTO `community` VALUES ('3', '3', 'userC', 'C1', '2018-08-07 08:30:00', '2018-08-07 08:30:00');
INSERT INTO `community` VALUES ('4', '4', 'userD', 'D1', '2019-01-01 08:00:00', '2019-01-01 08:00:00');

DROP TABLE IF EXISTS `post`;
CREATE TABLE `post` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `post_id` bigint(20) NOT NULL COMMENT '帖子id',
    `title` varchar(128) COLLATE utf8mb4_general_ci NOT NULL COMMENT '标题',
    `content` varchar(8192) COLLATE utf8mb4_general_ci NOT NULL COMMENT '内容',
    `author_id` bigint(20) NOT NULL COMMENT '作者的用户id',
    `community_id` bigint(20) NOT NULL COMMENT '所属社区',
    `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '帖子状态',
    `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_post_id` (`post_id`),
    KEY `idx_author_id` (`author_id`),
    KEY `idx_community_id` (`community_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;