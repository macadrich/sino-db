-- +migrate Up
CREATE TABLE IF NOT EXISTS `users` (
  `id` VARCHAR(28) NOT NULL, 
  `email` VARCHAR(100) NOT NULL,
  `name` VARCHAR(20) NOT NULL,
  `age` int(10) UNSIGNED,
  PRIMARY KEY (`id`),
  CONSTRAINT email_unique UNIQUE(email)
)ENGINE = InnoDB DEFAULT CHARSET=utf8mb4;

-- +migrate Down
DROP TABLE `users`;