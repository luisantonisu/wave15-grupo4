CREATE TABLE IF NOT EXISTS `grupo4`.`countries` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `country_name` VARCHAR(50) UNIQUE,
    PRIMARY KEY(`id`)
);

CREATE TABLE IF NOT EXISTS `grupo4`.`provinces`(
    `id` INT NOT NULL AUTO_INCREMENT,
    `province_name` VARCHAR(50),
    `country_id` INT,
    PRIMARY KEY(`id`),
    UNIQUE(`country_id`, `province_name`),
    KEY `idx_country_id` (`country_id`),
    CONSTRAINT `fx_country_id` FOREIGN KEY (`country_id`) 
    REFERENCES `grupo4`.`countries`(`id`) ON DELETE CASCADE ON UPDATE CASCADE    
);

CREATE TABLE IF NOT EXISTS `grupo4`.`localities` (
    `id` INT NOT NULL,
    `locality_name` VARCHAR(50),
    `province_id` INT,
    PRIMARY KEY(`id`),
    UNIQUE(`province_id`, `locality_name`),
    KEY `idx_province_id` (`province_id`),
    CONSTRAINT `fx_province_id` FOREIGN KEY (`province_id`) 
    REFERENCES `grupo4`.`provinces`(`id`) ON DELETE CASCADE ON UPDATE CASCADE    
);

CREATE TABLE IF NOT EXISTS `grupo4`.`carriers` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `carry_id` VARCHAR(50) UNIQUE NOT NULL,
	`company_name` VARCHAR(50),
	`address` VARCHAR(50),
	`telephone` INT,
	`locality_id` INT,
    PRIMARY KEY(`id`),
	KEY `idx_locality_id` (`locality_id`),
    CONSTRAINT `fx_carry_locality_id` FOREIGN KEY (`locality_id`) 
    REFERENCES `localities`(`id`) ON DELETE CASCADE ON UPDATE CASCADE    
);

CREATE TABLE IF NOT EXISTS `grupo4`.`sellers`(
	`id` INT AUTO_INCREMENT NOT NULL,
	`company_id` VARCHAR(50) UNIQUE NOT NULL,
	`company_name` VARCHAR(50),
	`address` VARCHAR(50),
	`telephone` VARCHAR(50),
	`locality_id` INT,
	PRIMARY KEY(`id`),
    KEY `idx_locality_id` (`locality_id`),
    CONSTRAINT `fx_seller_locality_id` FOREIGN KEY (`locality_id`) 
    REFERENCES `localities`(`id`) ON DELETE CASCADE ON UPDATE CASCADE    
);

CREATE TABLE IF NOT EXISTS `grupo4`.`product_types` (
    `id` INT AUTO_INCREMENT NOT NULL,
    `description` VARCHAR(255),
	PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS `grupo4`.`products`(
	`id` INT AUTO_INCREMENT NOT NULL,
	`product_code` VARCHAR(50) UNIQUE NOT NULL,
	`description` VARCHAR(255),
	`width` DECIMAL(10,2),
	`height` DECIMAL(10,2),
    `length` DECIMAL(10,2),
	`net_weight` DECIMAL(10,2),
	`expiration_rate` DECIMAL(10,2),
    `recommended_freezing_temperature`DECIMAL(10,2),
	`freezing_rate` DECIMAL(10,2),
	`product_type_id` INT,
	`seller_id` INT,
	PRIMARY KEY(id),
    KEY `idx_product_seller_id` (`seller_id`),
	CONSTRAINT `fx_product_seller_id` FOREIGN KEY (`seller_id`) REFERENCES `sellers`(`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    KEY `idx_product_type_id` (`product_type_id`),
	CONSTRAINT `fx_product_type_id` FOREIGN KEY (`product_type_id`) REFERENCES `product_types`(`id`) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS `grupo4`.`warehouses`(
	`id` INT AUTO_INCREMENT NOT NULL,
	`warehouse_code` VARCHAR(50) UNIQUE NOT NULL,
	`address` VARCHAR(50),
	`telephone` INT,
	`minimum_capacity` INT,
	`minimum_temperature` DECIMAL(10,2),
    `locality_id` INT,
    PRIMARY KEY(`id`),
    KEY `idx_locality_id`  (`locality_id`),
	CONSTRAINT `fx_warehouse_locality_id` FOREIGN KEY (`locality_id`) 
    REFERENCES `localities`(`id`) ON DELETE CASCADE ON UPDATE CASCADE    
);

CREATE TABLE IF NOT EXISTS `grupo4`.`employees`(
	`id` INT AUTO_INCREMENT NOT NULL,
	`card_number_id` VARCHAR(50) UNIQUE NOT NULL,
	`first_name` VARCHAR(50),
	`last_name` VARCHAR(50),
	`warehouse_id` INT,
    PRIMARY KEY(id),
	CONSTRAINT `fx_warehouse_id` FOREIGN KEY (`warehouse_id`) REFERENCES `warehouses`(`id`) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS `grupo4`.`buyers` (
    `id` INT AUTO_INCREMENT NOT NULL,
    `card_number_id` VARCHAR(50) UNIQUE NOT NULL,
    `first_name` VARCHAR(50),
    `last_name` VARCHAR(50),
	PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS `grupo4`.`order_status` (
    `id` INT AUTO_INCREMENT NOT NULL,
    `description` VARCHAR(255),
	PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS `grupo4`.`purchase_orders` (
    `id` INT AUTO_INCREMENT NOT NULL,
    `order_number` VARCHAR(50) UNIQUE NOT NULL,
    `order_date` DATETIME(6),
    `tracking_code` VARCHAR(50),
    `buyer_id` INT,
    `carrier_id` INT,
    `order_status_id` INT,
    `warehouse_id` INT,
	PRIMARY KEY(id),
	KEY `idx_buyer_id` (`buyer_id`),
    CONSTRAINT `fx_buyer_id` FOREIGN KEY (`buyer_id`) REFERENCES `buyers`(`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    KEY `idx_carrier_id` (`carrier_id`),
    CONSTRAINT `fx_carrier_id` FOREIGN KEY (`carrier_id`) REFERENCES `carriers`(`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    KEY `idx_order_status_id` (`order_status_id`),
    CONSTRAINT `fx_order_status_id` FOREIGN KEY (`order_status_id`) REFERENCES `order_status`(`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    KEY `idx_warehouse_id` (`warehouse_id`),
    CONSTRAINT `fx_purchase_orders_warehouse_id` FOREIGN KEY (`warehouse_id`) REFERENCES `warehouses`(`id`) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS `grupo4`.`sections`(
	`id` INT AUTO_INCREMENT NOT NULL,
	`section_number` VARCHAR(50) UNIQUE NOT NULL,
	`current_temperature` DECIMAL(10,2),
	`minimum_temperature` DECIMAL(10,2),
	`current_capacity` INT,
	`minimum_capacity` INT,
	`maximum_capacity` INT,
	`warehouse_id` INT,
	`product_type_id` INT,
	PRIMARY KEY(id),
	KEY `idx_section_warehouse_id` (`warehouse_id`),
    CONSTRAINT `fx_section_warehouse_id` FOREIGN KEY (`warehouse_id`) REFERENCES `warehouses`(`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    KEY `idx_product_type_id` (`product_type_id`),
	CONSTRAINT `fx_section_product_type_id` FOREIGN KEY (`product_type_id`) REFERENCES `product_types`(`id`) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS `grupo4`.`product_batches`(
	`id` INT AUTO_INCREMENT NOT NULL,
	`batch_number` VARCHAR(50) UNIQUE NOT NULL,
	`current_quantity` INT,
	`current_temperature` DECIMAL(19,2),
	`due_date` DATETIME(6),
	`intial_quantity` INT,
	`manufacturing_date` DATETIME(6),
	`manufacturing_hour` DATETIME(6),
	`minimum_temperature` DECIMAL(19,2),
    `product_id` INT,
	`section_id` INT,
    PRIMARY KEY(id),
	KEY `idx_product_id` (`product_id`),
    CONSTRAINT `fx_product_batches_product_id` FOREIGN KEY (`product_id`) REFERENCES `products`(`id`) ON DELETE CASCADE ON UPDATE CASCADE,
	KEY `idx_section_id` (`section_id`),
	CONSTRAINT `fx_section_id` FOREIGN KEY (`section_id`) REFERENCES `sections`(`id`) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS `grupo4`.`inbound_orders`(
	`id` INT AUTO_INCREMENT NOT NULL,
	`order_date` DATETIME(6),
	`order_number` VARCHAR(50) UNIQUE NOT NULL,
    `employee_id` INT,
	`product_batch_id` INT,
    `warehouse_id` INT,
    PRIMARY KEY(id),
	KEY `idx_employee_id` (`employee_id`),
    CONSTRAINT `fx_product_batches_employee_id` FOREIGN KEY (`employee_id`) REFERENCES `employees`(`id`) ON DELETE CASCADE ON UPDATE CASCADE,
	KEY `idx_product_batch_id` (`product_batch_id`),
	CONSTRAINT `fx_inbound_order_product_batch_id` FOREIGN KEY (`product_batch_id`) REFERENCES `product_batches`(`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    KEY `idx_warehouse_id` (`warehouse_id`),
	CONSTRAINT `fx_inbound_order_warehouse_id` FOREIGN KEY (`warehouse_id`) REFERENCES `warehouses`(`id`) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS `grupo4`.`product_records`(
	`id` INT AUTO_INCREMENT NOT NULL,
	`last_update_date` DATETIME(6),
	`purchase_price` DECIMAL(10,2),
	`sale_price` DECIMAL(10,2),
	`product_id` INT,
	PRIMARY KEY(id),
	KEY `idx_product_id` (`product_id`),
	CONSTRAINT `fx_product_id` FOREIGN KEY (`product_id`) REFERENCES `products`(`id`) 
	ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS `grupo4`.`order_details`(
	`id` INT AUTO_INCREMENT NOT NULL,
	`cleanliness_status` VARCHAR(50),
	`quantity` INT,
	`temperature` DECIMAL(10,2),
	`product_record_id` INT,
	`purchase_order_id` INT,
	PRIMARY KEY(id),
	KEY `idx_product_record_id` (`product_record_id`),
	CONSTRAINT `fx_product_record_id` FOREIGN KEY (`product_record_id`) REFERENCES `product_records`(`id`) ON DELETE CASCADE ON UPDATE CASCADE,
	KEY `idx_product_order_id` (`purchase_order_id`),
	CONSTRAINT `fx_purchase_order_id` FOREIGN KEY (`purchase_order_id`) REFERENCES `purchase_orders`(`id`) ON DELETE CASCADE ON UPDATE CASCADE
);