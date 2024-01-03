-- DDL for User table
CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    deposit_amount DECIMAL(10, 2) DEFAULT 0.00
);

-- DDL for Product table
CREATE TABLE products (
    product_id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    stock INT NOT NULL,
    price DECIMAL(10, 2) NOT NULL
);

-- DDL for Transaction table
CREATE TABLE transactions (
    transaction_id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(user_id),
    product_id INT REFERENCES products(product_id),
    quantity INT NOT NULL,
    total_amount DECIMAL(10, 2) NOT NULL
);

-- Seeder data for User table
-- pass: password
INSERT INTO users (username, password, deposit_amount) VALUES
    ('john_doe', '$2a$14$IOACGVsh/mVmUHzzlGPyDuubSQ/oQj3R.MqAGtpIHMKGbFjEbcfbi', 1000.00),
    ('jane_smith', '$2a$14$IOACGVsh/mVmUHzzlGPyDuubSQ/oQj3R.MqAGtpIHMKGbFjEbcfbi', 1500.50),
    ('admin_user', '$2a$14$IOACGVsh/mVmUHzzlGPyDuubSQ/oQj3R.MqAGtpIHMKGbFjEbcfbi', 2000.00),
    ('test_user1', '$2a$14$IOACGVsh/mVmUHzzlGPyDuubSQ/oQj3R.MqAGtpIHMKGbFjEbcfbi', 500.00);

-- Seeder data for Product table
INSERT INTO products (name, stock, price) VALUES
    ('Laptop', 50, 1200.00),
    ('Smartphone', 100, 500.00),
    ('Headphones', 30, 80.00),
    ('Camera', 20, 800.00);

-- Seeder data for Transaction table
INSERT INTO transactions (user_id, product_id, quantity, total_amount) VALUES
    (1, 1, 2, 2400.00),
    (2, 3, 1, 80.00),
    (3, 2, 3, 1500.00),
    (4, 4, 1, 800.00),
    (1, 2, 1, 500.00),
    (2, 1, 3, 3600.00),
    (3, 3, 2, 160.00),
    (4, 1, 2, 2400.00);
