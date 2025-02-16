DROP TABLE IF EXISTS relationships;
DROP TABLE IF EXISTS retweets;
DROP TABLE IF EXISTS messages;
DROP TABLE IF EXISTS users;

-- Crear tabla de usuarios
CREATE TABLE users (
    user_id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT UNIQUE NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Crear tabla de mensajes
CREATE TABLE messages (
    message_id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
);

-- Crear tabla de relaciones (seguidores/seguidos)
CREATE TABLE relationships (
    relationship_id INTEGER PRIMARY KEY AUTOINCREMENT,
    follower_id INTEGER NOT NULL,
    followee_id INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (follower_id) REFERENCES users(user_id) ON DELETE CASCADE,
    FOREIGN KEY (followee_id) REFERENCES users(user_id) ON DELETE CASCADE,
    UNIQUE (follower_id, followee_id)
);

-- Crear tabla de retweets
CREATE TABLE retweets (
    retweet_id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    message_id INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE,
    FOREIGN KEY (message_id) REFERENCES messages(message_id) ON DELETE CASCADE,
    UNIQUE (user_id, message_id)
);

-- Insertar datos de ejemplo en la tabla de usuarios
INSERT INTO users (username, email, password_hash) VALUES
('alice', 'alice@example.com', 'hashed_password_1'),
('bob', 'bob@example.com', 'hashed_password_2'),
('carol', 'carol@example.com', 'hashed_password_3');

-- Insertar datos de ejemplo en la tabla de mensajes
INSERT INTO messages (user_id, content) VALUES
(1, 'Hello, this is Alice!'),
(2, 'Hi, I am Bob and this is my first post.'),
(3, 'Carol here, exploring the platform!');

-- Insertar datos de ejemplo en la tabla de relaciones
INSERT INTO relationships (follower_id, followee_id) VALUES
(1, 2), -- Alice sigue a Bob
(2, 3), -- Bob sigue a Carol
(3, 1); -- Carol sigue a Alice

-- Insertar datos de ejemplo en la tabla de retweets
INSERT INTO retweets (user_id, message_id) VALUES
(2, 1), -- Bob retweetea el mensaje de Alice
(3, 2); -- Carol retweetea el mensaje de Bob

-- Verificar datos insertados
SELECT * FROM users;
SELECT * FROM messages;
SELECT * FROM relationships;
SELECT * FROM retweets;