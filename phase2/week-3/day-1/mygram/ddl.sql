CREATE DATABASE mygram;

-- User Table
CREATE TABLE User (
    id INT PRIMARY KEY,
    username VARCHAR(255),
    email VARCHAR(255),
    password VARCHAR(255),
    age INT,
    created_at DATE,
    updated_at DATE
);

-- Photo Table
CREATE TABLE Photo (
    id INT PRIMARY KEY,
    title VARCHAR(255),
    caption VARCHAR(255),
    photo_url VARCHAR(255),
    user_id INT,
    created_at DATE,
    updated_at DATE,
    FOREIGN KEY (user_id) REFERENCES User(id)
);

-- Comment Table
CREATE TABLE Comment (
    id INT PRIMARY KEY,
    user_id INT,
    photo_id INT,
    message VARCHAR(255),
    created_at DATE,
    updated_at DATE,
    FOREIGN KEY (user_id) REFERENCES User(id),
    FOREIGN KEY (photo_id) REFERENCES Photo(id)
);

-- SocialMedia Table
CREATE TABLE SocialMedia (
    id INT PRIMARY KEY,
    name VARCHAR(255),
    social_media_url TEXT,
    UserId INT,
    FOREIGN KEY (UserId) REFERENCES User(id)
);
