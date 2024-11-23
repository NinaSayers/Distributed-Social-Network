-- Mock Database for Social Media API distnet
CREATE DATABASE distnetdb CHARACTER SET utf8mb4 COLLATE utf8mb4
\c distnetdb

-- Drop existing tables if they exist
DROP TABLE IF EXISTS relationships;
DROP TABLE IF EXISTS messages;
DROP TABLE IF EXISTS users;

-- Users table: Stores user information
CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Messages table: Stores user messages
CREATE TABLE messages (
    message_id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    content TEXT NOT NULL CHECK (LENGTH(content) <= 280), -- Limit to 280 characters
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
);

-- Relationships table: Stores directional relationships (e.g., followers/following)
CREATE TABLE relationships (
    relationship_id SERIAL PRIMARY KEY,
    follower_id INT NOT NULL,
    followee_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (follower_id) REFERENCES users(user_id) ON DELETE CASCADE,
    FOREIGN KEY (followee_id) REFERENCES users(user_id) ON DELETE CASCADE,
    UNIQUE (follower_id, followee_id) -- Prevent duplicate relationships
);

-- Insert mock data for users
INSERT INTO users (username, email, password_hash) VALUES
('alice', 'alice@example.com', 'hashed_password_1'),
('bob', 'bob@example.com', 'hashed_password_2'),
('carol', 'carol@example.com', 'hashed_password_3');

-- Insert mock data for messages
INSERT INTO messages (user_id, content) VALUES
(1, 'Hello, this is Alice!'),
(2, 'Hi, I am Bob and this is my first post.'),
(3, 'Carol here, exploring the platform!');

-- Insert mock data for relationships
INSERT INTO relationships (follower_id, followee_id) VALUES
(1, 2), -- Alice follows Bob
(2, 3), -- Bob follows Carol
(3, 1); -- Carol follows Alice


-- Creating Database User
CREATE USER 'web'@'localhost';
GRANT SELECT, INSERT, UPDATE, DELETE ON disnetdb.
* TO 'web'@'localhost';
ALTER USER 'web'@'localhost' IDENTIFIED BY 'pass';

-- Select data to verify
SELECT * FROM users;
SELECT * FROM messages;
SELECT * FROM relationships;