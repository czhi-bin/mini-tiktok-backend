SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for `users`
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
    `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT 'USER ID',
    `user_name` VARCHAR(255) NOT NULL COMMENT 'USER NAME',
    `password` CHAR(60) NOT NULL COMMENT 'USER PASSWORD BCRYPT HASH',
    `avatar` VARCHAR(255) NOT NULL COMMENT 'USER AVATAR URL',
    `background_image` VARCHAR(255) NOT NULL COMMENT 'USER BACKGROUND IMAGE URL',
    `signature` VARCHAR(255) NOT NULL COMMENT 'USER BIO',
    PRIMARY KEY (`id`),
    UNIQUE KEY `user_name` (`user_name`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = 'USERS TABLE';

-- ----------------------------
-- Table structure for `videos`
-- ----------------------------
DROP TABLE IF EXISTS `videos`;
CREATE TABLE `videos` (
    `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT 'VIDEO ID',
    `author_id` BIGINT NOT NULL COMMENT 'AUTHOR ID',
    `cover_url` VARCHAR(255) NOT NULL COMMENT 'VIDEO COVER URL',
    `video_url` VARCHAR(255) NOT NULL COMMENT 'VIDEO URL',
    `title` VARCHAR(255) NOT NULL COMMENT 'VIDEO TITLE',
    `publish_time` TIMESTAMP NOT NULL COMMENT 'VIDEO PUBLISH TIME',
    PRIMARY KEY (`id`),
    FOREIGN KEY (`author_id`) REFERENCES `users` (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = 'VIDEOS TABLE';

-- ----------------------------
-- Table structure for `likes`
-- ----------------------------
DROP TABLE IF EXISTS `likes`;
CREATE TABLE `likes` (
    `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT 'LIKE ID',
    `user_id` BIGINT NOT NULL COMMENT 'USER ID',
    `video_id` BIGINT NOT NULL COMMENT 'VIDEO ID',
    `liked_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'LIKE TIME',
    `deleted_at` TIMESTAMP NULL DEFAULT NULL COMMENT 'UNLIKE TIME',
    PRIMARY KEY (`id`),
    FOREIGN KEY (`user_id`) REFERENCES `users` (`id`),
    FOREIGN KEY (`video_id`) REFERENCES `videos` (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = 'LIKES TABLE';

-- ----------------------------
-- Table structure for `follows`
-- ----------------------------
DROP TABLE IF EXISTS `follows`;
CREATE TABLE `follows` (
    `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT 'FOLLOW ID',
    `user_id` BIGINT NOT NULL COMMENT 'USER ID',
    `follower_id` BIGINT NOT NULL COMMENT 'USER ID OF THE FOLLOWER',
    `followed_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'FOLLOW TIME',
    `deleted_at` TIMESTAMP NULL DEFAULT NULL COMMENT 'UNFOLLOW TIME',
    PRIMARY KEY (`id`),
    FOREIGN KEY (`user_id`) REFERENCES `users` (`id`),
    FOREIGN KEY (`follower_id`) REFERENCES `users` (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = 'FOLLOWS TABLE';

SET FOREIGN_KEY_CHECKS = 1;
