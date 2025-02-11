USE `grupo4`;

-- COUNTRIES
insert into countries (id, country_name) values (1, 'China');
insert into countries (id, country_name) values (2, 'Indonesia');
insert into countries (id, country_name) values (3, 'Bosnia and Herzegovina');
insert into countries (id, country_name) values (4, 'Peru');
insert into countries (id, country_name) values (5, 'Russia');
insert into countries (id, country_name) values (6, 'Philippines');
insert into countries (id, country_name) values (7, 'Mexico');
insert into countries (id, country_name) values (8, 'Portugal');
insert into countries (id, country_name) values (9, 'France');
insert into countries (id, country_name) values (10, 'Ukraine');
insert into countries (id, country_name) values (11, 'Cape Verde');
insert into countries (id, country_name) values (12, 'Serbia');
insert into countries (id, country_name) values (13, 'Brazil');
insert into countries (id, country_name) values (14, 'Belarus');
insert into countries (id, country_name) values (15, 'Thailand');
insert into countries (id, country_name) values (16, 'Canada');

-- PROVINCES
insert into provinces (id, province_name, country_id) values (1, 'Republika Srpska', 3);
insert into provinces (id, province_name, country_id) values (2, 'Lima', 4);
insert into provinces (id, province_name, country_id) values (3, 'Siberia', 5); 
insert into provinces (id, province_name, country_id) values (4, 'Metro Manila', 7);
insert into provinces (id, province_name, country_id) values (5, 'Yucatán', 8);
insert into provinces (id, province_name, country_id) values (6, 'West Nusa Tenggara', 2);
insert into provinces (id, province_name, country_id) values (7, 'Papua', 2);
insert into provinces (id, province_name, country_id) values (8, 'Hunan', 1);
insert into provinces (id, province_name, country_id) values (9, 'Fujian', 1);

-- LOCALTIES
insert into localities (id, locality_name, province_id) values (1, 'Banja Luka', 1);
insert into localities (id, locality_name, province_id) values (2, 'Callao', 2);
insert into localities (id, locality_name, province_id) values (3, 'Novosibirsk', 3);
insert into localities (id, locality_name, province_id) values (4, 'Quezon City', 4);
insert into localities (id, locality_name, province_id) values (5, 'Mérida', 5);
insert into localities (id, locality_name, province_id) values (6, 'Mataram', 6);
insert into localities (id, locality_name, province_id) values (7, 'Jayapura', 7);
insert into localities (id, locality_name, province_id) values (8, 'Changsha', 8);
insert into localities (id, locality_name, province_id) values (9, 'Fuzhou', 9);

-- CARRIERS
insert into carriers (id, company_id, company_name, address, telephone, locality_id) values (1, 'OHDAS8943', 'Jaxbean', 'PO Box 30728', 376216700, 1);
insert into carriers (id, company_id, company_name, address, telephone, locality_id) values (2, 'HN9F8AD43', 'Vitz', 'Apt 1857', 375360656, 2);
insert into carriers (id, company_id, company_name, address, telephone, locality_id) values (3, 'JHFDS9834', 'Jazzy', 'Apt 1019', 322214906, 3);
insert into carriers (id, company_id, company_name, address, telephone, locality_id) values (4, 'DFSN93434', 'Zooveo', 'PO Box 75080', 367998758, 4);
insert into carriers (id, company_id, company_name, address, telephone, locality_id) values (5, 'KLDHGF983', 'Edgeify', 'Apt 630', 336619985, 5);
insert into carriers (id, company_id, company_name, address, telephone, locality_id) values (6, 'FDSUY9343', 'Kamba', 'Apt 75', 351145212, 6);
insert into carriers (id, company_id, company_name, address, telephone, locality_id) values (7, 'OIFG98432', 'Quinu', '6th Floor', 322106479, 7);
insert into carriers (id, company_id, company_name, address, telephone, locality_id) values (8, '9043MFOIS', 'Brightbean', 'Room 265', 386071621, 8);
insert into carriers (id, company_id, company_name, address, telephone, locality_id) values (9, 'HDF983LKS', 'Gigabox', 'PO Box 85319', 336403786, 9);

-- SELLERS
insert into sellers (id, company_id, company_name, address, telephone, locality_id) values (1, 'OHDAS8943', 'Yadel', 'Room 1596', 358236283, 1);
insert into sellers (id, company_id, company_name, address, telephone, locality_id) values (2, 'HN9F8AD43', 'Yakitri', 'Apt 1151', 339809427, 2);
insert into sellers (id, company_id, company_name, address, telephone, locality_id) values (3, 'JHFDS9834', 'Realmix', 'PO Box 68088', 336385310, 3);
insert into sellers (id, company_id, company_name, address, telephone, locality_id) values (4, 'DFSN93434', 'Jaxbean', 'Room 872', 369063749, 4);
insert into sellers (id, company_id, company_name, address, telephone, locality_id) values (5, 'KLDHGF983', 'Lazz', 'Apt 1689', 314735317, 5);
insert into sellers (id, company_id, company_name, address, telephone, locality_id) values (6, 'FDSUY9343', 'Pixoboo', 'Apt 1334', 327603406, 6);
insert into sellers (id, company_id, company_name, address, telephone, locality_id) values (7, 'OIFG98432', 'Oyoloo', 'Room 1935', 360540176, 7);
insert into sellers (id, company_id, company_name, address, telephone, locality_id) values (8, '9043MFOIS', 'Buzzbean', 'Suite 3', 384226814, 8);
insert into sellers (id, company_id, company_name, address, telephone, locality_id) values (9, 'HDF983LKS', 'Zoomcast', 'Room 1143', 323990798, 9);

-- PRODUCT TYPES
insert into product_types (id, description) values (1, 'leo odio porttitor id consequat in consequat');
insert into product_types (id, description) values (2, 'aliquam quis turpis eget elit sodales sceler');
insert into product_types (id, description) values (3, 'donec dapibus duis lacinia sapien quis liber');
insert into product_types (id, description) values (4, 'vestibulum ante ipsum primis in faucibus orc');
insert into product_types (id, description) values (5, 'amet lobortis sapien sapien non mi integer');
insert into product_types (id, description) values (6, 'dapibus augue vel accumsan tellus nisi eu');
insert into product_types (id, description) values (7, 'elementum ligula vehicula consequat morbi');
insert into product_types (id, description) values (8, 'ut rhoncus aliquet pulvinar sed nisl nunc');
insert into product_types (id, description) values (9, 'at nulla suspendisse potenti cras');
insert into product_types (id, description) values (10, 'lectus suspendisse potenti in eleifend qu');

-- PRODUCTS
insert into products (id, product_code, description, width, height, net_weight, expiration_rate, recommended_freezing_temperature, product_type_id, seller_id) values (1, '37000-164', 'blandit mi in porttitor pede justo eu massa donec', 54.05, 51.38, 74.97, 1, 4, 1, 1);
insert into products (id, product_code, description, width, height, net_weight, expiration_rate, recommended_freezing_temperature, product_type_id, seller_id) values (2, '0115-4411', 'vulputate ut ultrices vel augue vestibulum ant', 23.53, 61.94, 12.07, 5, 1, 2, 2);
insert into products (id, product_code, description, width, height, net_weight, expiration_rate, recommended_freezing_temperature, product_type_id, seller_id) values (3, '63629-2640', 'augue vestibulum ante ipsum primis in faucibus', 84.62, 90.73, 22.04, 3, 7, 3, 3);
insert into products (id, product_code, description, width, height, net_weight, expiration_rate, recommended_freezing_temperature, product_type_id, seller_id) values (4, '43063-028', 'sit amet consectetuer adipiscing elit proin rit', 46.33, 9.26, 58.14, 5, 4, 4, 4);
insert into products (id, product_code, description, width, height, net_weight, expiration_rate, recommended_freezing_temperature, product_type_id, seller_id) values (5, '0615-7661', 'ipsum primis in faucibus orci luctus et ultrices', 63.09, 85.03, 56.1, 1, 5, 5, 5);
insert into products (id, product_code, description, width, height, net_weight, expiration_rate, recommended_freezing_temperature, product_type_id, seller_id) values (6, '65435-0131', 'maecenas pulvinar lobortis est phasellus sit am', 2.06, 76.17, 70.02, 6, 1, 6, 6);
insert into products (id, product_code, description, width, height, net_weight, expiration_rate, recommended_freezing_temperature, product_type_id, seller_id) values (7, '65517-2021', 'gravida sem praesent id massa id nisl ', 8.27, 89.24, 56.95, 10, 8, 7, 7);
insert into products (id, product_code, description, width, height, net_weight, expiration_rate, recommended_freezing_temperature, product_type_id, seller_id) values (8, '49288-0383', 'vel est donec odio justo sollicitudin ut suscis', 37.01, 5.68, 79.37, 4, 3, 8, 8);
insert into products (id, product_code, description, width, height, net_weight, expiration_rate, recommended_freezing_temperature, product_type_id, seller_id) values (9, '11822-8200', 'montes nascetur ridiculus mus vivamus ves sit', 40.52, 24.09, 81.72, 1, 5, 9, 9);
insert into products (id, product_code, description, width, height, net_weight, expiration_rate, recommended_freezing_temperature, product_type_id, seller_id) values (10, '49580-0329', 'in felis eu sapien cursus vestibulum proin eu', 35.14, 43.13, 93.27, 7, 2, 10, 9);

-- WAREHOUSE
insert into warehouses (id, warehouse_code, address, telephone, minimum_capacity, minimum_temperature, locality_id) values (1, '63549-919', 'Suite 63', 355889454, 88, 26.81, 1);
insert into warehouses (id, warehouse_code, address, telephone, minimum_capacity, minimum_temperature, locality_id) values (2, '64725-0114', 'Apt 1720', 320051755, 100, 89.68, 2);
insert into warehouses (id, warehouse_code, address, telephone, minimum_capacity, minimum_temperature, locality_id) values (3, '61657-0966', 'PO Box 82008', 341140284, 38, 20.31, 3);
insert into warehouses (id, warehouse_code, address, telephone, minimum_capacity, minimum_temperature, locality_id) values (4, '37205-338', 'Room 1046', 302513730, 43, 38.3, 4);
insert into warehouses (id, warehouse_code, address, telephone, minimum_capacity, minimum_temperature, locality_id) values (5, '64159-7693', 'Room 1214', 307528335, 97, 46.49, 5);
insert into warehouses (id, warehouse_code, address, telephone, minimum_capacity, minimum_temperature, locality_id) values (6, '49967-206', 'Room 811', 334714960, 77, 84.51, 6);
insert into warehouses (id, warehouse_code, address, telephone, minimum_capacity, minimum_temperature, locality_id) values (7, '50682-507', 'PO Box 44536', 355312009, 35, 87.76, 7);
insert into warehouses (id, warehouse_code, address, telephone, minimum_capacity, minimum_temperature, locality_id) values (8, '50845-0048', '18th Floor', 339875560, 65, 29.66, 8);
insert into warehouses (id, warehouse_code, address, telephone, minimum_capacity, minimum_temperature, locality_id) values (9, '60505-6088', 'Room 39', 328253889, 4, 40.6, 9);
insert into warehouses (id, warehouse_code, address, telephone, minimum_capacity, minimum_temperature, locality_id) values (10, '43857-0170', 'Room 1208', 338368346, 81, 52.09, 1);


-- EMPLOYEES
insert into employees (id, card_number_id, first_name, last_name, warehouse_id) values (1, '085824075', 'Timmy', 'Durnall', 1);
insert into employees (id, card_number_id, first_name, last_name, warehouse_id) values (2, '358323332', 'Bunny', 'Bayless', 2);
insert into employees (id, card_number_id, first_name, last_name, warehouse_id) values (3, '116554553', 'Joellen', 'Fernley', 3);
insert into employees (id, card_number_id, first_name, last_name, warehouse_id) values (4, '270180781', 'Tine', 'Tocque', 4);
insert into employees (id, card_number_id, first_name, last_name, warehouse_id) values (5, '957197037', 'Washington', 'Miskimmon', 5);
insert into employees (id, card_number_id, first_name, last_name, warehouse_id) values (6, '940519619', 'Chick', 'Jakobsson', 6);
insert into employees (id, card_number_id, first_name, last_name, warehouse_id) values (7, '955402491', 'Pattie', 'Kliche', 7);
insert into employees (id, card_number_id, first_name, last_name, warehouse_id) values (8, '459536367', 'Ivor', 'Mahy', 8);
insert into employees (id, card_number_id, first_name, last_name, warehouse_id) values (9, '748887339', 'Pinchas', 'Barette', 9);
insert into employees (id, card_number_id, first_name, last_name, warehouse_id) values (10, '182838497', 'Tore', 'Levick', 10);

-- BUYERS
insert into buyers (id, card_number_id, first_name, last_name) values (1, '497784779', 'Gaspar', 'Urian');
insert into buyers (id, card_number_id, first_name, last_name) values (2, '950806862', 'Vonny', 'Fihelly');
insert into buyers (id, card_number_id, first_name, last_name) values (3, '962230493', 'Netta', 'Francie');
insert into buyers (id, card_number_id, first_name, last_name) values (4, '601419225', 'Aubine', 'Kerfod');
insert into buyers (id, card_number_id, first_name, last_name) values (5, '917411862', 'Marten', 'Kenelin');
insert into buyers (id, card_number_id, first_name, last_name) values (6, '027879388', 'Hollie', 'Padden');
insert into buyers (id, card_number_id, first_name, last_name) values (7, '532249856', 'Bogey', 'Rotham');
insert into buyers (id, card_number_id, first_name, last_name) values (8, '314718801', 'Onida', 'Sisnett');
insert into buyers (id, card_number_id, first_name, last_name) values (9, '659506981', 'Gloriane', 'Godin');
insert into buyers (id, card_number_id, first_name, last_name) values (10, '435252094', 'Stanfield', 'Toffetto');

-- ORDER STATUS
insert into order_status (id, description) values (1, 'pellentesque viverra pede ac diam cras');
insert into order_status (id, description) values (2, 'leo maecenas');
insert into order_status (id, description) values (3, 'magnis dis parturient montes nascetur');
insert into order_status (id, description) values (4, 'vestibulum proin eu');
insert into order_status (id, description) values (5, 'semper est quam pharetra');
insert into order_status (id, description) values (6, 'morbi a ipsum integer a nibh in quis');
insert into order_status (id, description) values (7, 'duis bibendum morbi non quam nec dui');
insert into order_status (id, description) values (8, 'a feugiat parturient monte');
insert into order_status (id, description) values (9, 'elementum eu interdum');
insert into order_status (id, description) values (10, 'convallis nunc proin at turpis a ped');

-- PURCHASE ORDERS
insert into purchase_orders (id, order_number, order_date, tracking_code, buyer_id, carrier_id, order_status_id, warehouse_id) values (1, '763116554', '2024-08-25 06:52:20', '68998-344', 1, 6, 9, 5);
insert into purchase_orders (id, order_number, order_date, tracking_code, buyer_id, carrier_id, order_status_id, warehouse_id) values (2, '680681805', '2024-06-29 05:58:56', '59316-104', 9, 6, 6, 1);
insert into purchase_orders (id, order_number, order_date, tracking_code, buyer_id, carrier_id, order_status_id, warehouse_id) values (3, '755843106', '2024-05-27 09:41:25', '49721-0003', 8, 3, 2, 3);
insert into purchase_orders (id, order_number, order_date, tracking_code, buyer_id, carrier_id, order_status_id, warehouse_id) values (4, '653473971', '2024-04-22 11:14:07', '0327-0011', 6, 7, 2, 7);
insert into purchase_orders (id, order_number, order_date, tracking_code, buyer_id, carrier_id, order_status_id, warehouse_id) values (5, '200035622', '2024-08-12 05:19:06', '62856-705', 1, 4, 5, 5);
insert into purchase_orders (id, order_number, order_date, tracking_code, buyer_id, carrier_id, order_status_id, warehouse_id) values (6, '705474753', '2024-08-31 10:53:04', '51346-131', 8, 2, 5, 7);
insert into purchase_orders (id, order_number, order_date, tracking_code, buyer_id, carrier_id, order_status_id, warehouse_id) values (7, '343449448', '2024-04-29 06:36:40', '57775-001', 8, 5, 7, 6);
insert into purchase_orders (id, order_number, order_date, tracking_code, buyer_id, carrier_id, order_status_id, warehouse_id) values (8, '980802180', '2024-04-08 15:51:48', '44924-007', 9, 3, 3, 2);
insert into purchase_orders (id, order_number, order_date, tracking_code, buyer_id, carrier_id, order_status_id, warehouse_id) values (9, '744251036', '2024-07-07 08:24:43', '36987-3046', 6, 5, 6, 4);
insert into purchase_orders (id, order_number, order_date, tracking_code, buyer_id, carrier_id, order_status_id, warehouse_id) values (10, '725216975', '2024-12-28 20:39:53', '59535-2301', 5, 7, 6, 4);

-- SECTIONS
insert into sections (id, section_number, current_temperature, minimum_temperature, current_capacity, minimum_capacity, maximum_capacity, warehouse_id, product_type_id) values (1, 2562, 5.68, 7.53, 10, 5, 9, 9, 2);
insert into sections (id, section_number, current_temperature, minimum_temperature, current_capacity, minimum_capacity, maximum_capacity, warehouse_id, product_type_id) values (2, 3290, 4.09, 5.15, 61, 10, 10, 6, 9);
insert into sections (id, section_number, current_temperature, minimum_temperature, current_capacity, minimum_capacity, maximum_capacity, warehouse_id, product_type_id) values (3, 8347, 5.76, 7.51, 48, 7, 4, 7, 1);
insert into sections (id, section_number, current_temperature, minimum_temperature, current_capacity, minimum_capacity, maximum_capacity, warehouse_id, product_type_id) values (4, 2152, 5.51, 2.53, 57, 2, 2, 4, 6);
insert into sections (id, section_number, current_temperature, minimum_temperature, current_capacity, minimum_capacity, maximum_capacity, warehouse_id, product_type_id) values (5, 9215, 8.17, 3.94, 2, 8, 7, 4, 4);
insert into sections (id, section_number, current_temperature, minimum_temperature, current_capacity, minimum_capacity, maximum_capacity, warehouse_id, product_type_id) values (6, 8269, 2.6, 3.64, 57, 1, 10, 9, 6);
insert into sections (id, section_number, current_temperature, minimum_temperature, current_capacity, minimum_capacity, maximum_capacity, warehouse_id, product_type_id) values (7, 3704, 6.27, 4.82, 33, 4, 3, 3, 2);
insert into sections (id, section_number, current_temperature, minimum_temperature, current_capacity, minimum_capacity, maximum_capacity, warehouse_id, product_type_id) values (8, 11041, 1.27, 8.52, 84, 3, 1, 6, 9);
insert into sections (id, section_number, current_temperature, minimum_temperature, current_capacity, minimum_capacity, maximum_capacity, warehouse_id, product_type_id) values (9, 2613, 5.64, 7.81, 89, 2, 4, 7, 8);
insert into sections (id, section_number, current_temperature, minimum_temperature, current_capacity, minimum_capacity, maximum_capacity, warehouse_id, product_type_id) values (10, 4792, 3.2, 4.24, 70, 4, 4, 7, 10);

-- PRODUCT BATCHES
insert into product_batches (id, batch_number, current_quantity, current_temperature, due_date, intial_quantity, manufacturing_date, manufacturing_hour, minimum_temperature, product_id, section_id) values (1, '13107-007', 29, 1.09, '2025-08-10 09:40:18', 7, '2024-08-29 04:04:17', '2024-03-20 08:12:38', -5.43, 1, 5);
insert into product_batches (id, batch_number, current_quantity, current_temperature, due_date, intial_quantity, manufacturing_date, manufacturing_hour, minimum_temperature, product_id, section_id) values (2, '65954-010', 52, 1.64, '2024-12-31 13:09:16', 6, '2024-04-10 01:25:46', '2025-01-30 09:41:18', 25.49, 7, 1);
insert into product_batches (id, batch_number, current_quantity, current_temperature, due_date, intial_quantity, manufacturing_date, manufacturing_hour, minimum_temperature, product_id, section_id) values (3, '42808-110', 19, 8.28, '2025-06-03 01:55:02', 3, '2025-02-04 17:17:34', '2024-07-15 01:59:32', 30.91, 7, 4);
insert into product_batches (id, batch_number, current_quantity, current_temperature, due_date, intial_quantity, manufacturing_date, manufacturing_hour, minimum_temperature, product_id, section_id) values (4, '61442-224', 54, 9.82, '2025-08-27 08:32:34', 10, '2024-06-30 17:38:53', '2024-12-18 11:54:40', -2.37, 6, 10);
insert into product_batches (id, batch_number, current_quantity, current_temperature, due_date, intial_quantity, manufacturing_date, manufacturing_hour, minimum_temperature, product_id, section_id) values (5, '0409-4170', 85, 5.96, '2024-11-23 12:15:22', 8, '2024-03-28 21:28:44', '2024-08-17 16:55:50', 28.76, 9, 6);
insert into product_batches (id, batch_number, current_quantity, current_temperature, due_date, intial_quantity, manufacturing_date, manufacturing_hour, minimum_temperature, product_id, section_id) values (6, '0378-6540', 54, 9.64, '2024-03-11 05:59:39', 6, '2025-05-29 04:35:05', '2025-01-11 16:45:49', 25.7, 5, 1);
insert into product_batches (id, batch_number, current_quantity, current_temperature, due_date, intial_quantity, manufacturing_date, manufacturing_hour, minimum_temperature, product_id, section_id) values (7, '68645-483', 92, 2.78, '2025-08-20 23:21:40', 6, '2024-12-31 09:22:17', '2024-10-30 05:45:58', 44.92, 1, 3);
insert into product_batches (id, batch_number, current_quantity, current_temperature, due_date, intial_quantity, manufacturing_date, manufacturing_hour, minimum_temperature, product_id, section_id) values (8, '59137-535', 75, 6.8, '2024-09-09 06:33:24', 2, '2024-09-23 22:28:36', '2024-07-30 11:55:11', -4.25, 7, 10);
insert into product_batches (id, batch_number, current_quantity, current_temperature, due_date, intial_quantity, manufacturing_date, manufacturing_hour, minimum_temperature, product_id, section_id) values (9, '16590-319', 4, 1.92, '2025-07-27 04:57:01', 8, '2025-03-19 03:59:48', '2025-06-19 03:17:34', -1.56, 1, 9);
insert into product_batches (id, batch_number, current_quantity, current_temperature, due_date, intial_quantity, manufacturing_date, manufacturing_hour, minimum_temperature, product_id, section_id) values (10, '43419-361', 10, 6.76, '2025-06-08 16:13:38', 10, '2024-12-22 03:12:07', '2024-03-26 01:29:23', -5.89, 10, 2);

-- INBOUND ORDERS
insert into inbound_orders (id, order_date, order_number, employee_id, product_batch_id, warehouse_id) values (1, '2025-02-19 07:24:19', 8, 6, 2, 10);
insert into inbound_orders (id, order_date, order_number, employee_id, product_batch_id, warehouse_id) values (2, '2024-07-03 09:17:51', 89, 9, 10, 10);
insert into inbound_orders (id, order_date, order_number, employee_id, product_batch_id, warehouse_id) values (3, '2025-01-06 19:52:31', 86, 1, 1, 7);
insert into inbound_orders (id, order_date, order_number, employee_id, product_batch_id, warehouse_id) values (4, '2024-08-10 13:38:51', 85, 10, 5, 9);
insert into inbound_orders (id, order_date, order_number, employee_id, product_batch_id, warehouse_id) values (5, '2025-06-20 15:53:32', 15, 1, 5, 6);
insert into inbound_orders (id, order_date, order_number, employee_id, product_batch_id, warehouse_id) values (6, '2025-02-24 15:27:48', 67, 1, 5, 9);
insert into inbound_orders (id, order_date, order_number, employee_id, product_batch_id, warehouse_id) values (7, '2024-11-25 06:28:37', 25, 5, 2, 4);
insert into inbound_orders (id, order_date, order_number, employee_id, product_batch_id, warehouse_id) values (8, '2025-04-09 07:59:38', 3, 7, 1, 4);
insert into inbound_orders (id, order_date, order_number, employee_id, product_batch_id, warehouse_id) values (9, '2025-03-02 15:55:11', 2, 7, 3, 9);
insert into inbound_orders (id, order_date, order_number, employee_id, product_batch_id, warehouse_id) values (10, '2024-03-06 05:19:44', 66, 6, 5, 9);

-- PRODUCT RECORDS
insert into product_records (id, last_update_date, purchase_price, sale_price, product_id) values (1, '2025-03-22 15:50:30', 89.51, 9.74, 7);
insert into product_records (id, last_update_date, purchase_price, sale_price, product_id) values (2, '2025-01-17 04:37:55', 95.71, 2.38, 2);
insert into product_records (id, last_update_date, purchase_price, sale_price, product_id) values (3, '2024-08-05 02:34:57', 64.26, 4.01, 1);
insert into product_records (id, last_update_date, purchase_price, sale_price, product_id) values (4, '2025-05-15 16:58:34', 77.38, 5.64, 1);
insert into product_records (id, last_update_date, purchase_price, sale_price, product_id) values (5, '2024-11-06 20:48:08', 89.03, 3.07, 5);
insert into product_records (id, last_update_date, purchase_price, sale_price, product_id) values (6, '2024-04-24 23:44:19', 99.48, 6.8, 8);
insert into product_records (id, last_update_date, purchase_price, sale_price, product_id) values (7, '2024-09-27 23:54:32', 37.97, 1.97, 10);
insert into product_records (id, last_update_date, purchase_price, sale_price, product_id) values (8, '2024-10-15 08:00:13', 33.38, 3.19, 5);
insert into product_records (id, last_update_date, purchase_price, sale_price, product_id) values (9, '2024-03-12 03:44:18', 1.55, 4.0, 7);
insert into product_records (id, last_update_date, purchase_price, sale_price, product_id) values (10, '2024-08-04 01:40:29', 5.96, 5.17, 10);

-- ORDER DETAILS
insert into order_details (id, cleanliness_status, quantity, temperature, product_record_id, purchase_order_id) values (1, 'mauris non ligula', 3, -1.85, 5, 7);
insert into order_details (id, cleanliness_status, quantity, temperature, product_record_id, purchase_order_id) values (2, 'suspendisse ornare consequat lectus', 90, 14.08, 9, 10);
insert into order_details (id, cleanliness_status, quantity, temperature, product_record_id, purchase_order_id) values (3, 'pede libero quis orci nullam molestie', 37, 26.61, 1, 9);
insert into order_details (id, cleanliness_status, quantity, temperature, product_record_id, purchase_order_id) values (4, 'suspendisse potenti', 31, 20.13, 7, 6);
insert into order_details (id, cleanliness_status, quantity, temperature, product_record_id, purchase_order_id) values (5, 'erat curabitur gravida', 42, 8.64, 7, 9);
insert into order_details (id, cleanliness_status, quantity, temperature, product_record_id, purchase_order_id) values (6, 'sed ante vivamus tortor duis mattis', 63, -1.84, 10, 5);
insert into order_details (id, cleanliness_status, quantity, temperature, product_record_id, purchase_order_id) values (7, 'rutrum at lorem', 2, 2.57, 7, 8);
insert into order_details (id, cleanliness_status, quantity, temperature, product_record_id, purchase_order_id) values (8, 'sed nisl', 74, 21.68, 10, 10);
insert into order_details (id, cleanliness_status, quantity, temperature, product_record_id, purchase_order_id) values (9, 'eros elementum pellentesque quisque porta volutpat', 46, 24.73, 5, 7);
insert into order_details (id, cleanliness_status, quantity, temperature, product_record_id, purchase_order_id) values (10, 'gravida nisi at', 77, 17.88, 1, 3);
