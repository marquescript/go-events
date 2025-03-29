CREATE DATABASE IF NOT EXISTS go_events;

CREATE TABLE IF NOT EXISTS users (
    id int auto_increment PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS events (
    id int auto_increment PRIMARY KEY,
    description TEXT NOT NULL,
    date DATE NOT NULL,
    address TEXT NOT NULL,
    user_id int NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)
); 