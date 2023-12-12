CREATE DATABASE marvel;

USE marvel;

-- Create the "heroes" table
CREATE TABLE heroes (
    ID INT PRIMARY KEY,
    Name VARCHAR(255) NOT NULL,
    Universe VARCHAR(255),
    Skill VARCHAR(255),
    ImageURL VARCHAR(255)
);

-- Create the "villains" table
CREATE TABLE villains (
    ID INT PRIMARY KEY,
    Name VARCHAR(255) NOT NULL,
    Universe VARCHAR(255),
    ImageURL VARCHAR(255)
);
