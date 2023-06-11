CREATE TABLE `agents_tag_link`
(
    `id`       int unsigned NOT NULL AUTO_INCREMENT,
    `agent_id` int unsigned NOT NULL,
    `tag_id`   int unsigned NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `agent_id_tag_id` (`agent_id`, `tag_id`),
    KEY `agent_id` (`agent_id`),
    KEY `tag_id` (`tag_id`),
    CONSTRAINT `agents_tag_link_ibfk_1` FOREIGN KEY (`agent_id`) REFERENCES `agents_access` (`id`),
    CONSTRAINT `agents_tag_link_ibfk_2` FOREIGN KEY (`agent_id`) REFERENCES `agents_access` (`id`),
    CONSTRAINT `agents_tag_link_ibfk_3` FOREIGN KEY (`tag_id`) REFERENCES `agents_tags` (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;