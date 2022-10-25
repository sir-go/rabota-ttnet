USE `work`;

DELIMITER ;;

CREATE PROCEDURE `get_by_phone`(IN `phone` varchar(20))
select * from profiles
where json_extract(profiles.profile, '$."основная информация"."телефон"') = phone;;

DELIMITER ;

CREATE TABLE IF NOT EXISTS `profiles` (
    `id` varchar(10) NOT NULL,
    `notified_at` datetime DEFAULT NULL,
    `created_at` datetime DEFAULT NULL,
    `updated_at` datetime DEFAULT NULL,
    `profile` json DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY `created_at` (`created_at`),
    KEY `updated_at` (`updated_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


DELIMITER ;;

CREATE TRIGGER `profiles_bi` BEFORE INSERT ON `profiles` FOR EACH ROW
begin
    set NEW.created_at = now();
end;;

CREATE TRIGGER `profiles_bu` BEFORE UPDATE ON `profiles` FOR EACH ROW
begin
    set NEW.updated_at = now();
end;;

DELIMITER ;
