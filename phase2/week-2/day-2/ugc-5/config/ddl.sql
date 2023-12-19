CREATE DATABASE data_center;

USE data_center;

CREATE TABLE user (
    id INT AUTO_INCREMENT PRIMARY KEY,
    email VARCHAR(255) UNIQUE,
    password VARCHAR(255),
    full_name VARCHAR(255),
    age INT,
    occupation VARCHAR(255),
    role VARCHAR(255)
);
