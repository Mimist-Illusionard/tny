
CREATE TABLE IF NOT EXISTS urls (
    id        SERIAL PRIMARY KEY,
    short     VARCHAR(10)   NOT NULL,
    original  VARCHAR(2048) NOT NULL,
    createdAt TIMESTAMP DEFAULT now(),
    expiresAt TIMESTAMP
);