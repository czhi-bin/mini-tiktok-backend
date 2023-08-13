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
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = 'USER TABLE';

SET FOREIGN_KEY_CHECKS = 1;
