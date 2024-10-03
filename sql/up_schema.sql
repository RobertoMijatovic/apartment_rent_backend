CREATE TABLE `roles` (
    `id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `title` VARCHAR(64) NOT NULL
);

CREATE TABLE `users` (
    `id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `email` VARCHAR(64) NOT NULL,
    `password` VARCHAR(64) NOT NULL,
    `firstname` VARCHAR(64) NOT NULL,
    `lastname` VARCHAR(64) NOT NULL,
    `roles_id` INT NOT NULL DEFAULT 3,
    FOREIGN KEY (`roles_id`) REFERENCES `roles`(`id`)
);

CREATE TABLE `daily_rents` (
    `id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `name` VARCHAR(128) NOT NULL,
    `address` VARCHAR(128) NOT NULL,
    `description` VARCHAR(128) NOT NULL,
    `price` DECIMAL(32) NOT NULL,
    `status` VARCHAR(128) NOT NULL
    
);

CREATE TABLE `long_term_rents` (
    `id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `name` VARCHAR(128) NOT NULL,
    `address` VARCHAR(128) NOT NULL,
    `description` VARCHAR(128) NOT NULL,
    `price` DECIMAL(32) NOT NULL,
    `status` VARCHAR(128) NOT NULL
);

CREATE TABLE `markets` (
    `id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `name` VARCHAR(128) NOT NULL,
    `address` VARCHAR(128) NOT NULL,
    `description` VARCHAR(128) NOT NULL,
    `owner_name` VARCHAR(128) NOT NULL,
    `rs_type` VARCHAR(128) NOT NULL,
    `size` DECIMAL(32) NOT NULL,
    `price` DECIMAL(32) NOT NULL,
    `status` VARCHAR(128) NOT NULL
);


INSERT INTO `roles` (`id`, `title`)
VALUES 
    (1, 'admin'),
    (2, 'moderator'),
    (3, 'customer');