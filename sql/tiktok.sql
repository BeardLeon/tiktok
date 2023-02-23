/*
 Navicat Premium Data Transfer

 Source Server         : mysql
 Source Server Type    : MySQL
 Source Server Version : 80031
 Source Host           : localhost:3306
 Source Schema         : tiktok

 Target Server Type    : MySQL
 Target Server Version : 80031
 File Encoding         : 65001

 Date: 15/02/2023 20:39:43
*/
/**
  适用于MySQL5.7
 */
SET NAMES utf8;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for comments
-- ----------------------------
DROP TABLE IF EXISTS `comments`;
CREATE TABLE `comments`
(
    `id`           int(0)                                                        NOT NULL AUTO_INCREMENT,
    `user_id`      int(0)                                                        NULL DEFAULT NULL COMMENT '评论发布用户id',
    `video_id`     int(0)                                                        NULL DEFAULT NULL COMMENT '评论视频id',
    `comment_text` text CHARACTER SET utf8 COLLATE utf8_general_ci         NULL COMMENT '评论内容',
    `cancel`       int(0)                                                        NULL DEFAULT 0 COMMENT '默认为0，取消后为1',
    `created_at`   datetime(0)                                                   NULL DEFAULT NULL,
    `updated_at`   datetime(0)                                                   NULL DEFAULT NULL,
    `deleted_at`   varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE,
    INDEX `user_id` (`user_id`) USING BTREE,
    INDEX `video_id` (`video_id`) USING BTREE,
    CONSTRAINT `comments_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
    CONSTRAINT `comments_ibfk_2` FOREIGN KEY (`video_id`) REFERENCES `videos` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  CHARACTER SET = utf8
  COLLATE = utf8_general_ci
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for follows
-- ----------------------------
DROP TABLE IF EXISTS `follows`;
CREATE TABLE `follows`
(
    `id`          int(0)      NOT NULL AUTO_INCREMENT,
    `user_id`     int(0)      NULL DEFAULT NULL COMMENT '用户id',
    `follower_id` int(0)      NULL DEFAULT NULL COMMENT '关注的用户id',
    `cancel`      int(0)      NULL DEFAULT 0 COMMENT '默认为0，取消后为1',
    `created_at`  datetime(0) NULL DEFAULT NULL,
    `updated_at`  datetime(0) NULL DEFAULT NULL,
    `deleted_at`  datetime(0) NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE,
    INDEX `user_id` (`user_id`) USING BTREE,
    INDEX `follower_id` (`follower_id`) USING BTREE,
    CONSTRAINT `follows_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
    CONSTRAINT `follows_ibfk_2` FOREIGN KEY (`follower_id`) REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  CHARACTER SET = utf8
  COLLATE = utf8_general_ci
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for likes
-- ----------------------------
DROP TABLE IF EXISTS `likes`;
CREATE TABLE `likes`
(
    `id`         int(0)      NOT NULL AUTO_INCREMENT,
    `user_id`    int(0)      NULL DEFAULT NULL COMMENT '点赞用户id',
    `video_id`   int(0)      NULL DEFAULT NULL COMMENT '被点赞的视频id',
    `cancel`     int(0)      NULL DEFAULT 0 COMMENT '默认为0，取消后为1',
    `created_at` datetime(0) NULL DEFAULT NULL,
    `updated_at` datetime(0) NULL DEFAULT NULL,
    `deleted_at` datetime(0) NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE,
    INDEX `user_id` (`user_id`) USING BTREE,
    INDEX `video_id` (`video_id`) USING BTREE,
    CONSTRAINT `likes_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
    CONSTRAINT `likes_ibfk_2` FOREIGN KEY (`video_id`) REFERENCES `videos` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  CHARACTER SET = utf8
  COLLATE = utf8_general_ci
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`
(
    `id`         int(0)                                                        NOT NULL AUTO_INCREMENT,
    `name`       varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '用户名',
    `password`   varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '密码',
    `created_at` datetime(0)                                                   NULL DEFAULT NULL COMMENT '创建时间',
    `updated_at` datetime(0)                                                   NULL DEFAULT NULL COMMENT '更新时间',
    `deleted_at` datetime(0)                                                   NULL DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE INDEX `unique_name` (`name`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  CHARACTER SET = utf8
  COLLATE = utf8_general_ci
  ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for videos
-- ----------------------------
DROP TABLE IF EXISTS `videos`;
CREATE TABLE `videos`
(
    `id`         int(0)                                                        NOT NULL AUTO_INCREMENT,
    `author_id`  int(0)                                                        NULL DEFAULT NULL COMMENT '作者id',
    `play_url`   varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '播放url',
    `cover_url`  varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '封面url',
    `title`      varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '视频标题',
    `created_at` datetime(0)                                                   NULL DEFAULT NULL,
    `updated_at` datetime(0)                                                   NULL DEFAULT NULL,
    `deleted_at` datetime(0)                                                   NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE,
    INDEX `author_id` (`author_id`) USING BTREE,
    CONSTRAINT `videos_ibfk_1` FOREIGN KEY (`author_id`) REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  CHARACTER SET = utf8
  COLLATE = utf8_general_ci
  ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
