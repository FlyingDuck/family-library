USE family_lib;

CREATE TABLE IF NOT EXISTS `book` (
  `id` BIGINT UNSIGNED AUTO_INCREMENT,
  `title` VARCHAR(100) NOT NULL COMMENT '标题',
  `sub_title` VARCHAR(100) NOT NULL COMMENT '副标题',
  `author` VARCHAR(30) NOT NULL COMMENT '作者',
  `press` VARCHAR(100) NOT NULL COMMENT '出版社',
  `pages` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '总页数',
  `desc` VARCHAR(200) COMMENT '描述 | 简述',
  `key_words` VARCHAR(150) COMMENT '关键词（技术|编程|数据库）',
  `create_time` DATETIME NOT NULL,
  `update_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

ALTER TABLE book MODIFY pages INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '总页数';
ALTER TABLE book ADD create_time DATETIME NOT NULL;
ALTER TABLE book ADD update_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;

INSERT INTO
book (title, sub_title, author, press, pages, `desc`, key_words, create_time)
VALUES ("国富论", "", "亚当·斯密", "商务印书馆", 398, "", "经济学|文学", NOW());

INSERT INTO
book (title, sub_title, author, press, pages, `desc`, key_words, create_time)
VALUES ("重构", "改善既有代码的设计", "Martin Fowler", "中国工信出版集团", 420, "", "计算机|编程|Java", NOW());
