DROP TABLE IF EXISTS relationship;
DROP TABLE IF EXISTS retweet;
DROP TABLE IF EXISTS post;
DROP TABLE IF EXISTS user;

-- Crear tabla de usuarios
CREATE TABLE user (
    user_id TEXT PRIMARY KEY,
    username TEXT UNIQUE NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Crear tabla de mensajes
CREATE TABLE post (
    post_id TEXT PRIMARY KEY,
    user_id TEXT NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
);

-- Crear tabla de relaciones (seguidores/seguidos)
CREATE TABLE relationship (
    relationship_id INTEGER PRIMARY KEY AUTOINCREMENT,
    follower_id TEXT NOT NULL,
    followee_id TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (follower_id) REFERENCES users(user_id) ON DELETE CASCADE,
    FOREIGN KEY (followee_id) REFERENCES users(user_id) ON DELETE CASCADE,
    UNIQUE (follower_id, followee_id)
);

-- Crear tabla de retweets
CREATE TABLE retweet (
    retweet_id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id TEXT NOT NULL,
    post_id TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE,
    FOREIGN KEY (post_id) REFERENCES post(post_id) ON DELETE CASCADE,
    UNIQUE (user_id, post_id)
);

-- Insertar datos de ejemplo en la tabla de usuarios
INSERT INTO user (user_id, username, email, password_hash) VALUES
('12345','alice', 'alice@example.com', 'hashed_password_1'),
('12347', 'bob', 'bob@example.com', 'hashed_password_2'),
('12346', 'carol', 'carol@example.com', 'hashed_password_3');

-- Insertar datos de ejemplo en la tabla de mensajes
INSERT INTO post (user_id, post_id, content) VALUES
('12345', 'asdasd','Hello, this is Alice!'),
('12346', 'xzcdfsd','Hi, I am Bob and this is my first post.'),
('12347', 'asdwqw','Carol here, exploring the platform!');

-- Insertar datos de ejemplo en la tabla de relaciones
INSERT INTO relationship (follower_id, followee_id) VALUES
(1, 2), -- Alice sigue a Bob
(2, 3), -- Bob sigue a Carol
(3, 1); -- Carol sigue a Alice

-- Insertar datos de ejemplo en la tabla de retweets
INSERT INTO retweet (user_id, post_id) VALUES
(2, 1), -- Bob retweetea el mensaje de Alice
(3, 2); -- Carol retweetea el mensaje de Bob

-- Verificar datos insertados
SELECT * FROM user;
SELECT * FROM post;
SELECT * FROM relationship;
SELECT * FROM retweet;