# DROP DATABASE simple_microservice;
CREATE DATABASE simple_microservice;
USE simple_microservice;

SHOW TABLES;

DESC orders;
DESC products;
DESC product_orders;

INSERT INTO `products` (`created_at`, `updated_at`, `deleted_at`, `name`, `stock`, `price_per_item`)
VALUES ('2025-02-01 12:00:00', '2025-02-01 12:00:00', NULL, 'Product 1', 100, 2000),
       ('2025-02-01 12:00:00', '2025-02-01 12:00:00', NULL, 'Product 2', 150, 3000),
       ('2025-02-01 12:00:00', '2025-02-01 12:00:00', NULL, 'Product 3', 200, 2500),
       ('2025-02-01 12:00:00', '2025-02-01 12:00:00', NULL, 'Product 4', 50, 1000),
       ('2025-02-01 12:00:00', '2025-02-01 12:00:00', NULL, 'Product 5', 300, 3500),
       ('2025-02-01 12:00:00', '2025-02-01 12:00:00', NULL, 'Product 6', 250, 2000),
       ('2025-02-01 12:00:00', '2025-02-01 12:00:00', NULL, 'Product 7', 120, 2200),
       ('2025-02-01 12:00:00', '2025-02-01 12:00:00', NULL, 'Product 8', 80, 1800),
       ('2025-02-01 12:00:00', '2025-02-01 12:00:00', NULL, 'Product 9', 60, 1500),
       ('2025-02-01 12:00:00', '2025-02-01 12:00:00', NULL, 'Product 10', 90, 2700);

SELECT *
FROM products;
SELECT *
FROM orders;
SELECT *
FROM products;


SELECT DISTINCT *
FROM orders o
         JOIN product_orders po ON o.id = po.order_id -- Menggunakan `order_id` di `product_orders`
         JOIN products p ON po.product_id = p.id -- Menghubungkan `product_orders` dengan `products`
WHERE o.id = 2; -- Filter untuk memilih order dengan ID 1

