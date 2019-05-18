SET NAMES 'utf8';

USE `shortLinks`;

CREATE TABLE IF NOT EXISTS `short_links` (
  `hash` varchar(32) NOT NULL,
  `short` varchar(10) NOT NULL,
  `link` text NOT NULL,
  PRIMARY KEY (`hash`),
  INDEX (`short`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
