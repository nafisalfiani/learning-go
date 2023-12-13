DROP DATABASE IF EXISTS marvel;

CREATE DATABASE marvel;

USE marvel;

-- Create the "heroes" table
CREATE TABLE IF NOT EXISTS heroes (
    ID INT AUTO_INCREMENT PRIMARY KEY,
    Name VARCHAR(255) NOT NULL,
    Universe VARCHAR(255),
    Skill VARCHAR(255),
    ImageURL VARCHAR(255)
);

-- Create the "villains" table
CREATE TABLE IF NOT EXISTS villains (
    ID INT AUTO_INCREMENT PRIMARY KEY,
    Name VARCHAR(255) NOT NULL,
    Universe VARCHAR(255),
    ImageURL VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS inventories (
    ID INT AUTO_INCREMENT PRIMARY KEY,
    Name VARCHAR(255) NOT NULL,
    Code VARCHAR(20) NOT NULL UNIQUE,
    Stock INT NOT NULL,
    Description TEXT,
    Status ENUM('active', 'broken') NOT NULL
);
