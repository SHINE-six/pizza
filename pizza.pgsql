SELECT table_name
FROM information_schema.tables
WHERE table_schema = 'public'
ORDER BY table_name;

-- Get the metadata for each column in the table orders
SELECT column_name, data_type, is_nullable
FROM information_schema.columns
WHERE table_name = 'orders';

SELECT *
FROM  pg_settings
WHERE  name = 'max_connections';

SELECT * FROM pg_stat_activity;
SELECT sum(numbackends) FROM pg_stat_database;

DROP TABLE IF EXISTS public.verification_tokens;
DROP TABLE IF EXISTS public.customer;

CREATE TABLE public.verification_tokens (
    id SERIAL PRIMARY KEY, -- serial = auto increment
    email VARCHAR (50) NOT NULL UNIQUE,
    token VARCHAR (50) NOT NULL,
    username VARCHAR (50) NOT NULL,
    password VARCHAR (128) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE public.customer (
    id SERIAL PRIMARY KEY,
    email VARCHAR (50) NOT NULL UNIQUE,
    username VARCHAR (50) NOT NULL,
    password VARCHAR (128) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

DELETE FROM public.verification_tokens;
DELETE FROM public.customers;





CREATE TABLE PizzaBase (
    id SERIAL PRIMARY KEY,
    name VARCHAR (50) NOT NULL UNIQUE,
    base_price DECIMAL (5, 2) NOT NULL
);



CREATE TABLE Size (
    id SERIAL PRIMARY KEY,
    name VARCHAR (50) NOT NULL UNIQUE,
    price_multiplier DECIMAL (3, 1) NOT NULL
);

SELECT * FROM "size" WHERE "size"."id" = '5' ORDER BY "size"."id" LIMIT 1


-- check column metadata for table pizza_bases to know auto increment column name


SELECT * FROM pizza_bases;
SELECT * FROM Sizes;
SELECT * FROM Toppings;
SELECT * FROM Pizzas;
SELECT * FROM pizza_Toppings;
SELECT * FROM orders;
SELECT * FROM delivery_staffs;
SELECT * FROM public.verification_tokens;
SELECT * FROM public.customers;

-- update the value of status column in orders
UPDATE orders
SET status = 'pending'
WHERE id = 1;

DELETE FROM pizza_Toppings;
DELETE FROM Pizzas;
DELETE FROM orders;

INSERT INTO delivery_staffs (name, email) VALUES 
('John Doe', 'johnny@mail.com');

-- Get the metadata for each column in the table orders
SELECT column_name, data_type, character_maximum_length, is_nullable 
FROM information_schema.columns 
WHERE table_name = 'orders';

INSERT INTO "orders" ("created_at","updated_at","deleted_at","customer_id","status","total_price") VALUES ('2024-07-22 22:32:14.063','2024-07-22 22:32:14.063',NULL,8,'pending',0) RETURNING "id"

CREATE TABLE Topping (
    id SERIAL PRIMARY KEY,
    name VARCHAR (50) NOT NULL UNIQUE,
    topping_price DECIMAL (5, 2) NOT NULL
);

DELETE FROM Bases 
WHERE name = 'Extra Thin';

INSERT INTO Bases (name, price) VALUES 
('Extra Thin', 5.00),
('Thin', 5.00),
('Thick', 6.00),
('Cheese Stuffed', 7.00);
INSERT INTO Sizes (name, multiplier) VALUES 
('Small', 1.0),
('Medium', 1.5),
('Large', 2.0);
INSERT INTO Toppings (name, price) VALUES 
('Pepperoni', 1.00),
('Mushrooms', 0.50),
('Onions', 0.50),
('Sausage', 1.00),
('Bacon', 1.00),
('Extra Cheese', 1.00),
('Black Olives', 0.50),
('Green Peppers', 0.50),
('Pineapple', 0.50),
('Spinach', 0.50),
('Broccoli', 0.50),
('Eggplant', 0.50),
('Sundried Tomatoes', 0.50),
('Artichoke Hearts', 0.50),
('Fresh Garlic', 0.50),
('Roasted Red Peppers', 0.50);

CREATE TABLE Pizza (
    id SERIAL PRIMARY KEY,
    pizza_base_id INT NOT NULL,
    size_id INT NOT NULL,
    calculated_price DECIMAL (5, 2) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    FOREIGN KEY (pizza_base_id) REFERENCES PizzaBase(id),
    FOREIGN KEY (size_id) REFERENCES Size(id),
    FOREIGN KEY (customer_id) REFERENCES Customer(id)
);


CREATE TABLE PizzaTopping (
    pizza_id INT NOT NULL,
    topping_id INT NOT NULL,
    PRIMARY KEY (pizza_id, topping_id),
    FOREIGN KEY (pizza_id) REFERENCES Pizza(id),
    FOREIGN KEY (topping_id) REFERENCES Topping(id)
);

CREATE TABLE "Order" (
    id SERIAL PRIMARY KEY,
    customer_id INT NOT NULL,
    total_price DECIMAL (5, 2) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    FOREIGN KEY (customer_id) REFERENCES Customer(id)
);

CREATE TABLE OrderPizza (
    order_id INT NOT NULL,
    pizza_id INT NOT NULL,
    PRIMARY KEY (order_id, pizza_id),
    FOREIGN KEY (order_id) REFERENCES "Order"(id),
    FOREIGN KEY (pizza_id) REFERENCES Pizza(id)
);

-- DROP ALL Public TABLES
DROP TABLE order_pizzas, orders, pizza_bases, pizza_toppings, pizzas, sizes, toppings;


CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- delete all rows from customers table ignore those foreign key constraints
TRUNCATE customers CASCADE;

SELECT * FROM "customers" WHERE "pizzas"."id" = '1' ORDER BY "pizzas"."id" LIMIT 1

SELECT * FROM "pizza_bases"