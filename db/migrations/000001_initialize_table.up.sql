CREATE TABLE IF NOT EXISTS `question_statuses` (
    `id` INT(11) NOT NULL AUTO_INCREMENT,
    `name` varchar(255) NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `questions` (
    `id` VARCHAR(36) NOT NULL,
    `user_id` VARCHAR(255) NOT NULL,
    `title` varchar(255) NOT NULL,
    `content` text NOT NULL,
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `status_id` int(11) NOT NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`status_id`) REFERENCES `question_statuses` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `tags` (
    `id` VARCHAR(36) NOT NULL,
    `name` varchar(255) NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `question_tags` (
    `question_id` VARCHAR(36) NOT NULL,
    `tag_id` VARCHAR(36) NOT NULL,
    PRIMARY KEY (`question_id`, `tag_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `answers` (
    `id` VARCHAR(36) NOT NULL,
    `user_id` VARCHAR(255) NOT NULL,
    `question_id` VARCHAR(36) NOT NULL,
    `content` text NOT NULL,
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`question_id`) REFERENCES `questions` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO `question_statuses` (`id`, `name`) VALUES
(1, 'open'),
(2, 'closed');
