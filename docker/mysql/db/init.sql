CREATE DATABASE IF NOT EXISTS peingo;
USE peingo;

CREATE TABLE IF NOT EXISTS qnas (
    id         INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
    question   TEXT NOT NULL,
    answer     TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8;