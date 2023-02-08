# session table
CREATE TABLE IF NOT EXISTS sessions (
    k  VARCHAR(64) PRIMARY KEY NOT NULL DEFAULT '',
    v  BLOB NOT NULL,
    e  BIGINT NOT NULL DEFAULT '0',
    u  TEXT
);

# user table
CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username VARCHAR(16) UNIQUE NOT NULL,
    password CHAR(32) NOT NULL
);
CREATE UNIQUE INDEX username_uiq on users (username);
INSERT INTO users(username, password) VALUES ("liuxu", "liuxu123");

# page url table
CREATE TABLE IF NOT EXISTS page_urls (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title VARCHAR(32) NOT NULL,
    page_url VARCHAR(512) NOT NULL
);