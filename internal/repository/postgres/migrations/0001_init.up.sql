
CREATE TABLE IF NOT EXISTS urls (
    id        SERIAL PRIMARY KEY,
    short     VARCHAR(10)   NOT NULL,
    original  VARCHAR(2048) NOT NULL,
    created_at TIMESTAMP DEFAULT now(),
    expires_at TIMESTAMP
);