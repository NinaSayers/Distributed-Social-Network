-- Seleccionar la base de datos antes de usarla o crearla
-- CREATE DATABASE IF NOT EXISTS distnetdb;
-- USE distnetdb;

-- Asegurarse de eliminar tablas existentes para evitar conflictos
DROP TABLE IF EXISTS relationships;
DROP TABLE IF EXISTS post;
DROP TABLE IF EXISTS user;

-- Crear tabla de usuarios
CREATE TABLE user (
    user_id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Crear tabla de mensajes
CREATE TABLE post (
    post_id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    content TEXT NOT NULL CHECK (CHAR_LENGTH(content) <= 280), -- Limitar a 280 caracteres
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
);

-- Crear tabla de relaciones (seguidores/seguidos)
CREATE TABLE relationships (
    relationship_id INT AUTO_INCREMENT PRIMARY KEY,
    follower_id INT NOT NULL,
    followee_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (follower_id) REFERENCES users(user_id) ON DELETE CASCADE,
    FOREIGN KEY (followee_id) REFERENCES users(user_id) ON DELETE CASCADE,
    UNIQUE (follower_id, followee_id) -- Evitar duplicados
);

-- Crear tabla de retweets
CREATE TABLE retweets (
    retweet_id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    post_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE,
    FOREIGN KEY (post_id) REFERENCES post(post_id) ON DELETE CASCADE,
    UNIQUE (user_id, post_id) -- Un usuario no puede retweetear el mismo mensaje dos veces
);

-- Insertar datos de ejemplo en la tabla de usuarios
INSERT INTO users (username, email, password_hash) VALUES
('alice', 'alice@example.com', 'hashed_password_1'),
('bob', 'bob@example.com', 'hashed_password_2'),
('carol', 'carol@example.com', 'hashed_password_3');

-- Insertar datos de ejemplo en la tabla de mensajes
INSERT INTO post (user_id, content) VALUES
(1, 'Hello, this is Alice!'),
(2, 'Hi, I am Bob and this is my first post.'),
(3, 'Carol here, exploring the platform!');

-- Insertar datos de ejemplo en la tabla de relaciones
INSERT INTO relationships (follower_id, followee_id) VALUES
(1, 2), -- Alice sigue a Bob
(2, 3), -- Bob sigue a Carol
(3, 1); -- Carol sigue a Alice

-- Insertar datos de ejemplo en la tabla de retweets (ejemplos)
INSERT INTO retweets (user_id, post_id) VALUES
(2,1), -- Bob retweetea el mensaje de Alice
(3,2); -- Carol retweetea el mensaje de Bob

-- Verificar datos insertados
SELECT * FROM user;
SELECT * FROM post;
SELECT * FROM relationships;
SELECT * FROM retweets;