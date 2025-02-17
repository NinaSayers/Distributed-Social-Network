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
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Crear tabla de mensajes
CREATE TABLE post (
    post_id TEXT PRIMARY KEY,
    user_id TEXT NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES user(user_id) ON DELETE CASCADE
);

-- Crear tabla de relaciones (seguidores/seguidos)
CREATE TABLE relationship (
    relationship_id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id TEXT NOT NULL,
    followee_id TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES user(user_id) ON DELETE CASCADE,
    FOREIGN KEY (followee_id) REFERENCES user(user_id) ON DELETE CASCADE,
    UNIQUE (follower_id, followee_id)
);

-- Crear tabla de retweets
CREATE TABLE retweet (
    retweet_id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id TEXT NOT NULL,
    post_id TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES user(user_id) ON DELETE CASCADE,
    FOREIGN KEY (post_id) REFERENCES post(post_id) ON DELETE CASCADE,
    UNIQUE (user_id, post_id)
);

-- Insertar datos de ejemplo en la tabla de usuarios
-- INSERT INTO user (user_id, username, email, password_hash) VALUES
-- ('482XpV3ptfU9n5P8DsX6Pi','Paco', 'paco@gmail.com', '$2a$12$P0HK5HWUlYUt/.lgJMMJn.L5t3aaLb7/K4knFjoyjabT3fxBIxzIO'),
-- ('5aNm6MLjdQnY6eoESE8L6d', 'Fulano', 'fulano@gmail.com', '$2a$12$vxlbBn0KLVidG.wnkEAPyelSvu.g8J7HOMxEIPBQzn0H/5BOU1wW.');


-- -- Insertar datos de ejemplo en la tabla de mensajes
-- INSERT INTO post (user_id, post_id, content) VALUES
-- ('482XpV3ptfU9n5P8DsX6Pi', 'asdasd','Hello, this is Paco!'),
-- ('5aNm6MLjdQnY6eoESE8L6d', 'xzcdfsd','Hi, I am Fulano and this is my first post.');

-- -- Insertar datos de ejemplo en la tabla de relaciones
-- INSERT INTO relationship (follower_id, followee_id) VALUES
-- ('482XpV3ptfU9n5P8DsX6Pi', '5aNm6MLjdQnY6eoESE8L6d'), -- Alice sigue a Bob
-- ('5aNm6MLjdQnY6eoESE8L6d', '482XpV3ptfU9n5P8DsX6Pi'); -- Bob sigue a Carol

-- -- Insertar datos de ejemplo en la tabla de retweets
-- INSERT INTO retweet (user_id, post_id) VALUES
-- ('482XpV3ptfU9n5P8DsX6Pi', 'xzcdfsd'), -- Bob retweetea el mensaje de Alice
-- ('5aNm6MLjdQnY6eoESE8L6d', 'asdasd'); -- Carol retweetea el mensaje de Bob

-- Verificar datos insertados
SELECT * FROM user;
SELECT * FROM post;
SELECT * FROM relationship;
SELECT * FROM retweet;