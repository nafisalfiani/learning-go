CREATE DATABASE game_store;

USE game_store;

CREATE TABLE game (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255),
    genre VARCHAR(255),
    price INT,
    stock INT
);

CREATE TABLE branch (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255),
    location VARCHAR(255)
);

CREATE TABLE sales (
    id INT AUTO_INCREMENT PRIMARY KEY,
    game_id INT,
    branch_id INT,
    date DATE,
    quantity INT,
    FOREIGN KEY (game_id) REFERENCES game(id),
    FOREIGN KEY (branch_id) REFERENCES branch(id)
);
